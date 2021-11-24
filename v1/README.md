# Welcome to ioextra!

## Introduction

_ioextra_ is a library that adds helpful Golang io interfaces, structs and utilities.

## At A Glance

### Interfaces

Interface         | Purpose
------------------|---------
`TextInput`       | Represents a text-oriented input source, such as stdin.
`TextOutput`      | Represents a text-oriented output source, such as stdout / stderr.
`TextInputOutput` | Represents a text-oriented input & output source.

### Utilities

Utility            | Purpose
-------------------|--------
`NewTextScanner()` | Creates a text-oriented input channel.