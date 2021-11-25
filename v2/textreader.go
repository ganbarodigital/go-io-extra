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

import "io"

// TextReader represents an input source that provides blocks of text
// for us to work with.
//
// It was originally designed to represent a UNIX process's stdin.
type TextReader interface {
	// io.Reader support gives us compatibility with the wider
	// Golang io ecosystem
	io.Reader

	// ParseInt returns the remaining data in this input source as an integer.
	//
	// If the input source contains anything other than a valid number, an
	// error is returned.
	ParseInt() (int, error)

	// ReadLine returns the next line of data from this input source, or
	// an error if that wasn't possible.
	ReadLine() (string, error)

	// ReadLines returns a channel that you can `range` over to get each
	// remaining line from this input source.
	ReadLines() <-chan string

	// ReadWords returns a channel that you can `range` over to get each
	// remaining word from this input source.
	ReadWords() <-chan string

	// String returns all of the remaining data in this input source as a
	// single (possibly multi-line) string.
	String() string

	// Strings returns all of the remaining data in this input source as an
	// array of strings, one line per array entry.
	Strings() []string

	// TrimmedString returns all of the data in this input source as a single
	// (possibly multi-line) string. Any leading and trailing whitespace
	// is removed.
	TrimmedString() string
}
