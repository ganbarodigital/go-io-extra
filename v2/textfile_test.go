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
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// createTestFile is a helper function. It gives us a file that we can
// test against
func createTestFile(content string) *os.File {
	tmpFile, err := ioutil.TempFile(os.TempDir(), "ioextra-textfile-*")
	if err != nil {
		log.Fatal(err)
	}

	// clean up after ourselves
	defer os.Remove(tmpFile.Name())

	// write the content into our new file
	tmpFile.WriteString(content)
	tmpFile.Seek(0, 0)

	// all done
	return tmpFile
}

// ================================================================
//
// Constructors
//
// ----------------------------------------------------------------

func TestNewTextFileCreatesATextFile(t *testing.T) {
	t.Parallel()

	// ----------------------------------------------------------------
	// setup your test

	testData := "hello world\nhave a nice day"

	// ----------------------------------------------------------------
	// perform the change

	dest := NewTextFile(
		createTestFile(testData),
	)

	// ----------------------------------------------------------------
	// test the results

	assert.NotNil(t, dest)
}

// ================================================================
//
// Interface compatibility
//
// ----------------------------------------------------------------

func TestTextFileImplementsTextReader(t *testing.T) {
	t.Parallel()

	// ----------------------------------------------------------------
	// setup your test

	unit := TextFile{}
	var i interface{} = &unit

	// ----------------------------------------------------------------
	// perform the change

	_, ok := i.(TextReader)

	// ----------------------------------------------------------------
	// test the results

	assert.True(t, ok)
}

func TestTextFileImplementsTextWriter(t *testing.T) {
	t.Parallel()

	// ----------------------------------------------------------------
	// setup your test

	unit := TextFile{}
	var i interface{} = &unit

	// ----------------------------------------------------------------
	// perform the change

	_, ok := i.(TextWriter)

	// ----------------------------------------------------------------
	// test the results

	assert.True(t, ok)
}

func TestTextFileImplementsTextReaderWriter(t *testing.T) {
	t.Parallel()

	// ----------------------------------------------------------------
	// setup your test

	unit := TextFile{}
	var i interface{} = &unit

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

func TestTextFileReadFetchesBytesFromTheUnderlyingFile(t *testing.T) {
	t.Parallel()

	// ----------------------------------------------------------------
	// setup your test

	testData := "hello world\nhave a nice day"
	unit := NewTextFile(
		createTestFile(testData),
	)

	expectedResult := []byte(testData)
	expectedLen := len(testData)

	// ----------------------------------------------------------------
	// perform the change

	actualResult := make([]byte, expectedLen)
	actualLen, err := unit.Read(actualResult)

	// ----------------------------------------------------------------
	// test the results

	assert.Nil(t, err)
	assert.Equal(t, expectedLen, actualLen)
	assert.Equal(t, expectedResult, actualResult)
}

func TestTextFileCloseClosesTheUnderlyingFile(t *testing.T) {
	t.Parallel()

	// ----------------------------------------------------------------
	// setup your test

	testData := "hello world\nhave a nice day"
	unit := NewTextFile(
		createTestFile(testData),
	)

	// ----------------------------------------------------------------
	// perform the change

	unit.Close()

	// we should not be able to read from this buffer any more
	actualResult := make([]byte, len(testData))
	actualLen, err := unit.Read(actualResult)

	// ----------------------------------------------------------------
	// test the results

	assert.NotNil(t, err)
	assert.Equal(t, 0, actualLen)
}

func TestTextFileParseIntReturnsValueOnSuccess(t *testing.T) {
	t.Parallel()

	// ----------------------------------------------------------------
	// setup your test

	testData := " 100 \n"
	unit := NewTextFile(
		createTestFile(testData),
	)

	expectedOutput := 100

	// ----------------------------------------------------------------
	// perform the change

	actualOutput, err := unit.ParseInt()

	// ----------------------------------------------------------------
	// test the results

	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, actualOutput)
}

func TestTextFileParseIntReturnsErrorOnParseFailure(t *testing.T) {
	t.Parallel()

	// ----------------------------------------------------------------
	// setup your test

	testData := " one hundred \n"
	unit := NewTextFile(
		createTestFile(testData),
	)

	expectedOutput := 0

	// ----------------------------------------------------------------
	// perform the change

	actualOutput, err := unit.ParseInt()

	// ----------------------------------------------------------------
	// test the results

	assert.NotNil(t, err)
	assert.Equal(t, expectedOutput, actualOutput)
}

func TestTextFileParseIntReturnsErrorWhenFileIsClosed(t *testing.T) {
	t.Parallel()

	// ----------------------------------------------------------------
	// setup your test

	testData := " 100 \n"
	unit := NewTextFile(
		createTestFile(testData),
	)
	unit.Close()

	expectedOutput := 0

	// ----------------------------------------------------------------
	// perform the change

	actualOutput, err := unit.ParseInt()

	// ----------------------------------------------------------------
	// test the results

	assert.NotNil(t, err)
	assert.Equal(t, expectedOutput, actualOutput)
}

func TestTextFileReadLinesIteratesOverUnderlyingFile(t *testing.T) {
	t.Parallel()

	// ----------------------------------------------------------------
	// setup your test

	testData := "hello world\nhave a nice day"
	unit := NewTextFile(
		createTestFile(testData),
	)

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

func TestTextFileReadLinesActsLikeStream(t *testing.T) {
	t.Parallel()

	// ----------------------------------------------------------------
	// setup your test

	testData := "hello world\nhave a nice day"
	unit := NewTextFile(
		createTestFile(testData),
	)

	expectedResult := []string{"hello world", "have a nice day"}

	// ----------------------------------------------------------------
	// perform the change

	// this first read moves us to the end of the file
	var actualResult []string
	for line := range unit.ReadLines() {
		actualResult = append(actualResult, line)
	}

	// this second read proves that ReadLines() acts like a stream,
	// and does not rewind to the beginning of the file
	secondResult := []string{}
	for line := range unit.ReadLines() {
		secondResult = append(secondResult, line)
	}

	// ----------------------------------------------------------------
	// test the results

	assert.Equal(t, expectedResult, actualResult)
	assert.Empty(t, secondResult)
}

func TestTextFileReadWordsIteratesOverTheUnderlyingFile(t *testing.T) {
	t.Parallel()

	// ----------------------------------------------------------------
	// setup your test

	testData := "hello world\nhave a nice day"
	unit := NewTextFile(
		createTestFile(testData),
	)

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

func TestTextFileReadWordsActsLikeAStream(t *testing.T) {
	t.Parallel()

	// ----------------------------------------------------------------
	// setup your test

	testData := "hello world\nhave a nice day"
	unit := NewTextFile(
		createTestFile(testData),
	)

	expectedResult := []string{"hello", "world", "have", "a", "nice", "day"}

	// ----------------------------------------------------------------
	// perform the change

	// this first read moves us to the end of the file
	var actualResult []string
	for word := range unit.ReadWords() {
		actualResult = append(actualResult, word)
	}

	// this second read proves that ReadWords() acts like a stream,
	// and does not rewind to the beginning of the file
	var secondResult []string
	for word := range unit.ReadWords() {
		secondResult = append(secondResult, word)
	}

	// ----------------------------------------------------------------
	// test the results

	assert.Equal(t, expectedResult, actualResult)
	assert.Empty(t, secondResult)
}

func TestTextFileStringReturnsTheUnderlyingFile(t *testing.T) {
	t.Parallel()

	// ----------------------------------------------------------------
	// setup your test

	testData := "hello world\nhave a nice day"
	unit := NewTextFile(
		createTestFile(testData),
	)

	expectedOutput := testData

	// ----------------------------------------------------------------
	// perform the change

	actualOutput := unit.String()

	// ----------------------------------------------------------------
	// test the results

	assert.Equal(t, expectedOutput, actualOutput)
}

func TestTextFileStringActsLikeAStream(t *testing.T) {
	t.Parallel()

	// ----------------------------------------------------------------
	// setup your test

	testData := "hello world\nhave a nice day"
	unit := NewTextFile(
		createTestFile(testData),
	)

	expectedOutput := testData

	// ----------------------------------------------------------------
	// perform the change

	// this first read moves us to the end of the file
	actualOutput := unit.String()

	// this second read proves that String() acts like a stream, and
	// does not rewind to the beginning of the file
	secondOutput := unit.String()

	// ----------------------------------------------------------------
	// test the results

	assert.Equal(t, expectedOutput, actualOutput)
	assert.Empty(t, secondOutput)
}

func TestTextFileStringsReturnsTheUnderlyingFile(t *testing.T) {
	t.Parallel()

	// ----------------------------------------------------------------
	// setup your test

	testData := "hello world\nhave a nice day"
	unit := NewTextFile(
		createTestFile(testData),
	)

	expectedOutput := []string{"hello world", "have a nice day"}

	// ----------------------------------------------------------------
	// perform the change

	actualOutput := unit.Strings()

	// ----------------------------------------------------------------
	// test the results

	assert.Equal(t, expectedOutput, actualOutput)
}

func TestTextFileStringsActsLikeAStream(t *testing.T) {
	t.Parallel()

	// ----------------------------------------------------------------
	// setup your test

	testData := "hello world\nhave a nice day"
	unit := NewTextFile(
		createTestFile(testData),
	)

	expectedOutput := []string{"hello world", "have a nice day"}

	// ----------------------------------------------------------------
	// perform the change

	// this first read moves us to the end of the file
	actualOutput := unit.Strings()

	// this second read proves that Strings() acts like a stream, and
	// does not rewind to the beginning of the file
	secondOutput := unit.Strings()

	// ----------------------------------------------------------------
	// test the results

	assert.Equal(t, expectedOutput, actualOutput)
	assert.Empty(t, secondOutput)
}

func TestTextFileTrimmedStringReturnsUnderlyingFileWithWhitespaceRemoved(t *testing.T) {
	t.Parallel()

	// ----------------------------------------------------------------
	// setup your test

	testData := " hello world\nhave a nice day\n "
	unit := NewTextFile(
		createTestFile(testData),
	)

	// NOTE: Golang treats trailing '\n' characters as white space :(
	expectedOutput := "hello world\nhave a nice day"

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

func TestTextFileWriteWritesToTheUnderlyingFile(t *testing.T) {
	t.Parallel()

	// ----------------------------------------------------------------
	// setup your test

	testData := "hello world\nhave a nice day"
	unit := NewTextFile(
		createTestFile(""),
	)

	expectedResult := testData

	// ----------------------------------------------------------------
	// perform the change

	unit.Write([]byte(testData))

	// ----------------------------------------------------------------
	// test the results

	unit.Rewind()
	actualResult := unit.String()

	assert.Equal(t, expectedResult, actualResult)
}

func TestTextFileWriteRuneWritesToTheUnderlyingFile(t *testing.T) {
	t.Parallel()

	// ----------------------------------------------------------------
	// setup your test

	testData := rune('🙂')
	unit := NewTextFile(
		createTestFile(""),
	)

	expectedResult := string(testData)

	// ----------------------------------------------------------------
	// perform the change

	unit.WriteRune(testData)

	// ----------------------------------------------------------------
	// test the results

	unit.Rewind()
	actualResult := unit.String()

	assert.Equal(t, expectedResult, actualResult)
}

func TestTextFileWriteStringWritesToTheUnderlyingFile(t *testing.T) {
	t.Parallel()

	// ----------------------------------------------------------------
	// setup your test

	testData := "hello world\nhave a nice day"
	unit := NewTextFile(
		createTestFile(""),
	)

	expectedResult := testData

	// ----------------------------------------------------------------
	// perform the change

	unit.WriteString(testData)

	// ----------------------------------------------------------------
	// test the results

	unit.Rewind()
	actualResult := unit.String()

	assert.Equal(t, expectedResult, actualResult)
}
