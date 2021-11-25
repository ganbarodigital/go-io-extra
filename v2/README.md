# Welcome to ioextra!

## Introduction

_ioextra_ is a package that adds helpful Golang io interfaces, structs and utilities.

## At A Glance

### Interfaces

Interface          | Purpose
-------------------|---------
`RuneWriter`       | Represents an output source that accepts unicode characters.
`TextReader`       | Represents a text-oriented input source, such as stdin.
`TextWriter`       | Represents a text-oriented output source, such as stdout / stderr.
`TextReaderWriter` | Represents a text-oriented input & output source.

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
`LogFatalf`        | How this package logs fatal errors.