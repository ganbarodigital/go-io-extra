# Welcome to ioextra!

## Introduction

_ioextra_ is a package that adds helpful Golang io interfaces, structs and utilities.

## At A Glance

### Interfaces

Interface          | Purpose
-------------------|---------
`LineReader`       | Represents an input source that has the ReadLine() function.
`LinesReader`      | Represents an input source that has the ReadLines() function.
`RuneWriter`       | Represents an output source that accepts unicode characters.
`StringReader`     | Represents an input source that has the String() function.
`TextReader`       | Represents a text-oriented input source, such as stdin.
`TextWriter`       | Represents a text-oriented output source, such as stdout / stderr.
`TextReaderWriter` | Represents a text-oriented input & output source.
`WordsReader`      | Represents an input source that has the ReadWords() function.

### Structs

Struct       | Purpose
-------------|--------
`DevNull`    | An io.ReadWriteCloser that emulates UNIX /dev/null behaviour.
`DevZero`    | An io.ReadWriteCloser that emulates UNIX /dev/zero behaviour.
`TextBuffer` | A bytes.Buffer with full `TextReader` and `TextWriter` support.
`TextFile`   | An os.File with full `TextReader` and `TextWriter` support.

### Utilities

Utility            | Purpose
-------------------|--------
`NewTextScanner()` | Creates a text-oriented input channel.
`ParseInt()`       | Returns the next line from the input channel as an int.
`ReadLine()`       | Returns the next line from the input channel, as a string.
`ReadLines()`      | Returns the remaining text from the input channel, one line at a time.
`ReadWords()`      | Returns the remaining text from the input channel, one word at a time.
`LogFatalf`        | How this package logs fatal errors.