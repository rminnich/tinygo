// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os_test

import (
	"io"
	. "os"
	"runtime"
	"testing"
)

// localTmp returns a local temporary directory not on NFS.
func localTmp() string {
	return TempDir()
}

func newFile(testName string, t *testing.T) (f *File) {
	// TODO: use CreateTemp when it lands
	f, err := OpenFile(TempDir()+"/_Go_"+testName, O_RDWR|O_CREATE, 0644)
	if err != nil {
		t.Fatalf("TempFile %s: %s", testName, err)
	}
	return
}

func TestReadAt(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Log("TODO: implement Pread for Windows")
		return
	}
	f := newFile("TestReadAt", t)
	defer Remove(f.Name())
	defer f.Close()

	const data = "hello, world\n"
	io.WriteString(f, data)

	b := make([]byte, 5)
	n, err := f.ReadAt(b, 7)
	if err != nil || n != len(b) {
		t.Fatalf("ReadAt 7: %d, %v", n, err)
	}
	if string(b) != "world" {
		t.Fatalf("ReadAt 7: have %q want %q", string(b), "world")
	}
}