package main

import (
	"fmt"
	"github.com/xackery/w32"
)

//Wait for a hotkey event to occur.
func waitForInputLoop() (err error) {
	var msg w32.MSG
	quit := false
	for !quit {
		msg, err = w32.GetMessage(0, 0, 0)
		if err != nil {
			fmt.Println("Error with GetMessage:", err.Error())
			return
		}

		switch msg.Message {
		case w32.WM_HOTKEY:
			if msg.WParam == 1 {
				fmt.Println("1 pressed")
			} else if msg.WParam == 2 {
				fmt.Println("2 pressed")
			} else if msg.WParam == 3 {
				fmt.Println("Q pressed")
				quit = true
			}
		}
	}
	return
}

//This shows how to register and unregister a hotkey
func main() {
	var err error
	err = w32.RegisterHotKey(0, 1, w32.MOD_NOREPEAT, 0x31) //Bind the 1 key to the global hotkey 1
	if err != nil {
		fmt.Println("Error registering 1:", err.Error())
		return
	}
	err = w32.RegisterHotKey(0, 2, w32.MOD_NOREPEAT, 0x32) //Bind the 2 key to the global hotkey 2
	if err != nil {
		fmt.Println("Error registering 2:", err.Error())
		return
	}
	err = w32.RegisterHotKey(0, 3, w32.MOD_NOREPEAT, 0x51) //Bind the Q key Globally to the global hotkey 3
	if err != nil {
		fmt.Println("Error registering Q:", err.Error())
		return
	}
	fmt.Println("Waiting for Hotkey of 1, 2 or Q (quit)...")
	err = waitForInputLoop()
	if err != nil {
		fmt.Println("Error registering Q:", err.Error())
		return
	}
	w32.UnregisterHotKey(0, 2) //remove the global hotkey 2, removing 2 key
	fmt.Println("Unregistered 2 key, 1 and Q (quit) should still work...")
	err = waitForInputLoop()
	if err != nil {
		fmt.Println("Error registering Q:", err.Error())
		return
	}
	fmt.Println("Exited!")
}
