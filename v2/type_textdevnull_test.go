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
	"testing"

	"github.com/stretchr/testify/assert"
)

// ================================================================
//
// Constructors
//
// ----------------------------------------------------------------

func TestNewTextDevNullWorks(t *testing.T) {
	t.Parallel()

	// ----------------------------------------------------------------
	// setup your test

	// ----------------------------------------------------------------
	// perform the change

	unit := NewTextDevNull()

	// ----------------------------------------------------------------
	// test the results

	assert.NotNil(t, unit)
}

// ================================================================
//
// Interface compatibility
//
// ----------------------------------------------------------------

func TestTextDevNullImplementsTextReader(t *testing.T) {
	t.Parallel()

	// ----------------------------------------------------------------
	// setup your test

	// this will not compile if our TextDevNull is not compatible
	// with the TextReader interface
	tmp := interfaceCompatibility{
		in:  NewTextDevNull(),
		out: NewTextDevNull(),
	}
	assert.NotNil(t, tmp)

	unit := NewTextDevNull()
	var i interface{} = unit

	// ----------------------------------------------------------------
	// perform the change

	_, ok := i.(TextReader)

	// ----------------------------------------------------------------
	// test the results

	assert.True(t, ok)
}

func TestTextDevNullImplementsTextWriter(t *testing.T) {
	t.Parallel()

	// ----------------------------------------------------------------
	// setup your test

	// this will not compile if our TextDevNull is not compatible
	// with the TextWriter interface
	tmp := interfaceCompatibility{
		in:  NewTextDevNull(),
		out: NewTextDevNull(),
	}
	assert.NotNil(t, tmp)

	unit := NewTextDevNull()
	var i interface{} = unit

	// ----------------------------------------------------------------
	// perform the change

	_, ok := i.(TextWriter)

	// ----------------------------------------------------------------
	// test the results

	assert.True(t, ok)
}

func TestTextDevNullImplementsTextReaderWriter(t *testing.T) {
	t.Parallel()

	// ----------------------------------------------------------------
	// setup your test

	unit := NewTextDevNull()
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

func TestTextDevNullParseIntAlwaysReturnsEOF(t *testing.T) {
	t.Parallel()

	// ----------------------------------------------------------------
	// setup your test

	testData := "100\n"
	expectedOutput := 0

	unit := NewTextDevNull()
	unit.WriteString(testData)

	// ----------------------------------------------------------------
	// perform the change

	actualOutput, err := unit.ParseInt()

	// ----------------------------------------------------------------
	// test the results

	assert.Equal(t, io.EOF, err)
	assert.Equal(t, expectedOutput, actualOutput)
}

func TestTextDevNullReadLineAlwaysReturnsEOF(t *testing.T) {
	t.Parallel()

	// ----------------------------------------------------------------
	// setup your test

	unit := NewTextDevNull()

	expectedResult := ""

	// ----------------------------------------------------------------
	// perform the change

	actualResult, err := unit.ReadLine()

	// ----------------------------------------------------------------
	// test the results

	assert.Equal(t, io.EOF, err)
	assert.Equal(t, expectedResult, actualResult)
}

func TestTextDevNullReadLinesAlwaysReturnsEmptySlice(t *testing.T) {
	t.Parallel()

	// ----------------------------------------------------------------
	// setup your test

	unit := NewTextDevNull()
	unit.WriteString("hello world\nhave a nice day")

	expectedResult := []string{}

	// ----------------------------------------------------------------
	// perform the change

	actualResult := []string{}
	for line := range unit.ReadLines() {
		actualResult = append(actualResult, line)
	}

	// ----------------------------------------------------------------
	// test the results

	assert.Equal(t, expectedResult, actualResult)
}

func TestTextDevNullReadWordsReturnsEmptySlice(t *testing.T) {
	t.Parallel()

	// ----------------------------------------------------------------
	// setup your test

	unit := NewTextDevNull()
	unit.WriteString("hello world\nhave a nice day")

	expectedResult := []string{}

	// ----------------------------------------------------------------
	// perform the change

	actualResult := []string{}
	for word := range unit.ReadWords() {
		actualResult = append(actualResult, word)
	}

	// ----------------------------------------------------------------
	// test the results

	assert.Equal(t, expectedResult, actualResult)
}

func TestTextDevNullStringReturnsEmptyString(t *testing.T) {
	t.Parallel()

	// ----------------------------------------------------------------
	// setup your test

	testData := "hello world\n"

	unit := NewTextDevNull()
	unit.WriteString(testData)

	expectedOutput := ""

	// ----------------------------------------------------------------
	// perform the change

	actualOutput := unit.String()

	// ----------------------------------------------------------------
	// test the results

	assert.Equal(t, expectedOutput, actualOutput)
}

func TestTextDevNullStringsReturnsEmptyString(t *testing.T) {
	t.Parallel()

	// ----------------------------------------------------------------
	// setup your test

	unit := NewTextDevNull()
	unit.WriteString("hello world\nhave a nice day\n")

	expectedOutput := []string{}

	// ----------------------------------------------------------------
	// perform the change

	actualOutput := unit.Strings()

	// ----------------------------------------------------------------
	// test the results

	assert.Equal(t, expectedOutput, actualOutput)
}

func TestTextDevNullTrimmedStringReturnsEmptyString(t *testing.T) {
	t.Parallel()

	// ----------------------------------------------------------------
	// setup your test

	unit := NewTextDevNull()
	unit.WriteString(" hello world\n")

	expectedOutput := ""

	// ----------------------------------------------------------------
	// perform the change

	actualOutput := unit.TrimmedString()

	// ----------------------------------------------------------------
	// test the results

	assert.Equal(t, expectedOutput, actualOutput)
}

// ================================================================
//
// TextWriter interface
//
// ----------------------------------------------------------------

func TestTextDevNullWriteDoesNotFail(t *testing.T) {
	t.Parallel()

	// ----------------------------------------------------------------
	// setup your test

	testData := "hello world\nhave a nice day"
	unit := NewTextDevNull()

	expectedLen := len(testData)

	// ----------------------------------------------------------------
	// perform the change

	actualLen, err := unit.Write([]byte(testData))

	// ----------------------------------------------------------------
	// test the results

	assert.Nil(t, err)
	assert.Equal(t, expectedLen, actualLen)
}

func TestTextDevNullWriteRuneDoesNotFail(t *testing.T) {
	t.Parallel()

	// ----------------------------------------------------------------
	// setup your test

	testData := rune('🙂')
	unit := NewTextDevNull()

	expectedLen := len([]byte(string(testData)))

	// ----------------------------------------------------------------
	// perform the change

	actualLen, err := unit.WriteRune(testData)

	// ----------------------------------------------------------------
	// test the results

	assert.Nil(t, err)
	assert.Equal(t, expectedLen, actualLen)
}

func TestTextDevNullWriteStringDoesNotFail(t *testing.T) {
	t.Parallel()

	// ----------------------------------------------------------------
	// setup your test

	testData := "hello world\nhave a nice day"
	unit := NewTextDevNull()

	expectedLen := len(testData)

	// ----------------------------------------------------------------
	// perform the change

	actualLen, err := unit.WriteString(testData)

	// ----------------------------------------------------------------
	// test the results

	assert.Nil(t, err)
	assert.Equal(t, expectedLen, actualLen)
}
