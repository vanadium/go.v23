// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Package: message

package message

import (
	"v.io/v23/context"
	"v.io/v23/i18n"
	"v.io/v23/verror"
)

var _ = __VDLInit() // Must be first; see __VDLInit comments for details.

//////////////////////////////////////////////////
// Error definitions

var (
	ErrInvalidMsg         = verror.Register("v.io/v23/flow/message.InvalidMsg", verror.NoRetry, "{1:}{2:} message of type {3} and size {4} failed decoding at field {5}{:6}.")
	ErrInvalidSetupOption = verror.Register("v.io/v23/flow/message.InvalidSetupOption", verror.NoRetry, "{1:}{2:} setup option{:3} failed decoding at field{:4}.")
	ErrUnknownMsg         = verror.Register("v.io/v23/flow/message.UnknownMsg", verror.NoRetry, "{1:}{2:} unknown message type{:3}.")
	ErrMissingBlessings   = verror.Register("v.io/v23/flow/message.MissingBlessings", verror.NoRetry, "{1:}{2:} {3} message received with no blessings.")
)

// NewErrInvalidMsg returns an error with the ErrInvalidMsg ID.
func NewErrInvalidMsg(ctx *context.T, typ byte, size uint64, field uint64, err error) error {
	return verror.New(ErrInvalidMsg, ctx, typ, size, field, err)
}

// NewErrInvalidSetupOption returns an error with the ErrInvalidSetupOption ID.
func NewErrInvalidSetupOption(ctx *context.T, option uint64, field uint64) error {
	return verror.New(ErrInvalidSetupOption, ctx, option, field)
}

// NewErrUnknownMsg returns an error with the ErrUnknownMsg ID.
func NewErrUnknownMsg(ctx *context.T, typ byte) error {
	return verror.New(ErrUnknownMsg, ctx, typ)
}

// NewErrMissingBlessings returns an error with the ErrMissingBlessings ID.
func NewErrMissingBlessings(ctx *context.T, typ byte) error {
	return verror.New(ErrMissingBlessings, ctx, typ)
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

	// Set error format strings.
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrInvalidMsg.ID), "{1:}{2:} message of type {3} and size {4} failed decoding at field {5}{:6}.")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrInvalidSetupOption.ID), "{1:}{2:} setup option{:3} failed decoding at field{:4}.")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrUnknownMsg.ID), "{1:}{2:} unknown message type{:3}.")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrMissingBlessings.ID), "{1:}{2:} {3} message received with no blessings.")

	return struct{}{}
}
