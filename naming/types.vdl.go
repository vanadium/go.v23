// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Source: types.vdl

package naming

import (
	// VDL system imports
	"v.io/v23/vdl"

	// VDL user imports
	"v.io/v23/vdlroot/time"
)

// MountFlag is a bit mask of options to the mount call.
type MountFlag uint32

func (MountFlag) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/naming.MountFlag"`
}) {
}

// MountedServer represents a server mounted on a specific name.
type MountedServer struct {
	// Server is the OA that's mounted.
	Server string
	// Deadline before the mount entry expires.
	Deadline time.Deadline
}

func (MountedServer) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/naming.MountedServer"`
}) {
}

// MountEntry represents a given name mounted in the mounttable.
type MountEntry struct {
	// Name is the mounted name.
	Name string
	// Servers (if present) specifies the mounted names.
	Servers []MountedServer
	// ServesMountTable is true if the servers represent mount tables.
	ServesMountTable bool
	// IsLeaf is true if this entry represents a leaf object.
	IsLeaf bool
}

func (MountEntry) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/naming.MountEntry"`
}) {
}

// GlobError is returned by namespace.Glob to indicate a subtree of the namespace
// that could not be traversed.
type GlobError struct {
	// Root of the subtree.
	Name string
	// The error that occurred fulfilling the request.
	Error error
}

func (GlobError) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/naming.GlobError"`
}) {
}

type (
	// GlobReply represents any single field of the GlobReply union type.
	//
	// GlobReply is the data type returned by Glob__.
	GlobReply interface {
		// Index returns the field index.
		Index() int
		// Interface returns the field value as an interface.
		Interface() interface{}
		// Name returns the field name.
		Name() string
		// __VDLReflect describes the GlobReply union type.
		__VDLReflect(__GlobReplyReflect)
	}
	// GlobReplyEntry represents field Entry of the GlobReply union type.
	GlobReplyEntry struct{ Value MountEntry }
	// GlobReplyError represents field Error of the GlobReply union type.
	GlobReplyError struct{ Value GlobError }
	// __GlobReplyReflect describes the GlobReply union type.
	__GlobReplyReflect struct {
		Name  string `vdl:"v.io/v23/naming.GlobReply"`
		Type  GlobReply
		Union struct {
			Entry GlobReplyEntry
			Error GlobReplyError
		}
	}
)

func (x GlobReplyEntry) Index() int                      { return 0 }
func (x GlobReplyEntry) Interface() interface{}          { return x.Value }
func (x GlobReplyEntry) Name() string                    { return "Entry" }
func (x GlobReplyEntry) __VDLReflect(__GlobReplyReflect) {}

func (x GlobReplyError) Index() int                      { return 1 }
func (x GlobReplyError) Interface() interface{}          { return x.Value }
func (x GlobReplyError) Name() string                    { return "Error" }
func (x GlobReplyError) __VDLReflect(__GlobReplyReflect) {}

type (
	// GlobChildrenReply represents any single field of the GlobChildrenReply union type.
	//
	// GlobChildrenReply is the data type returned by GlobChildren__.
	GlobChildrenReply interface {
		// Index returns the field index.
		Index() int
		// Interface returns the field value as an interface.
		Interface() interface{}
		// Name returns the field name.
		Name() string
		// __VDLReflect describes the GlobChildrenReply union type.
		__VDLReflect(__GlobChildrenReplyReflect)
	}
	// GlobChildrenReplyName represents field Name of the GlobChildrenReply union type.
	GlobChildrenReplyName struct{ Value string }
	// GlobChildrenReplyError represents field Error of the GlobChildrenReply union type.
	GlobChildrenReplyError struct{ Value GlobError }
	// __GlobChildrenReplyReflect describes the GlobChildrenReply union type.
	__GlobChildrenReplyReflect struct {
		Name  string `vdl:"v.io/v23/naming.GlobChildrenReply"`
		Type  GlobChildrenReply
		Union struct {
			Name  GlobChildrenReplyName
			Error GlobChildrenReplyError
		}
	}
)

func (x GlobChildrenReplyName) Index() int                              { return 0 }
func (x GlobChildrenReplyName) Interface() interface{}                  { return x.Value }
func (x GlobChildrenReplyName) Name() string                            { return "Name" }
func (x GlobChildrenReplyName) __VDLReflect(__GlobChildrenReplyReflect) {}

func (x GlobChildrenReplyError) Index() int                              { return 1 }
func (x GlobChildrenReplyError) Interface() interface{}                  { return x.Value }
func (x GlobChildrenReplyError) Name() string                            { return "Error" }
func (x GlobChildrenReplyError) __VDLReflect(__GlobChildrenReplyReflect) {}

func init() {
	vdl.Register((*MountFlag)(nil))
	vdl.Register((*MountedServer)(nil))
	vdl.Register((*MountEntry)(nil))
	vdl.Register((*GlobError)(nil))
	vdl.Register((*GlobReply)(nil))
	vdl.Register((*GlobChildrenReply)(nil))
}

const Replace = MountFlag(1) // Replace means the mount should replace what is currently at the mount point

const MT = MountFlag(2) // MT means that the target server is a mount table.

const Leaf = MountFlag(4) // Leaf means that the target server is a leaf.
