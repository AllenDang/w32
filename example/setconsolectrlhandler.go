package main

import (
	"fmt"
	"github.com/xackery/w32"
)

func CtrlHandler(fdwCtrlType w32.DWORD) int32 {

	fmt.Println("SomeHandler triggered")
	switch fdwCtrlType {
	case w32.CTRL_C_EVENT:
		fmt.Println("CTRL+C")
	case w32.CTRL_CLOSE_EVENT:
		fmt.Println("Close")
	case w32.CTRL_BREAK_EVENT:
		fmt.Println("Break")
	case w32.CTRL_LOGOFF_EVENT:
		fmt.Println("Logoff")
	case w32.CTRL_SHUTDOWN_EVENT:
		fmt.Println("Shutdown")
	default:
		fmt.Println("Default")
	}
	return 0
}

//This code is purely experimental, it is returning with an exit stauts instead of calling my ctrlhandler.
func main() {
	err := w32.SetConsoleCtrlHandler(CtrlHandler, 1)
	if err != nil {
		fmt.Println("Error with setconsole:", err.Error())
		return
	}
	fmt.Println("Successfully Registered CtrlHandler")
	select {}
}
