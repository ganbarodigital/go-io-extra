# CHANGELOG

## develop

### New

* Added `DevNull` struct
* Added `DevZero` struct
* Added `LineReader` interface
* Added `LinesReader` interface
* Added `StringReader` interface
* Added `StringsReader` interface
* Added `TrimmedStringReader` interface
* Added `WordsReader` interface
* Added `ParseInt()`
* Added `ReadLine()`
* Added `ReadLines()`
* Added `ReadWords()`
* Added `String()`
* Added `Strings()`
* Added `closed` flag

## v2.1.1

Released Wednesday, 24th November 2021.

### Fixes

* TextBuffer
  - TextBuffer.String() now returns only the remaining data (ie, non-read data)
* TextFile
  - TextFile now acts like a stream, no longer auto-rewinds on any reads
    - You can use `TextFile.Rewind()` if you *need* it to rewind for any reason

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
