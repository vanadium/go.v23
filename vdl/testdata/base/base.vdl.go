// This file was auto-generated by the veyron vdl tool.
// Source: base.vdl

// Package base is a simple single-file test of vdl functionality.
package base

import (
	// VDL system imports
	"fmt"
	"io"
	"v.io/core/veyron2"
	"v.io/core/veyron2/context"
	"v.io/core/veyron2/i18n"
	"v.io/core/veyron2/ipc"
	"v.io/core/veyron2/vdl"
	"v.io/core/veyron2/verror"
)

type NamedBool bool

func (NamedBool) __VDLReflect(struct {
	Name string "v.io/core/veyron2/vdl/testdata/base.NamedBool"
}) {
}

type NamedByte byte

func (NamedByte) __VDLReflect(struct {
	Name string "v.io/core/veyron2/vdl/testdata/base.NamedByte"
}) {
}

type NamedUint16 uint16

func (NamedUint16) __VDLReflect(struct {
	Name string "v.io/core/veyron2/vdl/testdata/base.NamedUint16"
}) {
}

type NamedUint32 uint32

func (NamedUint32) __VDLReflect(struct {
	Name string "v.io/core/veyron2/vdl/testdata/base.NamedUint32"
}) {
}

type NamedUint64 uint64

func (NamedUint64) __VDLReflect(struct {
	Name string "v.io/core/veyron2/vdl/testdata/base.NamedUint64"
}) {
}

type NamedInt16 int16

func (NamedInt16) __VDLReflect(struct {
	Name string "v.io/core/veyron2/vdl/testdata/base.NamedInt16"
}) {
}

type NamedInt32 int32

func (NamedInt32) __VDLReflect(struct {
	Name string "v.io/core/veyron2/vdl/testdata/base.NamedInt32"
}) {
}

type NamedInt64 int64

func (NamedInt64) __VDLReflect(struct {
	Name string "v.io/core/veyron2/vdl/testdata/base.NamedInt64"
}) {
}

type NamedFloat32 float32

func (NamedFloat32) __VDLReflect(struct {
	Name string "v.io/core/veyron2/vdl/testdata/base.NamedFloat32"
}) {
}

type NamedFloat64 float64

func (NamedFloat64) __VDLReflect(struct {
	Name string "v.io/core/veyron2/vdl/testdata/base.NamedFloat64"
}) {
}

type NamedComplex64 complex64

func (NamedComplex64) __VDLReflect(struct {
	Name string "v.io/core/veyron2/vdl/testdata/base.NamedComplex64"
}) {
}

type NamedComplex128 complex128

func (NamedComplex128) __VDLReflect(struct {
	Name string "v.io/core/veyron2/vdl/testdata/base.NamedComplex128"
}) {
}

type NamedString string

func (NamedString) __VDLReflect(struct {
	Name string "v.io/core/veyron2/vdl/testdata/base.NamedString"
}) {
}

type NamedEnum int

const (
	NamedEnumA NamedEnum = iota
	NamedEnumB
	NamedEnumC
)

// NamedEnumAll holds all labels for NamedEnum.
var NamedEnumAll = []NamedEnum{NamedEnumA, NamedEnumB, NamedEnumC}

// NamedEnumFromString creates a NamedEnum from a string label.
func NamedEnumFromString(label string) (x NamedEnum, err error) {
	err = x.Set(label)
	return
}

// Set assigns label to x.
func (x *NamedEnum) Set(label string) error {
	switch label {
	case "A", "a":
		*x = NamedEnumA
		return nil
	case "B", "b":
		*x = NamedEnumB
		return nil
	case "C", "c":
		*x = NamedEnumC
		return nil
	}
	*x = -1
	return fmt.Errorf("unknown label %q in base.NamedEnum", label)
}

// String returns the string label of x.
func (x NamedEnum) String() string {
	switch x {
	case NamedEnumA:
		return "A"
	case NamedEnumB:
		return "B"
	case NamedEnumC:
		return "C"
	}
	return ""
}

func (NamedEnum) __VDLReflect(struct {
	Name string "v.io/core/veyron2/vdl/testdata/base.NamedEnum"
	Enum struct{ A, B, C string }
}) {
}

type NamedArray [2]bool

func (NamedArray) __VDLReflect(struct {
	Name string "v.io/core/veyron2/vdl/testdata/base.NamedArray"
}) {
}

type NamedList []uint32

func (NamedList) __VDLReflect(struct {
	Name string "v.io/core/veyron2/vdl/testdata/base.NamedList"
}) {
}

type NamedSet map[string]struct{}

func (NamedSet) __VDLReflect(struct {
	Name string "v.io/core/veyron2/vdl/testdata/base.NamedSet"
}) {
}

type NamedMap map[string]float32

func (NamedMap) __VDLReflect(struct {
	Name string "v.io/core/veyron2/vdl/testdata/base.NamedMap"
}) {
}

type NamedStruct struct {
	A bool
	B string
	C int32
}

func (NamedStruct) __VDLReflect(struct {
	Name string "v.io/core/veyron2/vdl/testdata/base.NamedStruct"
}) {
}

type (
	// NamedUnion represents any single field of the NamedUnion union type.
	NamedUnion interface {
		// Index returns the field index.
		Index() int
		// Interface returns the field value as an interface.
		Interface() interface{}
		// Name returns the field name.
		Name() string
		// __VDLReflect describes the NamedUnion union type.
		__VDLReflect(__NamedUnionReflect)
	}
	// NamedUnionA represents field A of the NamedUnion union type.
	NamedUnionA struct{ Value bool }
	// NamedUnionB represents field B of the NamedUnion union type.
	NamedUnionB struct{ Value string }
	// NamedUnionC represents field C of the NamedUnion union type.
	NamedUnionC struct{ Value int32 }
	// __NamedUnionReflect describes the NamedUnion union type.
	__NamedUnionReflect struct {
		Name  string "v.io/core/veyron2/vdl/testdata/base.NamedUnion"
		Type  NamedUnion
		Union struct {
			A NamedUnionA
			B NamedUnionB
			C NamedUnionC
		}
	}
)

func (x NamedUnionA) Index() int                       { return 0 }
func (x NamedUnionA) Interface() interface{}           { return x.Value }
func (x NamedUnionA) Name() string                     { return "A" }
func (x NamedUnionA) __VDLReflect(__NamedUnionReflect) {}

func (x NamedUnionB) Index() int                       { return 1 }
func (x NamedUnionB) Interface() interface{}           { return x.Value }
func (x NamedUnionB) Name() string                     { return "B" }
func (x NamedUnionB) __VDLReflect(__NamedUnionReflect) {}

func (x NamedUnionC) Index() int                       { return 2 }
func (x NamedUnionC) Interface() interface{}           { return x.Value }
func (x NamedUnionC) Name() string                     { return "C" }
func (x NamedUnionC) __VDLReflect(__NamedUnionReflect) {}

type Scalars struct {
	A0  bool
	A1  byte
	A2  uint16
	A3  uint32
	A4  uint64
	A5  int16
	A6  int32
	A7  int64
	A8  float32
	A9  float64
	A10 complex64
	A11 complex128
	A12 string
	A13 error
	A14 vdl.AnyRep
	A15 *vdl.Type
	B0  NamedBool
	B1  NamedByte
	B2  NamedUint16
	B3  NamedUint32
	B4  NamedUint64
	B5  NamedInt16
	B6  NamedInt32
	B7  NamedInt64
	B8  NamedFloat32
	B9  NamedFloat64
	B10 NamedComplex64
	B11 NamedComplex128
	B12 NamedString
	B13 NamedEnum
	B14 NamedUnion
}

func (Scalars) __VDLReflect(struct {
	Name string "v.io/core/veyron2/vdl/testdata/base.Scalars"
}) {
}

// These are all scalars that may be used as map or set keys.
type KeyScalars struct {
	A0  bool
	A1  byte
	A2  uint16
	A3  uint32
	A4  uint64
	A5  int16
	A6  int32
	A7  int64
	A8  float32
	A9  float64
	A10 complex64
	A11 complex128
	A12 string
	B0  NamedBool
	B1  NamedByte
	B2  NamedUint16
	B3  NamedUint32
	B4  NamedUint64
	B5  NamedInt16
	B6  NamedInt32
	B7  NamedInt64
	B8  NamedFloat32
	B9  NamedFloat64
	B10 NamedComplex64
	B11 NamedComplex128
	B12 NamedString
}

func (KeyScalars) __VDLReflect(struct {
	Name string "v.io/core/veyron2/vdl/testdata/base.KeyScalars"
}) {
}

type ScalarsArray [2]Scalars

func (ScalarsArray) __VDLReflect(struct {
	Name string "v.io/core/veyron2/vdl/testdata/base.ScalarsArray"
}) {
}

type Composites struct {
	A0 Scalars
	A1 ScalarsArray
	A2 []Scalars
	A3 map[KeyScalars]struct{}
	A4 map[string]Scalars
	A5 map[KeyScalars][]map[string]complex128
}

func (Composites) __VDLReflect(struct {
	Name string "v.io/core/veyron2/vdl/testdata/base.Composites"
}) {
}

type CompositesArray [2]Composites

func (CompositesArray) __VDLReflect(struct {
	Name string "v.io/core/veyron2/vdl/testdata/base.CompositesArray"
}) {
}

type CompComp struct {
	A0 Composites
	A1 CompositesArray
	A2 []Composites
	A3 map[string]Composites
	A4 map[KeyScalars][]map[string]Composites
}

func (CompComp) __VDLReflect(struct {
	Name string "v.io/core/veyron2/vdl/testdata/base.CompComp"
}) {
}

// NestedArgs is defined before Args; that's allowed in regular Go, and also
// allowed in our vdl files.  The compiler will re-order dependent types to ease
// code generation in other languages.
type NestedArgs struct {
	Args Args
}

func (NestedArgs) __VDLReflect(struct {
	Name string "v.io/core/veyron2/vdl/testdata/base.NestedArgs"
}) {
}

// Args will be reordered to show up before NestedArgs in the generated output.
type Args struct {
	A int32
	B int32
}

func (Args) __VDLReflect(struct {
	Name string "v.io/core/veyron2/vdl/testdata/base.Args"
}) {
}

func init() {
	vdl.Register((*NamedBool)(nil))
	vdl.Register((*NamedByte)(nil))
	vdl.Register((*NamedUint16)(nil))
	vdl.Register((*NamedUint32)(nil))
	vdl.Register((*NamedUint64)(nil))
	vdl.Register((*NamedInt16)(nil))
	vdl.Register((*NamedInt32)(nil))
	vdl.Register((*NamedInt64)(nil))
	vdl.Register((*NamedFloat32)(nil))
	vdl.Register((*NamedFloat64)(nil))
	vdl.Register((*NamedComplex64)(nil))
	vdl.Register((*NamedComplex128)(nil))
	vdl.Register((*NamedString)(nil))
	vdl.Register((*NamedEnum)(nil))
	vdl.Register((*NamedArray)(nil))
	vdl.Register((*NamedList)(nil))
	vdl.Register((*NamedSet)(nil))
	vdl.Register((*NamedMap)(nil))
	vdl.Register((*NamedStruct)(nil))
	vdl.Register((*NamedUnion)(nil))
	vdl.Register((*Scalars)(nil))
	vdl.Register((*KeyScalars)(nil))
	vdl.Register((*ScalarsArray)(nil))
	vdl.Register((*Composites)(nil))
	vdl.Register((*CompositesArray)(nil))
	vdl.Register((*CompComp)(nil))
	vdl.Register((*NestedArgs)(nil))
	vdl.Register((*Args)(nil))
}

const Cbool = true

const Cbyte = byte(1)

const Cint32 = int32(2)

const Cint64 = int64(3)

const Cuint32 = uint32(4)

const Cuint64 = uint64(5)

const Cfloat32 = float32(6)

const Cfloat64 = float64(7)

const CNamedBool = NamedBool(true)

var CNamedStruct = NamedStruct{
	A: true,
	B: "test",
}

const Ccomplex64 = complex64(8 + 9i)

const Ccomplex128 = complex128(10 + 11i)

const Cstring = "foo"

const Cenum = NamedEnumA

var Cunion = NamedUnion(NamedUnionA{true})

var Carray = NamedArray{
	true,
	false,
}

var Clist = []int32{
	1,
	2,
	3,
}

var Cset = map[int32]struct{}{
	1: struct{}{},
	2: struct{}{},
	3: struct{}{},
}

var cmap = map[int32]string{
	1: "A",
	2: "B",
	3: "C",
}

var Cargs = Args{
	A: 1,
	B: 2,
}

const True = true

const Foo = "foo"

const Five = int32(5)

const Six = uint64(6)

const SixSquared = uint64(36)

const FiveSquared = int32(25)

var CTObool = vdl.TypeOf(false)

var CTOstring = vdl.TypeOf("")

var CTObytes = vdl.TypeOf([]byte(""))

var CTObyte = vdl.TypeOf(byte(0))

var CTOuint16 = vdl.TypeOf(uint16(0))

var CTOint16 = vdl.TypeOf(int16(0))

var CTOfloat32 = vdl.TypeOf(float32(0))

var CTOcomplex64 = vdl.TypeOf(complex64(0))

var CTOenum = vdl.TypeOf(NamedEnumA)

var CTOArray = vdl.TypeOf(NamedArray{})

var CTOList = vdl.TypeOf([]string(nil))

var CTOSet = vdl.TypeOf(map[string]struct{}(nil))

var CTOMap = vdl.TypeOf(map[string]int64(nil))

var CTOStruct = vdl.TypeOf(Scalars{
	A15: vdl.AnyType,
	B14: NamedUnionA{false},
})

var CTOUnion = vdl.TypeOf(NamedUnion(NamedUnionA{false}))

var CTOTypeObject = vdl.TypeObjectType

var CTOAny = vdl.AnyType

var (
	ErrNoParams1   = verror.Register("v.io/core/veyron2/vdl/testdata/base.NoParams1", verror.NoRetry, "{1:}{2:} en msg")
	ErrNoParams2   = verror.Register("v.io/core/veyron2/vdl/testdata/base.NoParams2", verror.RetryRefetch, "{1:}{2:} en msg")
	ErrWithParams1 = verror.Register("v.io/core/veyron2/vdl/testdata/base.WithParams1", verror.NoRetry, "{1:}{2:} en x={3} y={4}")
	ErrWithParams2 = verror.Register("v.io/core/veyron2/vdl/testdata/base.WithParams2", verror.RetryRefetch, "{1:}{2:} en x={3} y={4}")
	errNotExported = verror.Register("v.io/core/veyron2/vdl/testdata/base.notExported", verror.NoRetry, "{1:}{2:} en x={3} y={4}")
)

func init() {
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrNoParams1.ID), "{1:}{2:} en msg")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrNoParams2.ID), "{1:}{2:} en msg")
	i18n.Cat().SetWithBase(i18n.LangID("fr"), i18n.MsgID(ErrNoParams2.ID), "{1:}{2:} fr msg")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrWithParams1.ID), "{1:}{2:} en x={3} y={4}")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrWithParams2.ID), "{1:}{2:} en x={3} y={4}")
	i18n.Cat().SetWithBase(i18n.LangID("fr"), i18n.MsgID(ErrWithParams2.ID), "{1:}{2:} fr y={4} x={3}")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(errNotExported.ID), "{1:}{2:} en x={3} y={4}")
}

// NewErrNoParams1 returns an error with the ErrNoParams1 ID.
func NewErrNoParams1(ctx *context.T) error {
	return verror.New(ErrNoParams1, ctx)
}

// NewErrNoParams2 returns an error with the ErrNoParams2 ID.
func NewErrNoParams2(ctx *context.T) error {
	return verror.New(ErrNoParams2, ctx)
}

// NewErrWithParams1 returns an error with the ErrWithParams1 ID.
func NewErrWithParams1(ctx *context.T, x string, y int32) error {
	return verror.New(ErrWithParams1, ctx, x, y)
}

// NewErrWithParams2 returns an error with the ErrWithParams2 ID.
func NewErrWithParams2(ctx *context.T, x string, y int32) error {
	return verror.New(ErrWithParams2, ctx, x, y)
}

// newErrNotExported returns an error with the errNotExported ID.
func newErrNotExported(ctx *context.T, x string, y int32) error {
	return verror.New(errNotExported, ctx, x, y)
}

// ServiceAClientMethods is the client interface
// containing ServiceA methods.
type ServiceAClientMethods interface {
	MethodA1(*context.T, ...ipc.CallOpt) error
	MethodA2(ctx *context.T, a int32, b string, opts ...ipc.CallOpt) (s string, err error)
	MethodA3(ctx *context.T, a int32, opts ...ipc.CallOpt) (ServiceAMethodA3Call, error)
	MethodA4(ctx *context.T, a int32, opts ...ipc.CallOpt) (ServiceAMethodA4Call, error)
}

// ServiceAClientStub adds universal methods to ServiceAClientMethods.
type ServiceAClientStub interface {
	ServiceAClientMethods
	ipc.UniversalServiceMethods
}

// ServiceAClient returns a client stub for ServiceA.
func ServiceAClient(name string, opts ...ipc.BindOpt) ServiceAClientStub {
	var client ipc.Client
	for _, opt := range opts {
		if clientOpt, ok := opt.(ipc.Client); ok {
			client = clientOpt
		}
	}
	return implServiceAClientStub{name, client}
}

type implServiceAClientStub struct {
	name   string
	client ipc.Client
}

func (c implServiceAClientStub) c(ctx *context.T) ipc.Client {
	if c.client != nil {
		return c.client
	}
	return veyron2.GetClient(ctx)
}

func (c implServiceAClientStub) MethodA1(ctx *context.T, opts ...ipc.CallOpt) (err error) {
	var call ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "MethodA1", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (c implServiceAClientStub) MethodA2(ctx *context.T, i0 int32, i1 string, opts ...ipc.CallOpt) (o0 string, err error) {
	var call ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "MethodA2", []interface{}{i0, i1}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&o0, &err); ierr != nil {
		err = ierr
	}
	return
}

func (c implServiceAClientStub) MethodA3(ctx *context.T, i0 int32, opts ...ipc.CallOpt) (ocall ServiceAMethodA3Call, err error) {
	var call ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "MethodA3", []interface{}{i0}, opts...); err != nil {
		return
	}
	ocall = &implServiceAMethodA3Call{Call: call}
	return
}

func (c implServiceAClientStub) MethodA4(ctx *context.T, i0 int32, opts ...ipc.CallOpt) (ocall ServiceAMethodA4Call, err error) {
	var call ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "MethodA4", []interface{}{i0}, opts...); err != nil {
		return
	}
	ocall = &implServiceAMethodA4Call{Call: call}
	return
}

// ServiceAMethodA3ClientStream is the client stream for ServiceA.MethodA3.
type ServiceAMethodA3ClientStream interface {
	// RecvStream returns the receiver side of the ServiceA.MethodA3 client stream.
	RecvStream() interface {
		// Advance stages an item so that it may be retrieved via Value.  Returns
		// true iff there is an item to retrieve.  Advance must be called before
		// Value is called.  May block if an item is not available.
		Advance() bool
		// Value returns the item that was staged by Advance.  May panic if Advance
		// returned false or was not called.  Never blocks.
		Value() Scalars
		// Err returns any error encountered by Advance.  Never blocks.
		Err() error
	}
}

// ServiceAMethodA3Call represents the call returned from ServiceA.MethodA3.
type ServiceAMethodA3Call interface {
	ServiceAMethodA3ClientStream
	// Finish blocks until the server is done, and returns the positional return
	// values for call.
	//
	// Finish returns immediately if the call has been canceled; depending on the
	// timing the output could either be an error signaling cancelation, or the
	// valid positional return values from the server.
	//
	// Calling Finish is mandatory for releasing stream resources, unless the call
	// has been canceled or any of the other methods return an error.  Finish should
	// be called at most once.
	Finish() (s string, err error)
}

type implServiceAMethodA3Call struct {
	ipc.Call
	valRecv Scalars
	errRecv error
}

func (c *implServiceAMethodA3Call) RecvStream() interface {
	Advance() bool
	Value() Scalars
	Err() error
} {
	return implServiceAMethodA3CallRecv{c}
}

type implServiceAMethodA3CallRecv struct {
	c *implServiceAMethodA3Call
}

func (c implServiceAMethodA3CallRecv) Advance() bool {
	c.c.valRecv = Scalars{}
	c.c.errRecv = c.c.Recv(&c.c.valRecv)
	return c.c.errRecv == nil
}
func (c implServiceAMethodA3CallRecv) Value() Scalars {
	return c.c.valRecv
}
func (c implServiceAMethodA3CallRecv) Err() error {
	if c.c.errRecv == io.EOF {
		return nil
	}
	return c.c.errRecv
}
func (c *implServiceAMethodA3Call) Finish() (o0 string, err error) {
	if ierr := c.Call.Finish(&o0, &err); ierr != nil {
		err = ierr
	}
	return
}

// ServiceAMethodA4ClientStream is the client stream for ServiceA.MethodA4.
type ServiceAMethodA4ClientStream interface {
	// RecvStream returns the receiver side of the ServiceA.MethodA4 client stream.
	RecvStream() interface {
		// Advance stages an item so that it may be retrieved via Value.  Returns
		// true iff there is an item to retrieve.  Advance must be called before
		// Value is called.  May block if an item is not available.
		Advance() bool
		// Value returns the item that was staged by Advance.  May panic if Advance
		// returned false or was not called.  Never blocks.
		Value() string
		// Err returns any error encountered by Advance.  Never blocks.
		Err() error
	}
	// SendStream returns the send side of the ServiceA.MethodA4 client stream.
	SendStream() interface {
		// Send places the item onto the output stream.  Returns errors
		// encountered while sending, or if Send is called after Close or
		// the stream has been canceled.  Blocks if there is no buffer
		// space; will unblock when buffer space is available or after
		// the stream has been canceled.
		Send(item int32) error
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

// ServiceAMethodA4Call represents the call returned from ServiceA.MethodA4.
type ServiceAMethodA4Call interface {
	ServiceAMethodA4ClientStream
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
	Finish() error
}

type implServiceAMethodA4Call struct {
	ipc.Call
	valRecv string
	errRecv error
}

func (c *implServiceAMethodA4Call) RecvStream() interface {
	Advance() bool
	Value() string
	Err() error
} {
	return implServiceAMethodA4CallRecv{c}
}

type implServiceAMethodA4CallRecv struct {
	c *implServiceAMethodA4Call
}

func (c implServiceAMethodA4CallRecv) Advance() bool {
	c.c.errRecv = c.c.Recv(&c.c.valRecv)
	return c.c.errRecv == nil
}
func (c implServiceAMethodA4CallRecv) Value() string {
	return c.c.valRecv
}
func (c implServiceAMethodA4CallRecv) Err() error {
	if c.c.errRecv == io.EOF {
		return nil
	}
	return c.c.errRecv
}
func (c *implServiceAMethodA4Call) SendStream() interface {
	Send(item int32) error
	Close() error
} {
	return implServiceAMethodA4CallSend{c}
}

type implServiceAMethodA4CallSend struct {
	c *implServiceAMethodA4Call
}

func (c implServiceAMethodA4CallSend) Send(item int32) error {
	return c.c.Send(item)
}
func (c implServiceAMethodA4CallSend) Close() error {
	return c.c.CloseSend()
}
func (c *implServiceAMethodA4Call) Finish() (err error) {
	if ierr := c.Call.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

// ServiceAServerMethods is the interface a server writer
// implements for ServiceA.
type ServiceAServerMethods interface {
	MethodA1(ipc.ServerContext) error
	MethodA2(ctx ipc.ServerContext, a int32, b string) (s string, err error)
	MethodA3(ctx ServiceAMethodA3Context, a int32) (s string, err error)
	MethodA4(ctx ServiceAMethodA4Context, a int32) error
}

// ServiceAServerStubMethods is the server interface containing
// ServiceA methods, as expected by ipc.Server.
// The only difference between this interface and ServiceAServerMethods
// is the streaming methods.
type ServiceAServerStubMethods interface {
	MethodA1(ipc.ServerContext) error
	MethodA2(ctx ipc.ServerContext, a int32, b string) (s string, err error)
	MethodA3(ctx *ServiceAMethodA3ContextStub, a int32) (s string, err error)
	MethodA4(ctx *ServiceAMethodA4ContextStub, a int32) error
}

// ServiceAServerStub adds universal methods to ServiceAServerStubMethods.
type ServiceAServerStub interface {
	ServiceAServerStubMethods
	// Describe the ServiceA interfaces.
	Describe__() []ipc.InterfaceDesc
}

// ServiceAServer returns a server stub for ServiceA.
// It converts an implementation of ServiceAServerMethods into
// an object that may be used by ipc.Server.
func ServiceAServer(impl ServiceAServerMethods) ServiceAServerStub {
	stub := implServiceAServerStub{
		impl: impl,
	}
	// Initialize GlobState; always check the stub itself first, to handle the
	// case where the user has the Glob method defined in their VDL source.
	if gs := ipc.NewGlobState(stub); gs != nil {
		stub.gs = gs
	} else if gs := ipc.NewGlobState(impl); gs != nil {
		stub.gs = gs
	}
	return stub
}

type implServiceAServerStub struct {
	impl ServiceAServerMethods
	gs   *ipc.GlobState
}

func (s implServiceAServerStub) MethodA1(ctx ipc.ServerContext) error {
	return s.impl.MethodA1(ctx)
}

func (s implServiceAServerStub) MethodA2(ctx ipc.ServerContext, i0 int32, i1 string) (string, error) {
	return s.impl.MethodA2(ctx, i0, i1)
}

func (s implServiceAServerStub) MethodA3(ctx *ServiceAMethodA3ContextStub, i0 int32) (string, error) {
	return s.impl.MethodA3(ctx, i0)
}

func (s implServiceAServerStub) MethodA4(ctx *ServiceAMethodA4ContextStub, i0 int32) error {
	return s.impl.MethodA4(ctx, i0)
}

func (s implServiceAServerStub) Globber() *ipc.GlobState {
	return s.gs
}

func (s implServiceAServerStub) Describe__() []ipc.InterfaceDesc {
	return []ipc.InterfaceDesc{ServiceADesc}
}

// ServiceADesc describes the ServiceA interface.
var ServiceADesc ipc.InterfaceDesc = descServiceA

// descServiceA hides the desc to keep godoc clean.
var descServiceA = ipc.InterfaceDesc{
	Name:    "ServiceA",
	PkgPath: "v.io/core/veyron2/vdl/testdata/base",
	Methods: []ipc.MethodDesc{
		{
			Name: "MethodA1",
			OutArgs: []ipc.ArgDesc{
				{"", ``}, // error
			},
		},
		{
			Name: "MethodA2",
			InArgs: []ipc.ArgDesc{
				{"a", ``}, // int32
				{"b", ``}, // string
			},
			OutArgs: []ipc.ArgDesc{
				{"s", ``},   // string
				{"err", ``}, // error
			},
		},
		{
			Name: "MethodA3",
			InArgs: []ipc.ArgDesc{
				{"a", ``}, // int32
			},
			OutArgs: []ipc.ArgDesc{
				{"s", ``},   // string
				{"err", ``}, // error
			},
			Tags: []vdl.AnyRep{"tag", uint64(6)},
		},
		{
			Name: "MethodA4",
			InArgs: []ipc.ArgDesc{
				{"a", ``}, // int32
			},
			OutArgs: []ipc.ArgDesc{
				{"", ``}, // error
			},
		},
	},
}

// ServiceAMethodA3ServerStream is the server stream for ServiceA.MethodA3.
type ServiceAMethodA3ServerStream interface {
	// SendStream returns the send side of the ServiceA.MethodA3 server stream.
	SendStream() interface {
		// Send places the item onto the output stream.  Returns errors encountered
		// while sending.  Blocks if there is no buffer space; will unblock when
		// buffer space is available.
		Send(item Scalars) error
	}
}

// ServiceAMethodA3Context represents the context passed to ServiceA.MethodA3.
type ServiceAMethodA3Context interface {
	ipc.ServerContext
	ServiceAMethodA3ServerStream
}

// ServiceAMethodA3ContextStub is a wrapper that converts ipc.ServerCall into
// a typesafe stub that implements ServiceAMethodA3Context.
type ServiceAMethodA3ContextStub struct {
	ipc.ServerCall
}

// Init initializes ServiceAMethodA3ContextStub from ipc.ServerCall.
func (s *ServiceAMethodA3ContextStub) Init(call ipc.ServerCall) {
	s.ServerCall = call
}

// SendStream returns the send side of the ServiceA.MethodA3 server stream.
func (s *ServiceAMethodA3ContextStub) SendStream() interface {
	Send(item Scalars) error
} {
	return implServiceAMethodA3ContextSend{s}
}

type implServiceAMethodA3ContextSend struct {
	s *ServiceAMethodA3ContextStub
}

func (s implServiceAMethodA3ContextSend) Send(item Scalars) error {
	return s.s.Send(item)
}

// ServiceAMethodA4ServerStream is the server stream for ServiceA.MethodA4.
type ServiceAMethodA4ServerStream interface {
	// RecvStream returns the receiver side of the ServiceA.MethodA4 server stream.
	RecvStream() interface {
		// Advance stages an item so that it may be retrieved via Value.  Returns
		// true iff there is an item to retrieve.  Advance must be called before
		// Value is called.  May block if an item is not available.
		Advance() bool
		// Value returns the item that was staged by Advance.  May panic if Advance
		// returned false or was not called.  Never blocks.
		Value() int32
		// Err returns any error encountered by Advance.  Never blocks.
		Err() error
	}
	// SendStream returns the send side of the ServiceA.MethodA4 server stream.
	SendStream() interface {
		// Send places the item onto the output stream.  Returns errors encountered
		// while sending.  Blocks if there is no buffer space; will unblock when
		// buffer space is available.
		Send(item string) error
	}
}

// ServiceAMethodA4Context represents the context passed to ServiceA.MethodA4.
type ServiceAMethodA4Context interface {
	ipc.ServerContext
	ServiceAMethodA4ServerStream
}

// ServiceAMethodA4ContextStub is a wrapper that converts ipc.ServerCall into
// a typesafe stub that implements ServiceAMethodA4Context.
type ServiceAMethodA4ContextStub struct {
	ipc.ServerCall
	valRecv int32
	errRecv error
}

// Init initializes ServiceAMethodA4ContextStub from ipc.ServerCall.
func (s *ServiceAMethodA4ContextStub) Init(call ipc.ServerCall) {
	s.ServerCall = call
}

// RecvStream returns the receiver side of the ServiceA.MethodA4 server stream.
func (s *ServiceAMethodA4ContextStub) RecvStream() interface {
	Advance() bool
	Value() int32
	Err() error
} {
	return implServiceAMethodA4ContextRecv{s}
}

type implServiceAMethodA4ContextRecv struct {
	s *ServiceAMethodA4ContextStub
}

func (s implServiceAMethodA4ContextRecv) Advance() bool {
	s.s.errRecv = s.s.Recv(&s.s.valRecv)
	return s.s.errRecv == nil
}
func (s implServiceAMethodA4ContextRecv) Value() int32 {
	return s.s.valRecv
}
func (s implServiceAMethodA4ContextRecv) Err() error {
	if s.s.errRecv == io.EOF {
		return nil
	}
	return s.s.errRecv
}

// SendStream returns the send side of the ServiceA.MethodA4 server stream.
func (s *ServiceAMethodA4ContextStub) SendStream() interface {
	Send(item string) error
} {
	return implServiceAMethodA4ContextSend{s}
}

type implServiceAMethodA4ContextSend struct {
	s *ServiceAMethodA4ContextStub
}

func (s implServiceAMethodA4ContextSend) Send(item string) error {
	return s.s.Send(item)
}

// ServiceBClientMethods is the client interface
// containing ServiceB methods.
type ServiceBClientMethods interface {
	ServiceAClientMethods
	MethodB1(ctx *context.T, a Scalars, b Composites, opts ...ipc.CallOpt) (c CompComp, err error)
}

// ServiceBClientStub adds universal methods to ServiceBClientMethods.
type ServiceBClientStub interface {
	ServiceBClientMethods
	ipc.UniversalServiceMethods
}

// ServiceBClient returns a client stub for ServiceB.
func ServiceBClient(name string, opts ...ipc.BindOpt) ServiceBClientStub {
	var client ipc.Client
	for _, opt := range opts {
		if clientOpt, ok := opt.(ipc.Client); ok {
			client = clientOpt
		}
	}
	return implServiceBClientStub{name, client, ServiceAClient(name, client)}
}

type implServiceBClientStub struct {
	name   string
	client ipc.Client

	ServiceAClientStub
}

func (c implServiceBClientStub) c(ctx *context.T) ipc.Client {
	if c.client != nil {
		return c.client
	}
	return veyron2.GetClient(ctx)
}

func (c implServiceBClientStub) MethodB1(ctx *context.T, i0 Scalars, i1 Composites, opts ...ipc.CallOpt) (o0 CompComp, err error) {
	var call ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "MethodB1", []interface{}{i0, i1}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&o0, &err); ierr != nil {
		err = ierr
	}
	return
}

// ServiceBServerMethods is the interface a server writer
// implements for ServiceB.
type ServiceBServerMethods interface {
	ServiceAServerMethods
	MethodB1(ctx ipc.ServerContext, a Scalars, b Composites) (c CompComp, err error)
}

// ServiceBServerStubMethods is the server interface containing
// ServiceB methods, as expected by ipc.Server.
// The only difference between this interface and ServiceBServerMethods
// is the streaming methods.
type ServiceBServerStubMethods interface {
	ServiceAServerStubMethods
	MethodB1(ctx ipc.ServerContext, a Scalars, b Composites) (c CompComp, err error)
}

// ServiceBServerStub adds universal methods to ServiceBServerStubMethods.
type ServiceBServerStub interface {
	ServiceBServerStubMethods
	// Describe the ServiceB interfaces.
	Describe__() []ipc.InterfaceDesc
}

// ServiceBServer returns a server stub for ServiceB.
// It converts an implementation of ServiceBServerMethods into
// an object that may be used by ipc.Server.
func ServiceBServer(impl ServiceBServerMethods) ServiceBServerStub {
	stub := implServiceBServerStub{
		impl:               impl,
		ServiceAServerStub: ServiceAServer(impl),
	}
	// Initialize GlobState; always check the stub itself first, to handle the
	// case where the user has the Glob method defined in their VDL source.
	if gs := ipc.NewGlobState(stub); gs != nil {
		stub.gs = gs
	} else if gs := ipc.NewGlobState(impl); gs != nil {
		stub.gs = gs
	}
	return stub
}

type implServiceBServerStub struct {
	impl ServiceBServerMethods
	ServiceAServerStub
	gs *ipc.GlobState
}

func (s implServiceBServerStub) MethodB1(ctx ipc.ServerContext, i0 Scalars, i1 Composites) (CompComp, error) {
	return s.impl.MethodB1(ctx, i0, i1)
}

func (s implServiceBServerStub) Globber() *ipc.GlobState {
	return s.gs
}

func (s implServiceBServerStub) Describe__() []ipc.InterfaceDesc {
	return []ipc.InterfaceDesc{ServiceBDesc, ServiceADesc}
}

// ServiceBDesc describes the ServiceB interface.
var ServiceBDesc ipc.InterfaceDesc = descServiceB

// descServiceB hides the desc to keep godoc clean.
var descServiceB = ipc.InterfaceDesc{
	Name:    "ServiceB",
	PkgPath: "v.io/core/veyron2/vdl/testdata/base",
	Embeds: []ipc.EmbedDesc{
		{"ServiceA", "v.io/core/veyron2/vdl/testdata/base", ``},
	},
	Methods: []ipc.MethodDesc{
		{
			Name: "MethodB1",
			InArgs: []ipc.ArgDesc{
				{"a", ``}, // Scalars
				{"b", ``}, // Composites
			},
			OutArgs: []ipc.ArgDesc{
				{"c", ``},   // CompComp
				{"err", ``}, // error
			},
		},
	},
}
