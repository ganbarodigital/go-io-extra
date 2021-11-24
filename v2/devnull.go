// ioextra is a library that adds helpful io stuff
//
// Copyright 2019-present Ganbaro Digital Ltd
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions
// are met:
//
//   * Redistributions of source code must retain the above copyright
//     notice, this list of conditions and the following disclaimer.
//
//   * Redistributions in binary form must reproduce the above copyright
//     notice, this list of conditions and the following disclaimer in
//     the documentation and/or other materials provided with the
//     distribution.
//
//   * Neither the names of the copyright holders nor the names of his
//     contributors may be used to endorse or promote products derived
//     from this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS
// FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE
// COPYRIGHT OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT,
// INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING,
// BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
// LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT
// LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN
// ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
// POSSIBILITY OF SUCH DAMAGE.

package ioextra

import (
	"io"
)

// DevNull emulates UNIX /dev/null behaviour, with full support as
// io.Reader, io.Closer and io.Writer.
type DevNull struct {
	flags int
}

// ================================================================
//
// Constructors
//
// ----------------------------------------------------------------

// NewDevNull creates an emulation of /dev/null.
func NewDevNull() *DevNull {
	retval := DevNull{flags: 0}

	// all done
	return &retval
}

// ================================================================
//
// io.Reader interface
//
// ----------------------------------------------------------------

// Read emulates /dev/null: it returns zero bytes.
func (d *DevNull) Read(b []byte) (int, error) {
	if d.flags&closed == 0 {
		return 0, nil
	}

	return 0, io.ErrClosedPipe
}

// ================================================================
//
// io.Closer interface
//
// ----------------------------------------------------------------

// Close prevents all further reads and writes.
func (d *DevNull) Close() error {
	d.flags |= closed
	return nil
}

// ================================================================
//
// io.Writer interface
//
// ----------------------------------------------------------------

// Write emulates /dev/null: all writes succeed (as long as you haven't
// called Close) but do nothing.
func (d *DevNull) Write(p []byte) (int, error) {
	if d.flags&closed == 0 {
		return io.Discard.Write(p)
	}
	return 0, io.ErrClosedPipe
}
