package main

import (
	"fmt"
	"github.com/xackery/w32"
)

func main() {
	//To make this work, you need to find a process ID and a memory address to grab from
	processHex := "006B0E0"
	memoryHex := "00287B08"

	//Convert hex address to uint32
	processAddr, err := w32.HexToUint32(processHex)
	if err != nil {
		fmt.Println("Error with process address:", err.Error())
		return
	}

	//Open a process handle for reading
	handle, err := w32.OpenProcess(w32.PROCESS_VM_READ, false, processAddr)
	if err != nil {
		fmt.Println("Error opening process:", err.Error())
		return
	}

	//Convert hex address to uint32
	memoryAddr, err := w32.HexToUint32(memoryHex)
	if err != nil {
		fmt.Println("Error with memory address:", err.Error())
		return
	}

	//Read memory and place into data
	data, err := w32.ReadProcessMemoryAsUint32(handle, memoryAddr)
	if err != nil {
		fmt.Println("Error Reading memory:", err.Error())
		return
	} else {
		fmt.Println("Success! Read memory:", data)
	}
}
