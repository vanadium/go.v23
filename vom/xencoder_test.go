// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build newvdltests

package vom_test

import (
	"bytes"
	"fmt"
	"testing"

	"v.io/v23/vdl"
	"v.io/v23/vom"
	"v.io/v23/vom/vomtest"
)

func TestXEncoder(t *testing.T) {
	for _, test := range vomtest.AllPass() {
		testXEncoder(t, "[go value]", test, test.Value.Interface())
		vv, err := vdl.ValueFromReflect(test.Value)
		if err != nil {
			t.Errorf("%s: ValueFromReflect failed: %v", test.Name(), err)
			continue
		}
		testXEncoder(t, "[vdl.Value]", test, vv)
	}
}

func testXEncoder(t *testing.T, pre string, test vomtest.Entry, value interface{}) {
	// Test vom.NewXEncoder.
	{
		var buf bytes.Buffer
		name := fmt.Sprintf("%s %s", pre, test.Name())
		enc := vom.NewVersionedXEncoder(test.Version, &buf)
		if err := enc.Encode(value); err != nil {
			t.Errorf("%s: Encode failed: %v", name, err)
			return
		}
		if got, want := buf.Bytes(), test.Bytes(); !bytes.Equal(got, want) {
			t.Errorf("%s\nGOT  %x\nWANT %x", name, got, want)
			return
		}
	}
	// Test vom.NewXEncoderWithTypeEncoder.
	{
		var buf, bufT bytes.Buffer
		name := fmt.Sprintf("%s (with TypeEncoder) %s", pre, test.Name())
		encT := vom.NewVersionedTypeEncoder(test.Version, &bufT)
		enc := vom.NewVersionedXEncoderWithTypeEncoder(test.Version, &buf, encT)
		if err := enc.Encode(value); err != nil {
			t.Errorf("%s: Encode failed: %v", name, err)
			return
		}
		if got, want := bufT.Bytes(), test.TypeBytes(); !bytes.Equal(got, want) {
			t.Errorf("%s TYPE\nGOT  %x\nWANT %x", name, got, want)
			return
		}
		if got, want := buf.Bytes(), test.ValueBytes(); !bytes.Equal(got, want) {
			t.Errorf("%s VALUE\nGOT  %x\nWANT %x", name, got, want)
			return
		}
	}
	// Test single-shot vom.XEncode.
	{
		name := fmt.Sprintf("%s (single-shot) %s", pre, test.Name())
		buf, err := vom.VersionedXEncode(test.Version, value)
		if err != nil {
			t.Errorf("%s: Encode failed: %v", name, err)
			return
		}
		if got, want := buf, test.Bytes(); !bytes.Equal(got, want) {
			t.Errorf("%s\nGOT  %x\nWANT %x", name, got, want)
			return
		}
	}
}