// Copyright 2016 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package vdl

import (
	"errors"
	"fmt"
	"reflect"
)

var (
	errReadMustReflect       = errors.New("vdl: read must be handled via reflection")
	errReadIntoNilValue      = errors.New("vdl: read into nil value")
	errReadReflectCantSet    = errors.New("vdl: read into unsettable reflect.Value")
	errReadAnyAlreadyStarted = errors.New("vdl: read into any after StartValue called")
	errReadAnyInterfaceOnly  = errors.New("vdl: read into any only supported for interfaces")
)

// Read uses dec to decode a value into v, calling VDLRead methods and fast
// compiled readers when available, and using reflection otherwise.  This is
// basically an all-purpose VDLRead implementation.
func Read(dec Decoder, v interface{}) error {
	if v == nil {
		return errReadIntoNilValue
	}
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Ptr && !rv.IsNil() {
		// Fastpath check for non-reflect support.  Unfortunately we must use
		// reflection to detect the case where v is a nil pointer, which returns an
		// error in ReadReflect.
		//
		// TODO(toddw): If reflection is too slow, add the nil pointer check to all
		// VDLRead methods, as well as other readNonReflect cases below.
		if err := readNonReflect(dec, false, v); err != errReadMustReflect {
			return err
		}
	}
	return ReadReflect(dec, rv)
}

func readNonReflect(dec Decoder, calledStart bool, v interface{}) error {
	switch x := v.(type) {
	case Reader:
		// Reader handles the case where x has a code-generated decoder, and
		// special-cases such as vdl.Value and vom.RawBytes.
		if calledStart {
			dec.IgnoreNextStartValue()
		}
		return x.VDLRead(dec)
	case **Type:
		// Special-case type decoding, since we must assign the hash-consed pointer
		// for correctness, rather than filling in a newly-created Type.
		if !calledStart {
			if err := dec.StartValue(); err != nil {
				return err
			}
		}
		var err error
		if *x, err = dec.DecodeTypeObject(); err != nil {
			return err
		}
		return dec.FinishValue()

		// Cases after this point are purely performance optimizations.
		// TODO(toddw): Handle other common cases.
	case *[]byte:
		if !calledStart {
			if err := dec.StartValue(); err != nil {
				return err
			}
		}
		if err := dec.DecodeBytes(-1, x); err != nil {
			return err
		}
		return dec.FinishValue()
	}
	return errReadMustReflect
}

// ReadReflect is like Read, but takes a reflect.Value argument.  Use Read if
// performance is important and you have an interface{} handy.
func ReadReflect(dec Decoder, rv reflect.Value) error {
	if !rv.IsValid() {
		return errReadIntoNilValue
	}
	if !rv.CanSet() && rv.Kind() == reflect.Ptr && !rv.IsNil() {
		// Dereference the pointer a single time to make rv settable.
		rv = rv.Elem()
	}
	if !rv.CanSet() {
		return errReadReflectCantSet
	}
	tt, err := TypeFromReflect(rv.Type())
	if err != nil {
		return err
	}
	return readReflect(dec, false, rv, tt)
}

// readReflect uses dec to decode a value into rv, which has VDL type tt.  On
// success we guarantee that StartValue / FinishValue has been called on dec.
// If calledStart is true, StartValue has already been called.
func readReflect(dec Decoder, calledStart bool, rv reflect.Value, tt *Type) error {
	// Handle native types first, since they need the ToNative conversion.
	if ni := nativeInfoFromNative(rv.Type()); ni != nil {
		rvWire := reflect.New(ni.WireType).Elem()
		if err := readNonNative(dec, calledStart, rvWire, tt); err != nil {
			return err
		}
		return ni.ToNative(rvWire, rv.Addr())
	}
	return readNonNative(dec, calledStart, rv, tt)
}

func readNonNative(dec Decoder, calledStart bool, rv reflect.Value, tt *Type) error {
	// Any is handled first, since any(nil) is handled differently from ?T(nil)
	// contained in an any value, so this factoring makes things simpler.
	if tt == AnyType {
		return readAny(dec, calledStart, rv)
	}
	// Now we can start the decoder value, if we haven't already.
	if !calledStart {
		if err := dec.StartValue(); err != nil {
			return err
		}
		if err := decoderCompatible(dec, tt); err != nil {
			return err
		}
	}
	// Nil decoded values are handled next, to special-case the pointer handling;
	// we don't create pointers all the way down to the actual value.
	if dec.IsNil() {
		return readFromNil(dec, rv, tt)
	}
	// Now we know that the decoded value isn't nil.  Walk pointers and check for
	// faster non-reflect support.
	rv = readWalkPointers(rv)
	if err := readNonReflect(dec, true, rv.Addr().Interface()); err != errReadMustReflect {
		return err
	}
	// Special-case the error interface, and fill it in with the native error
	// representation verror.E.  Nil errors are handled above in readFromNil.
	if rv.Type() == rtError {
		return readNonNilError(dec, rv)
	}
	// Handle the non-nil decoded value.
	if err := readNonNilValue(dec, rv, tt.NonOptional()); err != nil {
		return err
	}
	return dec.FinishValue()
}

// readWalkPointers repeatedly dereferences pointers, creating new values if the
// pointer is nil, and returns the final non-pointer reflect value.
func readWalkPointers(rv reflect.Value) reflect.Value {
	for rv.Kind() == reflect.Ptr {
		// Special-case to stop at *Type, which is filled in via readNonReflect.
		if rv.Type() == rtPtrToType {
			return rv
		}
		if rv.IsNil() {
			rv.Set(reflect.New(rv.Type().Elem()))
		}
		rv = rv.Elem()
	}
	return rv
}

func readAny(dec Decoder, calledStart bool, rv reflect.Value) error {
	if calledStart {
		// The existing code ensures that calledStart is always false here, since
		// readReflect(dec, true, ...) is only called below in this function, which
		// never calls it with another any type.  If we did, we'd have a vdl
		// any(any), which isn't allowed.  This error tries to prevent future
		// changes that will break this requirement.
		//
		// Also note that the implementation of vom.RawBytes.VDLRead requires that
		// StartValue has not been called yet.
		return errReadAnyAlreadyStarted
	}
	// Walk pointers and check for faster non-reflect support, which handles
	// vdl.Value and vom.RawBytes, and any other special-cases.
	rv = readWalkPointers(rv)
	if err := readNonReflect(dec, false, rv.Addr().Interface()); err != errReadMustReflect {
		return err
	}
	// The only case left is to handle interfaces.  We allow decoding into
	// interface{}, as well as any other interface.
	if rv.Kind() != reflect.Interface {
		return errReadAnyInterfaceOnly
	}
	if err := dec.StartValue(); err != nil {
		return err
	}
	// Handle decoding any(nil) by setting the interface to nil.  Note that the
	// only case where dec.Type() is AnyType is when the value is any(nil).
	if dec.Type() == AnyType {
		rv.Set(reflect.Zero(rv.Type()))
		return dec.FinishValue()
	}
	// Lookup the reflect type based on the decoder type, and create a new value
	// to decode into.
	//
	// TODO(toddw): Replace typeToReflectFixed with TypeToReflect, after we've
	// fixed it to treat the error type correctly.
	rtDecode := typeToReflectFixed(dec.Type())
	if rtDecode == nil {
		return fmt.Errorf("vdl: %v not registered, call vdl.Register, or use vdl.Value or vom.RawBytes instead", dec.Type())
	}
	// If we decoded an optional type, ensure that it is a pointer.  Note that if
	// we decoded a nil, dec.Type() is already optional, so rtDecode will already
	// be a pointer.
	if dec.IsOptional() && !dec.IsNil() {
		rtDecode = reflect.PtrTo(rtDecode)
	}
	if !rtDecode.Implements(rv.Type()) {
		return fmt.Errorf("vdl: %v doesn't implement %v", rtDecode, rv.Type())
	}
	// Handle decoding optional(nil), by setting rv to a nil pointer of the
	// concrete type.  We know that rtDecode must be a pointer, since dec.Type()
	// is optional.
	if dec.Type().Kind() == Optional {
		rv.Set(reflect.Zero(rtDecode))
		return dec.FinishValue()
	}
	// Handle non-nil values by decoding into rvDecode, and setting rv.
	rvDecode := reflect.New(rtDecode).Elem()
	if err := readReflect(dec, true, rvDecode, dec.Type()); err != nil {
		return err
	}
	rv.Set(rvDecode)
	// Note that dec.FinishValue has already been called by readReflect.
	return nil
}

func readFromNil(dec Decoder, rv reflect.Value, tt *Type) error {
	if tt.Kind() != Optional {
		return fmt.Errorf("vdl: can't decode nil into non-optional %v", tt)
	}
	// Note that since tt is optional, we know that rv is always a pointer, or the
	// special-case error interface.
	rv.Set(reflect.Zero(rv.Type()))
	return dec.FinishValue()
}

func readNonNilError(dec Decoder, rv reflect.Value) error {
	// This function implements the equivalent of verror.VDLRead for the non-nil
	// case.  We can't call the verror function directly, since that would create
	// a dependency cycle between the vdl and verror packages.
	dec.IgnoreNextStartValue()
	var wire WireError
	if err := wire.VDLRead(dec); err != nil {
		return err
	}
	ni, err := nativeInfoForError()
	if err != nil {
		return err
	}
	rvNativePtr := reflect.New(ni.NativeType)
	if err := ni.ToNative(reflect.ValueOf(wire), rvNativePtr); err != nil {
		return err
	}
	rv.Set(rvNativePtr.Elem())
	// Note that dec.FinishValue has already been called by wire.VDLRead.
	return nil
}

func readNonNilValue(dec Decoder, rv reflect.Value, tt *Type) error {
	// Handle named and unnamed []byte and [N]byte, where the element type is the
	// unnamed byte type.  Cases like []MyByte fall through and are handled as
	// regular lists, since we can't easily convert []MyByte to []byte.
	switch {
	case tt.Kind() == Array && tt.Elem() == ByteType:
		bytes := rv.Slice(0, tt.Len()).Interface().([]byte)
		return dec.DecodeBytes(tt.Len(), &bytes)
	case tt.Kind() == List && tt.Elem() == ByteType:
		var bytes []byte
		if err := dec.DecodeBytes(-1, &bytes); err != nil {
			return err
		}
		rv.Set(reflect.ValueOf(bytes))
		return nil
	}
	// Handle regular non-nil values.
	switch kind := tt.Kind(); kind {
	case Bool:
		val, err := dec.DecodeBool()
		if err != nil {
			return err
		}
		rv.SetBool(val)
		return nil
	case String:
		val, err := dec.DecodeString()
		if err != nil {
			return err
		}
		rv.SetString(val)
		return nil
	case Enum:
		val, err := dec.DecodeString()
		if err != nil {
			return err
		}
		return rv.Addr().Interface().(settable).Set(val)
	case Byte, Uint16, Uint32, Uint64:
		val, err := dec.DecodeUint(kind.BitLen())
		if err != nil {
			return err
		}
		rv.SetUint(val)
		return nil
	case Int8, Int16, Int32, Int64:
		val, err := dec.DecodeInt(kind.BitLen())
		if err != nil {
			return err
		}
		rv.SetInt(val)
		return nil
	case Float32, Float64:
		val, err := dec.DecodeFloat(kind.BitLen())
		if err != nil {
			return err
		}
		rv.SetFloat(val)
		return nil
	case Array:
		return readArray(dec, rv, tt)
	case List:
		return readList(dec, rv, tt)
	case Set:
		return readSet(dec, rv, tt)
	case Map:
		return readMap(dec, rv, tt)
	case Struct:
		return readStruct(dec, rv, tt)
	case Union:
		return readUnion(dec, rv, tt)
	}
	// Note that Any was already handled via readAny, Optional was handled via
	// readFromNil (or stripped off for non-nil values), and TypeObject was
	// handled via the readNonReflect special-case.
	return fmt.Errorf("vdl: Read unhandled type %v %v", rv.Type(), tt)
}

func readArray(dec Decoder, rv reflect.Value, tt *Type) error {
	rt := rv.Type()
	index := 0
	for {
		switch done, err := dec.NextEntry(); {
		case err != nil:
			return err
		case done != (index >= rv.Len()):
			return fmt.Errorf("array len mismatch, got %d, want %v", index, rt)
		case done:
			return nil
		}
		if err := readReflect(dec, false, rv.Index(index), tt.Elem()); err != nil {
			return err
		}
		index++
	}
}

func readList(dec Decoder, rv reflect.Value, tt *Type) error {
	rt := rv.Type()
	switch len := dec.LenHint(); {
	case len > 0:
		rv.Set(reflect.MakeSlice(rt, 0, len))
	default:
		rv.Set(reflect.Zero(rt))
	}
	for {
		switch done, err := dec.NextEntry(); {
		case err != nil:
			return err
		case done:
			return nil
		}
		elem := reflect.New(rt.Elem()).Elem()
		if err := readReflect(dec, false, elem, tt.Elem()); err != nil {
			return err
		}
		rv.Set(reflect.Append(rv, elem))
	}
}

var rvEmptyStruct = reflect.ValueOf(struct{}{})

func readSet(dec Decoder, rv reflect.Value, tt *Type) error {
	rt := rv.Type()
	tmpSet, isNil := reflect.Zero(rt), true
	for {
		switch done, err := dec.NextEntry(); {
		case err != nil:
			return err
		case done:
			rv.Set(tmpSet)
			return nil
		}
		key := reflect.New(rt.Key()).Elem()
		if err := readReflect(dec, false, key, tt.Key()); err != nil {
			return err
		}
		if isNil {
			tmpSet, isNil = reflect.MakeMap(rt), false
		}
		tmpSet.SetMapIndex(key, rvEmptyStruct)
	}
}

func readMap(dec Decoder, rv reflect.Value, tt *Type) error {
	rt := rv.Type()
	tmpMap, isNil := reflect.Zero(rt), true
	for {
		switch done, err := dec.NextEntry(); {
		case err != nil:
			return err
		case done:
			rv.Set(tmpMap)
			return nil
		}
		key := reflect.New(rt.Key()).Elem()
		if err := readReflect(dec, false, key, tt.Key()); err != nil {
			return err
		}
		elem := reflect.New(rt.Elem()).Elem()
		if err := readReflect(dec, false, elem, tt.Elem()); err != nil {
			return err
		}
		if isNil {
			tmpMap, isNil = reflect.MakeMap(rt), false
		}
		tmpMap.SetMapIndex(key, elem)
	}
}

func readStruct(dec Decoder, rv reflect.Value, tt *Type) error {
	rt := rv.Type()
	// Reset to the zero struct, since fields may be missing.
	//
	// TODO(toddw): Avoid repeated zero-setting of nested structs.
	rvZero, err := rvZeroValue(rt, tt)
	if err != nil {
		return err
	}
	rv.Set(rvZero)
	for {
		name, err := dec.NextField()
		switch {
		case err != nil:
			return err
		case name == "":
			return nil
		}
		switch ttField, index := tt.FieldByName(name); {
		case index != -1:
			rvField := rv.FieldByName(name)
			if err := readReflect(dec, false, rvField, ttField.Type); err != nil {
				return err
			}
		default:
			if err := dec.SkipValue(); err != nil {
				return err
			}
		}
	}
}

func readUnion(dec Decoder, rv reflect.Value, tt *Type) error {
	rt := rv.Type()
	name, err := dec.NextField()
	switch {
	case err != nil:
		return err
	case name == "":
		return fmt.Errorf("missing field in union %v, from %v", rt, dec.Type())
	}
	ttField, index := tt.FieldByName(name)
	if index == -1 {
		return fmt.Errorf("field %q not in union %v, from %v", name, rt, dec.Type())
	}
	// We have a union interface.  Create a new field based on its rep type, fill
	// in its value, and assign the field to the interface.
	ri, _, err := deriveReflectInfo(rt)
	if err != nil {
		return err
	}
	rvField := reflect.New(ri.UnionFields[index].RepType).Elem()
	if err := readReflect(dec, false, rvField.Field(0), ttField.Type); err != nil {
		return err
	}
	rv.Set(rvField)
	switch name, err := dec.NextField(); {
	case err != nil:
		return err
	case name != "":
		return fmt.Errorf("extra field %q in union %v, from %v", name, rt, dec.Type())
	}
	return nil
}
