// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Package: binary

// Package binary defines types for describing executable binaries.
package binary

import (
	"fmt"
	"v.io/v23/vdl"
)

var _ = __VDLInit() // Must be first; see __VDLInit comments for details.

//////////////////////////////////////////////////
// Type definitions

// Description describes a binary. Binaries are named and have been
// determined to run on some set of profiles. The mechanism for
// determing profiles is specifically not specified and left to the
// implementation of the interface that generates the description.
type Description struct {
	// Name is the Object name of the application binary that can
	// be used to fetch the actual binary from a content server.
	Name string
	// Profiles is a set of names of compatible profiles.  Each
	// name can either be an Object name that resolves to a
	// Profile, or can be the profile's label, e.g.:
	//
	//   "profiles/google/cluster/diskfull"
	//   "linux-media"
	//
	// Application developers can specify compatible profiles by
	// hand, but we also want to be able to automatically derive
	// the matching profiles from examining the binary itself
	// (e.g. that's what Build.Describe() does).
	Profiles map[string]bool
}

func (Description) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/services/binary.Description"`
}) {
}

func (m *Description) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}
	var4 := (m.Name == "")
	if var4 {
		if err := fieldsTarget1.ZeroField("Name"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Name")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}
			if err := fieldTarget3.FromString(string(m.Name), tt.NonOptional().Field(0).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
				return err
			}
		}
	}
	var var7 bool
	if len(m.Profiles) == 0 {
		var7 = true
	}
	if var7 {
		if err := fieldsTarget1.ZeroField("Profiles"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget5, fieldTarget6, err := fieldsTarget1.StartField("Profiles")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}

			mapTarget8, err := fieldTarget6.StartMap(tt.NonOptional().Field(1).Type, len(m.Profiles))
			if err != nil {
				return err
			}
			for key10, value12 := range m.Profiles {
				keyTarget9, err := mapTarget8.StartKey()
				if err != nil {
					return err
				}
				if err := keyTarget9.FromString(string(key10), tt.NonOptional().Field(1).Type.Key()); err != nil {
					return err
				}
				valueTarget11, err := mapTarget8.FinishKeyStartField(keyTarget9)
				if err != nil {
					return err
				}
				if err := valueTarget11.FromBool(bool(value12), tt.NonOptional().Field(1).Type.Elem()); err != nil {
					return err
				}
				if err := mapTarget8.FinishField(keyTarget9, valueTarget11); err != nil {
					return err
				}
			}
			if err := fieldTarget6.FinishMap(mapTarget8); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget5, fieldTarget6); err != nil {
				return err
			}
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *Description) MakeVDLTarget() vdl.Target {
	return &DescriptionTarget{Value: m}
}

type DescriptionTarget struct {
	Value          *Description
	nameTarget     vdl.StringTarget
	profilesTarget __VDLTarget1_map
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *DescriptionTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {

	if ttWant := vdl.TypeOf((*Description)(nil)).Elem(); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	return t, nil
}
func (t *DescriptionTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "Name":
		t.nameTarget.Value = &t.Value.Name
		target, err := &t.nameTarget, error(nil)
		return nil, target, err
	case "Profiles":
		t.profilesTarget.Value = &t.Value.Profiles
		target, err := &t.profilesTarget, error(nil)
		return nil, target, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct v.io/v23/services/binary.Description", name)
	}
}
func (t *DescriptionTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *DescriptionTarget) ZeroField(name string) error {
	switch name {
	case "Name":
		t.Value.Name = ""
		return nil
	case "Profiles":
		t.Value.Profiles = map[string]bool(nil)
		return nil
	default:
		return fmt.Errorf("field %s not in struct v.io/v23/services/binary.Description", name)
	}
}
func (t *DescriptionTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

// map[string]bool
type __VDLTarget1_map struct {
	Value      *map[string]bool
	currKey    string
	currElem   bool
	keyTarget  vdl.StringTarget
	elemTarget vdl.BoolTarget
	vdl.TargetBase
	vdl.MapTargetBase
}

func (t *__VDLTarget1_map) StartMap(tt *vdl.Type, len int) (vdl.MapTarget, error) {

	if ttWant := vdl.TypeOf((*map[string]bool)(nil)); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	*t.Value = make(map[string]bool)
	return t, nil
}
func (t *__VDLTarget1_map) StartKey() (key vdl.Target, _ error) {
	t.currKey = ""
	t.keyTarget.Value = &t.currKey
	target, err := &t.keyTarget, error(nil)
	return target, err
}
func (t *__VDLTarget1_map) FinishKeyStartField(key vdl.Target) (field vdl.Target, _ error) {
	t.currElem = false
	t.elemTarget.Value = &t.currElem
	target, err := &t.elemTarget, error(nil)
	return target, err
}
func (t *__VDLTarget1_map) FinishField(key, field vdl.Target) error {
	(*t.Value)[t.currKey] = t.currElem
	return nil
}
func (t *__VDLTarget1_map) FinishMap(elem vdl.MapTarget) error {
	if len(*t.Value) == 0 {
		*t.Value = nil
	}

	return nil
}

func (x Description) VDLIsZero() (bool, error) {
	if x.Name != "" {
		return false, nil
	}
	if len(x.Profiles) != 0 {
		return false, nil
	}
	return true, nil
}

func (x Description) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(vdl.TypeOf((*Description)(nil)).Elem()); err != nil {
		return err
	}
	if x.Name != "" {
		if err := enc.NextField("Name"); err != nil {
			return err
		}
		if err := enc.StartValue(vdl.TypeOf((*string)(nil))); err != nil {
			return err
		}
		if err := enc.EncodeString(x.Name); err != nil {
			return err
		}
		if err := enc.FinishValue(); err != nil {
			return err
		}
	}
	if len(x.Profiles) != 0 {
		if err := enc.NextField("Profiles"); err != nil {
			return err
		}
		if err := __VDLWriteAnon_map_1(enc, x.Profiles); err != nil {
			return err
		}
	}
	if err := enc.NextField(""); err != nil {
		return err
	}
	return enc.FinishValue()
}

func __VDLWriteAnon_map_1(enc vdl.Encoder, x map[string]bool) error {
	if err := enc.StartValue(vdl.TypeOf((*map[string]bool)(nil))); err != nil {
		return err
	}
	if err := enc.SetLenHint(len(x)); err != nil {
		return err
	}
	for key, elem := range x {
		if err := enc.NextEntry(false); err != nil {
			return err
		}
		if err := enc.StartValue(vdl.TypeOf((*string)(nil))); err != nil {
			return err
		}
		if err := enc.EncodeString(key); err != nil {
			return err
		}
		if err := enc.FinishValue(); err != nil {
			return err
		}
		if err := enc.StartValue(vdl.TypeOf((*bool)(nil))); err != nil {
			return err
		}
		if err := enc.EncodeBool(elem); err != nil {
			return err
		}
		if err := enc.FinishValue(); err != nil {
			return err
		}
	}
	if err := enc.NextEntry(true); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *Description) VDLRead(dec vdl.Decoder) error {
	*x = Description{}
	if err := dec.StartValue(); err != nil {
		return err
	}
	if (dec.StackDepth() == 1 || dec.IsAny()) && !vdl.Compatible(vdl.TypeOf(*x), dec.Type()) {
		return fmt.Errorf("incompatible struct %T, from %v", *x, dec.Type())
	}
	for {
		f, err := dec.NextField()
		if err != nil {
			return err
		}
		switch f {
		case "":
			return dec.FinishValue()
		case "Name":
			if err := dec.StartValue(); err != nil {
				return err
			}
			var err error
			if x.Name, err = dec.DecodeString(); err != nil {
				return err
			}
			if err := dec.FinishValue(); err != nil {
				return err
			}
		case "Profiles":
			if err := __VDLReadAnon_map_1(dec, &x.Profiles); err != nil {
				return err
			}
		default:
			if err := dec.SkipValue(); err != nil {
				return err
			}
		}
	}
}

func __VDLReadAnon_map_1(dec vdl.Decoder, x *map[string]bool) error {
	if err := dec.StartValue(); err != nil {
		return err
	}
	if (dec.StackDepth() == 1 || dec.IsAny()) && !vdl.Compatible(vdl.TypeOf(*x), dec.Type()) {
		return fmt.Errorf("incompatible map %T, from %v", *x, dec.Type())
	}
	var tmpMap map[string]bool
	if len := dec.LenHint(); len > 0 {
		tmpMap = make(map[string]bool, len)
	}
	for {
		switch done, err := dec.NextEntry(); {
		case err != nil:
			return err
		case done:
			*x = tmpMap
			return dec.FinishValue()
		}
		var key string
		{
			if err := dec.StartValue(); err != nil {
				return err
			}
			var err error
			if key, err = dec.DecodeString(); err != nil {
				return err
			}
			if err := dec.FinishValue(); err != nil {
				return err
			}
		}
		var elem bool
		{
			if err := dec.StartValue(); err != nil {
				return err
			}
			var err error
			if elem, err = dec.DecodeBool(); err != nil {
				return err
			}
			if err := dec.FinishValue(); err != nil {
				return err
			}
		}
		if tmpMap == nil {
			tmpMap = make(map[string]bool)
		}
		tmpMap[key] = elem
	}
}

// PartInfo holds information describing a binary part.
type PartInfo struct {
	// Checksum holds the hex-encoded MD5 checksum of the binary part.
	Checksum string
	// Size holds the binary part size in bytes.
	Size int64
}

func (PartInfo) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/services/binary.PartInfo"`
}) {
}

func (m *PartInfo) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}
	var4 := (m.Checksum == "")
	if var4 {
		if err := fieldsTarget1.ZeroField("Checksum"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Checksum")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}
			if err := fieldTarget3.FromString(string(m.Checksum), tt.NonOptional().Field(0).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
				return err
			}
		}
	}
	var7 := (m.Size == int64(0))
	if var7 {
		if err := fieldsTarget1.ZeroField("Size"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget5, fieldTarget6, err := fieldsTarget1.StartField("Size")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}
			if err := fieldTarget6.FromInt(int64(m.Size), tt.NonOptional().Field(1).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget5, fieldTarget6); err != nil {
				return err
			}
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *PartInfo) MakeVDLTarget() vdl.Target {
	return &PartInfoTarget{Value: m}
}

type PartInfoTarget struct {
	Value          *PartInfo
	checksumTarget vdl.StringTarget
	sizeTarget     vdl.Int64Target
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *PartInfoTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {

	if ttWant := vdl.TypeOf((*PartInfo)(nil)).Elem(); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	return t, nil
}
func (t *PartInfoTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "Checksum":
		t.checksumTarget.Value = &t.Value.Checksum
		target, err := &t.checksumTarget, error(nil)
		return nil, target, err
	case "Size":
		t.sizeTarget.Value = &t.Value.Size
		target, err := &t.sizeTarget, error(nil)
		return nil, target, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct v.io/v23/services/binary.PartInfo", name)
	}
}
func (t *PartInfoTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *PartInfoTarget) ZeroField(name string) error {
	switch name {
	case "Checksum":
		t.Value.Checksum = ""
		return nil
	case "Size":
		t.Value.Size = int64(0)
		return nil
	default:
		return fmt.Errorf("field %s not in struct v.io/v23/services/binary.PartInfo", name)
	}
}
func (t *PartInfoTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

func (x PartInfo) VDLIsZero() (bool, error) {
	return x == PartInfo{}, nil
}

func (x PartInfo) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(vdl.TypeOf((*PartInfo)(nil)).Elem()); err != nil {
		return err
	}
	if x.Checksum != "" {
		if err := enc.NextField("Checksum"); err != nil {
			return err
		}
		if err := enc.StartValue(vdl.TypeOf((*string)(nil))); err != nil {
			return err
		}
		if err := enc.EncodeString(x.Checksum); err != nil {
			return err
		}
		if err := enc.FinishValue(); err != nil {
			return err
		}
	}
	if x.Size != 0 {
		if err := enc.NextField("Size"); err != nil {
			return err
		}
		if err := enc.StartValue(vdl.TypeOf((*int64)(nil))); err != nil {
			return err
		}
		if err := enc.EncodeInt(x.Size); err != nil {
			return err
		}
		if err := enc.FinishValue(); err != nil {
			return err
		}
	}
	if err := enc.NextField(""); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *PartInfo) VDLRead(dec vdl.Decoder) error {
	*x = PartInfo{}
	if err := dec.StartValue(); err != nil {
		return err
	}
	if (dec.StackDepth() == 1 || dec.IsAny()) && !vdl.Compatible(vdl.TypeOf(*x), dec.Type()) {
		return fmt.Errorf("incompatible struct %T, from %v", *x, dec.Type())
	}
	for {
		f, err := dec.NextField()
		if err != nil {
			return err
		}
		switch f {
		case "":
			return dec.FinishValue()
		case "Checksum":
			if err := dec.StartValue(); err != nil {
				return err
			}
			var err error
			if x.Checksum, err = dec.DecodeString(); err != nil {
				return err
			}
			if err := dec.FinishValue(); err != nil {
				return err
			}
		case "Size":
			if err := dec.StartValue(); err != nil {
				return err
			}
			var err error
			if x.Size, err = dec.DecodeInt(64); err != nil {
				return err
			}
			if err := dec.FinishValue(); err != nil {
				return err
			}
		default:
			if err := dec.SkipValue(); err != nil {
				return err
			}
		}
	}
}

//////////////////////////////////////////////////
// Const definitions

const MissingChecksum = ""
const MissingSize = int64(-1)

var __VDLInitCalled bool

// __VDLInit performs vdl initialization.  It is safe to call multiple times.
// If you have an init ordering issue, just insert the following line verbatim
// into your source files in this package, right after the "package foo" clause:
//
//    var _ = __VDLInit()
//
// The purpose of this function is to ensure that vdl initialization occurs in
// the right order, and very early in the init sequence.  In particular, vdl
// registration and package variable initialization needs to occur before
// functions like vdl.TypeOf will work properly.
//
// This function returns a dummy value, so that it can be used to initialize the
// first var in the file, to take advantage of Go's defined init order.
func __VDLInit() struct{} {
	if __VDLInitCalled {
		return struct{}{}
	}
	__VDLInitCalled = true

	// Register types.
	vdl.Register((*Description)(nil))
	vdl.Register((*PartInfo)(nil))

	return struct{}{}
}
