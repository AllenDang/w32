About w32
==========

w32 is a wrapper for a number of Windows APIs for the [Go Programming Language](https://golang.org/).

This library has no other dependencies and you need no C-compiler, it is written in pure Go. Some function signatures have been adapted to make them work in Go, e.g. functions that take a pointer to an array and a count of items in C will now simply accept a slice in Go.

If you miss any Windows API functions, just write an issue or create a pull request, the Win32 API is very large and has many rarely used functions in it so this wrapper is supposed to grow as needed by the Go community.

Installation
============

Get the latest version with:

	go get github.com/gonutz/w32/v2

Version 1
=========

This repository has moved to use Go modules. It is recommended that you not use version 1 (with import path `github.com/gonutz/w32`) as it has known bugs. Use the latest version (v2 as of this writing) for new projects and migrate to a new version in case you still use v1.
