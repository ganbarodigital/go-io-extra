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
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// TextFile is an os.File with TextReader and TextWriter compatibility.
type TextFile struct {
	os.File
}

// ===========================================================================
//
// Constructors
//
// ---------------------------------------------------------------------------

// NewTextFile creates a new output destination that reads from / writes to
// and underlying file.
func NewTextFile(f *os.File) *TextFile {
	retval := TextFile{*f}

	// all done
	return &retval
}

// ================================================================
//
// Helpers
//
// ----------------------------------------------------------------

// Rewind moves the read/write position to the start of the underlying file.
func (d *TextFile) Rewind() error {
	_, err := d.Seek(0, 0)
	return err
}

// MustRewind logs a fatal error if the Rewind operation fails
func (d *TextFile) MustRewind() error {
	err := d.Rewind()

	// did the rewind operation succeed?
	if err == nil {
		return nil
	}

	// if we get here, then no, it did not succeed
	LogFatalf("unable to rewind, err: %v", err)
	return err
}

// ===========================================================================
//
// TextReader interface
//
// ---------------------------------------------------------------------------

// ParseInt returns the data in our underlying file as an integer.
//
// If the file contains anything other than a valid number, an error
// is returned.
func (d *TextFile) ParseInt() (int, error) {
	// make sure we are at the start of the file
	err := d.Rewind()
	if err != nil {
		return 0, err
	}

	text := d.TrimmedString()
	return strconv.Atoi(text)
}

// ReadLines returns a channel that you can `range` over to get each
// line from our underlying file.
//
// If the underlying file supports it, ReadLines will always return
// a channel that starts from the beginning of the file.
func (d *TextFile) ReadLines() <-chan string {
	// make sure we are at the start of the file
	d.Rewind()

	return NewTextScanner(d, bufio.ScanLines)
}

// ReadWords returns a channel that you can `range` over to get each
// word from our underlying file.
//
// If the underlying file supports it, ReadWords will always return
// a channel that starts from the beginning of the file.
func (d *TextFile) ReadWords() <-chan string {
	// make sure we are at the start of the file
	d.Rewind()

	return NewTextScanner(d, bufio.ScanWords)
}

// String returns all of the data in our underlying file as a single
// (possibly multi-line) string.
//
// If the underlying file does not support seeking to the beginning
// of the data (ie, it is a stream of some kind), then we log a fatal
// error.
func (d *TextFile) String() string {
	// make sure we are at the start of the file
	rewindErr := d.MustRewind()

	// because we guarantee that String() always returns all of the
	// data in the file, we cannot continue if the Rewind operation
	// failed for any reason
	if rewindErr != nil {
		// this return statement will only be reachable inside
		// our unit tests
		return ""
	}

	retval, err := ioutil.ReadAll(d)
	if err != nil {
		return ""
	}

	return string(retval)
}

// Strings returns all of the data in our underlying file as an array of
// strings, one line per array entry.
//
// If the underlying file does not support seeking to the beginning
// of the data (ie, it is a stream of some kind), then we log a fatal
// error.
func (d *TextFile) Strings() []string {
	// make sure we are at the start of the file
	rewindErr := d.MustRewind()

	// because we guarantee that Strings() always returns all of the
	// data in the file, we cannot continue if the Rewind operation
	// failed for any reason
	if rewindErr != nil {
		// this return statement will only be reachable inside
		// our unit tests
		return []string{}
	}

	retval := []string{}
	for line := range d.ReadLines() {
		retval = append(retval, line)
	}

	return retval
}

// TrimmedString returns all of the data in our underlying file as a string,
// with any leading or trailing whitespace removed.
//
// If the underlying file does not support seeking to the beginning
// of the data (ie, it is a stream of some kind), then we log a fatal
// error.
func (d *TextFile) TrimmedString() string {
	return strings.TrimSpace(d.String())
}

// ===========================================================================
//
// TextWriter interface
//
// The majority of the TextWriter interface is already handled by the
// underlying os.File.
//
// ---------------------------------------------------------------------------

// WriteRune writes a single rune (a unicode character) to the underlying
// file. It returns the number of types written, and any error encountered
// that caused the write to file.
func (d *TextFile) WriteRune(r rune) (int, error) {
	return d.WriteString(string(r))
}
