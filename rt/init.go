// Package rt provides initialization of a specific instantiation of the
// runtime.
package rt

import (
	"fmt"
	"os"
	"sync"

	"veyron.io/veyron/veyron2"
	"veyron.io/veyron/veyron2/options"
	"veyron.io/veyron/veyron2/verror2"
)

type Factory func(opts ...veyron2.ROpt) (veyron2.Runtime, error)

var (
	config struct {
		sync.Mutex
		profile veyron2.Profile
		factory Factory
	}

	runtimes struct {
		sync.Mutex
		registered map[string]Factory
	}

	once       sync.Once
	verrorOnce sync.Once
	globalR    veyron2.Runtime
)

func init() {
	runtimes.registered = make(map[string]Factory)
}

// RegisterRuntime registers a runtime name and associated Factory.
// The google runtime is preregistered with the names "google" and "".
// Additional runtimes may be registered and selected via an appropriate
// profile.
func RegisterRuntime(name string, factory Factory) {
	runtimes.Lock()
	runtimes.registered[name] = factory
	runtimes.Unlock()
}

// New creates and initializes a new instance of the runtime. It should
// be used in unit tests and any situation where a single global runtime
// instance is inappropriate.
func New(opts ...veyron2.ROpt) (veyron2.Runtime, error) {
	profile, profileOpts, factory, err := configure(opts...)
	if err != nil {
		return nil, err
	}
	opts = append(profileOpts, opts...)
	r, err := factory(prependProfile(profile, opts...)...)
	if err == nil {
		verrorOnce.Do(func() {
			verror2.SetDefaultContext(r.NewContext())
		})
	}
	return r, err
}

// R returns the global Runtime instance. It can only be called after
// Init has been called.
func R() veyron2.Runtime {
	return globalR
}

// Init returns the initialized global instance of the runtime.
// Calling it multiple times will always return the result of the
// first call to Init (ignoring subsequently provided options).
// All Veyron apps should call Init as the first call in their main
// function, it will call flag.Parse internally. It will panic on
// encountering an error.
func Init(opts ...veyron2.ROpt) veyron2.Runtime {
	// TODO(cnicolaou): check that subsequent calls to Init use the same
	// or compatible options as the first one.
	once.Do(func() {
		var err error
		globalR, err = New(opts...)
		if err != nil {
			panic(fmt.Sprintf("%s: failed to initialize global runtime: %s", os.Args[0], err))
		}
	})
	return globalR
}

// RegisterProfile registers the specified Profile.
// It must be called before the Init or New functions in this package
// are called; typically it will be called by an init function. If called
// multiple times, the last call 'wins'.
func RegisterProfile(profile veyron2.Profile) {
	config.Lock()
	defer config.Unlock()
	config.profile = profile
}

func prependProfile(profile veyron2.Profile, opts ...veyron2.ROpt) []veyron2.ROpt {
	return append([]veyron2.ROpt{options.Profile{profile}}, opts...)
}

func configure(opts ...veyron2.ROpt) (veyron2.Profile, []veyron2.ROpt, Factory, error) {
	config.Lock()
	defer config.Unlock()
	for _, o := range opts {
		switch v := o.(type) {
		case options.Profile:
			// Can override a registered profile.
			config.profile = v.Profile
		}
	}
	runtimes.Lock()
	defer runtimes.Unlock()
	// Let the profile specify the runtime, use a default otherwise.
	ropts := []veyron2.ROpt{}
	name := ""
	if config.profile != nil {
		name, ropts = config.profile.Runtime()
	} else {
		name = veyron2.GoogleRuntimeName
	}
	config.factory = runtimes.registered[name]

	// We must have a factory, but not necessarily a profile.
	if config.factory == nil {
		return nil, nil, nil, fmt.Errorf("no runtime factory has been found for %q", name)
	}
	if config.profile == nil {
		return nil, nil, nil, fmt.Errorf("no profile has been registered nor specified")
	}
	return config.profile, ropts, config.factory, nil
}
