# Welcome to ioextra!

## Introduction

_ioextra_ is a library that adds helpful Golang io interfaces, structs and utilities.

## At A Glance

### Interfaces

Interface          | Purpose
-------------------|---------
`TextReader`       | Represents a text-oriented input source, such as stdin.
`TextWriter`       | Represents a text-oriented output source, such as stdout / stderr.
`TextReaderWriter` | Represents a text-oriented input & output source.

### Utilities

Utility            | Purpose
-------------------|--------
`NewTextScanner()` | Creates a text-oriented input channel.