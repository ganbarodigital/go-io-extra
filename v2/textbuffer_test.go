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
	"testing"

	"github.com/stretchr/testify/assert"
)

type interfaceCompatibility struct {
	in  TextReader
	out TextWriter
}

// ================================================================
//
// Constructors
//
// ----------------------------------------------------------------

func NewTextBufferReturnsAnEmptyTextBuffer(t *testing.T) {
	t.Parallel()

	// ----------------------------------------------------------------
	// setup your test

	unit := NewTextBuffer()
	expectedResult := ""

	// ----------------------------------------------------------------
	// perform the change

	actualResult := unit.String()

	// ----------------------------------------------------------------
	// test the results

	assert.Equal(t, expectedResult, actualResult)
}

// ================================================================
//
// Interface compatibility
//
// ----------------------------------------------------------------

func TestTextBufferImplementsTextReader(t *testing.T) {
	t.Parallel()

	// ----------------------------------------------------------------
	// setup your test

	// this will not compile if our TextBuffer is not compatible
	// with the TextReader interface
	tmp := interfaceCompatibility{
		in:  NewTextBuffer(),
		out: NewTextBuffer(),
	}
	assert.NotNil(t, tmp)

	unit := NewTextBuffer()
	var i interface{} = unit

	// ----------------------------------------------------------------
	// perform the change

	_, ok := i.(TextReader)

	// ----------------------------------------------------------------
	// test the results

	assert.True(t, ok)
}

func TestTextBufferImplementsTextWriter(t *testing.T) {
	t.Parallel()

	// ----------------------------------------------------------------
	// setup your test

	// this will not compile if our TextBuffer is not compatible
	// with the TextWriter interface
	tmp := interfaceCompatibility{
		in:  NewTextBuffer(),
		out: NewTextBuffer(),
	}
	assert.NotNil(t, tmp)

	unit := NewTextBuffer()
	var i interface{} = unit

	// ----------------------------------------------------------------
	// perform the change

	_, ok := i.(TextWriter)

	// ----------------------------------------------------------------
	// test the results

	assert.True(t, ok)
}

func TestTextBufferImplementsTextReaderWriter(t *testing.T) {
	t.Parallel()

	// ----------------------------------------------------------------
	// setup your test

	unit := NewTextBuffer()
	var i interface{} = unit

	// ----------------------------------------------------------------
	// perform the change

	_, ok := i.(TextReaderWriter)

	// ----------------------------------------------------------------
	// test the results

	assert.True(t, ok)
}

// ================================================================
//
// TextReader interface
//
// ----------------------------------------------------------------

func TestTextBufferParseIntReturnsValueOnSuccess(t *testing.T) {
	t.Parallel()

	// ----------------------------------------------------------------
	// setup your test

	testData := " 100 \n"
	expectedOutput := 100

	dest := NewTextBuffer()
	dest.WriteString(testData)

	// ----------------------------------------------------------------
	// perform the change

	actualOutput, err := dest.ParseInt()

	// ----------------------------------------------------------------
	// test the results

	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, actualOutput)
}

func TestTextBufferReadLinesIteratesOverBuffer(t *testing.T) {
	t.Parallel()

	// ----------------------------------------------------------------
	// setup your test

	unit := NewTextBuffer()
	unit.WriteString("hello world\nhave a nice day")

	expectedResult := []string{"hello world", "have a nice day"}

	// ----------------------------------------------------------------
	// perform the change

	var actualResult []string
	for line := range unit.ReadLines() {
		actualResult = append(actualResult, line)
	}

	// ----------------------------------------------------------------
	// test the results

	assert.Equal(t, expectedResult, actualResult)
}

func TestTextBufferReadLinesEmptiesTheBuffer(t *testing.T) {
	t.Parallel()

	// ----------------------------------------------------------------
	// setup your test

	unit := NewTextBuffer()
	unit.WriteString("hello world\nhave a nice day")

	expectedResult := []string{"hello world", "have a nice day"}

	// ----------------------------------------------------------------
	// perform the change

	var actualResult []string
	for line := range unit.ReadLines() {
		actualResult = append(actualResult, line)
	}

	extraOutput := []string{}
	for line := range unit.ReadLines() {
		extraOutput = append(extraOutput, line)
	}

	// ----------------------------------------------------------------
	// test the results

	assert.Equal(t, expectedResult, actualResult)
	assert.Empty(t, extraOutput)
}

func TestTextBufferReadWordsIteratesOverBuffer(t *testing.T) {
	t.Parallel()

	// ----------------------------------------------------------------
	// setup your test

	unit := NewTextBuffer()
	unit.WriteString("hello world\nhave a nice day")

	expectedResult := []string{"hello", "world", "have", "a", "nice", "day"}

	// ----------------------------------------------------------------
	// perform the change

	var actualResult []string
	for word := range unit.ReadWords() {
		actualResult = append(actualResult, word)
	}

	// ----------------------------------------------------------------
	// test the results

	assert.Equal(t, expectedResult, actualResult)
}

func TestTextBufferStringReturnsBuffer(t *testing.T) {
	t.Parallel()

	// ----------------------------------------------------------------
	// setup your test

	testData := "hello world\n"

	unit := NewTextBuffer()
	unit.WriteString(testData)

	expectedOutput := testData

	// ----------------------------------------------------------------
	// perform the change

	actualOutput := unit.String()

	// ----------------------------------------------------------------
	// test the results

	assert.Equal(t, expectedOutput, actualOutput)
}

func TestTextBufferStringEmptiesTheBuffer(t *testing.T) {
	t.Parallel()

	// ----------------------------------------------------------------
	// setup your test

	testData := "hello world\n"

	unit := NewTextBuffer()
	unit.WriteString(testData)

	expectedOutput := testData

	// ----------------------------------------------------------------
	// perform the change

	actualOutput := unit.String()
	secondOutput := unit.String()

	// ----------------------------------------------------------------
	// test the results

	assert.Equal(t, expectedOutput, actualOutput)
	assert.Empty(t, secondOutput)
}

func TestTextBufferStringsReturnsBuffer(t *testing.T) {
	t.Parallel()

	// ----------------------------------------------------------------
	// setup your test

	unit := NewTextBuffer()
	unit.WriteString("hello world\nhave a nice day\n")

	expectedOutput := []string{"hello world", "have a nice day"}

	// ----------------------------------------------------------------
	// perform the change

	actualOutput := unit.Strings()

	// ----------------------------------------------------------------
	// test the results

	assert.Equal(t, expectedOutput, actualOutput)
}

func TestTextBufferStringsEmptiesTheBufferBuffer(t *testing.T) {
	t.Parallel()

	// ----------------------------------------------------------------
	// setup your test

	unit := NewTextBuffer()
	unit.WriteString("hello world\nhave a nice day\n")

	expectedOutput := []string{"hello world", "have a nice day"}

	// ----------------------------------------------------------------
	// perform the change

	actualOutput := unit.Strings()
	secondOutput := unit.String()

	// ----------------------------------------------------------------
	// test the results

	assert.Equal(t, expectedOutput, actualOutput)
	assert.Empty(t, secondOutput)
}

func TestTextBufferTrimmedStringReturnsBufferWithWhitespaceRemoved(t *testing.T) {
	t.Parallel()

	// ----------------------------------------------------------------
	// setup your test

	unit := NewTextBuffer()
	unit.WriteString(" hello world\n")

	expectedOutput := "hello world"

	// ----------------------------------------------------------------
	// perform the change

	actualOutput := unit.TrimmedString()

	// ----------------------------------------------------------------
	// test the results

	assert.Equal(t, expectedOutput, actualOutput)
}
