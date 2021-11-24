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

func TestNewDevNullReturnsAnEmptyDevNull(t *testing.T) {
	t.Parallel()

	// ----------------------------------------------------------------
	// setup your test

	unit := NewDevNull()
	expectedResult := make([]byte, 10)

	// ----------------------------------------------------------------
	// perform the change

	actualResult := make([]byte, 10)
	unit.Read(actualResult)

	// ----------------------------------------------------------------
	// test the results

	assert.Equal(t, expectedResult, actualResult)
}

// ================================================================
//
// Interface compatibility
//
// ----------------------------------------------------------------

func TestDevNullImplementsIOReader(t *testing.T) {
	t.Parallel()

	// ----------------------------------------------------------------
	// setup your test

	unit := NewDevNull()
	var i interface{} = unit

	// ----------------------------------------------------------------
	// perform the change

	_, ok := i.(io.Reader)

	// ----------------------------------------------------------------
	// test the results

	assert.True(t, ok)
}

func TestDevNullImplementsIOWriter(t *testing.T) {
	t.Parallel()

	// ----------------------------------------------------------------
	// setup your test

	unit := NewDevNull()
	var i interface{} = unit

	// ----------------------------------------------------------------
	// perform the change

	_, ok := i.(io.Writer)

	// ----------------------------------------------------------------
	// test the results

	assert.True(t, ok)
}

func TestDevNullImplementsIOCloser(t *testing.T) {
	t.Parallel()

	// ----------------------------------------------------------------
	// setup your test

	unit := NewDevNull()
	var i interface{} = unit

	// ----------------------------------------------------------------
	// perform the change

	_, ok := i.(io.Closer)

	// ----------------------------------------------------------------
	// test the results

	assert.True(t, ok)
}

// ================================================================
//
// io.Reader interface
//
// ----------------------------------------------------------------

func TestDevNullReadReturnsNoData(t *testing.T) {
	t.Parallel()

	// ----------------------------------------------------------------
	// setup your test

	unit := NewDevNull()

	expectedOutput := make([]byte, 10)
	for i := range expectedOutput {
		expectedOutput[i] = 255
	}

	// ----------------------------------------------------------------
	// perform the change

	actualOutput := make([]byte, 10)
	for i := range actualOutput {
		actualOutput[i] = 255
	}
	actualLen, err := unit.Read(actualOutput)

	// ----------------------------------------------------------------
	// test the results

	assert.Nil(t, err)
	assert.Zero(t, actualLen)
	assert.Equal(t, expectedOutput, actualOutput)
}

// ================================================================
//
// io.Closer interface
//
// ----------------------------------------------------------------

func TestDevNullCloseShutsDownAdditionalReads(t *testing.T) {
	t.Parallel()

	// ----------------------------------------------------------------
	// setup your test

	unit := NewDevNull()

	expectedOutput := make([]byte, 10)
	for i := range expectedOutput {
		expectedOutput[i] = 255
	}

	actualOutput := make([]byte, 10)
	for i := range actualOutput {
		actualOutput[i] = 255
	}

	// ----------------------------------------------------------------
	// perform the change

	unit.Close()

	// ----------------------------------------------------------------
	// test the results

	actualLen, err := unit.Read(actualOutput)

	// prove that unit.Read() now returns errors
	assert.NotNil(t, err)

	// prove that unit.Read() still returns no bytes read
	assert.Zero(t, actualLen)

	// prove that unit.Read() still does not modify the byte slice
	assert.Equal(t, expectedOutput, actualOutput)
}

func TestDevNullCloseShutsDownAdditionalWrites(t *testing.T) {
	t.Parallel()

	// ----------------------------------------------------------------
	// setup your test

	unit := NewDevNull()

	testData := []byte("hello world!")

	// ----------------------------------------------------------------
	// perform the change

	unit.Close()

	// ----------------------------------------------------------------
	// test the results

	actualLen, err := unit.Write(testData)

	// prove that unit.Write() now returns errors
	assert.NotNil(t, err)

	// prove that unit.Write() now returns no bytes written
	assert.Zero(t, actualLen)
}

// ================================================================
//
// io.Writer interface
//
// ----------------------------------------------------------------

func TestDevNullWritePretendsTheWriteWorked(t *testing.T) {
	t.Parallel()

	// ----------------------------------------------------------------
	// setup your test

	unit := NewDevNull()

	// we use a unicode character to prove that we're getting the
	// right number of bytes back at the end
	testData := []byte("hello world 🙂!")
	expectedLen := len(testData)

	// ----------------------------------------------------------------
	// perform the change

	actualLen, err := unit.Write(testData)

	// ----------------------------------------------------------------
	// test the results

	// prove that unit.Write() tells you that it did work
	assert.Nil(t, err)

	// prove that unit.Write() tells you that it wrote the amount of
	// data you expected it to
	assert.Equal(t, expectedLen, actualLen)
}
