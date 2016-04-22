// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Package: build

// Package build defines interfaces for building executable binaries.
package build

import (
	"fmt"
	"io"
	"v.io/v23"
	"v.io/v23/context"
	"v.io/v23/rpc"
	"v.io/v23/services/binary"
	"v.io/v23/vdl"
)

var _ = __VDLInit() // Must be first; see __VDLInit comments for details.

//////////////////////////////////////////////////
// Type definitions

// Architecture specifies the hardware architecture of a host.
type Architecture int

const (
	ArchitectureAmd64 Architecture = iota
	ArchitectureArm
	ArchitectureX86
)

// ArchitectureAll holds all labels for Architecture.
var ArchitectureAll = [...]Architecture{ArchitectureAmd64, ArchitectureArm, ArchitectureX86}

// ArchitectureFromString creates a Architecture from a string label.
func ArchitectureFromString(label string) (x Architecture, err error) {
	err = x.Set(label)
	return
}

// Set assigns label to x.
func (x *Architecture) Set(label string) error {
	switch label {
	case "Amd64", "amd64":
		*x = ArchitectureAmd64
		return nil
	case "Arm", "arm":
		*x = ArchitectureArm
		return nil
	case "X86", "x86":
		*x = ArchitectureX86
		return nil
	}
	*x = -1
	return fmt.Errorf("unknown label %q in build.Architecture", label)
}

// String returns the string label of x.
func (x Architecture) String() string {
	switch x {
	case ArchitectureAmd64:
		return "Amd64"
	case ArchitectureArm:
		return "Arm"
	case ArchitectureX86:
		return "X86"
	}
	return ""
}

func (Architecture) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/services/build.Architecture"`
	Enum struct{ Amd64, Arm, X86 string }
}) {
}

func (m *Architecture) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	if err := t.FromEnumLabel((*m).String(), tt); err != nil {
		return err
	}
	return nil
}

func (m *Architecture) MakeVDLTarget() vdl.Target {
	return &ArchitectureTarget{Value: m}
}

type ArchitectureTarget struct {
	Value *Architecture
	vdl.TargetBase
}

func (t *ArchitectureTarget) FromEnumLabel(src string, tt *vdl.Type) error {

	if ttWant := vdl.TypeOf((*Architecture)(nil)); !vdl.Compatible(tt, ttWant) {
		return fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	switch src {
	case "Amd64":
		*t.Value = 0
	case "Arm":
		*t.Value = 1
	case "X86":
		*t.Value = 2
	default:
		return fmt.Errorf("label %s not in enum Architecture", src)
	}

	return nil
}

func (x Architecture) VDLIsZero() (bool, error) {
	return x == ArchitectureAmd64, nil
}

func (x Architecture) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(vdl.TypeOf((*Architecture)(nil))); err != nil {
		return err
	}
	if err := enc.EncodeString(x.String()); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *Architecture) VDLRead(dec vdl.Decoder) error {
	if err := dec.StartValue(); err != nil {
		return err
	}
	enum, err := dec.DecodeString()
	if err != nil {
		return err
	}
	if err := x.Set(enum); err != nil {
		return err
	}
	return dec.FinishValue()
}

// Format specifies the file format of a host.
type Format int

const (
	FormatElf Format = iota
	FormatMach
	FormatPe
)

// FormatAll holds all labels for Format.
var FormatAll = [...]Format{FormatElf, FormatMach, FormatPe}

// FormatFromString creates a Format from a string label.
func FormatFromString(label string) (x Format, err error) {
	err = x.Set(label)
	return
}

// Set assigns label to x.
func (x *Format) Set(label string) error {
	switch label {
	case "Elf", "elf":
		*x = FormatElf
		return nil
	case "Mach", "mach":
		*x = FormatMach
		return nil
	case "Pe", "pe":
		*x = FormatPe
		return nil
	}
	*x = -1
	return fmt.Errorf("unknown label %q in build.Format", label)
}

// String returns the string label of x.
func (x Format) String() string {
	switch x {
	case FormatElf:
		return "Elf"
	case FormatMach:
		return "Mach"
	case FormatPe:
		return "Pe"
	}
	return ""
}

func (Format) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/services/build.Format"`
	Enum struct{ Elf, Mach, Pe string }
}) {
}

func (m *Format) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	if err := t.FromEnumLabel((*m).String(), tt); err != nil {
		return err
	}
	return nil
}

func (m *Format) MakeVDLTarget() vdl.Target {
	return &FormatTarget{Value: m}
}

type FormatTarget struct {
	Value *Format
	vdl.TargetBase
}

func (t *FormatTarget) FromEnumLabel(src string, tt *vdl.Type) error {

	if ttWant := vdl.TypeOf((*Format)(nil)); !vdl.Compatible(tt, ttWant) {
		return fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	switch src {
	case "Elf":
		*t.Value = 0
	case "Mach":
		*t.Value = 1
	case "Pe":
		*t.Value = 2
	default:
		return fmt.Errorf("label %s not in enum Format", src)
	}

	return nil
}

func (x Format) VDLIsZero() (bool, error) {
	return x == FormatElf, nil
}

func (x Format) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(vdl.TypeOf((*Format)(nil))); err != nil {
		return err
	}
	if err := enc.EncodeString(x.String()); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *Format) VDLRead(dec vdl.Decoder) error {
	if err := dec.StartValue(); err != nil {
		return err
	}
	enum, err := dec.DecodeString()
	if err != nil {
		return err
	}
	if err := x.Set(enum); err != nil {
		return err
	}
	return dec.FinishValue()
}

// OperatingSystem specifies the operating system of a host.
type OperatingSystem int

const (
	OperatingSystemDarwin OperatingSystem = iota
	OperatingSystemLinux
	OperatingSystemWindows
	OperatingSystemAndroid
)

// OperatingSystemAll holds all labels for OperatingSystem.
var OperatingSystemAll = [...]OperatingSystem{OperatingSystemDarwin, OperatingSystemLinux, OperatingSystemWindows, OperatingSystemAndroid}

// OperatingSystemFromString creates a OperatingSystem from a string label.
func OperatingSystemFromString(label string) (x OperatingSystem, err error) {
	err = x.Set(label)
	return
}

// Set assigns label to x.
func (x *OperatingSystem) Set(label string) error {
	switch label {
	case "Darwin", "darwin":
		*x = OperatingSystemDarwin
		return nil
	case "Linux", "linux":
		*x = OperatingSystemLinux
		return nil
	case "Windows", "windows":
		*x = OperatingSystemWindows
		return nil
	case "Android", "android":
		*x = OperatingSystemAndroid
		return nil
	}
	*x = -1
	return fmt.Errorf("unknown label %q in build.OperatingSystem", label)
}

// String returns the string label of x.
func (x OperatingSystem) String() string {
	switch x {
	case OperatingSystemDarwin:
		return "Darwin"
	case OperatingSystemLinux:
		return "Linux"
	case OperatingSystemWindows:
		return "Windows"
	case OperatingSystemAndroid:
		return "Android"
	}
	return ""
}

func (OperatingSystem) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/services/build.OperatingSystem"`
	Enum struct{ Darwin, Linux, Windows, Android string }
}) {
}

func (m *OperatingSystem) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	if err := t.FromEnumLabel((*m).String(), tt); err != nil {
		return err
	}
	return nil
}

func (m *OperatingSystem) MakeVDLTarget() vdl.Target {
	return &OperatingSystemTarget{Value: m}
}

type OperatingSystemTarget struct {
	Value *OperatingSystem
	vdl.TargetBase
}

func (t *OperatingSystemTarget) FromEnumLabel(src string, tt *vdl.Type) error {

	if ttWant := vdl.TypeOf((*OperatingSystem)(nil)); !vdl.Compatible(tt, ttWant) {
		return fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	switch src {
	case "Darwin":
		*t.Value = 0
	case "Linux":
		*t.Value = 1
	case "Windows":
		*t.Value = 2
	case "Android":
		*t.Value = 3
	default:
		return fmt.Errorf("label %s not in enum OperatingSystem", src)
	}

	return nil
}

func (x OperatingSystem) VDLIsZero() (bool, error) {
	return x == OperatingSystemDarwin, nil
}

func (x OperatingSystem) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(vdl.TypeOf((*OperatingSystem)(nil))); err != nil {
		return err
	}
	if err := enc.EncodeString(x.String()); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *OperatingSystem) VDLRead(dec vdl.Decoder) error {
	if err := dec.StartValue(); err != nil {
		return err
	}
	enum, err := dec.DecodeString()
	if err != nil {
		return err
	}
	if err := x.Set(enum); err != nil {
		return err
	}
	return dec.FinishValue()
}

// File records the name and contents of a file.
type File struct {
	Name     string
	Contents []byte
}

func (File) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/services/build.File"`
}) {
}

func (m *File) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
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
	if len(m.Contents) == 0 {
		var7 = true
	}
	if var7 {
		if err := fieldsTarget1.ZeroField("Contents"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget5, fieldTarget6, err := fieldsTarget1.StartField("Contents")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}

			if err := fieldTarget6.FromBytes([]byte(m.Contents), tt.NonOptional().Field(1).Type); err != nil {
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

func (m *File) MakeVDLTarget() vdl.Target {
	return &FileTarget{Value: m}
}

type FileTarget struct {
	Value          *File
	nameTarget     vdl.StringTarget
	contentsTarget vdl.BytesTarget
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *FileTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {

	if ttWant := vdl.TypeOf((*File)(nil)).Elem(); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	return t, nil
}
func (t *FileTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "Name":
		t.nameTarget.Value = &t.Value.Name
		target, err := &t.nameTarget, error(nil)
		return nil, target, err
	case "Contents":
		t.contentsTarget.Value = &t.Value.Contents
		target, err := &t.contentsTarget, error(nil)
		return nil, target, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct v.io/v23/services/build.File", name)
	}
}
func (t *FileTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *FileTarget) ZeroField(name string) error {
	switch name {
	case "Name":
		t.Value.Name = ""
		return nil
	case "Contents":
		t.Value.Contents = []byte(nil)
		return nil
	default:
		return fmt.Errorf("field %s not in struct v.io/v23/services/build.File", name)
	}
}
func (t *FileTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

func (x File) VDLIsZero() (bool, error) {
	if x.Name != "" {
		return false, nil
	}
	if len(x.Contents) != 0 {
		return false, nil
	}
	return true, nil
}

func (x File) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(vdl.TypeOf((*File)(nil)).Elem()); err != nil {
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
	if len(x.Contents) != 0 {
		if err := enc.NextField("Contents"); err != nil {
			return err
		}
		if err := enc.StartValue(vdl.TypeOf((*[]byte)(nil))); err != nil {
			return err
		}
		if err := enc.EncodeBytes(x.Contents); err != nil {
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

func (x *File) VDLRead(dec vdl.Decoder) error {
	*x = File{}
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
		case "Contents":
			if err := dec.StartValue(); err != nil {
				return err
			}
			if err := dec.DecodeBytes(-1, &x.Contents); err != nil {
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
// Interface definitions

// BuilderClientMethods is the client interface
// containing Builder methods.
//
// Builder describes an interface for building binaries from source.
type BuilderClientMethods interface {
	// Build streams sources to the build server, which then attempts to
	// build the sources and streams back the compiled binaries.
	Build(_ *context.T, arch Architecture, os OperatingSystem, _ ...rpc.CallOpt) (BuilderBuildClientCall, error)
	// Describe generates a description for a binary identified by
	// the given Object name.
	Describe(_ *context.T, name string, _ ...rpc.CallOpt) (binary.Description, error)
}

// BuilderClientStub adds universal methods to BuilderClientMethods.
type BuilderClientStub interface {
	BuilderClientMethods
	rpc.UniversalServiceMethods
}

// BuilderClient returns a client stub for Builder.
func BuilderClient(name string) BuilderClientStub {
	return implBuilderClientStub{name}
}

type implBuilderClientStub struct {
	name string
}

func (c implBuilderClientStub) Build(ctx *context.T, i0 Architecture, i1 OperatingSystem, opts ...rpc.CallOpt) (ocall BuilderBuildClientCall, err error) {
	var call rpc.ClientCall
	if call, err = v23.GetClient(ctx).StartCall(ctx, c.name, "Build", []interface{}{i0, i1}, opts...); err != nil {
		return
	}
	ocall = &implBuilderBuildClientCall{ClientCall: call}
	return
}

func (c implBuilderClientStub) Describe(ctx *context.T, i0 string, opts ...rpc.CallOpt) (o0 binary.Description, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "Describe", []interface{}{i0}, []interface{}{&o0}, opts...)
	return
}

// BuilderBuildClientStream is the client stream for Builder.Build.
type BuilderBuildClientStream interface {
	// RecvStream returns the receiver side of the Builder.Build client stream.
	RecvStream() interface {
		// Advance stages an item so that it may be retrieved via Value.  Returns
		// true iff there is an item to retrieve.  Advance must be called before
		// Value is called.  May block if an item is not available.
		Advance() bool
		// Value returns the item that was staged by Advance.  May panic if Advance
		// returned false or was not called.  Never blocks.
		Value() File
		// Err returns any error encountered by Advance.  Never blocks.
		Err() error
	}
	// SendStream returns the send side of the Builder.Build client stream.
	SendStream() interface {
		// Send places the item onto the output stream.  Returns errors
		// encountered while sending, or if Send is called after Close or
		// the stream has been canceled.  Blocks if there is no buffer
		// space; will unblock when buffer space is available or after
		// the stream has been canceled.
		Send(item File) error
		// Close indicates to the server that no more items will be sent;
		// server Recv calls will receive io.EOF after all sent items.
		// This is an optional call - e.g. a client might call Close if it
		// needs to continue receiving items from the server after it's
		// done sending.  Returns errors encountered while closing, or if
		// Close is called after the stream has been canceled.  Like Send,
		// blocks if there is no buffer space available.
		Close() error
	}
}

// BuilderBuildClientCall represents the call returned from Builder.Build.
type BuilderBuildClientCall interface {
	BuilderBuildClientStream
	// Finish performs the equivalent of SendStream().Close, then blocks until
	// the server is done, and returns the positional return values for the call.
	//
	// Finish returns immediately if the call has been canceled; depending on the
	// timing the output could either be an error signaling cancelation, or the
	// valid positional return values from the server.
	//
	// Calling Finish is mandatory for releasing stream resources, unless the call
	// has been canceled or any of the other methods return an error.  Finish should
	// be called at most once.
	Finish() ([]byte, error)
}

type implBuilderBuildClientCall struct {
	rpc.ClientCall
	valRecv File
	errRecv error
}

func (c *implBuilderBuildClientCall) RecvStream() interface {
	Advance() bool
	Value() File
	Err() error
} {
	return implBuilderBuildClientCallRecv{c}
}

type implBuilderBuildClientCallRecv struct {
	c *implBuilderBuildClientCall
}

func (c implBuilderBuildClientCallRecv) Advance() bool {
	c.c.valRecv = File{}
	c.c.errRecv = c.c.Recv(&c.c.valRecv)
	return c.c.errRecv == nil
}
func (c implBuilderBuildClientCallRecv) Value() File {
	return c.c.valRecv
}
func (c implBuilderBuildClientCallRecv) Err() error {
	if c.c.errRecv == io.EOF {
		return nil
	}
	return c.c.errRecv
}
func (c *implBuilderBuildClientCall) SendStream() interface {
	Send(item File) error
	Close() error
} {
	return implBuilderBuildClientCallSend{c}
}

type implBuilderBuildClientCallSend struct {
	c *implBuilderBuildClientCall
}

func (c implBuilderBuildClientCallSend) Send(item File) error {
	return c.c.Send(item)
}
func (c implBuilderBuildClientCallSend) Close() error {
	return c.c.CloseSend()
}
func (c *implBuilderBuildClientCall) Finish() (o0 []byte, err error) {
	err = c.ClientCall.Finish(&o0)
	return
}

// BuilderServerMethods is the interface a server writer
// implements for Builder.
//
// Builder describes an interface for building binaries from source.
type BuilderServerMethods interface {
	// Build streams sources to the build server, which then attempts to
	// build the sources and streams back the compiled binaries.
	Build(_ *context.T, _ BuilderBuildServerCall, arch Architecture, os OperatingSystem) ([]byte, error)
	// Describe generates a description for a binary identified by
	// the given Object name.
	Describe(_ *context.T, _ rpc.ServerCall, name string) (binary.Description, error)
}

// BuilderServerStubMethods is the server interface containing
// Builder methods, as expected by rpc.Server.
// The only difference between this interface and BuilderServerMethods
// is the streaming methods.
type BuilderServerStubMethods interface {
	// Build streams sources to the build server, which then attempts to
	// build the sources and streams back the compiled binaries.
	Build(_ *context.T, _ *BuilderBuildServerCallStub, arch Architecture, os OperatingSystem) ([]byte, error)
	// Describe generates a description for a binary identified by
	// the given Object name.
	Describe(_ *context.T, _ rpc.ServerCall, name string) (binary.Description, error)
}

// BuilderServerStub adds universal methods to BuilderServerStubMethods.
type BuilderServerStub interface {
	BuilderServerStubMethods
	// Describe the Builder interfaces.
	Describe__() []rpc.InterfaceDesc
}

// BuilderServer returns a server stub for Builder.
// It converts an implementation of BuilderServerMethods into
// an object that may be used by rpc.Server.
func BuilderServer(impl BuilderServerMethods) BuilderServerStub {
	stub := implBuilderServerStub{
		impl: impl,
	}
	// Initialize GlobState; always check the stub itself first, to handle the
	// case where the user has the Glob method defined in their VDL source.
	if gs := rpc.NewGlobState(stub); gs != nil {
		stub.gs = gs
	} else if gs := rpc.NewGlobState(impl); gs != nil {
		stub.gs = gs
	}
	return stub
}

type implBuilderServerStub struct {
	impl BuilderServerMethods
	gs   *rpc.GlobState
}

func (s implBuilderServerStub) Build(ctx *context.T, call *BuilderBuildServerCallStub, i0 Architecture, i1 OperatingSystem) ([]byte, error) {
	return s.impl.Build(ctx, call, i0, i1)
}

func (s implBuilderServerStub) Describe(ctx *context.T, call rpc.ServerCall, i0 string) (binary.Description, error) {
	return s.impl.Describe(ctx, call, i0)
}

func (s implBuilderServerStub) Globber() *rpc.GlobState {
	return s.gs
}

func (s implBuilderServerStub) Describe__() []rpc.InterfaceDesc {
	return []rpc.InterfaceDesc{BuilderDesc}
}

// BuilderDesc describes the Builder interface.
var BuilderDesc rpc.InterfaceDesc = descBuilder

// descBuilder hides the desc to keep godoc clean.
var descBuilder = rpc.InterfaceDesc{
	Name:    "Builder",
	PkgPath: "v.io/v23/services/build",
	Doc:     "// Builder describes an interface for building binaries from source.",
	Methods: []rpc.MethodDesc{
		{
			Name: "Build",
			Doc:  "// Build streams sources to the build server, which then attempts to\n// build the sources and streams back the compiled binaries.",
			InArgs: []rpc.ArgDesc{
				{"arch", ``}, // Architecture
				{"os", ``},   // OperatingSystem
			},
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // []byte
			},
		},
		{
			Name: "Describe",
			Doc:  "// Describe generates a description for a binary identified by\n// the given Object name.",
			InArgs: []rpc.ArgDesc{
				{"name", ``}, // string
			},
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // binary.Description
			},
		},
	},
}

// BuilderBuildServerStream is the server stream for Builder.Build.
type BuilderBuildServerStream interface {
	// RecvStream returns the receiver side of the Builder.Build server stream.
	RecvStream() interface {
		// Advance stages an item so that it may be retrieved via Value.  Returns
		// true iff there is an item to retrieve.  Advance must be called before
		// Value is called.  May block if an item is not available.
		Advance() bool
		// Value returns the item that was staged by Advance.  May panic if Advance
		// returned false or was not called.  Never blocks.
		Value() File
		// Err returns any error encountered by Advance.  Never blocks.
		Err() error
	}
	// SendStream returns the send side of the Builder.Build server stream.
	SendStream() interface {
		// Send places the item onto the output stream.  Returns errors encountered
		// while sending.  Blocks if there is no buffer space; will unblock when
		// buffer space is available.
		Send(item File) error
	}
}

// BuilderBuildServerCall represents the context passed to Builder.Build.
type BuilderBuildServerCall interface {
	rpc.ServerCall
	BuilderBuildServerStream
}

// BuilderBuildServerCallStub is a wrapper that converts rpc.StreamServerCall into
// a typesafe stub that implements BuilderBuildServerCall.
type BuilderBuildServerCallStub struct {
	rpc.StreamServerCall
	valRecv File
	errRecv error
}

// Init initializes BuilderBuildServerCallStub from rpc.StreamServerCall.
func (s *BuilderBuildServerCallStub) Init(call rpc.StreamServerCall) {
	s.StreamServerCall = call
}

// RecvStream returns the receiver side of the Builder.Build server stream.
func (s *BuilderBuildServerCallStub) RecvStream() interface {
	Advance() bool
	Value() File
	Err() error
} {
	return implBuilderBuildServerCallRecv{s}
}

type implBuilderBuildServerCallRecv struct {
	s *BuilderBuildServerCallStub
}

func (s implBuilderBuildServerCallRecv) Advance() bool {
	s.s.valRecv = File{}
	s.s.errRecv = s.s.Recv(&s.s.valRecv)
	return s.s.errRecv == nil
}
func (s implBuilderBuildServerCallRecv) Value() File {
	return s.s.valRecv
}
func (s implBuilderBuildServerCallRecv) Err() error {
	if s.s.errRecv == io.EOF {
		return nil
	}
	return s.s.errRecv
}

// SendStream returns the send side of the Builder.Build server stream.
func (s *BuilderBuildServerCallStub) SendStream() interface {
	Send(item File) error
} {
	return implBuilderBuildServerCallSend{s}
}

type implBuilderBuildServerCallSend struct {
	s *BuilderBuildServerCallStub
}

func (s implBuilderBuildServerCallSend) Send(item File) error {
	return s.s.Send(item)
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
	vdl.Register((*Architecture)(nil))
	vdl.Register((*Format)(nil))
	vdl.Register((*OperatingSystem)(nil))
	vdl.Register((*File)(nil))

	return struct{}{}
}
