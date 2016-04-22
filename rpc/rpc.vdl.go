// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Package: rpc

package rpc

import (
	"fmt"
	"v.io/v23/security"
	"v.io/v23/uniqueid"
	"v.io/v23/vdl"
	"v.io/v23/vdlroot/time"
	"v.io/v23/verror"
	"v.io/v23/vtrace"
)

var _ = __VDLInit() // Must be first; see __VDLInit comments for details.

//////////////////////////////////////////////////
// Type definitions

// Request describes the request header sent by the client to the server.  A
// non-zero request header is sent at the beginning of the RPC call, followed by
// the positional args.  Thereafter a zero request header is sent before each
// streaming arg, terminated by a non-zero request header with EndStreamArgs set
// to true.
type Request struct {
	// Suffix of the name used to identify the object hosting the service.
	Suffix string
	// Method to invoke on the service.
	Method string
	// NumPosArgs is the number of positional arguments, which follow this message
	// (and any blessings) on the request stream.
	NumPosArgs uint64
	// EndStreamArgs is true iff no more streaming arguments will be sent.  No
	// more data will be sent on the request stream.
	//
	// NOTE(bprosnitz): We can support multiple stream values per request (+response) header
	// efficiently by adding a NumExtraStreamArgs (+NumExtraStreamResults to response) field
	// that is the uint64 (number of stream args to send) - 1. The request is then zero when
	// exactly one streaming arg is sent. Since the request and response headers are small,
	// this is only likely necessary for frequently streaming small values.
	// See implementation in CL: 3913
	EndStreamArgs bool
	// Deadline after which the request should be cancelled.  This is a hint to
	// the server, to avoid wasted work.
	Deadline time.Deadline
	// GrantedBlessings are blessings bound to the principal running the server,
	// provided by the client.
	GrantedBlessings security.Blessings
	// TraceRequest maintains the vtrace context between clients and servers
	// and specifies additional parameters that control how tracing behaves.
	TraceRequest vtrace.Request
	// Language indicates the language of the instegator of the RPC.
	// By convention it should be an IETF language tag:
	// http://en.wikipedia.org/wiki/IETF_language_tag
	Language string
}

func (Request) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/rpc.Request"`
}) {
}

func (m *Request) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}
	var4 := (m.Suffix == "")
	if var4 {
		if err := fieldsTarget1.ZeroField("Suffix"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Suffix")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}
			if err := fieldTarget3.FromString(string(m.Suffix), tt.NonOptional().Field(0).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
				return err
			}
		}
	}
	var7 := (m.Method == "")
	if var7 {
		if err := fieldsTarget1.ZeroField("Method"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget5, fieldTarget6, err := fieldsTarget1.StartField("Method")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}
			if err := fieldTarget6.FromString(string(m.Method), tt.NonOptional().Field(1).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget5, fieldTarget6); err != nil {
				return err
			}
		}
	}
	var10 := (m.NumPosArgs == uint64(0))
	if var10 {
		if err := fieldsTarget1.ZeroField("NumPosArgs"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget8, fieldTarget9, err := fieldsTarget1.StartField("NumPosArgs")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}
			if err := fieldTarget9.FromUint(uint64(m.NumPosArgs), tt.NonOptional().Field(2).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget8, fieldTarget9); err != nil {
				return err
			}
		}
	}
	var13 := (m.EndStreamArgs == false)
	if var13 {
		if err := fieldsTarget1.ZeroField("EndStreamArgs"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget11, fieldTarget12, err := fieldsTarget1.StartField("EndStreamArgs")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}
			if err := fieldTarget12.FromBool(bool(m.EndStreamArgs), tt.NonOptional().Field(3).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget11, fieldTarget12); err != nil {
				return err
			}
		}
	}
	var wireValue14 time.WireDeadline
	if err := time.WireDeadlineFromNative(&wireValue14, m.Deadline); err != nil {
		return err
	}

	var17 := (wireValue14 == time.WireDeadline{})
	if var17 {
		if err := fieldsTarget1.ZeroField("Deadline"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget15, fieldTarget16, err := fieldsTarget1.StartField("Deadline")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}

			if err := wireValue14.FillVDLTarget(fieldTarget16, tt.NonOptional().Field(4).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget15, fieldTarget16); err != nil {
				return err
			}
		}
	}
	var wireValue18 security.WireBlessings
	if err := security.WireBlessingsFromNative(&wireValue18, m.GrantedBlessings); err != nil {
		return err
	}

	var21 := true
	var var22 bool
	if len(wireValue18.CertificateChains) == 0 {
		var22 = true
	}
	var21 = var21 && var22
	if var21 {
		if err := fieldsTarget1.ZeroField("GrantedBlessings"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget19, fieldTarget20, err := fieldsTarget1.StartField("GrantedBlessings")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}

			if err := wireValue18.FillVDLTarget(fieldTarget20, tt.NonOptional().Field(5).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget19, fieldTarget20); err != nil {
				return err
			}
		}
	}
	var25 := (m.TraceRequest == vtrace.Request{})
	if var25 {
		if err := fieldsTarget1.ZeroField("TraceRequest"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget23, fieldTarget24, err := fieldsTarget1.StartField("TraceRequest")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}

			if err := m.TraceRequest.FillVDLTarget(fieldTarget24, tt.NonOptional().Field(6).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget23, fieldTarget24); err != nil {
				return err
			}
		}
	}
	var28 := (m.Language == "")
	if var28 {
		if err := fieldsTarget1.ZeroField("Language"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget26, fieldTarget27, err := fieldsTarget1.StartField("Language")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}
			if err := fieldTarget27.FromString(string(m.Language), tt.NonOptional().Field(7).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget26, fieldTarget27); err != nil {
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
	Value                  *Request
	suffixTarget           vdl.StringTarget
	methodTarget           vdl.StringTarget
	numPosArgsTarget       vdl.Uint64Target
	endStreamArgsTarget    vdl.BoolTarget
	deadlineTarget         time.WireDeadlineTarget
	grantedBlessingsTarget security.WireBlessingsTarget
	traceRequestTarget     vtrace.RequestTarget
	languageTarget         vdl.StringTarget
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
	case "Suffix":
		t.suffixTarget.Value = &t.Value.Suffix
		target, err := &t.suffixTarget, error(nil)
		return nil, target, err
	case "Method":
		t.methodTarget.Value = &t.Value.Method
		target, err := &t.methodTarget, error(nil)
		return nil, target, err
	case "NumPosArgs":
		t.numPosArgsTarget.Value = &t.Value.NumPosArgs
		target, err := &t.numPosArgsTarget, error(nil)
		return nil, target, err
	case "EndStreamArgs":
		t.endStreamArgsTarget.Value = &t.Value.EndStreamArgs
		target, err := &t.endStreamArgsTarget, error(nil)
		return nil, target, err
	case "Deadline":
		t.deadlineTarget.Value = &t.Value.Deadline
		target, err := &t.deadlineTarget, error(nil)
		return nil, target, err
	case "GrantedBlessings":
		t.grantedBlessingsTarget.Value = &t.Value.GrantedBlessings
		target, err := &t.grantedBlessingsTarget, error(nil)
		return nil, target, err
	case "TraceRequest":
		t.traceRequestTarget.Value = &t.Value.TraceRequest
		target, err := &t.traceRequestTarget, error(nil)
		return nil, target, err
	case "Language":
		t.languageTarget.Value = &t.Value.Language
		target, err := &t.languageTarget, error(nil)
		return nil, target, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct v.io/v23/rpc.Request", name)
	}
}
func (t *RequestTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *RequestTarget) ZeroField(name string) error {
	switch name {
	case "Suffix":
		t.Value.Suffix = ""
		return nil
	case "Method":
		t.Value.Method = ""
		return nil
	case "NumPosArgs":
		t.Value.NumPosArgs = uint64(0)
		return nil
	case "EndStreamArgs":
		t.Value.EndStreamArgs = false
		return nil
	case "Deadline":
		t.Value.Deadline = func() time.Deadline {
			var native time.Deadline
			if err := vdl.Convert(&native, time.WireDeadline{}); err != nil {
				panic(err)
			}
			return native
		}()
		return nil
	case "GrantedBlessings":
		t.Value.GrantedBlessings = func() security.Blessings {
			var native security.Blessings
			if err := vdl.Convert(&native, security.WireBlessings{}); err != nil {
				panic(err)
			}
			return native
		}()
		return nil
	case "TraceRequest":
		t.Value.TraceRequest = vtrace.Request{}
		return nil
	case "Language":
		t.Value.Language = ""
		return nil
	default:
		return fmt.Errorf("field %s not in struct v.io/v23/rpc.Request", name)
	}
}
func (t *RequestTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

func (x Request) VDLIsZero() (bool, error) {
	if x.Suffix != "" {
		return false, nil
	}
	if x.Method != "" {
		return false, nil
	}
	if x.NumPosArgs != 0 {
		return false, nil
	}
	if x.EndStreamArgs {
		return false, nil
	}
	var wireDeadline time.WireDeadline
	if err := time.WireDeadlineFromNative(&wireDeadline, x.Deadline); err != nil {
		return false, err
	}
	isZeroDeadline, err := wireDeadline.VDLIsZero()
	if err != nil {
		return false, err
	}
	if !isZeroDeadline {
		return false, nil
	}
	var wireGrantedBlessings security.WireBlessings
	if err := security.WireBlessingsFromNative(&wireGrantedBlessings, x.GrantedBlessings); err != nil {
		return false, err
	}
	isZeroGrantedBlessings, err := wireGrantedBlessings.VDLIsZero()
	if err != nil {
		return false, err
	}
	if !isZeroGrantedBlessings {
		return false, nil
	}
	if x.TraceRequest != (vtrace.Request{}) {
		return false, nil
	}
	if x.Language != "" {
		return false, nil
	}
	return true, nil
}

func (x Request) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(vdl.TypeOf((*Request)(nil)).Elem()); err != nil {
		return err
	}
	if x.Suffix != "" {
		if err := enc.NextField("Suffix"); err != nil {
			return err
		}
		if err := enc.StartValue(vdl.TypeOf((*string)(nil))); err != nil {
			return err
		}
		if err := enc.EncodeString(x.Suffix); err != nil {
			return err
		}
		if err := enc.FinishValue(); err != nil {
			return err
		}
	}
	if x.Method != "" {
		if err := enc.NextField("Method"); err != nil {
			return err
		}
		if err := enc.StartValue(vdl.TypeOf((*string)(nil))); err != nil {
			return err
		}
		if err := enc.EncodeString(x.Method); err != nil {
			return err
		}
		if err := enc.FinishValue(); err != nil {
			return err
		}
	}
	if x.NumPosArgs != 0 {
		if err := enc.NextField("NumPosArgs"); err != nil {
			return err
		}
		if err := enc.StartValue(vdl.TypeOf((*uint64)(nil))); err != nil {
			return err
		}
		if err := enc.EncodeUint(x.NumPosArgs); err != nil {
			return err
		}
		if err := enc.FinishValue(); err != nil {
			return err
		}
	}
	if x.EndStreamArgs {
		if err := enc.NextField("EndStreamArgs"); err != nil {
			return err
		}
		if err := enc.StartValue(vdl.TypeOf((*bool)(nil))); err != nil {
			return err
		}
		if err := enc.EncodeBool(x.EndStreamArgs); err != nil {
			return err
		}
		if err := enc.FinishValue(); err != nil {
			return err
		}
	}
	var wireDeadline time.WireDeadline
	if err := time.WireDeadlineFromNative(&wireDeadline, x.Deadline); err != nil {
		return err
	}
	isZeroDeadline, err := wireDeadline.VDLIsZero()
	if err != nil {
		return err
	}
	if !isZeroDeadline {
		if err := enc.NextField("Deadline"); err != nil {
			return err
		}
		if err := wireDeadline.VDLWrite(enc); err != nil {
			return err
		}
	}
	var wireGrantedBlessings security.WireBlessings
	if err := security.WireBlessingsFromNative(&wireGrantedBlessings, x.GrantedBlessings); err != nil {
		return err
	}
	isZeroGrantedBlessings, err := wireGrantedBlessings.VDLIsZero()
	if err != nil {
		return err
	}
	if !isZeroGrantedBlessings {
		if err := enc.NextField("GrantedBlessings"); err != nil {
			return err
		}
		if err := wireGrantedBlessings.VDLWrite(enc); err != nil {
			return err
		}
	}
	if x.TraceRequest != (vtrace.Request{}) {
		if err := enc.NextField("TraceRequest"); err != nil {
			return err
		}
		if err := x.TraceRequest.VDLWrite(enc); err != nil {
			return err
		}
	}
	if x.Language != "" {
		if err := enc.NextField("Language"); err != nil {
			return err
		}
		if err := enc.StartValue(vdl.TypeOf((*string)(nil))); err != nil {
			return err
		}
		if err := enc.EncodeString(x.Language); err != nil {
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
		case "Suffix":
			if err := dec.StartValue(); err != nil {
				return err
			}
			var err error
			if x.Suffix, err = dec.DecodeString(); err != nil {
				return err
			}
			if err := dec.FinishValue(); err != nil {
				return err
			}
		case "Method":
			if err := dec.StartValue(); err != nil {
				return err
			}
			var err error
			if x.Method, err = dec.DecodeString(); err != nil {
				return err
			}
			if err := dec.FinishValue(); err != nil {
				return err
			}
		case "NumPosArgs":
			if err := dec.StartValue(); err != nil {
				return err
			}
			var err error
			if x.NumPosArgs, err = dec.DecodeUint(64); err != nil {
				return err
			}
			if err := dec.FinishValue(); err != nil {
				return err
			}
		case "EndStreamArgs":
			if err := dec.StartValue(); err != nil {
				return err
			}
			var err error
			if x.EndStreamArgs, err = dec.DecodeBool(); err != nil {
				return err
			}
			if err := dec.FinishValue(); err != nil {
				return err
			}
		case "Deadline":
			var wire time.WireDeadline
			if err := wire.VDLRead(dec); err != nil {
				return err
			}
			if err := time.WireDeadlineToNative(wire, &x.Deadline); err != nil {
				return err
			}
		case "GrantedBlessings":
			var wire security.WireBlessings
			if err := wire.VDLRead(dec); err != nil {
				return err
			}
			if err := security.WireBlessingsToNative(wire, &x.GrantedBlessings); err != nil {
				return err
			}
		case "TraceRequest":
			if err := x.TraceRequest.VDLRead(dec); err != nil {
				return err
			}
		case "Language":
			if err := dec.StartValue(); err != nil {
				return err
			}
			var err error
			if x.Language, err = dec.DecodeString(); err != nil {
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

// Response describes the response header sent by the server to the client.  A
// zero response header is sent before each streaming arg.  Thereafter a
// non-zero response header is sent at the end of the RPC call, right before
// the positional results.
type Response struct {
	// Error in processing the RPC at the server. Implies EndStreamResults.
	Error error
	// EndStreamResults is true iff no more streaming results will be sent; the
	// remainder of the stream consists of NumPosResults positional results.
	EndStreamResults bool
	// NumPosResults is the number of positional results, which immediately follow
	// on the response stream.  After these results, no further data will be sent
	// on the response stream.
	NumPosResults uint64
	// TraceResponse maintains the vtrace context between clients and servers.
	// In some cases trace data will be included in this response as well.
	TraceResponse vtrace.Response
	// AckBlessings is true if the server successfully recevied the client's
	// blessings and stored them in the server's blessings cache.
	AckBlessings bool
}

func (Response) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/rpc.Response"`
}) {
}

func (m *Response) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}
	var4 := (m.Error == (error)(nil))
	if var4 {
		if err := fieldsTarget1.ZeroField("Error"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Error")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}

			var wireError5 vdl.WireError
			if err := verror.WireFromNative(&wireError5, m.Error); err != nil {
				return err
			}
			if err := wireError5.FillVDLTarget(fieldTarget3, vdl.ErrorType); err != nil {
				return err
			}

			if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
				return err
			}
		}
	}
	var8 := (m.EndStreamResults == false)
	if var8 {
		if err := fieldsTarget1.ZeroField("EndStreamResults"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget6, fieldTarget7, err := fieldsTarget1.StartField("EndStreamResults")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}
			if err := fieldTarget7.FromBool(bool(m.EndStreamResults), tt.NonOptional().Field(1).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget6, fieldTarget7); err != nil {
				return err
			}
		}
	}
	var11 := (m.NumPosResults == uint64(0))
	if var11 {
		if err := fieldsTarget1.ZeroField("NumPosResults"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget9, fieldTarget10, err := fieldsTarget1.StartField("NumPosResults")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}
			if err := fieldTarget10.FromUint(uint64(m.NumPosResults), tt.NonOptional().Field(2).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget9, fieldTarget10); err != nil {
				return err
			}
		}
	}
	var14 := true
	var15 := (m.TraceResponse.Flags == vtrace.TraceFlags(0))
	var14 = var14 && var15
	var16 := true
	var17 := (m.TraceResponse.Trace.Id == uniqueid.Id{})
	var16 = var16 && var17
	var var18 bool
	if len(m.TraceResponse.Trace.Spans) == 0 {
		var18 = true
	}
	var16 = var16 && var18
	var14 = var14 && var16
	if var14 {
		if err := fieldsTarget1.ZeroField("TraceResponse"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget12, fieldTarget13, err := fieldsTarget1.StartField("TraceResponse")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}

			if err := m.TraceResponse.FillVDLTarget(fieldTarget13, tt.NonOptional().Field(3).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget12, fieldTarget13); err != nil {
				return err
			}
		}
	}
	var21 := (m.AckBlessings == false)
	if var21 {
		if err := fieldsTarget1.ZeroField("AckBlessings"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget19, fieldTarget20, err := fieldsTarget1.StartField("AckBlessings")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}
			if err := fieldTarget20.FromBool(bool(m.AckBlessings), tt.NonOptional().Field(4).Type); err != nil {
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

func (m *Response) MakeVDLTarget() vdl.Target {
	return &ResponseTarget{Value: m}
}

type ResponseTarget struct {
	Value                  *Response
	errorTarget            verror.ErrorTarget
	endStreamResultsTarget vdl.BoolTarget
	numPosResultsTarget    vdl.Uint64Target
	traceResponseTarget    vtrace.ResponseTarget
	ackBlessingsTarget     vdl.BoolTarget
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
	case "Error":
		t.errorTarget.Value = &t.Value.Error
		target, err := &t.errorTarget, error(nil)
		return nil, target, err
	case "EndStreamResults":
		t.endStreamResultsTarget.Value = &t.Value.EndStreamResults
		target, err := &t.endStreamResultsTarget, error(nil)
		return nil, target, err
	case "NumPosResults":
		t.numPosResultsTarget.Value = &t.Value.NumPosResults
		target, err := &t.numPosResultsTarget, error(nil)
		return nil, target, err
	case "TraceResponse":
		t.traceResponseTarget.Value = &t.Value.TraceResponse
		target, err := &t.traceResponseTarget, error(nil)
		return nil, target, err
	case "AckBlessings":
		t.ackBlessingsTarget.Value = &t.Value.AckBlessings
		target, err := &t.ackBlessingsTarget, error(nil)
		return nil, target, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct v.io/v23/rpc.Response", name)
	}
}
func (t *ResponseTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *ResponseTarget) ZeroField(name string) error {
	switch name {
	case "Error":
		t.Value.Error = (error)(nil)
		return nil
	case "EndStreamResults":
		t.Value.EndStreamResults = false
		return nil
	case "NumPosResults":
		t.Value.NumPosResults = uint64(0)
		return nil
	case "TraceResponse":
		t.Value.TraceResponse = vtrace.Response{}
		return nil
	case "AckBlessings":
		t.Value.AckBlessings = false
		return nil
	default:
		return fmt.Errorf("field %s not in struct v.io/v23/rpc.Response", name)
	}
}
func (t *ResponseTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

func (x Response) VDLIsZero() (bool, error) {
	if x.Error != nil {
		return false, nil
	}
	if x.EndStreamResults {
		return false, nil
	}
	if x.NumPosResults != 0 {
		return false, nil
	}
	isZeroTraceResponse, err := x.TraceResponse.VDLIsZero()
	if err != nil {
		return false, err
	}
	if !isZeroTraceResponse {
		return false, nil
	}
	if x.AckBlessings {
		return false, nil
	}
	return true, nil
}

func (x Response) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(vdl.TypeOf((*Response)(nil)).Elem()); err != nil {
		return err
	}
	if x.Error != nil {
		if err := enc.NextField("Error"); err != nil {
			return err
		}
		if err := verror.VDLWrite(enc, x.Error); err != nil {
			return err
		}
	}
	if x.EndStreamResults {
		if err := enc.NextField("EndStreamResults"); err != nil {
			return err
		}
		if err := enc.StartValue(vdl.TypeOf((*bool)(nil))); err != nil {
			return err
		}
		if err := enc.EncodeBool(x.EndStreamResults); err != nil {
			return err
		}
		if err := enc.FinishValue(); err != nil {
			return err
		}
	}
	if x.NumPosResults != 0 {
		if err := enc.NextField("NumPosResults"); err != nil {
			return err
		}
		if err := enc.StartValue(vdl.TypeOf((*uint64)(nil))); err != nil {
			return err
		}
		if err := enc.EncodeUint(x.NumPosResults); err != nil {
			return err
		}
		if err := enc.FinishValue(); err != nil {
			return err
		}
	}
	isZeroTraceResponse, err := x.TraceResponse.VDLIsZero()
	if err != nil {
		return err
	}
	if !isZeroTraceResponse {
		if err := enc.NextField("TraceResponse"); err != nil {
			return err
		}
		if err := x.TraceResponse.VDLWrite(enc); err != nil {
			return err
		}
	}
	if x.AckBlessings {
		if err := enc.NextField("AckBlessings"); err != nil {
			return err
		}
		if err := enc.StartValue(vdl.TypeOf((*bool)(nil))); err != nil {
			return err
		}
		if err := enc.EncodeBool(x.AckBlessings); err != nil {
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
		case "Error":
			if err := verror.VDLRead(dec, &x.Error); err != nil {
				return err
			}
		case "EndStreamResults":
			if err := dec.StartValue(); err != nil {
				return err
			}
			var err error
			if x.EndStreamResults, err = dec.DecodeBool(); err != nil {
				return err
			}
			if err := dec.FinishValue(); err != nil {
				return err
			}
		case "NumPosResults":
			if err := dec.StartValue(); err != nil {
				return err
			}
			var err error
			if x.NumPosResults, err = dec.DecodeUint(64); err != nil {
				return err
			}
			if err := dec.FinishValue(); err != nil {
				return err
			}
		case "TraceResponse":
			if err := x.TraceResponse.VDLRead(dec); err != nil {
				return err
			}
		case "AckBlessings":
			if err := dec.StartValue(); err != nil {
				return err
			}
			var err error
			if x.AckBlessings, err = dec.DecodeBool(); err != nil {
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

// TODO(toddw): Rename GlobMethod to ReservedGlob.
const GlobMethod = "__Glob"
const ReservedSignature = "__Signature"
const ReservedMethodSignature = "__MethodSignature"

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
	vdl.Register((*Request)(nil))
	vdl.Register((*Response)(nil))

	return struct{}{}
}
