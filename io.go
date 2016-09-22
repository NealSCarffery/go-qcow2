// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package qcow2

import "os"

// bdrvPread reads the child qcow2 image file.
// Return nil on success, err on error.
//
// NOTE: This function does not use the pread syscall.
// The function name only of compatible for qemu intelnal source.
func bdrvPread(child *BdrvChild, offset int64, res interface{}, byt uintptr) error {
	f, err := os.Open(child.Name)
	if err != nil {
		return err
	}

	if offset != 0 {
		f.Seek(offset, 0)
	}

	var buf []byte
	if _, err := f.ReadAt(buf, int64(byt)); err != nil {
		return err
	}

	res = &buf
	return nil
}
