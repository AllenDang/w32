/* add_basic.go

   Demonstrates creating a DLL with an exported function, the inflexible way.
*/

package main

import "C"

//go tool cgo -gccgo add_basic.go
//TODO: Figure out how to create .o file and declspec
