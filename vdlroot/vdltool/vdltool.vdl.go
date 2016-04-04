// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Package: vdltool

// Package vdltool defines types used by the vdl tool itself, including the
// format of vdl.config files.
package vdltool

import (
	"fmt"
	"reflect"
	"v.io/v23/vdl"
)

var _ = __VDLInit() // Must be first; see __VDLInit comments for details.

//////////////////////////////////////////////////
// Type definitions

// GenLanguage enumerates the known code generation languages.
type GenLanguage int

const (
	GenLanguageGo GenLanguage = iota
	GenLanguageJava
	GenLanguageJavascript
)

// GenLanguageAll holds all labels for GenLanguage.
var GenLanguageAll = [...]GenLanguage{GenLanguageGo, GenLanguageJava, GenLanguageJavascript}

// GenLanguageFromString creates a GenLanguage from a string label.
func GenLanguageFromString(label string) (x GenLanguage, err error) {
	err = x.Set(label)
	return
}

// Set assigns label to x.
func (x *GenLanguage) Set(label string) error {
	switch label {
	case "Go", "go":
		*x = GenLanguageGo
		return nil
	case "Java", "java":
		*x = GenLanguageJava
		return nil
	case "Javascript", "javascript":
		*x = GenLanguageJavascript
		return nil
	}
	*x = -1
	return fmt.Errorf("unknown label %q in vdltool.GenLanguage", label)
}

// String returns the string label of x.
func (x GenLanguage) String() string {
	switch x {
	case GenLanguageGo:
		return "Go"
	case GenLanguageJava:
		return "Java"
	case GenLanguageJavascript:
		return "Javascript"
	}
	return ""
}

func (GenLanguage) __VDLReflect(struct {
	Name string `vdl:"vdltool.GenLanguage"`
	Enum struct{ Go, Java, Javascript string }
}) {
}

func (m *GenLanguage) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	if err := t.FromEnumLabel((*m).String(), tt); err != nil {
		return err
	}
	return nil
}

func (m *GenLanguage) MakeVDLTarget() vdl.Target {
	return nil
}

type GenLanguageTarget struct {
	Value *GenLanguage
	vdl.TargetBase
}

func (t *GenLanguageTarget) FromEnumLabel(src string, tt *vdl.Type) error {

	if ttWant := vdl.TypeOf((*GenLanguage)(nil)); !vdl.Compatible(tt, ttWant) {
		return fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	switch src {
	case "Go":
		*t.Value = 0
	case "Java":
		*t.Value = 1
	case "Javascript":
		*t.Value = 2
	default:
		return fmt.Errorf("label %s not in enum GenLanguage", src)
	}

	return nil
}

// GoImport describes Go import information.
type GoImport struct {
	// Path is the package path that uniquely identifies the imported package.
	Path string
	// Name is the name of the package identified by Path.  Due to Go conventions,
	// it is typically just the basename of Path, but may be set to something
	// different if the imported package doesn't follow Go conventions.
	Name string
}

func (GoImport) __VDLReflect(struct {
	Name string `vdl:"vdltool.GoImport"`
}) {
}

func (m *GoImport) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}
	var4 := (m.Path == "")
	if var4 {
		if err := fieldsTarget1.ZeroField("Path"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Path")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}
			if err := fieldTarget3.FromString(string(m.Path), tt.NonOptional().Field(0).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
				return err
			}
		}
	}
	var7 := (m.Name == "")
	if var7 {
		if err := fieldsTarget1.ZeroField("Name"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget5, fieldTarget6, err := fieldsTarget1.StartField("Name")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}
			if err := fieldTarget6.FromString(string(m.Name), tt.NonOptional().Field(1).Type); err != nil {
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

func (m *GoImport) MakeVDLTarget() vdl.Target {
	return nil
}

type GoImportTarget struct {
	Value *GoImport

	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *GoImportTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {

	if ttWant := vdl.TypeOf((*GoImport)(nil)).Elem(); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	return t, nil
}
func (t *GoImportTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "Path":
		target, err := vdl.ReflectTarget(reflect.ValueOf(&t.Value.Path))
		return nil, target, err
	case "Name":
		target, err := vdl.ReflectTarget(reflect.ValueOf(&t.Value.Name))
		return nil, target, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct vdltool.GoImport", name)
	}
}
func (t *GoImportTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *GoImportTarget) ZeroField(name string) error {
	switch name {
	case "Path":
		t.Value.Path = ""
		return nil
	case "Name":
		t.Value.Name = ""
		return nil
	default:
		return fmt.Errorf("field %s not in struct vdltool.GoImport", name)
	}
}
func (t *GoImportTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

// GoType describes the Go type information associated with a VDL type.
// See v.io/x/ref/lib/vdl/testdata/native for examples.
type GoType struct {
	// Type is the Go type to use in generated code, instead of the VDL type.  If
	// the Go type requires additional imports, specify the type using the
	// standard local package name here, and also specify the import package in
	// Imports.
	Type string
	// Imports are the Go imports to use in generated code, required by the Type.
	Imports []GoImport
}

func (GoType) __VDLReflect(struct {
	Name string `vdl:"vdltool.GoType"`
}) {
}

func (m *GoType) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}
	var4 := (m.Type == "")
	if var4 {
		if err := fieldsTarget1.ZeroField("Type"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Type")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}
			if err := fieldTarget3.FromString(string(m.Type), tt.NonOptional().Field(0).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
				return err
			}
		}
	}
	var var7 bool
	if len(m.Imports) == 0 {
		var7 = true
	}
	if var7 {
		if err := fieldsTarget1.ZeroField("Imports"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget5, fieldTarget6, err := fieldsTarget1.StartField("Imports")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}

			listTarget8, err := fieldTarget6.StartList(tt.NonOptional().Field(1).Type, len(m.Imports))
			if err != nil {
				return err
			}
			for i, elem10 := range m.Imports {
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

func (m *GoType) MakeVDLTarget() vdl.Target {
	return nil
}

type GoTypeTarget struct {
	Value *GoType

	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *GoTypeTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {

	if ttWant := vdl.TypeOf((*GoType)(nil)).Elem(); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	return t, nil
}
func (t *GoTypeTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "Type":
		target, err := vdl.ReflectTarget(reflect.ValueOf(&t.Value.Type))
		return nil, target, err
	case "Imports":
		target, err := vdl.ReflectTarget(reflect.ValueOf(&t.Value.Imports))
		return nil, target, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct vdltool.GoType", name)
	}
}
func (t *GoTypeTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *GoTypeTarget) ZeroField(name string) error {
	switch name {
	case "Type":
		t.Value.Type = ""
		return nil
	case "Imports":
		t.Value.Imports = []GoImport(nil)
		return nil
	default:
		return fmt.Errorf("field %s not in struct vdltool.GoType", name)
	}
}
func (t *GoTypeTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

// GoConfig specifies go specific configuration.
type GoConfig struct {
	// WireToNativeTypes specifies the mapping from a VDL wire type to its Go
	// native type representation.  This is rarely used and easy to configure
	// incorrectly; usage is currently restricted to packages that are explicitly
	// whitelisted.
	//
	// WireToNativeTypes are meant for scenarios where there is an idiomatic Go
	// type used in your code, but you need a standard VDL representation for wire
	// compatibility.  E.g. the VDL time package defines Duration and Time for
	// wire compatibility, but we want the generated code to use the standard Go
	// time package.
	//
	// The key of the map is the name of the VDL type (aka WireType), which must
	// be defined in the vdl package associated with the vdl.config file.
	//
	// The code generator assumes the existence of a pair of conversion functions
	// converting between the wire and native types, and will automatically call
	// vdl.RegisterNative with these function names.
	//
	// Assuming the name of the WireType is Foo:
	//   func fooToNative(x Foo, n *Native) error
	//   func fooFromNative(x *Foo, n Native) error
	WireToNativeTypes map[string]GoType
}

func (GoConfig) __VDLReflect(struct {
	Name string `vdl:"vdltool.GoConfig"`
}) {
}

func (m *GoConfig) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}
	var var4 bool
	if len(m.WireToNativeTypes) == 0 {
		var4 = true
	}
	if var4 {
		if err := fieldsTarget1.ZeroField("WireToNativeTypes"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("WireToNativeTypes")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}

			mapTarget5, err := fieldTarget3.StartMap(tt.NonOptional().Field(0).Type, len(m.WireToNativeTypes))
			if err != nil {
				return err
			}
			for key7, value9 := range m.WireToNativeTypes {
				keyTarget6, err := mapTarget5.StartKey()
				if err != nil {
					return err
				}
				if err := keyTarget6.FromString(string(key7), tt.NonOptional().Field(0).Type.Key()); err != nil {
					return err
				}
				valueTarget8, err := mapTarget5.FinishKeyStartField(keyTarget6)
				if err != nil {
					return err
				}

				if err := value9.FillVDLTarget(valueTarget8, tt.NonOptional().Field(0).Type.Elem()); err != nil {
					return err
				}
				if err := mapTarget5.FinishField(keyTarget6, valueTarget8); err != nil {
					return err
				}
			}
			if err := fieldTarget3.FinishMap(mapTarget5); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
				return err
			}
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *GoConfig) MakeVDLTarget() vdl.Target {
	return nil
}

type GoConfigTarget struct {
	Value *GoConfig

	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *GoConfigTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {

	if ttWant := vdl.TypeOf((*GoConfig)(nil)).Elem(); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	return t, nil
}
func (t *GoConfigTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "WireToNativeTypes":
		target, err := vdl.ReflectTarget(reflect.ValueOf(&t.Value.WireToNativeTypes))
		return nil, target, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct vdltool.GoConfig", name)
	}
}
func (t *GoConfigTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *GoConfigTarget) ZeroField(name string) error {
	switch name {
	case "WireToNativeTypes":
		t.Value.WireToNativeTypes = map[string]GoType(nil)
		return nil
	default:
		return fmt.Errorf("field %s not in struct vdltool.GoConfig", name)
	}
}
func (t *GoConfigTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

// JavaConfig specifies java specific configuration.
type JavaConfig struct {
	// WireToNativeTypes specifies the mapping from a VDL wire type to its Java
	// native type representation.  This is rarely used and easy to configure
	// incorrectly; usage is currently restricted to packages that are explicitly
	// whitelisted.
	//
	// WireToNativeTypes are meant for scenarios where there is an idiomatic Java
	// type used in your code, but you need a standard VDL representation for wire
	// compatibility.  E.g. the VDL time package defines Duration and Time for
	// wire compatibility, but we want the generated code to use the org.joda.time
	// package.
	//
	// The key of the map is the name of the VDL type (aka WireType), which must
	// be defined in the vdl package associated with the vdl.config file.
	//
	// The code generator assumes that the conversion functions will be registered
	// in java vdl package.
	WireToNativeTypes map[string]string
	// WireTypeRenames specifies the mapping from a VDL wire type name to its
	// Java native type name.
	//
	// WireTypeRenames are meant for scenarios where the VDL wire name
	// conflicts in some way with the Java native names, e.g., a VDL Integer
	// type could be named VInteger for clarity.
	//
	// When combined with WireToNativeTypes, this feature allows us to attach
	// functions to VDL types.  For example, we may rename AccessList VDL type
	// into WireAccessList and then map WireAccessList to our Java native type
	// AccessList which defines functions on the VDL data.
	//
	// The key of the map is the name of the VDL wire type, which must be
	// defined in the vdl package associated with the vdl.config file.
	WireTypeRenames map[string]string
}

func (JavaConfig) __VDLReflect(struct {
	Name string `vdl:"vdltool.JavaConfig"`
}) {
}

func (m *JavaConfig) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}
	var var4 bool
	if len(m.WireToNativeTypes) == 0 {
		var4 = true
	}
	if var4 {
		if err := fieldsTarget1.ZeroField("WireToNativeTypes"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("WireToNativeTypes")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}

			mapTarget5, err := fieldTarget3.StartMap(tt.NonOptional().Field(0).Type, len(m.WireToNativeTypes))
			if err != nil {
				return err
			}
			for key7, value9 := range m.WireToNativeTypes {
				keyTarget6, err := mapTarget5.StartKey()
				if err != nil {
					return err
				}
				if err := keyTarget6.FromString(string(key7), tt.NonOptional().Field(0).Type.Key()); err != nil {
					return err
				}
				valueTarget8, err := mapTarget5.FinishKeyStartField(keyTarget6)
				if err != nil {
					return err
				}
				if err := valueTarget8.FromString(string(value9), tt.NonOptional().Field(0).Type.Elem()); err != nil {
					return err
				}
				if err := mapTarget5.FinishField(keyTarget6, valueTarget8); err != nil {
					return err
				}
			}
			if err := fieldTarget3.FinishMap(mapTarget5); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
				return err
			}
		}
	}
	var var12 bool
	if len(m.WireTypeRenames) == 0 {
		var12 = true
	}
	if var12 {
		if err := fieldsTarget1.ZeroField("WireTypeRenames"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget10, fieldTarget11, err := fieldsTarget1.StartField("WireTypeRenames")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}

			mapTarget13, err := fieldTarget11.StartMap(tt.NonOptional().Field(1).Type, len(m.WireTypeRenames))
			if err != nil {
				return err
			}
			for key15, value17 := range m.WireTypeRenames {
				keyTarget14, err := mapTarget13.StartKey()
				if err != nil {
					return err
				}
				if err := keyTarget14.FromString(string(key15), tt.NonOptional().Field(1).Type.Key()); err != nil {
					return err
				}
				valueTarget16, err := mapTarget13.FinishKeyStartField(keyTarget14)
				if err != nil {
					return err
				}
				if err := valueTarget16.FromString(string(value17), tt.NonOptional().Field(1).Type.Elem()); err != nil {
					return err
				}
				if err := mapTarget13.FinishField(keyTarget14, valueTarget16); err != nil {
					return err
				}
			}
			if err := fieldTarget11.FinishMap(mapTarget13); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget10, fieldTarget11); err != nil {
				return err
			}
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *JavaConfig) MakeVDLTarget() vdl.Target {
	return nil
}

type JavaConfigTarget struct {
	Value *JavaConfig

	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *JavaConfigTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {

	if ttWant := vdl.TypeOf((*JavaConfig)(nil)).Elem(); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	return t, nil
}
func (t *JavaConfigTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "WireToNativeTypes":
		target, err := vdl.ReflectTarget(reflect.ValueOf(&t.Value.WireToNativeTypes))
		return nil, target, err
	case "WireTypeRenames":
		target, err := vdl.ReflectTarget(reflect.ValueOf(&t.Value.WireTypeRenames))
		return nil, target, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct vdltool.JavaConfig", name)
	}
}
func (t *JavaConfigTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *JavaConfigTarget) ZeroField(name string) error {
	switch name {
	case "WireToNativeTypes":
		t.Value.WireToNativeTypes = map[string]string(nil)
		return nil
	case "WireTypeRenames":
		t.Value.WireTypeRenames = map[string]string(nil)
		return nil
	default:
		return fmt.Errorf("field %s not in struct vdltool.JavaConfig", name)
	}
}
func (t *JavaConfigTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

// JavascriptConfig specifies javascript specific configuration.
type JavascriptConfig struct {
}

func (JavascriptConfig) __VDLReflect(struct {
	Name string `vdl:"vdltool.JavascriptConfig"`
}) {
}

func (m *JavascriptConfig) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *JavascriptConfig) MakeVDLTarget() vdl.Target {
	return nil
}

type JavascriptConfigTarget struct {
	Value *JavascriptConfig
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *JavascriptConfigTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {

	if ttWant := vdl.TypeOf((*JavascriptConfig)(nil)).Elem(); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	return t, nil
}
func (t *JavascriptConfigTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	default:
		return nil, nil, fmt.Errorf("field %s not in struct vdltool.JavascriptConfig", name)
	}
}
func (t *JavascriptConfigTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *JavascriptConfigTarget) ZeroField(name string) error {
	switch name {
	default:
		return fmt.Errorf("field %s not in struct vdltool.JavascriptConfig", name)
	}
}
func (t *JavascriptConfigTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

// Config specifies the configuration for the vdl tool.  This is typically
// represented in optional "vdl.config" files in each vdl source package.  Each
// vdl.config file implicitly imports this package.  E.g. you may refer to
// vdltool.Config in the "vdl.config" file without explicitly importing vdltool.
type Config struct {
	// GenLanguages restricts the set of code generation languages.  If the set is
	// empty, all supported languages are allowed to be generated.
	GenLanguages map[GenLanguage]struct{}
	// Language-specific configurations.
	Go         GoConfig
	Java       JavaConfig
	Javascript JavascriptConfig
}

func (Config) __VDLReflect(struct {
	Name string `vdl:"vdltool.Config"`
}) {
}

func (m *Config) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}
	var var4 bool
	if len(m.GenLanguages) == 0 {
		var4 = true
	}
	if var4 {
		if err := fieldsTarget1.ZeroField("GenLanguages"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("GenLanguages")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}

			setTarget5, err := fieldTarget3.StartSet(tt.NonOptional().Field(0).Type, len(m.GenLanguages))
			if err != nil {
				return err
			}
			for key7 := range m.GenLanguages {
				keyTarget6, err := setTarget5.StartKey()
				if err != nil {
					return err
				}

				if err := key7.FillVDLTarget(keyTarget6, tt.NonOptional().Field(0).Type.Key()); err != nil {
					return err
				}
				if err := setTarget5.FinishKey(keyTarget6); err != nil {
					return err
				}
			}
			if err := fieldTarget3.FinishSet(setTarget5); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
				return err
			}
		}
	}
	var10 := true
	var var11 bool
	if len(m.Go.WireToNativeTypes) == 0 {
		var11 = true
	}
	var10 = var10 && var11
	if var10 {
		if err := fieldsTarget1.ZeroField("Go"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget8, fieldTarget9, err := fieldsTarget1.StartField("Go")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}

			if err := m.Go.FillVDLTarget(fieldTarget9, tt.NonOptional().Field(1).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget8, fieldTarget9); err != nil {
				return err
			}
		}
	}
	var14 := true
	var var15 bool
	if len(m.Java.WireToNativeTypes) == 0 {
		var15 = true
	}
	var14 = var14 && var15
	var var16 bool
	if len(m.Java.WireTypeRenames) == 0 {
		var16 = true
	}
	var14 = var14 && var16
	if var14 {
		if err := fieldsTarget1.ZeroField("Java"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget12, fieldTarget13, err := fieldsTarget1.StartField("Java")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}

			if err := m.Java.FillVDLTarget(fieldTarget13, tt.NonOptional().Field(2).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget12, fieldTarget13); err != nil {
				return err
			}
		}
	}
	var19 := (m.Javascript == JavascriptConfig{})
	if var19 {
		if err := fieldsTarget1.ZeroField("Javascript"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget17, fieldTarget18, err := fieldsTarget1.StartField("Javascript")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}

			if err := m.Javascript.FillVDLTarget(fieldTarget18, tt.NonOptional().Field(3).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget17, fieldTarget18); err != nil {
				return err
			}
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *Config) MakeVDLTarget() vdl.Target {
	return nil
}

type ConfigTarget struct {
	Value *Config

	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *ConfigTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {

	if ttWant := vdl.TypeOf((*Config)(nil)).Elem(); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	return t, nil
}
func (t *ConfigTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "GenLanguages":
		target, err := vdl.ReflectTarget(reflect.ValueOf(&t.Value.GenLanguages))
		return nil, target, err
	case "Go":
		target, err := vdl.ReflectTarget(reflect.ValueOf(&t.Value.Go))
		return nil, target, err
	case "Java":
		target, err := vdl.ReflectTarget(reflect.ValueOf(&t.Value.Java))
		return nil, target, err
	case "Javascript":
		target, err := vdl.ReflectTarget(reflect.ValueOf(&t.Value.Javascript))
		return nil, target, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct vdltool.Config", name)
	}
}
func (t *ConfigTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *ConfigTarget) ZeroField(name string) error {
	switch name {
	case "GenLanguages":
		t.Value.GenLanguages = map[GenLanguage]struct{}(nil)
		return nil
	case "Go":
		t.Value.Go = GoConfig{}
		return nil
	case "Java":
		t.Value.Java = JavaConfig{}
		return nil
	case "Javascript":
		t.Value.Javascript = JavascriptConfig{}
		return nil
	default:
		return fmt.Errorf("field %s not in struct vdltool.Config", name)
	}
}
func (t *ConfigTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

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
	vdl.Register((*GenLanguage)(nil))
	vdl.Register((*GoImport)(nil))
	vdl.Register((*GoType)(nil))
	vdl.Register((*GoConfig)(nil))
	vdl.Register((*JavaConfig)(nil))
	vdl.Register((*JavascriptConfig)(nil))
	vdl.Register((*Config)(nil))

	return struct{}{}
}
