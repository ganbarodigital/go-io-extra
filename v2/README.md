# Welcome to ioextra!

## Introduction

_ioextra_ is a package that adds helpful Golang io interfaces, structs and utilities.

## At A Glance

### Interfaces

Read Interface        | Purpose
----------------------|---------
`LineReader`          | Represents an input source that has the ReadLine() function.
`LinesReader`         | Represents an input source that has the ReadLines() function.
`StringReader`        | Represents an input source that has the String() function.
`StringsReader`       | Represents an input source that has the Strings() function.
`TrimmedStringReader` | Represents an input source that has the TrimmedString() function.
`TextReader`          | Represents a text-oriented input source, such as stdin.
`WordsReader`         | Represents an input source that has the ReadWords() function.

Write Interface    | Purpose
-------------------|---------
`RuneWriter`       | Represents an output source that accepts unicode characters.
`TextWriter`       | Represents a text-oriented output source, such as stdout / stderr.
`TextReaderWriter` | Represents a text-oriented input & output source.

### Structs

Struct          | Purpose
----------------|--------
`DevNull`       | An io.ReadWriteCloser that emulates UNIX /dev/null behaviour.
`DevZero`       | An io.ReadWriteCloser that emulates UNIX /dev/zero behaviour.
`TextBuffer`    | A bytes.Buffer with full `TextReader` and `TextWriter` support.
`TextFile`      | An os.File with full `TextReader` and `TextWriter` support.
`TextIOWrapper` | An io.ReadWriteCloser with full `TextReader` and `TextWriter` support.

### Utilities

Utility                | Purpose
-----------------------|--------
`NewTextScanner()`     | Creates a text-oriented input channel.
`NopReadWriteCloser()` | Adds io.Closer compatibility to an io.ReadWriter
`ParseInt()`           | Returns the next line from the input channel as an int.
`ReadLine()`           | Returns the next line from the input channel, as a string.
`ReadLines()`          | Returns the remaining text from the input channel, one line at a time.
`ReadWords()`          | Returns the remaining text from the input channel, one word at a time.
`String()`             | Returns the remaining text from the input channel, as a string.
`Strings()`            | Returns the remaining text from the input channel, as an array of strings.
`TrimmedString()`      | Returns the remaining text from the input channel, as a string with leading/trailing whitespace removed.
`WriteRune()`          | Writes a unicode character to the output channel.
`WriteString()`        | Writes the given string to the output channel.
`LogFatalf`            | How this package logs fatal errors.