// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Package: vtrace

package vtrace

import (
	"fmt"
	"time"
	"v.io/v23/uniqueid"
	"v.io/v23/vdl"
	"v.io/v23/vdl/vdlconv"
	time_2 "v.io/v23/vdlroot/time"
)

var _ = __VDLInit() // Must be first; see __VDLInit comments for details.

//////////////////////////////////////////////////
// Type definitions

// An Annotation represents data that is relevant at a specific moment.
// They can be attached to spans to add useful debugging information.
type Annotation struct {
	// When the annotation was added.
	When time.Time
	// The annotation message.
	// TODO(mattr): Allow richer annotations.
	Message string
}

func (Annotation) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/vtrace.Annotation"`
}) {
}

func (m *Annotation) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}
	var wireValue2 time_2.Time
	if err := time_2.TimeFromNative(&wireValue2, m.When); err != nil {
		return err
	}

	var5 := (wireValue2 == time_2.Time{})
	if var5 {
		if err := fieldsTarget1.ZeroField("When"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget3, fieldTarget4, err := fieldsTarget1.StartField("When")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}

			if err := wireValue2.FillVDLTarget(fieldTarget4, tt.NonOptional().Field(0).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget3, fieldTarget4); err != nil {
				return err
			}
		}
	}
	var8 := (m.Message == "")
	if var8 {
		if err := fieldsTarget1.ZeroField("Message"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget6, fieldTarget7, err := fieldsTarget1.StartField("Message")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}
			if err := fieldTarget7.FromString(string(m.Message), tt.NonOptional().Field(1).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget6, fieldTarget7); err != nil {
				return err
			}
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *Annotation) MakeVDLTarget() vdl.Target {
	return &AnnotationTarget{Value: m}
}

type AnnotationTarget struct {
	Value         *Annotation
	whenTarget    time_2.TimeTarget
	messageTarget vdl.StringTarget
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *AnnotationTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {

	if ttWant := vdl.TypeOf((*Annotation)(nil)).Elem(); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	return t, nil
}
func (t *AnnotationTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "When":
		t.whenTarget.Value = &t.Value.When
		target, err := &t.whenTarget, error(nil)
		return nil, target, err
	case "Message":
		t.messageTarget.Value = &t.Value.Message
		target, err := &t.messageTarget, error(nil)
		return nil, target, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct v.io/v23/vtrace.Annotation", name)
	}
}
func (t *AnnotationTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *AnnotationTarget) ZeroField(name string) error {
	switch name {
	case "When":
		t.Value.When = func() time.Time {
			var native time.Time
			if err := vdl.Convert(&native, time_2.Time{}); err != nil {
				panic(err)
			}
			return native
		}()
		return nil
	case "Message":
		t.Value.Message = ""
		return nil
	default:
		return fmt.Errorf("field %s not in struct v.io/v23/vtrace.Annotation", name)
	}
}
func (t *AnnotationTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

func (x Annotation) VDLIsZero() (bool, error) {
	var wireWhen time_2.Time
	if err := time_2.TimeFromNative(&wireWhen, x.When); err != nil {
		return false, err
	}
	if wireWhen != (time_2.Time{}) {
		return false, nil
	}
	if x.Message != "" {
		return false, nil
	}
	return true, nil
}

func (x Annotation) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(vdl.TypeOf((*Annotation)(nil)).Elem()); err != nil {
		return err
	}
	var wireWhen time_2.Time
	if err := time_2.TimeFromNative(&wireWhen, x.When); err != nil {
		return err
	}
	if wireWhen != (time_2.Time{}) {
		if err := enc.NextField("When"); err != nil {
			return err
		}
		if err := wireWhen.VDLWrite(enc); err != nil {
			return err
		}
	}
	if x.Message != "" {
		if err := enc.NextField("Message"); err != nil {
			return err
		}
		if err := enc.StartValue(vdl.TypeOf((*string)(nil))); err != nil {
			return err
		}
		if err := enc.EncodeString(x.Message); err != nil {
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

func (x *Annotation) VDLRead(dec vdl.Decoder) error {
	*x = Annotation{}
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
		case "When":
			var wire time_2.Time
			if err := wire.VDLRead(dec); err != nil {
				return err
			}
			if err := time_2.TimeToNative(wire, &x.When); err != nil {
				return err
			}
		case "Message":
			if err := dec.StartValue(); err != nil {
				return err
			}
			var err error
			if x.Message, err = dec.DecodeString(); err != nil {
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

// A SpanRecord is the wire format for a Span.
type SpanRecord struct {
	Id     uniqueid.Id // The Id of the Span.
	Parent uniqueid.Id // The Id of this Span's parent.
	Name   string      // The Name of this span.
	Start  time.Time   // The start time of this span.
	End    time.Time   // The end time of this span.
	// A series of annotations.
	Annotations []Annotation
}

func (SpanRecord) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/vtrace.SpanRecord"`
}) {
}

func (m *SpanRecord) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}
	var4 := (m.Id == uniqueid.Id{})
	if var4 {
		if err := fieldsTarget1.ZeroField("Id"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Id")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}

			if err := m.Id.FillVDLTarget(fieldTarget3, tt.NonOptional().Field(0).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
				return err
			}
		}
	}
	var7 := (m.Parent == uniqueid.Id{})
	if var7 {
		if err := fieldsTarget1.ZeroField("Parent"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget5, fieldTarget6, err := fieldsTarget1.StartField("Parent")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}

			if err := m.Parent.FillVDLTarget(fieldTarget6, tt.NonOptional().Field(1).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget5, fieldTarget6); err != nil {
				return err
			}
		}
	}
	var10 := (m.Name == "")
	if var10 {
		if err := fieldsTarget1.ZeroField("Name"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget8, fieldTarget9, err := fieldsTarget1.StartField("Name")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}
			if err := fieldTarget9.FromString(string(m.Name), tt.NonOptional().Field(2).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget8, fieldTarget9); err != nil {
				return err
			}
		}
	}
	var wireValue11 time_2.Time
	if err := time_2.TimeFromNative(&wireValue11, m.Start); err != nil {
		return err
	}

	var14 := (wireValue11 == time_2.Time{})
	if var14 {
		if err := fieldsTarget1.ZeroField("Start"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget12, fieldTarget13, err := fieldsTarget1.StartField("Start")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}

			if err := wireValue11.FillVDLTarget(fieldTarget13, tt.NonOptional().Field(3).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget12, fieldTarget13); err != nil {
				return err
			}
		}
	}
	var wireValue15 time_2.Time
	if err := time_2.TimeFromNative(&wireValue15, m.End); err != nil {
		return err
	}

	var18 := (wireValue15 == time_2.Time{})
	if var18 {
		if err := fieldsTarget1.ZeroField("End"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget16, fieldTarget17, err := fieldsTarget1.StartField("End")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}

			if err := wireValue15.FillVDLTarget(fieldTarget17, tt.NonOptional().Field(4).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget16, fieldTarget17); err != nil {
				return err
			}
		}
	}
	var var21 bool
	if len(m.Annotations) == 0 {
		var21 = true
	}
	if var21 {
		if err := fieldsTarget1.ZeroField("Annotations"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget19, fieldTarget20, err := fieldsTarget1.StartField("Annotations")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}

			listTarget22, err := fieldTarget20.StartList(tt.NonOptional().Field(5).Type, len(m.Annotations))
			if err != nil {
				return err
			}
			for i, elem24 := range m.Annotations {
				elemTarget23, err := listTarget22.StartElem(i)
				if err != nil {
					return err
				}

				if err := elem24.FillVDLTarget(elemTarget23, tt.NonOptional().Field(5).Type.Elem()); err != nil {
					return err
				}
				if err := listTarget22.FinishElem(elemTarget23); err != nil {
					return err
				}
			}
			if err := fieldTarget20.FinishList(listTarget22); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget19, fieldTarget20); err != nil {
				return err
			}
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *SpanRecord) MakeVDLTarget() vdl.Target {
	return &SpanRecordTarget{Value: m}
}

type SpanRecordTarget struct {
	Value             *SpanRecord
	idTarget          uniqueid.IdTarget
	parentTarget      uniqueid.IdTarget
	nameTarget        vdl.StringTarget
	startTarget       time_2.TimeTarget
	endTarget         time_2.TimeTarget
	annotationsTarget __VDLTarget1_list
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *SpanRecordTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {

	if ttWant := vdl.TypeOf((*SpanRecord)(nil)).Elem(); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	return t, nil
}
func (t *SpanRecordTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "Id":
		t.idTarget.Value = &t.Value.Id
		target, err := &t.idTarget, error(nil)
		return nil, target, err
	case "Parent":
		t.parentTarget.Value = &t.Value.Parent
		target, err := &t.parentTarget, error(nil)
		return nil, target, err
	case "Name":
		t.nameTarget.Value = &t.Value.Name
		target, err := &t.nameTarget, error(nil)
		return nil, target, err
	case "Start":
		t.startTarget.Value = &t.Value.Start
		target, err := &t.startTarget, error(nil)
		return nil, target, err
	case "End":
		t.endTarget.Value = &t.Value.End
		target, err := &t.endTarget, error(nil)
		return nil, target, err
	case "Annotations":
		t.annotationsTarget.Value = &t.Value.Annotations
		target, err := &t.annotationsTarget, error(nil)
		return nil, target, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct v.io/v23/vtrace.SpanRecord", name)
	}
}
func (t *SpanRecordTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *SpanRecordTarget) ZeroField(name string) error {
	switch name {
	case "Id":
		t.Value.Id = uniqueid.Id{}
		return nil
	case "Parent":
		t.Value.Parent = uniqueid.Id{}
		return nil
	case "Name":
		t.Value.Name = ""
		return nil
	case "Start":
		t.Value.Start = func() time.Time {
			var native time.Time
			if err := vdl.Convert(&native, time_2.Time{}); err != nil {
				panic(err)
			}
			return native
		}()
		return nil
	case "End":
		t.Value.End = func() time.Time {
			var native time.Time
			if err := vdl.Convert(&native, time_2.Time{}); err != nil {
				panic(err)
			}
			return native
		}()
		return nil
	case "Annotations":
		t.Value.Annotations = []Annotation(nil)
		return nil
	default:
		return fmt.Errorf("field %s not in struct v.io/v23/vtrace.SpanRecord", name)
	}
}
func (t *SpanRecordTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

// []Annotation
type __VDLTarget1_list struct {
	Value      *[]Annotation
	elemTarget AnnotationTarget
	vdl.TargetBase
	vdl.ListTargetBase
}

func (t *__VDLTarget1_list) StartList(tt *vdl.Type, len int) (vdl.ListTarget, error) {

	if ttWant := vdl.TypeOf((*[]Annotation)(nil)); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	if cap(*t.Value) < len {
		*t.Value = make([]Annotation, len)
	} else {
		*t.Value = (*t.Value)[:len]
	}
	return t, nil
}
func (t *__VDLTarget1_list) StartElem(index int) (elem vdl.Target, _ error) {
	t.elemTarget.Value = &(*t.Value)[index]
	target, err := &t.elemTarget, error(nil)
	return target, err
}
func (t *__VDLTarget1_list) FinishElem(elem vdl.Target) error {
	return nil
}
func (t *__VDLTarget1_list) FinishList(elem vdl.ListTarget) error {

	return nil
}

func (x SpanRecord) VDLIsZero() (bool, error) {
	if x.Id != (uniqueid.Id{}) {
		return false, nil
	}
	if x.Parent != (uniqueid.Id{}) {
		return false, nil
	}
	if x.Name != "" {
		return false, nil
	}
	var wireStart time_2.Time
	if err := time_2.TimeFromNative(&wireStart, x.Start); err != nil {
		return false, err
	}
	if wireStart != (time_2.Time{}) {
		return false, nil
	}
	var wireEnd time_2.Time
	if err := time_2.TimeFromNative(&wireEnd, x.End); err != nil {
		return false, err
	}
	if wireEnd != (time_2.Time{}) {
		return false, nil
	}
	if len(x.Annotations) != 0 {
		return false, nil
	}
	return true, nil
}

func (x SpanRecord) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(vdl.TypeOf((*SpanRecord)(nil)).Elem()); err != nil {
		return err
	}
	if x.Id != (uniqueid.Id{}) {
		if err := enc.NextField("Id"); err != nil {
			return err
		}
		if err := x.Id.VDLWrite(enc); err != nil {
			return err
		}
	}
	if x.Parent != (uniqueid.Id{}) {
		if err := enc.NextField("Parent"); err != nil {
			return err
		}
		if err := x.Parent.VDLWrite(enc); err != nil {
			return err
		}
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
	var wireStart time_2.Time
	if err := time_2.TimeFromNative(&wireStart, x.Start); err != nil {
		return err
	}
	if wireStart != (time_2.Time{}) {
		if err := enc.NextField("Start"); err != nil {
			return err
		}
		if err := wireStart.VDLWrite(enc); err != nil {
			return err
		}
	}
	var wireEnd time_2.Time
	if err := time_2.TimeFromNative(&wireEnd, x.End); err != nil {
		return err
	}
	if wireEnd != (time_2.Time{}) {
		if err := enc.NextField("End"); err != nil {
			return err
		}
		if err := wireEnd.VDLWrite(enc); err != nil {
			return err
		}
	}
	if len(x.Annotations) != 0 {
		if err := enc.NextField("Annotations"); err != nil {
			return err
		}
		if err := __VDLWriteAnon_list_1(enc, x.Annotations); err != nil {
			return err
		}
	}
	if err := enc.NextField(""); err != nil {
		return err
	}
	return enc.FinishValue()
}

func __VDLWriteAnon_list_1(enc vdl.Encoder, x []Annotation) error {
	if err := enc.StartValue(vdl.TypeOf((*[]Annotation)(nil))); err != nil {
		return err
	}
	if err := enc.SetLenHint(len(x)); err != nil {
		return err
	}
	for i := 0; i < len(x); i++ {
		if err := enc.NextEntry(false); err != nil {
			return err
		}
		if err := x[i].VDLWrite(enc); err != nil {
			return err
		}
	}
	if err := enc.NextEntry(true); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *SpanRecord) VDLRead(dec vdl.Decoder) error {
	*x = SpanRecord{}
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
		case "Id":
			if err := x.Id.VDLRead(dec); err != nil {
				return err
			}
		case "Parent":
			if err := x.Parent.VDLRead(dec); err != nil {
				return err
			}
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
		case "Start":
			var wire time_2.Time
			if err := wire.VDLRead(dec); err != nil {
				return err
			}
			if err := time_2.TimeToNative(wire, &x.Start); err != nil {
				return err
			}
		case "End":
			var wire time_2.Time
			if err := wire.VDLRead(dec); err != nil {
				return err
			}
			if err := time_2.TimeToNative(wire, &x.End); err != nil {
				return err
			}
		case "Annotations":
			if err := __VDLReadAnon_list_1(dec, &x.Annotations); err != nil {
				return err
			}
		default:
			if err := dec.SkipValue(); err != nil {
				return err
			}
		}
	}
}

func __VDLReadAnon_list_1(dec vdl.Decoder, x *[]Annotation) error {
	if err := dec.StartValue(); err != nil {
		return err
	}
	if (dec.StackDepth() == 1 || dec.IsAny()) && !vdl.Compatible(vdl.TypeOf(*x), dec.Type()) {
		return fmt.Errorf("incompatible list %T, from %v", *x, dec.Type())
	}
	switch len := dec.LenHint(); {
	case len > 0:
		*x = make([]Annotation, 0, len)
	default:
		*x = nil
	}
	for {
		switch done, err := dec.NextEntry(); {
		case err != nil:
			return err
		case done:
			return dec.FinishValue()
		}
		var elem Annotation
		if err := elem.VDLRead(dec); err != nil {
			return err
		}
		*x = append(*x, elem)
	}
}

type TraceRecord struct {
	Id    uniqueid.Id
	Spans []SpanRecord
}

func (TraceRecord) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/vtrace.TraceRecord"`
}) {
}

func (m *TraceRecord) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}
	var4 := (m.Id == uniqueid.Id{})
	if var4 {
		if err := fieldsTarget1.ZeroField("Id"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Id")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}

			if err := m.Id.FillVDLTarget(fieldTarget3, tt.NonOptional().Field(0).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
				return err
			}
		}
	}
	var var7 bool
	if len(m.Spans) == 0 {
		var7 = true
	}
	if var7 {
		if err := fieldsTarget1.ZeroField("Spans"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget5, fieldTarget6, err := fieldsTarget1.StartField("Spans")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}

			listTarget8, err := fieldTarget6.StartList(tt.NonOptional().Field(1).Type, len(m.Spans))
			if err != nil {
				return err
			}
			for i, elem10 := range m.Spans {
				elemTarget9, err := listTarget8.StartElem(i)
				if err != nil {
					return err
				}

				if err := elem10.FillVDLTarget(elemTarget9, tt.NonOptional().Field(1).Type.Elem()); err != nil {
					return err
				}
				if err := listTarget8.FinishElem(elemTarget9); err != nil {
					return err
				}
			}
			if err := fieldTarget6.FinishList(listTarget8); err != nil {
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

func (m *TraceRecord) MakeVDLTarget() vdl.Target {
	return &TraceRecordTarget{Value: m}
}

type TraceRecordTarget struct {
	Value       *TraceRecord
	idTarget    uniqueid.IdTarget
	spansTarget __VDLTarget2_list
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *TraceRecordTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {

	if ttWant := vdl.TypeOf((*TraceRecord)(nil)).Elem(); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	return t, nil
}
func (t *TraceRecordTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "Id":
		t.idTarget.Value = &t.Value.Id
		target, err := &t.idTarget, error(nil)
		return nil, target, err
	case "Spans":
		t.spansTarget.Value = &t.Value.Spans
		target, err := &t.spansTarget, error(nil)
		return nil, target, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct v.io/v23/vtrace.TraceRecord", name)
	}
}
func (t *TraceRecordTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *TraceRecordTarget) ZeroField(name string) error {
	switch name {
	case "Id":
		t.Value.Id = uniqueid.Id{}
		return nil
	case "Spans":
		t.Value.Spans = []SpanRecord(nil)
		return nil
	default:
		return fmt.Errorf("field %s not in struct v.io/v23/vtrace.TraceRecord", name)
	}
}
func (t *TraceRecordTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

// []SpanRecord
type __VDLTarget2_list struct {
	Value      *[]SpanRecord
	elemTarget SpanRecordTarget
	vdl.TargetBase
	vdl.ListTargetBase
}

func (t *__VDLTarget2_list) StartList(tt *vdl.Type, len int) (vdl.ListTarget, error) {

	if ttWant := vdl.TypeOf((*[]SpanRecord)(nil)); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	if cap(*t.Value) < len {
		*t.Value = make([]SpanRecord, len)
	} else {
		*t.Value = (*t.Value)[:len]
	}
	return t, nil
}
func (t *__VDLTarget2_list) StartElem(index int) (elem vdl.Target, _ error) {
	t.elemTarget.Value = &(*t.Value)[index]
	target, err := &t.elemTarget, error(nil)
	return target, err
}
func (t *__VDLTarget2_list) FinishElem(elem vdl.Target) error {
	return nil
}
func (t *__VDLTarget2_list) FinishList(elem vdl.ListTarget) error {

	return nil
}

func (x TraceRecord) VDLIsZero() (bool, error) {
	if x.Id != (uniqueid.Id{}) {
		return false, nil
	}
	if len(x.Spans) != 0 {
		return false, nil
	}
	return true, nil
}

func (x TraceRecord) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(vdl.TypeOf((*TraceRecord)(nil)).Elem()); err != nil {
		return err
	}
	if x.Id != (uniqueid.Id{}) {
		if err := enc.NextField("Id"); err != nil {
			return err
		}
		if err := x.Id.VDLWrite(enc); err != nil {
			return err
		}
	}
	if len(x.Spans) != 0 {
		if err := enc.NextField("Spans"); err != nil {
			return err
		}
		if err := __VDLWriteAnon_list_2(enc, x.Spans); err != nil {
			return err
		}
	}
	if err := enc.NextField(""); err != nil {
		return err
	}
	return enc.FinishValue()
}

func __VDLWriteAnon_list_2(enc vdl.Encoder, x []SpanRecord) error {
	if err := enc.StartValue(vdl.TypeOf((*[]SpanRecord)(nil))); err != nil {
		return err
	}
	if err := enc.SetLenHint(len(x)); err != nil {
		return err
	}
	for i := 0; i < len(x); i++ {
		if err := enc.NextEntry(false); err != nil {
			return err
		}
		if err := x[i].VDLWrite(enc); err != nil {
			return err
		}
	}
	if err := enc.NextEntry(true); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *TraceRecord) VDLRead(dec vdl.Decoder) error {
	*x = TraceRecord{}
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
		case "Id":
			if err := x.Id.VDLRead(dec); err != nil {
				return err
			}
		case "Spans":
			if err := __VDLReadAnon_list_2(dec, &x.Spans); err != nil {
				return err
			}
		default:
			if err := dec.SkipValue(); err != nil {
				return err
			}
		}
	}
}

func __VDLReadAnon_list_2(dec vdl.Decoder, x *[]SpanRecord) error {
	if err := dec.StartValue(); err != nil {
		return err
	}
	if (dec.StackDepth() == 1 || dec.IsAny()) && !vdl.Compatible(vdl.TypeOf(*x), dec.Type()) {
		return fmt.Errorf("incompatible list %T, from %v", *x, dec.Type())
	}
	switch len := dec.LenHint(); {
	case len > 0:
		*x = make([]SpanRecord, 0, len)
	default:
		*x = nil
	}
	for {
		switch done, err := dec.NextEntry(); {
		case err != nil:
			return err
		case done:
			return dec.FinishValue()
		}
		var elem SpanRecord
		if err := elem.VDLRead(dec); err != nil {
			return err
		}
		*x = append(*x, elem)
	}
}

type TraceFlags int32

func (TraceFlags) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/vtrace.TraceFlags"`
}) {
}

func (m *TraceFlags) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	if err := t.FromInt(int64((*m)), tt); err != nil {
		return err
	}
	return nil
}

func (m *TraceFlags) MakeVDLTarget() vdl.Target {
	return &TraceFlagsTarget{Value: m}
}

type TraceFlagsTarget struct {
	Value *TraceFlags
	vdl.TargetBase
}

func (t *TraceFlagsTarget) FromUint(src uint64, tt *vdl.Type) error {

	val, err := vdlconv.Uint64ToInt32(src)
	if err != nil {
		return err
	}
	*t.Value = TraceFlags(val)

	return nil
}
func (t *TraceFlagsTarget) FromInt(src int64, tt *vdl.Type) error {

	val, err := vdlconv.Int64ToInt32(src)
	if err != nil {
		return err
	}
	*t.Value = TraceFlags(val)

	return nil
}
func (t *TraceFlagsTarget) FromFloat(src float64, tt *vdl.Type) error {

	val, err := vdlconv.Float64ToInt32(src)
	if err != nil {
		return err
	}
	*t.Value = TraceFlags(val)

	return nil
}

func (x TraceFlags) VDLIsZero() (bool, error) {
	return x == 0, nil
}

func (x TraceFlags) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(vdl.TypeOf((*TraceFlags)(nil))); err != nil {
		return err
	}
	if err := enc.EncodeInt(int64(x)); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *TraceFlags) VDLRead(dec vdl.Decoder) error {
	if err := dec.StartValue(); err != nil {
		return err
	}
	tmp, err := dec.DecodeInt(32)
	if err != nil {
		return err
	}
	*x = TraceFlags(tmp)
	return dec.FinishValue()
}

// Request is the object that carries trace informtion between processes.
type Request struct {
	SpanId   uniqueid.Id // The Id of the span that originated the RPC call.
	TraceId  uniqueid.Id // The Id of the trace this call is a part of.
	Flags    TraceFlags
	LogLevel int32
}

func (Request) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/vtrace.Request"`
}) {
}

func (m *Request) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}
	var4 := (m.SpanId == uniqueid.Id{})
	if var4 {
		if err := fieldsTarget1.ZeroField("SpanId"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("SpanId")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}

			if err := m.SpanId.FillVDLTarget(fieldTarget3, tt.NonOptional().Field(0).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
				return err
			}
		}
	}
	var7 := (m.TraceId == uniqueid.Id{})
	if var7 {
		if err := fieldsTarget1.ZeroField("TraceId"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget5, fieldTarget6, err := fieldsTarget1.StartField("TraceId")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}

			if err := m.TraceId.FillVDLTarget(fieldTarget6, tt.NonOptional().Field(1).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget5, fieldTarget6); err != nil {
				return err
			}
		}
	}
	var10 := (m.Flags == TraceFlags(0))
	if var10 {
		if err := fieldsTarget1.ZeroField("Flags"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget8, fieldTarget9, err := fieldsTarget1.StartField("Flags")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}

			if err := m.Flags.FillVDLTarget(fieldTarget9, tt.NonOptional().Field(2).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget8, fieldTarget9); err != nil {
				return err
			}
		}
	}
	var13 := (m.LogLevel == int32(0))
	if var13 {
		if err := fieldsTarget1.ZeroField("LogLevel"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget11, fieldTarget12, err := fieldsTarget1.StartField("LogLevel")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}
			if err := fieldTarget12.FromInt(int64(m.LogLevel), tt.NonOptional().Field(3).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget11, fieldTarget12); err != nil {
				return err
			}
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *Request) MakeVDLTarget() vdl.Target {
	return &RequestTarget{Value: m}
}

type RequestTarget struct {
	Value          *Request
	spanIdTarget   uniqueid.IdTarget
	traceIdTarget  uniqueid.IdTarget
	flagsTarget    TraceFlagsTarget
	logLevelTarget vdl.Int32Target
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *RequestTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {

	if ttWant := vdl.TypeOf((*Request)(nil)).Elem(); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	return t, nil
}
func (t *RequestTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "SpanId":
		t.spanIdTarget.Value = &t.Value.SpanId
		target, err := &t.spanIdTarget, error(nil)
		return nil, target, err
	case "TraceId":
		t.traceIdTarget.Value = &t.Value.TraceId
		target, err := &t.traceIdTarget, error(nil)
		return nil, target, err
	case "Flags":
		t.flagsTarget.Value = &t.Value.Flags
		target, err := &t.flagsTarget, error(nil)
		return nil, target, err
	case "LogLevel":
		t.logLevelTarget.Value = &t.Value.LogLevel
		target, err := &t.logLevelTarget, error(nil)
		return nil, target, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct v.io/v23/vtrace.Request", name)
	}
}
func (t *RequestTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *RequestTarget) ZeroField(name string) error {
	switch name {
	case "SpanId":
		t.Value.SpanId = uniqueid.Id{}
		return nil
	case "TraceId":
		t.Value.TraceId = uniqueid.Id{}
		return nil
	case "Flags":
		t.Value.Flags = TraceFlags(0)
		return nil
	case "LogLevel":
		t.Value.LogLevel = int32(0)
		return nil
	default:
		return fmt.Errorf("field %s not in struct v.io/v23/vtrace.Request", name)
	}
}
func (t *RequestTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

func (x Request) VDLIsZero() (bool, error) {
	return x == Request{}, nil
}

func (x Request) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(vdl.TypeOf((*Request)(nil)).Elem()); err != nil {
		return err
	}
	if x.SpanId != (uniqueid.Id{}) {
		if err := enc.NextField("SpanId"); err != nil {
			return err
		}
		if err := x.SpanId.VDLWrite(enc); err != nil {
			return err
		}
	}
	if x.TraceId != (uniqueid.Id{}) {
		if err := enc.NextField("TraceId"); err != nil {
			return err
		}
		if err := x.TraceId.VDLWrite(enc); err != nil {
			return err
		}
	}
	if x.Flags != 0 {
		if err := enc.NextField("Flags"); err != nil {
			return err
		}
		if err := x.Flags.VDLWrite(enc); err != nil {
			return err
		}
	}
	if x.LogLevel != 0 {
		if err := enc.NextField("LogLevel"); err != nil {
			return err
		}
		if err := enc.StartValue(vdl.TypeOf((*int32)(nil))); err != nil {
			return err
		}
		if err := enc.EncodeInt(int64(x.LogLevel)); err != nil {
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

func (x *Request) VDLRead(dec vdl.Decoder) error {
	*x = Request{}
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
		case "SpanId":
			if err := x.SpanId.VDLRead(dec); err != nil {
				return err
			}
		case "TraceId":
			if err := x.TraceId.VDLRead(dec); err != nil {
				return err
			}
		case "Flags":
			if err := x.Flags.VDLRead(dec); err != nil {
				return err
			}
		case "LogLevel":
			if err := dec.StartValue(); err != nil {
				return err
			}
			tmp, err := dec.DecodeInt(32)
			if err != nil {
				return err
			}
			x.LogLevel = int32(tmp)
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

type Response struct {
	// Flags give options for trace collection, the client should alter its
	// collection for this trace according to the flags sent back from the
	// server.
	Flags TraceFlags
	// Trace is collected trace data.  This may be empty.
	Trace TraceRecord
}

func (Response) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/vtrace.Response"`
}) {
}

func (m *Response) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}
	var4 := (m.Flags == TraceFlags(0))
	if var4 {
		if err := fieldsTarget1.ZeroField("Flags"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Flags")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}

			if err := m.Flags.FillVDLTarget(fieldTarget3, tt.NonOptional().Field(0).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
				return err
			}
		}
	}
	var7 := true
	var8 := (m.Trace.Id == uniqueid.Id{})
	var7 = var7 && var8
	var var9 bool
	if len(m.Trace.Spans) == 0 {
		var9 = true
	}
	var7 = var7 && var9
	if var7 {
		if err := fieldsTarget1.ZeroField("Trace"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget5, fieldTarget6, err := fieldsTarget1.StartField("Trace")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}

			if err := m.Trace.FillVDLTarget(fieldTarget6, tt.NonOptional().Field(1).Type); err != nil {
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

func (m *Response) MakeVDLTarget() vdl.Target {
	return &ResponseTarget{Value: m}
}

type ResponseTarget struct {
	Value       *Response
	flagsTarget TraceFlagsTarget
	traceTarget TraceRecordTarget
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *ResponseTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {

	if ttWant := vdl.TypeOf((*Response)(nil)).Elem(); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	return t, nil
}
func (t *ResponseTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "Flags":
		t.flagsTarget.Value = &t.Value.Flags
		target, err := &t.flagsTarget, error(nil)
		return nil, target, err
	case "Trace":
		t.traceTarget.Value = &t.Value.Trace
		target, err := &t.traceTarget, error(nil)
		return nil, target, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct v.io/v23/vtrace.Response", name)
	}
}
func (t *ResponseTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *ResponseTarget) ZeroField(name string) error {
	switch name {
	case "Flags":
		t.Value.Flags = TraceFlags(0)
		return nil
	case "Trace":
		t.Value.Trace = TraceRecord{}
		return nil
	default:
		return fmt.Errorf("field %s not in struct v.io/v23/vtrace.Response", name)
	}
}
func (t *ResponseTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

func (x Response) VDLIsZero() (bool, error) {
	if x.Flags != 0 {
		return false, nil
	}
	isZeroTrace, err := x.Trace.VDLIsZero()
	if err != nil {
		return false, err
	}
	if !isZeroTrace {
		return false, nil
	}
	return true, nil
}

func (x Response) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(vdl.TypeOf((*Response)(nil)).Elem()); err != nil {
		return err
	}
	if x.Flags != 0 {
		if err := enc.NextField("Flags"); err != nil {
			return err
		}
		if err := x.Flags.VDLWrite(enc); err != nil {
			return err
		}
	}
	isZeroTrace, err := x.Trace.VDLIsZero()
	if err != nil {
		return err
	}
	if !isZeroTrace {
		if err := enc.NextField("Trace"); err != nil {
			return err
		}
		if err := x.Trace.VDLWrite(enc); err != nil {
			return err
		}
	}
	if err := enc.NextField(""); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *Response) VDLRead(dec vdl.Decoder) error {
	*x = Response{}
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
		case "Flags":
			if err := x.Flags.VDLRead(dec); err != nil {
				return err
			}
		case "Trace":
			if err := x.Trace.VDLRead(dec); err != nil {
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

const Empty = TraceFlags(0)
const CollectInMemory = TraceFlags(1)

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
	vdl.Register((*Annotation)(nil))
	vdl.Register((*SpanRecord)(nil))
	vdl.Register((*TraceRecord)(nil))
	vdl.Register((*TraceFlags)(nil))
	vdl.Register((*Request)(nil))
	vdl.Register((*Response)(nil))

	return struct{}{}
}
