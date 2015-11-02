About w32
==========

w32 is a wrapper of windows apis for the Go Programming Language.

It wraps win32 apis to "Go style" to make them easier to use.

Example
=====
```
package main

import (
	"github.com/xackery/w32"
)

func main() {
	w32.MessageBox(0, "Hello World!", "Hello, World!", 0)
}
```

Setup
=====

1. [Install Go](https://golang.org/dl/). I recommend 32bit aka i386 due to GCC 64bit issues on windows.
2. Get a GCC compiler. I recommend the [WinBuilds version](http://win-builds.org/doku.php/download_and_installation_from_windows).
3. In command line, type `go get github.com/xackery/w32`
4. Create a new file, and try the example above.

Contribute
==========

Contributions in form of design, code, documentation, bug reporting or other
ways you see fit are very welcome.

Thank You!
