package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/xackery/w32"
	"strings"
	"time"
	"unsafe"
)

type Player struct {
	SitState      uint16 //37F7E868
	UnknownChunk1 [422]byte
	FirstName     [65]byte  //37F7EA10 //+424
	LastName      [41]byte  //37F7EA51 //65
	UnknownChunk2 [246]byte //
	CurHP         uint32    //37F7EB70
	MaxHP         uint32    //37F7EB74
	CurSP         uint16    //37F7EB78
	MaxSP         uint16    //37F7EB7A
}

func (p *Player) GetFirstName() string {
	return strings.TrimSpace(string(p.FirstName[:]))
}
func (p *Player) GetLastName() string {
	return strings.TrimSpace(string(p.LastName[:]))
}

func GetPlayerMemory(handle w32.HANDLE, addr string) (player Player) {
	memoryAddr, err := w32.HexToUint32(addr)
	if err != nil {
		fmt.Println("Error with memory address:", err.Error())
		return
	}

	//fmt.Println("Grabbing size:", unsafe.Sizeof(player))
	data, err := w32.ReadProcessMemory(handle, memoryAddr, uint(unsafe.Sizeof(player)))
	if err != nil {
		fmt.Println("Error Reading memory:", err.Error())
		return
	} else {
		fmt.Println("Success! Read memory:", data)
	}

	buffer := bytes.NewBuffer(data)
	binary.Read(buffer, binary.LittleEndian, &player)
	return
}

//This is an example to show how to read process memory and convert it to a struct
func main() {
	for {
		processAddr, err := w32.HexToUint32("00BEF28")
		if err != nil {
			fmt.Println("Error with process address:", err.Error())
			return
		}

		handle, err := w32.OpenProcess(w32.PROCESS_VM_READ, false, processAddr)
		if err != nil {
			fmt.Println("Error opening process:", err.Error())
			return
		}
		player := GetPlayerMemory(handle, "37F7E868")
		fmt.Println("State:", player.SitState, "Name:", player.GetFirstName(), "LastName:", player.GetLastName(), "HP:", player.CurHP, "/", player.MaxHP, "SP:", player.CurSP, "/", player.MaxSP)
		time.Sleep(1 * time.Second)
	}
}
