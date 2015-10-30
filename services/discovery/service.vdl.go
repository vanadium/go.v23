// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Source: service.vdl

// Package discovery defines the wire interfaces for discovering services.
package discovery

import (
	// VDL system imports
	"io"
	"v.io/v23"
	"v.io/v23/context"
	"v.io/v23/rpc"
	"v.io/v23/vdl"

	// VDL user imports
	"v.io/v23/discovery"
	"v.io/v23/security"
	"v.io/v23/security/access"
)

// Used to unregister a service.
type ServiceHandle uint32

func (ServiceHandle) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/services/discovery.ServiceHandle"`
}) {
}

func init() {
	vdl.Register((*ServiceHandle)(nil))
}

// AdvertiserClientMethods is the client interface
// containing Advertiser methods.
//
// Advertiser is the interface for advertising services.
type AdvertiserClientMethods interface {
	// RegisterService registers a service to be discovered by "Scanner" implementations.
	// visibility is used to limit the principals that can see the advertisement. An empty
	// set means that there are no restrictions on visibility (i.e, equivalent to
	// []security.BlessingPattern{security.AllPrincipals}).
	RegisterService(ctx *context.T, service discovery.Service, visibility []security.BlessingPattern, opts ...rpc.CallOpt) (ServiceHandle, error)
	// UnregisterService unregisters a registered service from advertising.
	UnregisterService(ctx *context.T, handle ServiceHandle, opts ...rpc.CallOpt) error
}

// AdvertiserClientStub adds universal methods to AdvertiserClientMethods.
type AdvertiserClientStub interface {
	AdvertiserClientMethods
	rpc.UniversalServiceMethods
}

// AdvertiserClient returns a client stub for Advertiser.
func AdvertiserClient(name string) AdvertiserClientStub {
	return implAdvertiserClientStub{name}
}

type implAdvertiserClientStub struct {
	name string
}

func (c implAdvertiserClientStub) RegisterService(ctx *context.T, i0 discovery.Service, i1 []security.BlessingPattern, opts ...rpc.CallOpt) (o0 ServiceHandle, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "RegisterService", []interface{}{i0, i1}, []interface{}{&o0}, opts...)
	return
}

func (c implAdvertiserClientStub) UnregisterService(ctx *context.T, i0 ServiceHandle, opts ...rpc.CallOpt) (err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "UnregisterService", []interface{}{i0}, nil, opts...)
	return
}

// AdvertiserServerMethods is the interface a server writer
// implements for Advertiser.
//
// Advertiser is the interface for advertising services.
type AdvertiserServerMethods interface {
	// RegisterService registers a service to be discovered by "Scanner" implementations.
	// visibility is used to limit the principals that can see the advertisement. An empty
	// set means that there are no restrictions on visibility (i.e, equivalent to
	// []security.BlessingPattern{security.AllPrincipals}).
	RegisterService(ctx *context.T, call rpc.ServerCall, service discovery.Service, visibility []security.BlessingPattern) (ServiceHandle, error)
	// UnregisterService unregisters a registered service from advertising.
	UnregisterService(ctx *context.T, call rpc.ServerCall, handle ServiceHandle) error
}

// AdvertiserServerStubMethods is the server interface containing
// Advertiser methods, as expected by rpc.Server.
// There is no difference between this interface and AdvertiserServerMethods
// since there are no streaming methods.
type AdvertiserServerStubMethods AdvertiserServerMethods

// AdvertiserServerStub adds universal methods to AdvertiserServerStubMethods.
type AdvertiserServerStub interface {
	AdvertiserServerStubMethods
	// Describe the Advertiser interfaces.
	Describe__() []rpc.InterfaceDesc
}

// AdvertiserServer returns a server stub for Advertiser.
// It converts an implementation of AdvertiserServerMethods into
// an object that may be used by rpc.Server.
func AdvertiserServer(impl AdvertiserServerMethods) AdvertiserServerStub {
	stub := implAdvertiserServerStub{
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

type implAdvertiserServerStub struct {
	impl AdvertiserServerMethods
	gs   *rpc.GlobState
}

func (s implAdvertiserServerStub) RegisterService(ctx *context.T, call rpc.ServerCall, i0 discovery.Service, i1 []security.BlessingPattern) (ServiceHandle, error) {
	return s.impl.RegisterService(ctx, call, i0, i1)
}

func (s implAdvertiserServerStub) UnregisterService(ctx *context.T, call rpc.ServerCall, i0 ServiceHandle) error {
	return s.impl.UnregisterService(ctx, call, i0)
}

func (s implAdvertiserServerStub) Globber() *rpc.GlobState {
	return s.gs
}

func (s implAdvertiserServerStub) Describe__() []rpc.InterfaceDesc {
	return []rpc.InterfaceDesc{AdvertiserDesc}
}

// AdvertiserDesc describes the Advertiser interface.
var AdvertiserDesc rpc.InterfaceDesc = descAdvertiser

// descAdvertiser hides the desc to keep godoc clean.
var descAdvertiser = rpc.InterfaceDesc{
	Name:    "Advertiser",
	PkgPath: "v.io/v23/services/discovery",
	Doc:     "// Advertiser is the interface for advertising services.",
	Methods: []rpc.MethodDesc{
		{
			Name: "RegisterService",
			Doc:  "// RegisterService registers a service to be discovered by \"Scanner\" implementations.\n// visibility is used to limit the principals that can see the advertisement. An empty\n// set means that there are no restrictions on visibility (i.e, equivalent to\n// []security.BlessingPattern{security.AllPrincipals}).",
			InArgs: []rpc.ArgDesc{
				{"service", ``},    // discovery.Service
				{"visibility", ``}, // []security.BlessingPattern
			},
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // ServiceHandle
			},
			Tags: []*vdl.Value{vdl.ValueOf(access.Tag("Write"))},
		},
		{
			Name: "UnregisterService",
			Doc:  "// UnregisterService unregisters a registered service from advertising.",
			InArgs: []rpc.ArgDesc{
				{"handle", ``}, // ServiceHandle
			},
			Tags: []*vdl.Value{vdl.ValueOf(access.Tag("Write"))},
		},
	},
}

// ScannerClientMethods is the client interface
// containing Scanner methods.
//
// Scanner is the interface for scanning services.
type ScannerClientMethods interface {
	// Scan scans services that match the query and returns the stream of discovered
	// services. Scanning will continue until the client cancels the call.
	//
	// The query is a WHERE expression of syncQL query against scanned services, where
	// keys are InstanceUuids and values are Service.
	//
	// Examples
	//
	//    v.InstanceName = "v.io/i"
	//    v.InstanceName = "v.io/i" AND v.Attrs["a"] = "v"
	//    v.Attrs["a"] = "v1" OR v.Attrs["a"] = "v2"
	//
	// SyncQL tutorial at:
	//    https://github.com/vanadium/docs/blob/master/tutorials/syncql-tutorial.md
	Scan(ctx *context.T, query string, opts ...rpc.CallOpt) (ScannerScanClientCall, error)
}

// ScannerClientStub adds universal methods to ScannerClientMethods.
type ScannerClientStub interface {
	ScannerClientMethods
	rpc.UniversalServiceMethods
}

// ScannerClient returns a client stub for Scanner.
func ScannerClient(name string) ScannerClientStub {
	return implScannerClientStub{name}
}

type implScannerClientStub struct {
	name string
}

func (c implScannerClientStub) Scan(ctx *context.T, i0 string, opts ...rpc.CallOpt) (ocall ScannerScanClientCall, err error) {
	var call rpc.ClientCall
	if call, err = v23.GetClient(ctx).StartCall(ctx, c.name, "Scan", []interface{}{i0}, opts...); err != nil {
		return
	}
	ocall = &implScannerScanClientCall{ClientCall: call}
	return
}

// ScannerScanClientStream is the client stream for Scanner.Scan.
type ScannerScanClientStream interface {
	// RecvStream returns the receiver side of the Scanner.Scan client stream.
	RecvStream() interface {
		// Advance stages an item so that it may be retrieved via Value.  Returns
		// true iff there is an item to retrieve.  Advance must be called before
		// Value is called.  May block if an item is not available.
		Advance() bool
		// Value returns the item that was staged by Advance.  May panic if Advance
		// returned false or was not called.  Never blocks.
		Value() discovery.Update
		// Err returns any error encountered by Advance.  Never blocks.
		Err() error
	}
}

// ScannerScanClientCall represents the call returned from Scanner.Scan.
type ScannerScanClientCall interface {
	ScannerScanClientStream
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
	Finish() error
}

type implScannerScanClientCall struct {
	rpc.ClientCall
	valRecv discovery.Update
	errRecv error
}

func (c *implScannerScanClientCall) RecvStream() interface {
	Advance() bool
	Value() discovery.Update
	Err() error
} {
	return implScannerScanClientCallRecv{c}
}

type implScannerScanClientCallRecv struct {
	c *implScannerScanClientCall
}

func (c implScannerScanClientCallRecv) Advance() bool {
	c.c.errRecv = c.c.Recv(&c.c.valRecv)
	return c.c.errRecv == nil
}
func (c implScannerScanClientCallRecv) Value() discovery.Update {
	return c.c.valRecv
}
func (c implScannerScanClientCallRecv) Err() error {
	if c.c.errRecv == io.EOF {
		return nil
	}
	return c.c.errRecv
}
func (c *implScannerScanClientCall) Finish() (err error) {
	err = c.ClientCall.Finish()
	return
}

// ScannerServerMethods is the interface a server writer
// implements for Scanner.
//
// Scanner is the interface for scanning services.
type ScannerServerMethods interface {
	// Scan scans services that match the query and returns the stream of discovered
	// services. Scanning will continue until the client cancels the call.
	//
	// The query is a WHERE expression of syncQL query against scanned services, where
	// keys are InstanceUuids and values are Service.
	//
	// Examples
	//
	//    v.InstanceName = "v.io/i"
	//    v.InstanceName = "v.io/i" AND v.Attrs["a"] = "v"
	//    v.Attrs["a"] = "v1" OR v.Attrs["a"] = "v2"
	//
	// SyncQL tutorial at:
	//    https://github.com/vanadium/docs/blob/master/tutorials/syncql-tutorial.md
	Scan(ctx *context.T, call ScannerScanServerCall, query string) error
}

// ScannerServerStubMethods is the server interface containing
// Scanner methods, as expected by rpc.Server.
// The only difference between this interface and ScannerServerMethods
// is the streaming methods.
type ScannerServerStubMethods interface {
	// Scan scans services that match the query and returns the stream of discovered
	// services. Scanning will continue until the client cancels the call.
	//
	// The query is a WHERE expression of syncQL query against scanned services, where
	// keys are InstanceUuids and values are Service.
	//
	// Examples
	//
	//    v.InstanceName = "v.io/i"
	//    v.InstanceName = "v.io/i" AND v.Attrs["a"] = "v"
	//    v.Attrs["a"] = "v1" OR v.Attrs["a"] = "v2"
	//
	// SyncQL tutorial at:
	//    https://github.com/vanadium/docs/blob/master/tutorials/syncql-tutorial.md
	Scan(ctx *context.T, call *ScannerScanServerCallStub, query string) error
}

// ScannerServerStub adds universal methods to ScannerServerStubMethods.
type ScannerServerStub interface {
	ScannerServerStubMethods
	// Describe the Scanner interfaces.
	Describe__() []rpc.InterfaceDesc
}

// ScannerServer returns a server stub for Scanner.
// It converts an implementation of ScannerServerMethods into
// an object that may be used by rpc.Server.
func ScannerServer(impl ScannerServerMethods) ScannerServerStub {
	stub := implScannerServerStub{
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

type implScannerServerStub struct {
	impl ScannerServerMethods
	gs   *rpc.GlobState
}

func (s implScannerServerStub) Scan(ctx *context.T, call *ScannerScanServerCallStub, i0 string) error {
	return s.impl.Scan(ctx, call, i0)
}

func (s implScannerServerStub) Globber() *rpc.GlobState {
	return s.gs
}

func (s implScannerServerStub) Describe__() []rpc.InterfaceDesc {
	return []rpc.InterfaceDesc{ScannerDesc}
}

// ScannerDesc describes the Scanner interface.
var ScannerDesc rpc.InterfaceDesc = descScanner

// descScanner hides the desc to keep godoc clean.
var descScanner = rpc.InterfaceDesc{
	Name:    "Scanner",
	PkgPath: "v.io/v23/services/discovery",
	Doc:     "// Scanner is the interface for scanning services.",
	Methods: []rpc.MethodDesc{
		{
			Name: "Scan",
			Doc:  "// Scan scans services that match the query and returns the stream of discovered\n// services. Scanning will continue until the client cancels the call.\n//\n// The query is a WHERE expression of syncQL query against scanned services, where\n// keys are InstanceUuids and values are Service.\n//\n// Examples\n//\n//    v.InstanceName = \"v.io/i\"\n//    v.InstanceName = \"v.io/i\" AND v.Attrs[\"a\"] = \"v\"\n//    v.Attrs[\"a\"] = \"v1\" OR v.Attrs[\"a\"] = \"v2\"\n//\n// SyncQL tutorial at:\n//    https://github.com/vanadium/docs/blob/master/tutorials/syncql-tutorial.md",
			InArgs: []rpc.ArgDesc{
				{"query", ``}, // string
			},
			Tags: []*vdl.Value{vdl.ValueOf(access.Tag("Read"))},
		},
	},
}

// ScannerScanServerStream is the server stream for Scanner.Scan.
type ScannerScanServerStream interface {
	// SendStream returns the send side of the Scanner.Scan server stream.
	SendStream() interface {
		// Send places the item onto the output stream.  Returns errors encountered
		// while sending.  Blocks if there is no buffer space; will unblock when
		// buffer space is available.
		Send(item discovery.Update) error
	}
}

// ScannerScanServerCall represents the context passed to Scanner.Scan.
type ScannerScanServerCall interface {
	rpc.ServerCall
	ScannerScanServerStream
}

// ScannerScanServerCallStub is a wrapper that converts rpc.StreamServerCall into
// a typesafe stub that implements ScannerScanServerCall.
type ScannerScanServerCallStub struct {
	rpc.StreamServerCall
}

// Init initializes ScannerScanServerCallStub from rpc.StreamServerCall.
func (s *ScannerScanServerCallStub) Init(call rpc.StreamServerCall) {
	s.StreamServerCall = call
}

// SendStream returns the send side of the Scanner.Scan server stream.
func (s *ScannerScanServerCallStub) SendStream() interface {
	Send(item discovery.Update) error
} {
	return implScannerScanServerCallSend{s}
}

type implScannerScanServerCallSend struct {
	s *ScannerScanServerCallStub
}

func (s implScannerScanServerCallSend) Send(item discovery.Update) error {
	return s.s.Send(item)
}

// DiscoveryClientMethods is the client interface
// containing Discovery methods.
//
// Discovery is the interface for discovery operations.
type DiscoveryClientMethods interface {
	// Advertiser is the interface for advertising services.
	AdvertiserClientMethods
	// Scanner is the interface for scanning services.
	ScannerClientMethods
}

// DiscoveryClientStub adds universal methods to DiscoveryClientMethods.
type DiscoveryClientStub interface {
	DiscoveryClientMethods
	rpc.UniversalServiceMethods
}

// DiscoveryClient returns a client stub for Discovery.
func DiscoveryClient(name string) DiscoveryClientStub {
	return implDiscoveryClientStub{name, AdvertiserClient(name), ScannerClient(name)}
}

type implDiscoveryClientStub struct {
	name string

	AdvertiserClientStub
	ScannerClientStub
}

// DiscoveryServerMethods is the interface a server writer
// implements for Discovery.
//
// Discovery is the interface for discovery operations.
type DiscoveryServerMethods interface {
	// Advertiser is the interface for advertising services.
	AdvertiserServerMethods
	// Scanner is the interface for scanning services.
	ScannerServerMethods
}

// DiscoveryServerStubMethods is the server interface containing
// Discovery methods, as expected by rpc.Server.
// The only difference between this interface and DiscoveryServerMethods
// is the streaming methods.
type DiscoveryServerStubMethods interface {
	// Advertiser is the interface for advertising services.
	AdvertiserServerStubMethods
	// Scanner is the interface for scanning services.
	ScannerServerStubMethods
}

// DiscoveryServerStub adds universal methods to DiscoveryServerStubMethods.
type DiscoveryServerStub interface {
	DiscoveryServerStubMethods
	// Describe the Discovery interfaces.
	Describe__() []rpc.InterfaceDesc
}

// DiscoveryServer returns a server stub for Discovery.
// It converts an implementation of DiscoveryServerMethods into
// an object that may be used by rpc.Server.
func DiscoveryServer(impl DiscoveryServerMethods) DiscoveryServerStub {
	stub := implDiscoveryServerStub{
		impl:                 impl,
		AdvertiserServerStub: AdvertiserServer(impl),
		ScannerServerStub:    ScannerServer(impl),
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

type implDiscoveryServerStub struct {
	impl DiscoveryServerMethods
	AdvertiserServerStub
	ScannerServerStub
	gs *rpc.GlobState
}

func (s implDiscoveryServerStub) Globber() *rpc.GlobState {
	return s.gs
}

func (s implDiscoveryServerStub) Describe__() []rpc.InterfaceDesc {
	return []rpc.InterfaceDesc{DiscoveryDesc, AdvertiserDesc, ScannerDesc}
}

// DiscoveryDesc describes the Discovery interface.
var DiscoveryDesc rpc.InterfaceDesc = descDiscovery

// descDiscovery hides the desc to keep godoc clean.
var descDiscovery = rpc.InterfaceDesc{
	Name:    "Discovery",
	PkgPath: "v.io/v23/services/discovery",
	Doc:     "// Discovery is the interface for discovery operations.",
	Embeds: []rpc.EmbedDesc{
		{"Advertiser", "v.io/v23/services/discovery", "// Advertiser is the interface for advertising services."},
		{"Scanner", "v.io/v23/services/discovery", "// Scanner is the interface for scanning services."},
	},
}
