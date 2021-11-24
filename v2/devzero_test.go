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

func TestNewDevZeroWorks(t *testing.T) {
	t.Parallel()

	// ----------------------------------------------------------------
	// setup your test

	// ----------------------------------------------------------------
	// perform the change

	unit := NewDevZero()

	// ----------------------------------------------------------------
	// test the results

	assert.NotNil(t, unit)
}

// ================================================================
//
// Interface compatibility
//
// ----------------------------------------------------------------

func TestDevZeroImplementsIOReader(t *testing.T) {
	t.Parallel()

	// ----------------------------------------------------------------
	// setup your test

	unit := NewDevZero()
	var i interface{} = unit

	// ----------------------------------------------------------------
	// perform the change

	_, ok := i.(io.Reader)

	// ----------------------------------------------------------------
	// test the results

	assert.True(t, ok)
}

func TestDevZeroImplementsIOWriter(t *testing.T) {
	t.Parallel()

	// ----------------------------------------------------------------
	// setup your test

	unit := NewDevZero()
	var i interface{} = unit

	// ----------------------------------------------------------------
	// perform the change

	_, ok := i.(io.Writer)

	// ----------------------------------------------------------------
	// test the results

	assert.True(t, ok)
}

func TestDevZeroImplementsIOCloser(t *testing.T) {
	t.Parallel()

	// ----------------------------------------------------------------
	// setup your test

	unit := NewDevZero()
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

func TestDevZeroReadZerosTheInputBuffer(t *testing.T) {
	t.Parallel()

	// ----------------------------------------------------------------
	// setup your test

	unit := NewDevZero()

	expectedOutput := make([]byte, 10)

	// ----------------------------------------------------------------
	// perform the change

	actualOutput := make([]byte, 10)
	for i := range actualOutput {
		actualOutput[i] = 255
	}
	unit.Read(actualOutput)

	// ----------------------------------------------------------------
	// test the results

	assert.Equal(t, expectedOutput, actualOutput)
}

// ================================================================
//
// io.Closer interface
//
// ----------------------------------------------------------------

func TestDevZeroCloseShutsDownAdditionalReads(t *testing.T) {
	// although DevZero.Close() is inherited from DevNull.Close(),
	// we still need to prove that DevZero.Read() honours any
	// calls to DevZero.Close()
	t.Parallel()

	// ----------------------------------------------------------------
	// setup your test

	unit := NewDevZero()

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
