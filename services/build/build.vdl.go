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
	if err := t.FromEnumLabel((*m).String(), __VDLType_v_io_v23_services_build_Architecture); err != nil {
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
	if !vdl.Compatible(tt, __VDLType_v_io_v23_services_build_Architecture) {
		return fmt.Errorf("type %v incompatible with %v", tt, __VDLType_v_io_v23_services_build_Architecture)
	}
	switch src {
	case "Amd64":
		*t.Value = 0
	case "Arm":
		*t.Value = 1
	case "X86":
		*t.Value = 2
	default:
		return fmt.Errorf("label %s not in enum %v", src, __VDLType_v_io_v23_services_build_Architecture)
	}

	return nil
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
	if err := t.FromEnumLabel((*m).String(), __VDLType_v_io_v23_services_build_Format); err != nil {
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
	if !vdl.Compatible(tt, __VDLType_v_io_v23_services_build_Format) {
		return fmt.Errorf("type %v incompatible with %v", tt, __VDLType_v_io_v23_services_build_Format)
	}
	switch src {
	case "Elf":
		*t.Value = 0
	case "Mach":
		*t.Value = 1
	case "Pe":
		*t.Value = 2
	default:
		return fmt.Errorf("label %s not in enum %v", src, __VDLType_v_io_v23_services_build_Format)
	}

	return nil
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
	if err := t.FromEnumLabel((*m).String(), __VDLType_v_io_v23_services_build_OperatingSystem); err != nil {
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
	if !vdl.Compatible(tt, __VDLType_v_io_v23_services_build_OperatingSystem) {
		return fmt.Errorf("type %v incompatible with %v", tt, __VDLType_v_io_v23_services_build_OperatingSystem)
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
		return fmt.Errorf("label %s not in enum %v", src, __VDLType_v_io_v23_services_build_OperatingSystem)
	}

	return nil
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
	if __VDLType_v_io_v23_services_build_File == nil || __VDLType0 == nil {
		panic("Initialization order error: types generated for FillVDLTarget not initialized. Consider moving caller to an init() block.")
	}
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}

	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Name")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget3.FromString(string(m.Name), vdl.StringType); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
			return err
		}
	}
	keyTarget4, fieldTarget5, err := fieldsTarget1.StartField("Contents")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		if err := fieldTarget5.FromBytes([]byte(m.Contents), __VDLType1); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget4, fieldTarget5); err != nil {
			return err
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
	if !vdl.Compatible(tt, __VDLType_v_io_v23_services_build_File) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, __VDLType_v_io_v23_services_build_File)
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
		return nil, nil, fmt.Errorf("field %s not in struct %v", name, __VDLType_v_io_v23_services_build_File)
	}
}
func (t *FileTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *FileTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

func init() {
	vdl.Register((*Architecture)(nil))
	vdl.Register((*Format)(nil))
	vdl.Register((*OperatingSystem)(nil))
	vdl.Register((*File)(nil))
}

var __VDLType0 *vdl.Type = vdl.TypeOf((*File)(nil))
var __VDLType1 *vdl.Type = vdl.TypeOf([]byte(nil))
var __VDLType_v_io_v23_services_build_Architecture *vdl.Type = vdl.TypeOf(ArchitectureAmd64)
var __VDLType_v_io_v23_services_build_File *vdl.Type = vdl.TypeOf(File{})
var __VDLType_v_io_v23_services_build_Format *vdl.Type = vdl.TypeOf(FormatElf)
var __VDLType_v_io_v23_services_build_OperatingSystem *vdl.Type = vdl.TypeOf(OperatingSystemDarwin)

func __VDLEnsureNativeBuilt() {
}

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
