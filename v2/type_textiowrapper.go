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

// TextIOWrapper adds TextReader / TextWriter support to anything that
// supports io.ReadWriteCloser.
type TextIOWrapper struct {
	io.ReadWriteCloser
}

// ================================================================
//
// Constructors
//
// ----------------------------------------------------------------

// NewTextIOWrapper wraps your io.ReadWriteCloser with full support
// for the TextReader / TextWriter interfaces.
func NewTextIOWrapper(i io.ReadWriteCloser) *TextIOWrapper {
	retval := TextIOWrapper{i}

	// all done
	return &retval
}

// ================================================================
//
// TextReader
//
// ----------------------------------------------------------------

// ParseInt returns the next line in our underlying io.Reader as an
// integer.
//
// If the underlying source contains anything other than a valid number,
// an error is returned.
func (d *TextIOWrapper) ParseInt() (int, error) {
	return ParseInt(d)
}

// ReadLine returns the next line from our underlying io.Reader.
func (d *TextIOWrapper) ReadLine() (string, error) {
	return ReadLine(d)
}

// ReadLines returns a channel that you can `range` over to get each
// remaining line from our underlying io.Reader.
func (d *TextIOWrapper) ReadLines() <-chan string {
	return ReadLines(d)
}

// ReadWords returns a channel that you can `range` over to get each
// remaining word from our underlying io.Reader.
func (d *TextIOWrapper) ReadWords() <-chan string {
	return ReadWords(d)
}

// String returns all the remaining data in our underlying io.Reader
// as a single string.
func (d *TextIOWrapper) String() string {
	return String(d)
}

// Strings returns all of the remaining data in our underlying io.Reader
// as an array of strings, one line per array entry.
func (d *TextIOWrapper) Strings() []string {
	return Strings(d)
}

// TrimmedString returns all of the remaining data in our underlying
// input source as a string, with any leading or trailing whitespace
// removed.
func (d *TextIOWrapper) TrimmedString() string {
	return TrimmedString(d)
}

// ================================================================
//
// TextWriter
//
// ----------------------------------------------------------------

// WriteRune writes a single rune (a unicode character) to the underlying
// io.Writer. It returns the number of types written, and any error
// encountered that caused the write to fail.
func (d *TextIOWrapper) WriteRune(r rune) (int, error) {
	return WriteRune(d, r)
}

// WriteString writes a string to the underlying io.Writer. It returns
// the number of bytes written, and any error encountered that may
// have caused the write to fail.
func (d *TextIOWrapper) WriteString(s string) (int, error) {
	return WriteString(d, s)
}
