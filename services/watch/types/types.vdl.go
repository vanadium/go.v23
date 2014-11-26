// This file was auto-generated by the veyron vdl tool.
// Source: types.vdl

// Package types defines the types used by the watch interface.
package types

import (
	// The non-user imports are prefixed with "__" to prevent collisions.
	__vdl "veyron.io/veyron/veyron2/vdl"
	__vdlutil "veyron.io/veyron/veyron2/vdl/vdlutil"
)

// GlobRequest specifies which entities should be watched and, optionally,
// how to resume from a previous Watch call.
type GlobRequest struct {
	// Pattern specifies the subset of the children of the root entity
	// for which the client wants updates.
	Pattern string
	// ResumeMarker specifies how to resume from a previous Watch call.
	// See the ResumeMarker type for detailed comments.
	ResumeMarker ResumeMarker
}

func (GlobRequest) __VDLReflect(struct {
	Name string "veyron.io/veyron/veyron2/services/watch/types.GlobRequest"
}) {
}

// ResumeMarker specifies how much of the existing underlying state
// is delivered to the client when the watch request is received by
// the system. The client can set this marker in one of the
// following ways to get different semantics:
//
// (A) Parameter is left empty.
//     Semantics: Fetch initial state.
//     The client wants the entities' initial states to be delivered.
//     See the description in "Initial State".
//
// (B) Parameter is set to the string "now" (UTF-8 encoding).
//     Semantics: Fetch new changes only.
//     The client just wants to get the changes received by the
//     system after the watch point. The system may deliver changes
//     from before the watch point as well.
//
// (C) Parameter is set to a value received in an earlier
//     Change.ResumeMarker field while watching the same entity with
//     the same query.
//     Semantics: Resume from a specific point.
//     The client wants to receive the changes from a specific point
//     - this value must correspond to a value received in the
//     Change.ResumeMarker field. The system may deliver changes
//     from before the ResumeMarker as well.  If the system cannot
//     resume the stream from this point (e.g., if it is too far
//     behind in the stream), it can return the
//     ErrUnknownResumeMarker error.
//
// An implementation MUST support the empty string "" marker
// (initial state fetching) and the "now" marker. It need not
// support resuming from a specific point.
type ResumeMarker []byte

func (ResumeMarker) __VDLReflect(struct {
	Name string "veyron.io/veyron/veyron2/services/watch/types.ResumeMarker"
}) {
}

// Change is the new value for a watched entity.
type Change struct {
	// Name is the Object name of the entity that changed.  This name is relative
	// to the root entity (i.e. the name of the Watcher service).
	Name string
	// State must be one of Exists, DoesNotExist, or InitialStateSkipped.
	State int32
	// Value contains the service-specific data for the entity.
	Value __vdlutil.Any
	// If present, provides a compact representation of all the messages
	// that have been received by the caller for the given Watch call.
	// For example, it could be a sequence number or a multi-part
	// timestamp/version vector. This marker can be provided in the
	// Request message to allow the caller to resume the stream watching
	// at a specific point without fetching the initial state.
	ResumeMarker ResumeMarker
	// If true, this Change is followed by more Changes that are in the
	// same group as this Change.
	Continued bool
}

func (Change) __VDLReflect(struct {
	Name string "veyron.io/veyron/veyron2/services/watch/types.Change"
}) {
}

func init() {
	__vdl.Register(GlobRequest{})
	__vdl.Register(ResumeMarker(""))
	__vdl.Register(Change{})
}

// The entity exists and its full value is included in Value.
const Exists = int32(0)

// The entity does not exist.
const DoesNotExist = int32(1)

// The root entity and its children may or may not exist. Used only
// for initial state delivery when the client is not interested in
// fetching the initial state. See the "Initial State" section
// above.
const InitialStateSkipped = int32(2)
