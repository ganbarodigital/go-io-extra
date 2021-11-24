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

###

Struct     | Purpose
-----------|--------
`TextFile` | An os.File with full `TextReader` and `TextWriter` support.

### Utilities

Utility            | Purpose
-------------------|--------
`NewTextScanner()` | Creates a text-oriented input channel.
`LogFatalf`        | How this package logs fatal errors.