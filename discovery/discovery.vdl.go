// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Package: discovery

package discovery

import (
	"fmt"
	"v.io/v23/vdl"
)

// Advertisement represents a feed into advertiser to broadcast its contents
// to scanners.
//
// A large advertisement may require additional RPC calls causing delay in
// discovery. We limit the maximum size of an advertisement to 512 bytes
// excluding id and attachments.
type Advertisement struct {
	// Universal unique identifier of the advertisement.
	// If this is not specified, a random unique identifier will be assigned.
	Id AdId
	// Interface name that the advertised service implements.
	// E.g., 'v.io/v23/services/vtrace.Store'.
	InterfaceName string
	// Addresses (vanadium object names) that the advertised service is served on.
	// E.g., '/host:port/a/b/c', '/ns.dev.v.io:8101/blah/blah'.
	Addresses []string
	// Attributes as a key/value pair.
	// E.g., {'resolution': '1024x768'}.
	//
	// The key must be US-ASCII printable characters, excluding the '=' character
	// and should not start with '_' character.
	Attributes Attributes
	// Attachments as a key/value pair.
	// E.g., {'thumbnail': binary_data }.
	//
	// Unlike attributes, attachments are for binary data and they are not queryable.
	// We limit the maximum size of a single attachment to 4K bytes.
	//
	// The key must be US-ASCII printable characters, excluding the '=' character
	// and should not start with '_' character.
	Attachments Attachments
}

func (Advertisement) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/discovery.Advertisement"`
}) {
}

func (m *Advertisement) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	if __VDLType_v_io_v23_discovery_Advertisement == nil || __VDLType0 == nil {
		panic("Initialization order error: types generated for FillVDLTarget not initialized. Consider moving caller to an init() block.")
	}
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}

	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Id")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		if err := m.Id.FillVDLTarget(fieldTarget3, __VDLType_v_io_v23_discovery_AdId); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
			return err
		}
	}
	keyTarget4, fieldTarget5, err := fieldsTarget1.StartField("InterfaceName")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget5.FromString(string(m.InterfaceName), vdl.StringType); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget4, fieldTarget5); err != nil {
			return err
		}
	}
	keyTarget6, fieldTarget7, err := fieldsTarget1.StartField("Addresses")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		listTarget8, err := fieldTarget7.StartList(__VDLType1, len(m.Addresses))
		if err != nil {
			return err
		}
		for i, elem10 := range m.Addresses {
			elemTarget9, err := listTarget8.StartElem(i)
			if err != nil {
				return err
			}
			if err := elemTarget9.FromString(string(elem10), vdl.StringType); err != nil {
				return err
			}
			if err := listTarget8.FinishElem(elemTarget9); err != nil {
				return err
			}
		}
		if err := fieldTarget7.FinishList(listTarget8); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget6, fieldTarget7); err != nil {
			return err
		}
	}
	keyTarget11, fieldTarget12, err := fieldsTarget1.StartField("Attributes")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		if err := m.Attributes.FillVDLTarget(fieldTarget12, __VDLType_v_io_v23_discovery_Attributes); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget11, fieldTarget12); err != nil {
			return err
		}
	}
	keyTarget13, fieldTarget14, err := fieldsTarget1.StartField("Attachments")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		if err := m.Attachments.FillVDLTarget(fieldTarget14, __VDLType_v_io_v23_discovery_Attachments); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget13, fieldTarget14); err != nil {
			return err
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *Advertisement) MakeVDLTarget() vdl.Target {
	return &AdvertisementTarget{Value: m}
}

type AdvertisementTarget struct {
	Value               *Advertisement
	idTarget            AdIdTarget
	interfaceNameTarget vdl.StringTarget
	addressesTarget     vdl.StringSliceTarget
	attributesTarget    AttributesTarget
	attachmentsTarget   AttachmentsTarget
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *AdvertisementTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {
	if !vdl.Compatible(tt, __VDLType_v_io_v23_discovery_Advertisement) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, __VDLType_v_io_v23_discovery_Advertisement)
	}
	return t, nil
}
func (t *AdvertisementTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "Id":
		t.idTarget.Value = &t.Value.Id
		target, err := &t.idTarget, error(nil)
		return nil, target, err
	case "InterfaceName":
		t.interfaceNameTarget.Value = &t.Value.InterfaceName
		target, err := &t.interfaceNameTarget, error(nil)
		return nil, target, err
	case "Addresses":
		t.addressesTarget.Value = &t.Value.Addresses
		target, err := &t.addressesTarget, error(nil)
		return nil, target, err
	case "Attributes":
		t.attributesTarget.Value = &t.Value.Attributes
		target, err := &t.attributesTarget, error(nil)
		return nil, target, err
	case "Attachments":
		t.attachmentsTarget.Value = &t.Value.Attachments
		target, err := &t.attachmentsTarget, error(nil)
		return nil, target, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct %v", name, __VDLType_v_io_v23_discovery_Advertisement)
	}
}
func (t *AdvertisementTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *AdvertisementTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

type AdIdTarget struct {
	Value *AdId
	vdl.TargetBase
}

func (t *AdIdTarget) FromBytes(src []byte, tt *vdl.Type) error {
	if !vdl.Compatible(tt, __VDLType_v_io_v23_discovery_AdId) {
		return fmt.Errorf("type %v incompatible with %v", tt, __VDLType_v_io_v23_discovery_AdId)
	}
	copy((*t.Value)[:], src)

	return nil
}

type AttributesTarget struct {
	Value      *Attributes
	currKey    string
	currElem   string
	keyTarget  vdl.StringTarget
	elemTarget vdl.StringTarget
	vdl.TargetBase
	vdl.MapTargetBase
}

func (t *AttributesTarget) StartMap(tt *vdl.Type, len int) (vdl.MapTarget, error) {
	if !vdl.Compatible(tt, __VDLType_v_io_v23_discovery_Attributes) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, __VDLType_v_io_v23_discovery_Attributes)
	}
	*t.Value = make(Attributes)
	return t, nil
}
func (t *AttributesTarget) StartKey() (key vdl.Target, _ error) {
	t.currKey = ""
	t.keyTarget.Value = &t.currKey
	target, err := &t.keyTarget, error(nil)
	return target, err
}
func (t *AttributesTarget) FinishKeyStartField(key vdl.Target) (field vdl.Target, _ error) {
	t.currElem = ""
	t.elemTarget.Value = &t.currElem
	target, err := &t.elemTarget, error(nil)
	return target, err
}
func (t *AttributesTarget) FinishField(key, field vdl.Target) error {
	(*t.Value)[t.currKey] = t.currElem
	return nil
}
func (t *AttributesTarget) FinishMap(elem vdl.MapTarget) error {
	if len(*t.Value) == 0 {
		*t.Value = nil
	}

	return nil
}

type AttachmentsTarget struct {
	Value      *Attachments
	currKey    string
	currElem   []byte
	keyTarget  vdl.StringTarget
	elemTarget vdl.BytesTarget
	vdl.TargetBase
	vdl.MapTargetBase
}

func (t *AttachmentsTarget) StartMap(tt *vdl.Type, len int) (vdl.MapTarget, error) {
	if !vdl.Compatible(tt, __VDLType_v_io_v23_discovery_Attachments) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, __VDLType_v_io_v23_discovery_Attachments)
	}
	*t.Value = make(Attachments)
	return t, nil
}
func (t *AttachmentsTarget) StartKey() (key vdl.Target, _ error) {
	t.currKey = ""
	t.keyTarget.Value = &t.currKey
	target, err := &t.keyTarget, error(nil)
	return target, err
}
func (t *AttachmentsTarget) FinishKeyStartField(key vdl.Target) (field vdl.Target, _ error) {
	t.currElem = []byte(nil)
	t.elemTarget.Value = &t.currElem
	target, err := &t.elemTarget, error(nil)
	return target, err
}
func (t *AttachmentsTarget) FinishField(key, field vdl.Target) error {
	(*t.Value)[t.currKey] = t.currElem
	return nil
}
func (t *AttachmentsTarget) FinishMap(elem vdl.MapTarget) error {
	if len(*t.Value) == 0 {
		*t.Value = nil
	}

	return nil
}

// An AdId is a globally unique identifier of an advertisement.
type AdId [16]byte

func (AdId) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/discovery.AdId"`
}) {
}

func (m *AdId) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	if err := t.FromBytes([]byte((*m)[:]), __VDLType_v_io_v23_discovery_AdId); err != nil {
		return err
	}
	return nil
}

func (m *AdId) MakeVDLTarget() vdl.Target {
	return &AdIdTarget{Value: m}
}

// Attributes represents service attributes as a key/value pair.
type Attributes map[string]string

func (Attributes) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/discovery.Attributes"`
}) {
}

func (m *Attributes) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	mapTarget1, err := t.StartMap(__VDLType_v_io_v23_discovery_Attributes, len((*m)))
	if err != nil {
		return err
	}
	for key3, value5 := range *m {
		keyTarget2, err := mapTarget1.StartKey()
		if err != nil {
			return err
		}
		if err := keyTarget2.FromString(string(key3), vdl.StringType); err != nil {
			return err
		}
		valueTarget4, err := mapTarget1.FinishKeyStartField(keyTarget2)
		if err != nil {
			return err
		}
		if err := valueTarget4.FromString(string(value5), vdl.StringType); err != nil {
			return err
		}
		if err := mapTarget1.FinishField(keyTarget2, valueTarget4); err != nil {
			return err
		}
	}
	if err := t.FinishMap(mapTarget1); err != nil {
		return err
	}
	return nil
}

func (m *Attributes) MakeVDLTarget() vdl.Target {
	return &AttributesTarget{Value: m}
}

// Attachments represents service attachments as a key/value pair.
type Attachments map[string][]byte

func (Attachments) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/discovery.Attachments"`
}) {
}

func (m *Attachments) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	mapTarget1, err := t.StartMap(__VDLType_v_io_v23_discovery_Attachments, len((*m)))
	if err != nil {
		return err
	}
	for key3, value5 := range *m {
		keyTarget2, err := mapTarget1.StartKey()
		if err != nil {
			return err
		}
		if err := keyTarget2.FromString(string(key3), vdl.StringType); err != nil {
			return err
		}
		valueTarget4, err := mapTarget1.FinishKeyStartField(keyTarget2)
		if err != nil {
			return err
		}

		if err := valueTarget4.FromBytes([]byte(value5), __VDLType2); err != nil {
			return err
		}
		if err := mapTarget1.FinishField(keyTarget2, valueTarget4); err != nil {
			return err
		}
	}
	if err := t.FinishMap(mapTarget1); err != nil {
		return err
	}
	return nil
}

func (m *Attachments) MakeVDLTarget() vdl.Target {
	return &AttachmentsTarget{Value: m}
}

func init() {
	vdl.Register((*Advertisement)(nil))
	vdl.Register((*AdId)(nil))
	vdl.Register((*Attributes)(nil))
	vdl.Register((*Attachments)(nil))
}

var __VDLType0 *vdl.Type = vdl.TypeOf((*Advertisement)(nil))
var __VDLType2 *vdl.Type = vdl.TypeOf([]byte(nil))
var __VDLType1 *vdl.Type = vdl.TypeOf([]string(nil))
var __VDLType_v_io_v23_discovery_AdId *vdl.Type = vdl.TypeOf(AdId{})
var __VDLType_v_io_v23_discovery_Advertisement *vdl.Type = vdl.TypeOf(Advertisement{})
var __VDLType_v_io_v23_discovery_Attachments *vdl.Type = vdl.TypeOf(Attachments(nil))
var __VDLType_v_io_v23_discovery_Attributes *vdl.Type = vdl.TypeOf(Attributes(nil))

func __VDLEnsureNativeBuilt() {
}
