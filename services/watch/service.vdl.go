// This file was auto-generated by the veyron vdl tool.
// Source: service.vdl

// Package watch defines an API for watching updates that match a query.
//
// API Overview
//
// Watcher service allows a client to watch for updates that match a
// query.  For each watched query, the client will receive a reliable
// stream of watch events without re-ordering.
//
// The watching is done by starting a streaming RPC. The argument to
// the RPC contains the query. The result stream consists of a
// never-ending sequence of Change messages (until the call fails or
// is cancelled).
//
// Root Entity
//
// The Object name that receives the Watch RPC is called the root
// entity.  The root entity is the parent of all entities that the
// client cares about.  Therefore, the query is confined to children
// of the root entity, and the names in the Change messages are all
// relative to the root entity.
//
// Watch Request
//
// When a client makes a watch request, it can indicate whether it
// wants to receive the initial states of the entities that match the
// query, just new changes to the entities, or resume watching from a
// particular point in a previous watch stream.  On receiving a watch
// request, the server sends one or more messages to the client. The
// first message informs the client that the server has registered the
// client's request; the instant of time when the client receives the
// event is referred to as the client's "watch point" for that query.
//
// Atomic Delivery
//
// The response stream consists of a sequence of Change messages. Each
// Change message contains an optional continued bit
// (default=false). A sub-sequence of Change messages with
// continued=true followed by a Change message with continued=false
// forms an "atomic group". Systems that support multi-entity atomic
// updates may guarantee that all changes resulting from a single
// atomic update are delivered in the same "atomic group". It is up to
// the documentation of a particular system that implements the Watch
// API to document whether or not it supports such grouping. We expect
// that most callers will ignore the notion of atomic delivery and the
// continued bit, i.e., they will just process each Change message as
// it is received.
//
// Batching
//
// Multiple Change messages may be grouped into a single ChangeBatch message to
// reduce message transfer overhead. A single ChangeBatch may contain many
// atomic groups or a single atomic group may be split across many
// ChangeBatches.
//
// Initial State
//
// The first atomic group delivered by a watch call is special. It is
// delivered as soon as possible and contains the initial state of the
// entities being watched.  The client should consider itself caught up
// after processing this first atomic group.  The messages in this first
// atomic group depend on the value of ResumeMarker.
//
//   (1) ResumeMarker is "" or not specified: For every entity P that
//       matches the query and exists, there will be at least one message
//       delivered with entity == P and the last such message will contain
//       the current state of P.  For every entity Q (including the entity
//       itself) that matches the query but does not exist, either no
//       message will be delivered, or the last message for Q will have
//       state == DOES_NOT_EXIST. At least one message for entity="" will
//       be delivered.
//
//   (2) ResumeMarker == "now": there will be exactly one message with
//       entity = "" and state INITIAL_STATE_SKIPPED.  The client cannot
//       assume whether or not the entity exists after receiving this
//       message.
//
//   (3) ResumeMarker has a value R from a preceding watch call on this
//       entity: The same messages as described in (1) will be delivered
//       to the client except that any information implied by messages
//       received on the preceding call up to and including R may not be
//       delivered. The expectation is that the client will start with
//       state it had built up from the preceding watch call, apply the
//       changes received from this call and build an up-to-date view of
//       the entities without having to fetch a potentially large amount
//       of information that has not changed.  Note that some information
//       that had already been delivered by the preceding call might be
//       delivered again.
//
// Ordering and Reliability
//
// The Change messages that apply to a particular element of the
// entity will be delivered eventually in order without loss for the
// duration of the RPC. Note however that if multiple Changes apply to
// the same element, the implementation is free to suppress them and
// deliver just the last one.  The underlying system must provide the
// guarantee that any relevant update received for an entity E after a
// client's watch point for E MUST be delivered to that client.
//
// These tight guarantees allow for the following simplifications in
// the client:
//
//   (1) The client does not need to have a separate polling loop to
//       make up for missed updates.
//
//   (2) The client does not need to manage timestamps/versions
//       manually; the last update delivered corresponds to the
//       eventual state of the entity.
package watch

import (
	"veyron2/query"

	// The non-user imports are prefixed with "_gen_" to prevent collisions.
	_gen_io "io"
	_gen_context "veyron2/context"
	_gen_ipc "veyron2/ipc"
	_gen_naming "veyron2/naming"
	_gen_rt "veyron2/rt"
	_gen_vdlutil "veyron2/vdl/vdlutil"
	_gen_wiretype "veyron2/wiretype"
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

// QueryRequest specifies which entities should be watched and, optionally,
// how to resume from a previous Watch call.
type QueryRequest struct {
	// Query specifies the subset of the children of the root entity
	// for which the client wants updates.
	Query query.Query
	// ResumeMarker specifies how to resume from a previous Watch call.
	// See the ResumeMarker type for detailed comments.
	ResumeMarker ResumeMarker
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
//     ResumeMarkers are received in lexicographical order.
//
// An implementation MUST support the empty string "" marker
// (initial state fetching) and the "now" marker. It need not
// support resuming from a specific point.
type ResumeMarker []byte

// ChangeBatch is a batch of Change messages.
type ChangeBatch struct {
	Changes []Change
}

// Change is the new value for a watched entity.
type Change struct {
	// Name is the Object name of the entity that changed.  This name is relative
	// to the root entity (i.e. the name of the Watcher service).
	Name string
	// State must be one of Exists, DoesNotExist, or InitialStateSkipped.
	State int32
	// Value contains the service-specific data for the entity.
	Value _gen_vdlutil.Any
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

const (
	// The entity exists and its full value is included in Value.
	Exists = int32(0)

	// The entity does not exist.
	DoesNotExist = int32(1)

	// The root entity and its children may or may not exist. Used only
	// for initial state delivery when the client is not interested in
	// fetching the initial state. See the "Initial State" section
	// above.
	InitialStateSkipped = int32(2)
)

// TODO(bprosnitz) Remove this line once signatures are updated to use typevals.
// It corrects a bug where _gen_wiretype is unused in VDL pacakges where only bootstrap types are used on interfaces.
const _ = _gen_wiretype.TypeIDInvalid

// GlobWatcher allows a client to receive updates for changes to objects
// that match a pattern.  See the package comments for details.
// GlobWatcher is the interface the client binds and uses.
// GlobWatcher_ExcludingUniversal is the interface without internal framework-added methods
// to enable embedding without method collisions.  Not to be used directly by clients.
type GlobWatcher_ExcludingUniversal interface {
	// WatchGlob returns a stream of changes that match a pattern.
	WatchGlob(ctx _gen_context.T, Req GlobRequest, opts ..._gen_ipc.CallOpt) (reply GlobWatcherWatchGlobCall, err error)
}
type GlobWatcher interface {
	_gen_ipc.UniversalServiceMethods
	GlobWatcher_ExcludingUniversal
}

// GlobWatcherService is the interface the server implements.
type GlobWatcherService interface {

	// WatchGlob returns a stream of changes that match a pattern.
	WatchGlob(context _gen_ipc.ServerContext, Req GlobRequest, stream GlobWatcherServiceWatchGlobStream) (err error)
}

// GlobWatcherWatchGlobCall is the interface for call object of the method
// WatchGlob in the service interface GlobWatcher.
type GlobWatcherWatchGlobCall interface {
	// RecvStream returns the recv portion of the stream
	RecvStream() interface {
		// Advance stages an element so the client can retrieve it
		// with Value.  Advance returns true iff there is an
		// element to retrieve.  The client must call Advance before
		// calling Value. Advance may block if an element is not
		// immediately available.
		Advance() bool

		// Value returns the element that was staged by Advance.
		// Value may panic if Advance returned false or was not
		// called at all.  Value does not block.
		Value() ChangeBatch

		// Err returns a non-nil error iff the stream encountered
		// any errors.  Err does not block.
		Err() error
	}

	// Finish blocks until the server is done and returns the positional
	// return values for call.
	//
	// If Cancel has been called, Finish will return immediately; the output of
	// Finish could either be an error signalling cancelation, or the correct
	// positional return values from the server depending on the timing of the
	// call.
	//
	// Calling Finish is mandatory for releasing stream resources, unless Cancel
	// has been called or any of the other methods return an error.
	// Finish should be called at most once.
	Finish() (err error)

	// Cancel cancels the RPC, notifying the server to stop processing.  It
	// is safe to call Cancel concurrently with any of the other stream methods.
	// Calling Cancel after Finish has returned is a no-op.
	Cancel()
}

type implGlobWatcherWatchGlobStreamIterator struct {
	clientCall _gen_ipc.Call
	val        ChangeBatch
	err        error
}

func (c *implGlobWatcherWatchGlobStreamIterator) Advance() bool {
	c.val = ChangeBatch{}
	c.err = c.clientCall.Recv(&c.val)
	return c.err == nil
}

func (c *implGlobWatcherWatchGlobStreamIterator) Value() ChangeBatch {
	return c.val
}

func (c *implGlobWatcherWatchGlobStreamIterator) Err() error {
	if c.err == _gen_io.EOF {
		return nil
	}
	return c.err
}

// Implementation of the GlobWatcherWatchGlobCall interface that is not exported.
type implGlobWatcherWatchGlobCall struct {
	clientCall _gen_ipc.Call
	readStream implGlobWatcherWatchGlobStreamIterator
}

func (c *implGlobWatcherWatchGlobCall) RecvStream() interface {
	Advance() bool
	Value() ChangeBatch
	Err() error
} {
	return &c.readStream
}

func (c *implGlobWatcherWatchGlobCall) Finish() (err error) {
	if ierr := c.clientCall.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (c *implGlobWatcherWatchGlobCall) Cancel() {
	c.clientCall.Cancel()
}

type implGlobWatcherServiceWatchGlobStreamSender struct {
	serverCall _gen_ipc.ServerCall
}

func (s *implGlobWatcherServiceWatchGlobStreamSender) Send(item ChangeBatch) error {
	return s.serverCall.Send(item)
}

// GlobWatcherServiceWatchGlobStream is the interface for streaming responses of the method
// WatchGlob in the service interface GlobWatcher.
type GlobWatcherServiceWatchGlobStream interface {
	// SendStream returns the send portion of the stream.
	SendStream() interface {
		// Send places the item onto the output stream, blocking if there is no buffer
		// space available.  If the client has canceled, an error is returned.
		Send(item ChangeBatch) error
	}
}

// Implementation of the GlobWatcherServiceWatchGlobStream interface that is not exported.
type implGlobWatcherServiceWatchGlobStream struct {
	writer implGlobWatcherServiceWatchGlobStreamSender
}

func (s *implGlobWatcherServiceWatchGlobStream) SendStream() interface {
	// Send places the item onto the output stream, blocking if there is no buffer
	// space available.  If the client has canceled, an error is returned.
	Send(item ChangeBatch) error
} {
	return &s.writer
}

// BindGlobWatcher returns the client stub implementing the GlobWatcher
// interface.
//
// If no _gen_ipc.Client is specified, the default _gen_ipc.Client in the
// global Runtime is used.
func BindGlobWatcher(name string, opts ..._gen_ipc.BindOpt) (GlobWatcher, error) {
	var client _gen_ipc.Client
	switch len(opts) {
	case 0:
		client = _gen_rt.R().Client()
	case 1:
		switch o := opts[0].(type) {
		case _gen_ipc.Client:
			client = o
		default:
			return nil, _gen_vdlutil.ErrUnrecognizedOption
		}
	default:
		return nil, _gen_vdlutil.ErrTooManyOptionsToBind
	}
	stub := &clientStubGlobWatcher{client: client, name: name}

	return stub, nil
}

// NewServerGlobWatcher creates a new server stub.
//
// It takes a regular server implementing the GlobWatcherService
// interface, and returns a new server stub.
func NewServerGlobWatcher(server GlobWatcherService) interface{} {
	return &ServerStubGlobWatcher{
		service: server,
	}
}

// clientStubGlobWatcher implements GlobWatcher.
type clientStubGlobWatcher struct {
	client _gen_ipc.Client
	name   string
}

func (__gen_c *clientStubGlobWatcher) WatchGlob(ctx _gen_context.T, Req GlobRequest, opts ..._gen_ipc.CallOpt) (reply GlobWatcherWatchGlobCall, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "WatchGlob", []interface{}{Req}, opts...); err != nil {
		return
	}
	reply = &implGlobWatcherWatchGlobCall{clientCall: call, readStream: implGlobWatcherWatchGlobStreamIterator{clientCall: call}}
	return
}

func (__gen_c *clientStubGlobWatcher) UnresolveStep(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply []string, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "UnresolveStep", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubGlobWatcher) Signature(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply _gen_ipc.ServiceSignature, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "Signature", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubGlobWatcher) GetMethodTags(ctx _gen_context.T, method string, opts ..._gen_ipc.CallOpt) (reply []interface{}, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "GetMethodTags", []interface{}{method}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

// ServerStubGlobWatcher wraps a server that implements
// GlobWatcherService and provides an object that satisfies
// the requirements of veyron2/ipc.ReflectInvoker.
type ServerStubGlobWatcher struct {
	service GlobWatcherService
}

func (__gen_s *ServerStubGlobWatcher) GetMethodTags(call _gen_ipc.ServerCall, method string) ([]interface{}, error) {
	// TODO(bprosnitz) GetMethodTags() will be replaces with Signature().
	// Note: This exhibits some weird behavior like returning a nil error if the method isn't found.
	// This will change when it is replaced with Signature().
	switch method {
	case "WatchGlob":
		return []interface{}{}, nil
	default:
		return nil, nil
	}
}

func (__gen_s *ServerStubGlobWatcher) Signature(call _gen_ipc.ServerCall) (_gen_ipc.ServiceSignature, error) {
	result := _gen_ipc.ServiceSignature{Methods: make(map[string]_gen_ipc.MethodSignature)}
	result.Methods["WatchGlob"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{
			{Name: "Req", Type: 67},
		},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 68},
		},

		OutStream: 72,
	}

	result.TypeDefs = []_gen_vdlutil.Any{
		_gen_wiretype.NamedPrimitiveType{Type: 0x32, Name: "byte", Tags: []string(nil)}, _gen_wiretype.SliceType{Elem: 0x41, Name: "veyron2/services/watch.ResumeMarker", Tags: []string(nil)}, _gen_wiretype.StructType{
			[]_gen_wiretype.FieldType{
				_gen_wiretype.FieldType{Type: 0x3, Name: "Pattern"},
				_gen_wiretype.FieldType{Type: 0x42, Name: "ResumeMarker"},
			},
			"veyron2/services/watch.GlobRequest", []string(nil)},
		_gen_wiretype.NamedPrimitiveType{Type: 0x1, Name: "error", Tags: []string(nil)}, _gen_wiretype.NamedPrimitiveType{Type: 0x1, Name: "anydata", Tags: []string(nil)}, _gen_wiretype.StructType{
			[]_gen_wiretype.FieldType{
				_gen_wiretype.FieldType{Type: 0x3, Name: "Name"},
				_gen_wiretype.FieldType{Type: 0x24, Name: "State"},
				_gen_wiretype.FieldType{Type: 0x45, Name: "Value"},
				_gen_wiretype.FieldType{Type: 0x42, Name: "ResumeMarker"},
				_gen_wiretype.FieldType{Type: 0x2, Name: "Continued"},
			},
			"veyron2/services/watch.Change", []string(nil)},
		_gen_wiretype.SliceType{Elem: 0x46, Name: "", Tags: []string(nil)}, _gen_wiretype.StructType{
			[]_gen_wiretype.FieldType{
				_gen_wiretype.FieldType{Type: 0x47, Name: "Changes"},
			},
			"veyron2/services/watch.ChangeBatch", []string(nil)},
	}

	return result, nil
}

func (__gen_s *ServerStubGlobWatcher) UnresolveStep(call _gen_ipc.ServerCall) (reply []string, err error) {
	if unresolver, ok := __gen_s.service.(_gen_ipc.Unresolver); ok {
		return unresolver.UnresolveStep(call)
	}
	if call.Server() == nil {
		return
	}
	var published []string
	if published, err = call.Server().Published(); err != nil || published == nil {
		return
	}
	reply = make([]string, len(published))
	for i, p := range published {
		reply[i] = _gen_naming.Join(p, call.Name())
	}
	return
}

func (__gen_s *ServerStubGlobWatcher) WatchGlob(call _gen_ipc.ServerCall, Req GlobRequest) (err error) {
	stream := &implGlobWatcherServiceWatchGlobStream{writer: implGlobWatcherServiceWatchGlobStreamSender{serverCall: call}}
	err = __gen_s.service.WatchGlob(call, Req, stream)
	return
}

// QueryWatcher allows a client to receive updates for changes to objects
// that match a query.  See the package comments for details.
// QueryWatcher is the interface the client binds and uses.
// QueryWatcher_ExcludingUniversal is the interface without internal framework-added methods
// to enable embedding without method collisions.  Not to be used directly by clients.
type QueryWatcher_ExcludingUniversal interface {
	// WatchQuery returns a stream of changes that satisy a query.
	WatchQuery(ctx _gen_context.T, Req QueryRequest, opts ..._gen_ipc.CallOpt) (reply QueryWatcherWatchQueryCall, err error)
}
type QueryWatcher interface {
	_gen_ipc.UniversalServiceMethods
	QueryWatcher_ExcludingUniversal
}

// QueryWatcherService is the interface the server implements.
type QueryWatcherService interface {

	// WatchQuery returns a stream of changes that satisy a query.
	WatchQuery(context _gen_ipc.ServerContext, Req QueryRequest, stream QueryWatcherServiceWatchQueryStream) (err error)
}

// QueryWatcherWatchQueryCall is the interface for call object of the method
// WatchQuery in the service interface QueryWatcher.
type QueryWatcherWatchQueryCall interface {
	// RecvStream returns the recv portion of the stream
	RecvStream() interface {
		// Advance stages an element so the client can retrieve it
		// with Value.  Advance returns true iff there is an
		// element to retrieve.  The client must call Advance before
		// calling Value. Advance may block if an element is not
		// immediately available.
		Advance() bool

		// Value returns the element that was staged by Advance.
		// Value may panic if Advance returned false or was not
		// called at all.  Value does not block.
		Value() ChangeBatch

		// Err returns a non-nil error iff the stream encountered
		// any errors.  Err does not block.
		Err() error
	}

	// Finish blocks until the server is done and returns the positional
	// return values for call.
	//
	// If Cancel has been called, Finish will return immediately; the output of
	// Finish could either be an error signalling cancelation, or the correct
	// positional return values from the server depending on the timing of the
	// call.
	//
	// Calling Finish is mandatory for releasing stream resources, unless Cancel
	// has been called or any of the other methods return an error.
	// Finish should be called at most once.
	Finish() (err error)

	// Cancel cancels the RPC, notifying the server to stop processing.  It
	// is safe to call Cancel concurrently with any of the other stream methods.
	// Calling Cancel after Finish has returned is a no-op.
	Cancel()
}

type implQueryWatcherWatchQueryStreamIterator struct {
	clientCall _gen_ipc.Call
	val        ChangeBatch
	err        error
}

func (c *implQueryWatcherWatchQueryStreamIterator) Advance() bool {
	c.val = ChangeBatch{}
	c.err = c.clientCall.Recv(&c.val)
	return c.err == nil
}

func (c *implQueryWatcherWatchQueryStreamIterator) Value() ChangeBatch {
	return c.val
}

func (c *implQueryWatcherWatchQueryStreamIterator) Err() error {
	if c.err == _gen_io.EOF {
		return nil
	}
	return c.err
}

// Implementation of the QueryWatcherWatchQueryCall interface that is not exported.
type implQueryWatcherWatchQueryCall struct {
	clientCall _gen_ipc.Call
	readStream implQueryWatcherWatchQueryStreamIterator
}

func (c *implQueryWatcherWatchQueryCall) RecvStream() interface {
	Advance() bool
	Value() ChangeBatch
	Err() error
} {
	return &c.readStream
}

func (c *implQueryWatcherWatchQueryCall) Finish() (err error) {
	if ierr := c.clientCall.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (c *implQueryWatcherWatchQueryCall) Cancel() {
	c.clientCall.Cancel()
}

type implQueryWatcherServiceWatchQueryStreamSender struct {
	serverCall _gen_ipc.ServerCall
}

func (s *implQueryWatcherServiceWatchQueryStreamSender) Send(item ChangeBatch) error {
	return s.serverCall.Send(item)
}

// QueryWatcherServiceWatchQueryStream is the interface for streaming responses of the method
// WatchQuery in the service interface QueryWatcher.
type QueryWatcherServiceWatchQueryStream interface {
	// SendStream returns the send portion of the stream.
	SendStream() interface {
		// Send places the item onto the output stream, blocking if there is no buffer
		// space available.  If the client has canceled, an error is returned.
		Send(item ChangeBatch) error
	}
}

// Implementation of the QueryWatcherServiceWatchQueryStream interface that is not exported.
type implQueryWatcherServiceWatchQueryStream struct {
	writer implQueryWatcherServiceWatchQueryStreamSender
}

func (s *implQueryWatcherServiceWatchQueryStream) SendStream() interface {
	// Send places the item onto the output stream, blocking if there is no buffer
	// space available.  If the client has canceled, an error is returned.
	Send(item ChangeBatch) error
} {
	return &s.writer
}

// BindQueryWatcher returns the client stub implementing the QueryWatcher
// interface.
//
// If no _gen_ipc.Client is specified, the default _gen_ipc.Client in the
// global Runtime is used.
func BindQueryWatcher(name string, opts ..._gen_ipc.BindOpt) (QueryWatcher, error) {
	var client _gen_ipc.Client
	switch len(opts) {
	case 0:
		client = _gen_rt.R().Client()
	case 1:
		switch o := opts[0].(type) {
		case _gen_ipc.Client:
			client = o
		default:
			return nil, _gen_vdlutil.ErrUnrecognizedOption
		}
	default:
		return nil, _gen_vdlutil.ErrTooManyOptionsToBind
	}
	stub := &clientStubQueryWatcher{client: client, name: name}

	return stub, nil
}

// NewServerQueryWatcher creates a new server stub.
//
// It takes a regular server implementing the QueryWatcherService
// interface, and returns a new server stub.
func NewServerQueryWatcher(server QueryWatcherService) interface{} {
	return &ServerStubQueryWatcher{
		service: server,
	}
}

// clientStubQueryWatcher implements QueryWatcher.
type clientStubQueryWatcher struct {
	client _gen_ipc.Client
	name   string
}

func (__gen_c *clientStubQueryWatcher) WatchQuery(ctx _gen_context.T, Req QueryRequest, opts ..._gen_ipc.CallOpt) (reply QueryWatcherWatchQueryCall, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "WatchQuery", []interface{}{Req}, opts...); err != nil {
		return
	}
	reply = &implQueryWatcherWatchQueryCall{clientCall: call, readStream: implQueryWatcherWatchQueryStreamIterator{clientCall: call}}
	return
}

func (__gen_c *clientStubQueryWatcher) UnresolveStep(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply []string, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "UnresolveStep", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubQueryWatcher) Signature(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply _gen_ipc.ServiceSignature, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "Signature", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubQueryWatcher) GetMethodTags(ctx _gen_context.T, method string, opts ..._gen_ipc.CallOpt) (reply []interface{}, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "GetMethodTags", []interface{}{method}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

// ServerStubQueryWatcher wraps a server that implements
// QueryWatcherService and provides an object that satisfies
// the requirements of veyron2/ipc.ReflectInvoker.
type ServerStubQueryWatcher struct {
	service QueryWatcherService
}

func (__gen_s *ServerStubQueryWatcher) GetMethodTags(call _gen_ipc.ServerCall, method string) ([]interface{}, error) {
	// TODO(bprosnitz) GetMethodTags() will be replaces with Signature().
	// Note: This exhibits some weird behavior like returning a nil error if the method isn't found.
	// This will change when it is replaced with Signature().
	switch method {
	case "WatchQuery":
		return []interface{}{}, nil
	default:
		return nil, nil
	}
}

func (__gen_s *ServerStubQueryWatcher) Signature(call _gen_ipc.ServerCall) (_gen_ipc.ServiceSignature, error) {
	result := _gen_ipc.ServiceSignature{Methods: make(map[string]_gen_ipc.MethodSignature)}
	result.Methods["WatchQuery"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{
			{Name: "Req", Type: 68},
		},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 69},
		},

		OutStream: 73,
	}

	result.TypeDefs = []_gen_vdlutil.Any{
		_gen_wiretype.StructType{
			[]_gen_wiretype.FieldType{
				_gen_wiretype.FieldType{Type: 0x3, Name: "Stmt"},
			},
			"veyron2/query.Query", []string(nil)},
		_gen_wiretype.NamedPrimitiveType{Type: 0x32, Name: "byte", Tags: []string(nil)}, _gen_wiretype.SliceType{Elem: 0x42, Name: "veyron2/services/watch.ResumeMarker", Tags: []string(nil)}, _gen_wiretype.StructType{
			[]_gen_wiretype.FieldType{
				_gen_wiretype.FieldType{Type: 0x41, Name: "Query"},
				_gen_wiretype.FieldType{Type: 0x43, Name: "ResumeMarker"},
			},
			"veyron2/services/watch.QueryRequest", []string(nil)},
		_gen_wiretype.NamedPrimitiveType{Type: 0x1, Name: "error", Tags: []string(nil)}, _gen_wiretype.NamedPrimitiveType{Type: 0x1, Name: "anydata", Tags: []string(nil)}, _gen_wiretype.StructType{
			[]_gen_wiretype.FieldType{
				_gen_wiretype.FieldType{Type: 0x3, Name: "Name"},
				_gen_wiretype.FieldType{Type: 0x24, Name: "State"},
				_gen_wiretype.FieldType{Type: 0x46, Name: "Value"},
				_gen_wiretype.FieldType{Type: 0x43, Name: "ResumeMarker"},
				_gen_wiretype.FieldType{Type: 0x2, Name: "Continued"},
			},
			"veyron2/services/watch.Change", []string(nil)},
		_gen_wiretype.SliceType{Elem: 0x47, Name: "", Tags: []string(nil)}, _gen_wiretype.StructType{
			[]_gen_wiretype.FieldType{
				_gen_wiretype.FieldType{Type: 0x48, Name: "Changes"},
			},
			"veyron2/services/watch.ChangeBatch", []string(nil)},
	}

	return result, nil
}

func (__gen_s *ServerStubQueryWatcher) UnresolveStep(call _gen_ipc.ServerCall) (reply []string, err error) {
	if unresolver, ok := __gen_s.service.(_gen_ipc.Unresolver); ok {
		return unresolver.UnresolveStep(call)
	}
	if call.Server() == nil {
		return
	}
	var published []string
	if published, err = call.Server().Published(); err != nil || published == nil {
		return
	}
	reply = make([]string, len(published))
	for i, p := range published {
		reply[i] = _gen_naming.Join(p, call.Name())
	}
	return
}

func (__gen_s *ServerStubQueryWatcher) WatchQuery(call _gen_ipc.ServerCall, Req QueryRequest) (err error) {
	stream := &implQueryWatcherServiceWatchQueryStream{writer: implQueryWatcherServiceWatchQueryStreamSender{serverCall: call}}
	err = __gen_s.service.WatchQuery(call, Req, stream)
	return
}
