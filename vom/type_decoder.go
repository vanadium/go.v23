package vom

import (
	"fmt"
	"io"
	"sync"

	"v.io/v23/vdl"
)

// TypeDecoder manages the receipt and unmarshalling of types from the other
// side of a connection.
type TypeDecoder struct {
	// The type encoder uses a 2-lock strategy for decoding. We use typeMu to lock
	// type definitions, and use buildMu to allow only one worker to build types at
	// a time. This is for simplifying the workflow and avoid unnecessary blocking
	// for type lookups.
	typeMu   sync.RWMutex
	idToType map[typeId]*vdl.Type // GUARDED_BY(typeMu)

	buildMu  sync.Mutex
	err      error               // GUARDED_BY(buildMu)
	idToWire map[typeId]wireType // GUARDED_BY(buildMu)
	dec      *Decoder            // GUARDED_BY(buildMu)
}

// NewTypeDecoder returns a new TypeDecoder that reads from the given reader.
// The TypeDecoder understands all wire type formats generated by the TypeEncoder.
func NewTypeDecoder(r io.Reader) (*TypeDecoder, error) {
	buf := newDecbuf(r)
	if err := readMagicByte(buf); err != nil {
		return nil, err
	}
	return newTypeDecoder(buf), nil
}

func newTypeDecoder(buf *decbuf) *TypeDecoder {
	d := &TypeDecoder{
		idToType: make(map[typeId]*vdl.Type),
		idToWire: make(map[typeId]wireType),
		dec:      newDecoder(buf, nil),
	}
	return d
}

// LookupType returns the type for tid. If the type is not yet available,
// this will wait until it arrives and is built.
func (d *TypeDecoder) lookupType(tid typeId) (*vdl.Type, error) {
	if tt := d.lookupKnownType(tid); tt != nil {
		return tt, nil
	}

	d.buildMu.Lock()
	defer d.buildMu.Unlock()

	// Check again to avoid a race.
	if tt := d.lookupKnownType(tid); tt != nil {
		return tt, nil
	}

	// If an error occurred during previous decoding, we do not process
	// any more. Just return the error.
	if d.err != nil {
		return nil, d.err
	}

	// If the wire type is not available yet, read it now.
	if _, exists := d.idToWire[tid]; !exists {
		if err := d.readWireType(tid); err != nil {
			d.err = err
			return nil, err
		}
	}

	// Now we have the wire type corresponding to tid, build it now.
	if err := d.buildType(tid); err != nil {
		d.err = err
		return nil, err
	}
	return d.lookupKnownType(tid), nil
}

// addWireType adds the wire type wt with the type id tid.
func (d *TypeDecoder) addWireType(tid typeId, wt wireType) error {
	d.buildMu.Lock()
	err := d.addWireTypeBuildLocked(tid, wt)
	d.buildMu.Unlock()
	return err
}

func (d *TypeDecoder) addWireTypeBuildLocked(tid typeId, wt wireType) error {
	if tid < WireIdFirstUserType {
		return fmt.Errorf("vom: type %q id %d invalid, the min user type id is %d", wt, tid, WireIdFirstUserType)
	}
	// TODO(toddw): Allow duplicates according to some heuristic (e.g. only
	// identical, or only if the later one is a "superset", etc).
	if dup := d.lookupKnownType(tid); dup != nil {
		return fmt.Errorf("vom: type %q id %d already defined as %q", wt, tid, dup)
	}
	if dup := d.idToWire[tid]; dup != nil {
		return fmt.Errorf("vom: type %q id %d already defined as %q", wt, tid, dup)
	}
	d.idToWire[tid] = wt
	return nil
}

func (d *TypeDecoder) lookupKnownType(tid typeId) *vdl.Type {
	if tt := bootstrapIdToType[tid]; tt != nil {
		return tt
	}
	d.typeMu.RLock()
	tt := d.idToType[tid]
	d.typeMu.RUnlock()
	return tt
}

// readWireType reads and decode wire types until it meets the corresponding
// wire type to tid.
func (d *TypeDecoder) readWireType(tid typeId) error {
	for {
		var wt wireType
		curTypeId, err := d.dec.decodeWireType(&wt)
		if err != nil {
			return err
		}
		// Add the wire type and wake up waiters.
		if err := d.addWireTypeBuildLocked(curTypeId, wt); err != nil {
			return err
		}
		if curTypeId == tid {
			return nil
		}
	}
}

// buildType builds the type from the given wire type.
func (d *TypeDecoder) buildType(tid typeId) error {
	builder := vdl.TypeBuilder{}
	pending := make(map[typeId]vdl.PendingType)
	pt, err := d.makeType(tid, &builder, pending)
	if err != nil {
		return err
	}
	if !builder.Build() {
		// TODO(toddw): Change TypeBuilder.Build() to directly return the error.
		_, err := pt.Built()
		return err
	}
	types := make(map[typeId]*vdl.Type)
	for tid, pt := range pending {
		tt, err := pt.Built()
		if err != nil {
			return err
		}
		types[tid] = tt
	}
	// Add the types to idToType map.
	d.typeMu.Lock()
	for tid, tt := range types {
		d.idToType[tid] = tt
	}
	d.typeMu.Unlock()
	return nil
}

// makeType makes the pending type from its wire type representation.
func (d *TypeDecoder) makeType(tid typeId, builder *vdl.TypeBuilder, pending map[typeId]vdl.PendingType) (vdl.PendingType, error) {
	wt := d.idToWire[tid]
	if wt == nil {
		return nil, fmt.Errorf("vom: unknown type id %d", tid)
	}
	// Make the type from its wireType representation. First remove it from
	// dt.idToWire, and add it to pending, so that subsequent lookups will get the
	// pending type. Eventually the built type will be added to dt.idToType.
	delete(d.idToWire, tid)
	if name := wt.(wireTypeGeneric).TypeName(); name != "" {
		// Named types may be recursive, so we must create the named type first and
		// add it to pending, before we make the base type. The base type may refer
		// back to this named type, and will find it in pending.
		namedType := builder.Named(name)
		pending[tid] = namedType
		if wtNamed, ok := wt.(wireTypeNamedT); ok {
			// This is a NamedType pointing at a base type.
			baseType, err := d.lookupOrMakeType(wtNamed.Value.Base, builder, pending)
			if err != nil {
				return nil, err
			}
			namedType.AssignBase(baseType)
			return namedType, nil
		}
		// This isn't NamedType, but has a non-empty name.
		baseType, err := d.makeBaseType(wt, builder, pending)
		if err != nil {
			return nil, err
		}
		namedType.AssignBase(baseType)
		return namedType, nil
	}
	// Unnamed types are made directly from their base type.  It's fine to update
	// pending after making the base type, since there's no way to create a
	// recursive type based solely on unnamed vdl.
	baseType, err := d.makeBaseType(wt, builder, pending)
	if err != nil {
		return nil, err
	}
	pending[tid] = baseType
	return baseType, nil
}

func (d *TypeDecoder) makeBaseType(wt wireType, builder *vdl.TypeBuilder, pending map[typeId]vdl.PendingType) (vdl.PendingType, error) {
	switch wt := wt.(type) {
	case wireTypeNamedT:
		return nil, fmt.Errorf("vom: NamedType has empty name: %v", wt)
	case wireTypeEnumT:
		enumType := builder.Enum()
		for _, label := range wt.Value.Labels {
			enumType.AppendLabel(label)
		}
		return enumType, nil
	case wireTypeArrayT:
		elemType, err := d.lookupOrMakeType(wt.Value.Elem, builder, pending)
		if err != nil {
			return nil, err
		}
		return builder.Array().AssignElem(elemType).AssignLen(int(wt.Value.Len)), nil
	case wireTypeListT:
		elemType, err := d.lookupOrMakeType(wt.Value.Elem, builder, pending)
		if err != nil {
			return nil, err
		}
		return builder.List().AssignElem(elemType), nil
	case wireTypeSetT:
		keyType, err := d.lookupOrMakeType(wt.Value.Key, builder, pending)
		if err != nil {
			return nil, err
		}
		return builder.Set().AssignKey(keyType), nil
	case wireTypeMapT:
		keyType, err := d.lookupOrMakeType(wt.Value.Key, builder, pending)
		if err != nil {
			return nil, err
		}
		elemType, err := d.lookupOrMakeType(wt.Value.Elem, builder, pending)
		if err != nil {
			return nil, err
		}
		return builder.Map().AssignKey(keyType).AssignElem(elemType), nil
	case wireTypeStructT:
		structType := builder.Struct()
		for _, field := range wt.Value.Fields {
			fieldType, err := d.lookupOrMakeType(field.Type, builder, pending)
			if err != nil {
				return nil, err
			}
			structType.AppendField(field.Name, fieldType)
		}
		return structType, nil
	case wireTypeUnionT:
		unionType := builder.Union()
		for _, field := range wt.Value.Fields {
			fieldType, err := d.lookupOrMakeType(field.Type, builder, pending)
			if err != nil {
				return nil, err
			}
			unionType.AppendField(field.Name, fieldType)
		}
		return unionType, nil
	case wireTypeOptionalT:
		elemType, err := d.lookupOrMakeType(wt.Value.Elem, builder, pending)
		if err != nil {
			return nil, err
		}
		return builder.Optional().AssignElem(elemType), nil
	default:
		return nil, fmt.Errorf("vom: unknown wire type definition %v", wt)
	}
}

func (d *TypeDecoder) lookupOrMakeType(tid typeId, builder *vdl.TypeBuilder, pending map[typeId]vdl.PendingType) (vdl.TypeOrPending, error) {
	if tt := d.lookupKnownType(tid); tt != nil {
		return tt, nil
	}
	if p, ok := pending[tid]; ok {
		return p, nil
	}
	return d.makeType(tid, builder, pending)
}
