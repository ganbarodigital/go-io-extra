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
	"bufio"
	"bytes"
	"io/ioutil"
	"strconv"
	"strings"
)

// TextBuffer is a bytes.Buffer with full TextReader / TextWriter support.
type TextBuffer struct {
	bytes.Buffer
}

// ================================================================
//
// Constructors
//
// ----------------------------------------------------------------

// NewTextBuffer creates a new bytes.Buffer that also supports
// the TextReader / TextWriter interfaces.
func NewTextBuffer() *TextBuffer {
	retval := TextBuffer{}

	// all done
	return &retval
}

// ================================================================
//
// TextReader
//
// ----------------------------------------------------------------

// ParseInt returns the data in our buffer as an integer.
//
// If the buffer contains anything other than a valid number, an error
// is returned.
func (d *TextBuffer) ParseInt() (int, error) {
	text := d.TrimmedString()
	return strconv.Atoi(text)
}

// ReadLines returns a channel that you can `range` over to get each
// line from our buffer
func (d *TextBuffer) ReadLines() <-chan string {
	return NewTextScanner(d, bufio.ScanLines)
}

// ReadWords returns a channel that you can `range` over to get each
// word from our buffer
func (d *TextBuffer) ReadWords() <-chan string {
	return NewTextScanner(d, bufio.ScanWords)
}

// String returns all the remaining data in our buffer as a single string.
func (d *TextBuffer) String() string {
	data, _ := ioutil.ReadAll(d)
	return string(data)
}

// Strings returns all of the data in our buffer as an array of
// strings, one line per array entry
func (d *TextBuffer) Strings() []string {
	retval := []string{}
	for line := range d.ReadLines() {
		retval = append(retval, line)
	}

	return retval
}

// TrimmedString returns all of the data in our buffer as a string,
// with any leading or trailing whitespace removed.
func (d *TextBuffer) TrimmedString() string {
	return strings.TrimSpace(d.String())
}
