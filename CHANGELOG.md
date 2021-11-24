# CHANGELOG

## develop

### Fixes

* TextBuffer
  - TextBuffer.String() now returns only the remaining data (ie, non-read data)

## v2.1.0

Released Wednesday, 24th November 2021.

### New

* Interfaces:
  - `TextReader` now extends `io.Reader`, for better compatibility with the wider Golang io ecosystem.
  - `TextWriter` now extends `io.Writer`, for better compatibility with the wider Golang io ecosystem.

## v2.0.1

Released Wednesday, 24th November 2021.

### New

The following items have been extracted from my `go_pipe/v5` package, and then refactored into something more sensible :)

* Interfaces:
  - Added `RuneWriter`
  - Added `TextReader`
  - Added `TextWriter`
  - Added `TextReaderWriter`
* Structs
  - Added `TextBuffer`
  - Added `TextFile`
* Utilities:
  - Added `NewTextScanner`
