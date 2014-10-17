package build_test

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"veyron.io/veyron/veyron2/vdl/build"
	"veyron.io/veyron/veyron2/vdl/vdltest"
	"veyron.io/veyron/veyron2/vdl/vdlutil"
)

func init() {
	// Uncomment this to enable verbose logs for debugging.
	//vdlutil.SetVerbose()
}

func TestSrcDirs(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Getwd() failed: %v", err)
	}
	abs := func(relative string) string {
		return filepath.Join(cwd, relative)
	}
	tests := []struct {
		VDLPath string
		Want    []string
	}{
		{"", nil},
		// Test absolute paths.
		{"/a", []string{"/a/src"}},
		{"/a/b", []string{"/a/b/src"}},
		{"/a:/b", []string{"/a/src", "/b/src"}},
		{"/a/1:/b/2", []string{"/a/1/src", "/b/2/src"}},
		{"/a/1:/b/2:/c/3", []string{"/a/1/src", "/b/2/src", "/c/3/src"}},
		{":::/a/1::::/b/2::::/c/3:::", []string{"/a/1/src", "/b/2/src", "/c/3/src"}},
		// Test relative paths.
		{"a", []string{abs("a/src")}},
		{"a/b", []string{abs("a/b/src")}},
		{"a:b", []string{abs("a/src"), abs("b/src")}},
		{"a/1:b/2", []string{abs("a/1/src"), abs("b/2/src")}},
		{"a/1:b/2:c/3", []string{abs("a/1/src"), abs("b/2/src"), abs("c/3/src")}},
		{":::a/1::::b/2::::c/3:::", []string{abs("a/1/src"), abs("b/2/src"), abs("c/3/src")}},
		// Test mixed absolute / relative paths.
		{"a:/b", []string{abs("a/src"), "/b/src"}},
		{"/a/1:b/2", []string{"/a/1/src", abs("b/2/src")}},
		{"/a/1:b/2:/c/3", []string{"/a/1/src", abs("b/2/src"), "/c/3/src"}},
		{":::/a/1::::b/2::::/c/3:::", []string{"/a/1/src", abs("b/2/src"), "/c/3/src"}},
	}
	for _, test := range tests {
		if err := os.Setenv("VDLPATH", test.VDLPath); err != nil {
			t.Errorf("Setenv(VDLPATH, %q) failed: %v", err)
			continue
		}
		if got, want := build.SrcDirs(), test.Want; !reflect.DeepEqual(got, want) {
			t.Errorf("SrcDirs(%q) got %v, want %v", test.VDLPath, got, want)
		}
	}
}

func TestIsDirImportPath(t *testing.T) {
	tests := []struct {
		Path  string
		IsDir bool
	}{
		// Import paths.
		{"", false},
		{"...", false},
		{".../", false},
		{"all", false},
		{"foo", false},
		{"foo/", false},
		{"foo...", false},
		{"foo/...", false},
		{"a/b/c", false},
		{"a/b/c/", false},
		{"a/b/c...", false},
		{"a/b/c/...", false},
		{"...a/b/c...", false},
		{"...a/b/c/...", false},
		{".../a/b/c/...", false},
		{".../a/b/c...", false},
		// Dir paths.
		{".", true},
		{"..", true},
		{"./", true},
		{"../", true},
		{"./...", true},
		{"../...", true},
		{".././.././...", true},
		{"/", true},
		{"/.", true},
		{"/..", true},
		{"/...", true},
		{"/./...", true},
		{"/foo", true},
		{"/foo/", true},
		{"/foo...", true},
		{"/foo/...", true},
		{"/a/b/c", true},
		{"/a/b/c/", true},
		{"/a/b/c...", true},
		{"/a/b/c/...", true},
		{"/a/b/c/../../...", true},
	}
	for _, test := range tests {
		if got, want := build.IsDirPath(test.Path), test.IsDir; got != want {
			t.Errorf("IsDirPath(%q) want %v", want)
		}
		if got, want := build.IsImportPath(test.Path), !test.IsDir; got != want {
			t.Errorf("IsImportPath(%q) want %v", want)
		}
	}
}

// The cwd is set to the directory containing this file.  Currently we have the
// following directory structure:
//   .../veyron/go/src/veyron.io/veyron/veyron2/vdl/build/build_test.go
// So by backtracking a few times, we end up at the top:
//   .../veyron/go
const vdlpath = "../../../../../.."

var allModes = []build.UnknownPathMode{
	build.UnknownPathIsIgnored,
	build.UnknownPathIsError,
}

func TestTransitivePackages(t *testing.T) {
	if err := os.Setenv("VDLPATH", vdlpath); err != nil {
		t.Fatalf("Setenv(VDLPATH, %q) failed: %v", vdlpath, err)
	}
	tests := []struct {
		InPaths, OutPaths []string
	}{
		{nil, nil},
		{[]string{}, nil},
		// Single-package, both import and dir path.
		{
			[]string{"veyron.io/veyron/veyron2/vdl/testdata/base"},
			[]string{"veyron.io/veyron/veyron2/vdl/testdata/base"},
		},
		{
			[]string{"../testdata/base"},
			[]string{"veyron.io/veyron/veyron2/vdl/testdata/base"},
		},
		// Single-package with wildcard, both import and dir path.
		{
			[]string{"veyron.io/veyron/veyron2/vdl/testdata/base..."},
			[]string{"veyron.io/veyron/veyron2/vdl/testdata/base"},
		},
		{
			[]string{"veyron.io/veyron/veyron2/vdl/testdata/base/..."},
			[]string{"veyron.io/veyron/veyron2/vdl/testdata/base"},
		},
		{
			[]string{"../testdata/base..."},
			[]string{"veyron.io/veyron/veyron2/vdl/testdata/base"},
		},
		{
			[]string{"../testdata/base/..."},
			[]string{"veyron.io/veyron/veyron2/vdl/testdata/base"},
		},
		// Redundant specification as both import and dir path.
		{
			[]string{"veyron.io/veyron/veyron2/vdl/testdata/base", "../testdata/base"},
			[]string{"veyron.io/veyron/veyron2/vdl/testdata/base"},
		},
		{
			[]string{"veyron.io/veyron/veyron2/vdl/testdata/arith", "../testdata/arith"},
			[]string{
				"veyron.io/veyron/veyron2/vdl/testdata/arith/exp",
				"veyron.io/veyron/veyron2/vdl/testdata/base",
				"veyron.io/veyron/veyron2/vdl/testdata/arith",
			},
		},
		// Wildcards as both import and dir path.
		{
			[]string{"veyron.io/veyron/veyron2/vdl/testdata..."},
			[]string{
				"veyron.io/veyron/veyron2/vdl/testdata/arith/exp",
				"veyron.io/veyron/veyron2/vdl/testdata/base",
				"veyron.io/veyron/veyron2/vdl/testdata/arith",
			},
		},
		{
			[]string{"veyron.io/veyron/veyron2/vdl/testdata/..."},
			[]string{
				"veyron.io/veyron/veyron2/vdl/testdata/arith/exp",
				"veyron.io/veyron/veyron2/vdl/testdata/base",
				"veyron.io/veyron/veyron2/vdl/testdata/arith",
			},
		},
		{
			[]string{"../testdata..."},
			[]string{
				"veyron.io/veyron/veyron2/vdl/testdata/arith/exp",
				"veyron.io/veyron/veyron2/vdl/testdata/base",
				"veyron.io/veyron/veyron2/vdl/testdata/arith",
			},
		},
		{
			[]string{"../testdata/..."},
			[]string{
				"veyron.io/veyron/veyron2/vdl/testdata/arith/exp",
				"veyron.io/veyron/veyron2/vdl/testdata/base",
				"veyron.io/veyron/veyron2/vdl/testdata/arith",
			},
		},
		// Multi-Wildcards as both import and dir path.
		{
			[]string{"v...vdl/testdata/..."},
			[]string{
				"veyron.io/veyron/veyron2/vdl/testdata/arith/exp",
				"veyron.io/veyron/veyron2/vdl/testdata/base",
				"veyron.io/veyron/veyron2/vdl/testdata/arith",
			},
		},
		{
			[]string{"../../../...vdl/testdata/..."},
			[]string{
				"veyron.io/veyron/veyron2/vdl/testdata/arith/exp",
				"veyron.io/veyron/veyron2/vdl/testdata/base",
				"veyron.io/veyron/veyron2/vdl/testdata/arith",
			},
		},
		// Multi-Wildcards as both import and dir path.
		{
			[]string{"v...vdl/testdata/...exp"},
			[]string{"veyron.io/veyron/veyron2/vdl/testdata/arith/exp"},
		},
		{
			[]string{"../../../...vdl/testdata/...exp"},
			[]string{"veyron.io/veyron/veyron2/vdl/testdata/arith/exp"},
		},
	}
	exts := []string{".vdl"}
	for _, test := range tests {
		// All modes should result in the same successful output.
		for _, mode := range allModes {
			name := fmt.Sprintf("%v %v", mode, test.InPaths)
			errs := vdlutil.NewErrors(-1)
			pkgs := build.TransitivePackages(test.InPaths, exts, mode, errs)
			vdltest.ExpectResult(t, errs, name, "")
			var got []string
			for _, pkg := range pkgs {
				got = append(got, pkg.Path)
			}
			if want := []string(test.OutPaths); !reflect.DeepEqual(got, want) {
				t.Errorf("%v got %v, want %v", name, got, want)
			}
		}
	}
}

func TestTransitivePackagesUnknownPathError(t *testing.T) {
	if err := os.Setenv("VDLPATH", vdlpath); err != nil {
		t.Fatalf("Setenv(VDLPATH, %q) failed: %v", vdlpath, err)
	}
	tests := []struct {
		InPaths []string
		ErrRE   string
	}{
		// Non-existent as both import and dir path.
		{
			[]string{"noexist"},
			`Can't resolve "noexist" to any packages`,
		},
		{
			[]string{"./noexist"},
			`Can't resolve "./noexist" to any packages`,
		},
		// Invalid package path, as both import and dir path.
		{
			[]string{".foo"},
			`Import path ".foo" is invalid`,
		},
		{
			[]string{"foo/.bar"},
			`Import path "foo/.bar" is invalid`,
		},
		{
			[]string{"_foo"},
			`Import path "_foo" is invalid`,
		},
		{
			[]string{"foo/_bar"},
			`Import path "foo/_bar" is invalid`,
		},
		{
			[]string{"../../../../../.foo"},
			`package path ".foo" is invalid`,
		},
		{
			[]string{"../../../../../foo/.bar"},
			`package path "foo/.bar" is invalid`,
		},
		{
			[]string{"../../../../../_foo"},
			`package path "_foo" is invalid`,
		},
		{
			[]string{"../../../../../foo/_bar"},
			`package path "foo/_bar" is invalid`,
		},
	}
	exts := []string{".vdl"}
	for _, test := range tests {
		for _, mode := range allModes {
			name := fmt.Sprintf("%v %v", mode, test.InPaths)
			errs := vdlutil.NewErrors(-1)
			pkgs := build.TransitivePackages(test.InPaths, exts, mode, errs)
			errRE := test.ErrRE
			if mode == build.UnknownPathIsIgnored {
				// Ignore mode returns success, while error mode returns error.
				errRE = ""
			}
			vdltest.ExpectResult(t, errs, name, errRE)
			if pkgs != nil {
				t.Errorf("%v got unexpected packages %v", name, pkgs)
			}
		}
	}
}
