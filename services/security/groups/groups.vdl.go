// This file was auto-generated by the vanadium vdl tool.
// Source: groups.vdl

// Package groups defines types and interfaces pertaining to groups, which can
// be referenced by BlessingPatterns (e.g. in AccessLists).
//
// TODO(sadovsky): Write a detailed description of this package and add a
// reference to the (forthcoming) design doc.
package groups

import (
	// VDL system imports
	"v.io/v23"
	"v.io/v23/context"
	"v.io/v23/i18n"
	"v.io/v23/rpc"
	"v.io/v23/vdl"
	"v.io/v23/verror"

	// VDL user imports
	"v.io/v23/services/security/access"
	"v.io/v23/services/security/access/object"
)

// BlessingPatternChunk is a substring of a BlessingPattern. As with
// BlessingPatterns, BlessingPatternChunks may contain references to
// groups. However, they may be restricted in other ways. For example, in the
// future BlessingPatterns may support "$" terminators, but these may be
// disallowed for BlessingPatternChunks.
type BlessingPatternChunk string

func (BlessingPatternChunk) __VDLReflect(struct {
	Name string "v.io/v23/services/security/groups.BlessingPatternChunk"
}) {
}

type GetRequest struct {
}

func (GetRequest) __VDLReflect(struct {
	Name string "v.io/v23/services/security/groups.GetRequest"
}) {
}

type GetResponse struct {
	Entries map[BlessingPatternChunk]struct{}
}

func (GetResponse) __VDLReflect(struct {
	Name string "v.io/v23/services/security/groups.GetResponse"
}) {
}

type RestRequest struct {
}

func (RestRequest) __VDLReflect(struct {
	Name string "v.io/v23/services/security/groups.RestRequest"
}) {
}

type RestResponse struct {
}

func (RestResponse) __VDLReflect(struct {
	Name string "v.io/v23/services/security/groups.RestResponse"
}) {
}

func init() {
	vdl.Register((*BlessingPatternChunk)(nil))
	vdl.Register((*GetRequest)(nil))
	vdl.Register((*GetResponse)(nil))
	vdl.Register((*RestRequest)(nil))
	vdl.Register((*RestResponse)(nil))
}

var (
	ErrNoBlessings         = verror.Register("v.io/v23/services/security/groups.NoBlessings", verror.NoRetry, "{1:}{2:} No blessings recognized; cannot create group AccessList")
	ErrExcessiveContention = verror.Register("v.io/v23/services/security/groups.ExcessiveContention", verror.RetryBackoff, "{1:}{2:} Gave up after encountering excessive contention; try again later")
)

func init() {
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrNoBlessings.ID), "{1:}{2:} No blessings recognized; cannot create group AccessList")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrExcessiveContention.ID), "{1:}{2:} Gave up after encountering excessive contention; try again later")
}

// NewErrNoBlessings returns an error with the ErrNoBlessings ID.
func NewErrNoBlessings(ctx *context.T) error {
	return verror.New(ErrNoBlessings, ctx)
}

// NewErrExcessiveContention returns an error with the ErrExcessiveContention ID.
func NewErrExcessiveContention(ctx *context.T) error {
	return verror.New(ErrExcessiveContention, ctx)
}

// GroupClientMethods is the client interface
// containing Group methods.
//
// A group's etag covers its AccessList as well as any other data stored in the group.
// Clients should treat etags as opaque identifiers. For both Get and Rest, if
// etag is set and matches the Group's current etag, the response will indicate
// that fact but will otherwise be empty.
type GroupClientMethods interface {
	// Object provides access control for Vanadium objects.
	//
	// Vanadium services implementing dynamic access control would typically embed
	// this interface and tag additional methods defined by the service with one of
	// Admin, Read, Write, Resolve etc. For example, the VDL definition of the
	// object would be:
	//
	//   package mypackage
	//
	//   import "v.io/v23/security/access"
	//   import "v.io/v23/security/access/object"
	//
	//   type MyObject interface {
	//     object.Object
	//     MyRead() (string, error) {access.Read}
	//     MyWrite(string) error    {access.Write}
	//   }
	//
	// If the set of pre-defined tags is insufficient, services may define their
	// own tag type and annotate all methods with this new type.
	//
	// Instead of embedding this Object interface, define SetPermissions and
	// GetPermissions in their own interface. Authorization policies will typically
	// respect annotations of a single type. For example, the VDL definition of an
	// object would be:
	//
	//  package mypackage
	//
	//  import "v.io/v23/security/access"
	//
	//  type MyTag string
	//
	//  const (
	//    Blue = MyTag("Blue")
	//    Red  = MyTag("Red")
	//  )
	//
	//  type MyObject interface {
	//    MyMethod() (string, error) {Blue}
	//
	//    // Allow clients to change access via the access.Object interface:
	//    SetPermissions(acl access.Permissions, etag string) error         {Red}
	//    GetPermissions() (acl access.Permissions, etag string, err error) {Blue}
	//  }
	object.ObjectClientMethods
	// Create creates a new group if it doesn't already exist.
	// If acl is nil, a default Permissions is used, providing Admin access to
	// the caller.
	// Create requires the caller to have Write permission at the GroupServer.
	Create(ctx *context.T, acl access.Permissions, entries []BlessingPatternChunk, opts ...rpc.CallOpt) error
	// Delete deletes the group.
	// Permissions for all group-related methods except Create() are checked
	// against the Group object.
	Delete(ctx *context.T, etag string, opts ...rpc.CallOpt) error
	// Add adds an entry to the group.
	Add(ctx *context.T, entry BlessingPatternChunk, etag string, opts ...rpc.CallOpt) error
	// Remove removes an entry from the group.
	Remove(ctx *context.T, entry BlessingPatternChunk, etag string, opts ...rpc.CallOpt) error
	// Get returns all entries in the group.
	// TODO(sadovsky): Flesh out this API.
	Get(ctx *context.T, req GetRequest, reqEtag string, opts ...rpc.CallOpt) (res GetResponse, etag string, err error)
	// Rest returns information sufficient for the client to perform its AccessList
	// checks.
	// TODO(sadovsky): Flesh out this API.
	Rest(ctx *context.T, req RestRequest, reqEtag string, opts ...rpc.CallOpt) (res RestResponse, etag string, err error)
}

// GroupClientStub adds universal methods to GroupClientMethods.
type GroupClientStub interface {
	GroupClientMethods
	rpc.UniversalServiceMethods
}

// GroupClient returns a client stub for Group.
func GroupClient(name string, opts ...rpc.BindOpt) GroupClientStub {
	var client rpc.Client
	for _, opt := range opts {
		if clientOpt, ok := opt.(rpc.Client); ok {
			client = clientOpt
		}
	}
	return implGroupClientStub{name, client, object.ObjectClient(name, client)}
}

type implGroupClientStub struct {
	name   string
	client rpc.Client

	object.ObjectClientStub
}

func (c implGroupClientStub) c(ctx *context.T) rpc.Client {
	if c.client != nil {
		return c.client
	}
	return v23.GetClient(ctx)
}

func (c implGroupClientStub) Create(ctx *context.T, i0 access.Permissions, i1 []BlessingPatternChunk, opts ...rpc.CallOpt) (err error) {
	var call rpc.ClientCall
	if call, err = c.c(ctx).StartCall(ctx, c.name, "Create", []interface{}{i0, i1}, opts...); err != nil {
		return
	}
	err = call.Finish()
	return
}

func (c implGroupClientStub) Delete(ctx *context.T, i0 string, opts ...rpc.CallOpt) (err error) {
	var call rpc.ClientCall
	if call, err = c.c(ctx).StartCall(ctx, c.name, "Delete", []interface{}{i0}, opts...); err != nil {
		return
	}
	err = call.Finish()
	return
}

func (c implGroupClientStub) Add(ctx *context.T, i0 BlessingPatternChunk, i1 string, opts ...rpc.CallOpt) (err error) {
	var call rpc.ClientCall
	if call, err = c.c(ctx).StartCall(ctx, c.name, "Add", []interface{}{i0, i1}, opts...); err != nil {
		return
	}
	err = call.Finish()
	return
}

func (c implGroupClientStub) Remove(ctx *context.T, i0 BlessingPatternChunk, i1 string, opts ...rpc.CallOpt) (err error) {
	var call rpc.ClientCall
	if call, err = c.c(ctx).StartCall(ctx, c.name, "Remove", []interface{}{i0, i1}, opts...); err != nil {
		return
	}
	err = call.Finish()
	return
}

func (c implGroupClientStub) Get(ctx *context.T, i0 GetRequest, i1 string, opts ...rpc.CallOpt) (o0 GetResponse, o1 string, err error) {
	var call rpc.ClientCall
	if call, err = c.c(ctx).StartCall(ctx, c.name, "Get", []interface{}{i0, i1}, opts...); err != nil {
		return
	}
	err = call.Finish(&o0, &o1)
	return
}

func (c implGroupClientStub) Rest(ctx *context.T, i0 RestRequest, i1 string, opts ...rpc.CallOpt) (o0 RestResponse, o1 string, err error) {
	var call rpc.ClientCall
	if call, err = c.c(ctx).StartCall(ctx, c.name, "Rest", []interface{}{i0, i1}, opts...); err != nil {
		return
	}
	err = call.Finish(&o0, &o1)
	return
}

// GroupServerMethods is the interface a server writer
// implements for Group.
//
// A group's etag covers its AccessList as well as any other data stored in the group.
// Clients should treat etags as opaque identifiers. For both Get and Rest, if
// etag is set and matches the Group's current etag, the response will indicate
// that fact but will otherwise be empty.
type GroupServerMethods interface {
	// Object provides access control for Vanadium objects.
	//
	// Vanadium services implementing dynamic access control would typically embed
	// this interface and tag additional methods defined by the service with one of
	// Admin, Read, Write, Resolve etc. For example, the VDL definition of the
	// object would be:
	//
	//   package mypackage
	//
	//   import "v.io/v23/security/access"
	//   import "v.io/v23/security/access/object"
	//
	//   type MyObject interface {
	//     object.Object
	//     MyRead() (string, error) {access.Read}
	//     MyWrite(string) error    {access.Write}
	//   }
	//
	// If the set of pre-defined tags is insufficient, services may define their
	// own tag type and annotate all methods with this new type.
	//
	// Instead of embedding this Object interface, define SetPermissions and
	// GetPermissions in their own interface. Authorization policies will typically
	// respect annotations of a single type. For example, the VDL definition of an
	// object would be:
	//
	//  package mypackage
	//
	//  import "v.io/v23/security/access"
	//
	//  type MyTag string
	//
	//  const (
	//    Blue = MyTag("Blue")
	//    Red  = MyTag("Red")
	//  )
	//
	//  type MyObject interface {
	//    MyMethod() (string, error) {Blue}
	//
	//    // Allow clients to change access via the access.Object interface:
	//    SetPermissions(acl access.Permissions, etag string) error         {Red}
	//    GetPermissions() (acl access.Permissions, etag string, err error) {Blue}
	//  }
	object.ObjectServerMethods
	// Create creates a new group if it doesn't already exist.
	// If acl is nil, a default Permissions is used, providing Admin access to
	// the caller.
	// Create requires the caller to have Write permission at the GroupServer.
	Create(call rpc.ServerCall, acl access.Permissions, entries []BlessingPatternChunk) error
	// Delete deletes the group.
	// Permissions for all group-related methods except Create() are checked
	// against the Group object.
	Delete(call rpc.ServerCall, etag string) error
	// Add adds an entry to the group.
	Add(call rpc.ServerCall, entry BlessingPatternChunk, etag string) error
	// Remove removes an entry from the group.
	Remove(call rpc.ServerCall, entry BlessingPatternChunk, etag string) error
	// Get returns all entries in the group.
	// TODO(sadovsky): Flesh out this API.
	Get(call rpc.ServerCall, req GetRequest, reqEtag string) (res GetResponse, etag string, err error)
	// Rest returns information sufficient for the client to perform its AccessList
	// checks.
	// TODO(sadovsky): Flesh out this API.
	Rest(call rpc.ServerCall, req RestRequest, reqEtag string) (res RestResponse, etag string, err error)
}

// GroupServerStubMethods is the server interface containing
// Group methods, as expected by rpc.Server.
// There is no difference between this interface and GroupServerMethods
// since there are no streaming methods.
type GroupServerStubMethods GroupServerMethods

// GroupServerStub adds universal methods to GroupServerStubMethods.
type GroupServerStub interface {
	GroupServerStubMethods
	// Describe the Group interfaces.
	Describe__() []rpc.InterfaceDesc
}

// GroupServer returns a server stub for Group.
// It converts an implementation of GroupServerMethods into
// an object that may be used by rpc.Server.
func GroupServer(impl GroupServerMethods) GroupServerStub {
	stub := implGroupServerStub{
		impl:             impl,
		ObjectServerStub: object.ObjectServer(impl),
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

type implGroupServerStub struct {
	impl GroupServerMethods
	object.ObjectServerStub
	gs *rpc.GlobState
}

func (s implGroupServerStub) Create(call rpc.ServerCall, i0 access.Permissions, i1 []BlessingPatternChunk) error {
	return s.impl.Create(call, i0, i1)
}

func (s implGroupServerStub) Delete(call rpc.ServerCall, i0 string) error {
	return s.impl.Delete(call, i0)
}

func (s implGroupServerStub) Add(call rpc.ServerCall, i0 BlessingPatternChunk, i1 string) error {
	return s.impl.Add(call, i0, i1)
}

func (s implGroupServerStub) Remove(call rpc.ServerCall, i0 BlessingPatternChunk, i1 string) error {
	return s.impl.Remove(call, i0, i1)
}

func (s implGroupServerStub) Get(call rpc.ServerCall, i0 GetRequest, i1 string) (GetResponse, string, error) {
	return s.impl.Get(call, i0, i1)
}

func (s implGroupServerStub) Rest(call rpc.ServerCall, i0 RestRequest, i1 string) (RestResponse, string, error) {
	return s.impl.Rest(call, i0, i1)
}

func (s implGroupServerStub) Globber() *rpc.GlobState {
	return s.gs
}

func (s implGroupServerStub) Describe__() []rpc.InterfaceDesc {
	return []rpc.InterfaceDesc{GroupDesc, object.ObjectDesc}
}

// GroupDesc describes the Group interface.
var GroupDesc rpc.InterfaceDesc = descGroup

// descGroup hides the desc to keep godoc clean.
var descGroup = rpc.InterfaceDesc{
	Name:    "Group",
	PkgPath: "v.io/v23/services/security/groups",
	Doc:     "// A group's etag covers its AccessList as well as any other data stored in the group.\n// Clients should treat etags as opaque identifiers. For both Get and Rest, if\n// etag is set and matches the Group's current etag, the response will indicate\n// that fact but will otherwise be empty.",
	Embeds: []rpc.EmbedDesc{
		{"Object", "v.io/v23/services/security/access/object", "// Object provides access control for Vanadium objects.\n//\n// Vanadium services implementing dynamic access control would typically embed\n// this interface and tag additional methods defined by the service with one of\n// Admin, Read, Write, Resolve etc. For example, the VDL definition of the\n// object would be:\n//\n//   package mypackage\n//\n//   import \"v.io/v23/security/access\"\n//   import \"v.io/v23/security/access/object\"\n//\n//   type MyObject interface {\n//     object.Object\n//     MyRead() (string, error) {access.Read}\n//     MyWrite(string) error    {access.Write}\n//   }\n//\n// If the set of pre-defined tags is insufficient, services may define their\n// own tag type and annotate all methods with this new type.\n//\n// Instead of embedding this Object interface, define SetPermissions and\n// GetPermissions in their own interface. Authorization policies will typically\n// respect annotations of a single type. For example, the VDL definition of an\n// object would be:\n//\n//  package mypackage\n//\n//  import \"v.io/v23/security/access\"\n//\n//  type MyTag string\n//\n//  const (\n//    Blue = MyTag(\"Blue\")\n//    Red  = MyTag(\"Red\")\n//  )\n//\n//  type MyObject interface {\n//    MyMethod() (string, error) {Blue}\n//\n//    // Allow clients to change access via the access.Object interface:\n//    SetPermissions(acl access.Permissions, etag string) error         {Red}\n//    GetPermissions() (acl access.Permissions, etag string, err error) {Blue}\n//  }"},
	},
	Methods: []rpc.MethodDesc{
		{
			Name: "Create",
			Doc:  "// Create creates a new group if it doesn't already exist.\n// If acl is nil, a default Permissions is used, providing Admin access to\n// the caller.\n// Create requires the caller to have Write permission at the GroupServer.",
			InArgs: []rpc.ArgDesc{
				{"acl", ``},     // access.Permissions
				{"entries", ``}, // []BlessingPatternChunk
			},
			Tags: []*vdl.Value{vdl.ValueOf(access.Tag("Write"))},
		},
		{
			Name: "Delete",
			Doc:  "// Delete deletes the group.\n// Permissions for all group-related methods except Create() are checked\n// against the Group object.",
			InArgs: []rpc.ArgDesc{
				{"etag", ``}, // string
			},
			Tags: []*vdl.Value{vdl.ValueOf(access.Tag("Write"))},
		},
		{
			Name: "Add",
			Doc:  "// Add adds an entry to the group.",
			InArgs: []rpc.ArgDesc{
				{"entry", ``}, // BlessingPatternChunk
				{"etag", ``},  // string
			},
			Tags: []*vdl.Value{vdl.ValueOf(access.Tag("Write"))},
		},
		{
			Name: "Remove",
			Doc:  "// Remove removes an entry from the group.",
			InArgs: []rpc.ArgDesc{
				{"entry", ``}, // BlessingPatternChunk
				{"etag", ``},  // string
			},
			Tags: []*vdl.Value{vdl.ValueOf(access.Tag("Write"))},
		},
		{
			Name: "Get",
			Doc:  "// Get returns all entries in the group.\n// TODO(sadovsky): Flesh out this API.",
			InArgs: []rpc.ArgDesc{
				{"req", ``},     // GetRequest
				{"reqEtag", ``}, // string
			},
			OutArgs: []rpc.ArgDesc{
				{"res", ``},  // GetResponse
				{"etag", ``}, // string
			},
			Tags: []*vdl.Value{vdl.ValueOf(access.Tag("Read"))},
		},
		{
			Name: "Rest",
			Doc:  "// Rest returns information sufficient for the client to perform its AccessList\n// checks.\n// TODO(sadovsky): Flesh out this API.",
			InArgs: []rpc.ArgDesc{
				{"req", ``},     // RestRequest
				{"reqEtag", ``}, // string
			},
			OutArgs: []rpc.ArgDesc{
				{"res", ``},  // RestResponse
				{"etag", ``}, // string
			},
			Tags: []*vdl.Value{vdl.ValueOf(access.Tag("Resolve"))},
		},
	},
}
