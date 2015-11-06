/* addtest_basic.go

   Demonstrates using the function imported from the DLL, the inelegant way.
*/
package main

import (
	"fmt"
	"syscall"
)

/* Declare imported function so that we can actually use it. */
func main() {
	addbasic := syscall.NewLazyDLL("add_basic.dll")
	procAdd := addbasic.NewProc("Add")
	val, _, _ := procAdd.Call(6, 23)

	fmt.Println(val)
}
