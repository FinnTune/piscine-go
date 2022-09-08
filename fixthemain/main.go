package main

import "github.com/01-edu/z01"

type Door struct {
	state bool
}

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func CloseDoor(ptrDoor *Door) {
	PrintStr("Door Closing...")
	ptrDoor.state = false
}

func OpenDoor(ptrDoor *Door) {
	PrintStr("Door Opening...")
	ptrDoor.state = true
}

func IsDoorOpen(ptrDoor *Door) bool {
	PrintStr("is the Door opened ?")
	ptrDoor.state = true
	return true
}

func IsDoorClose(ptrDoor *Door) bool {
	PrintStr("is the Door closed ?")
	ptrDoor.state = false
	return true
}

func main() {
	door := &Door{}

	OpenDoor(door)
	PrintStr("\n")
	if IsDoorClose(door) {
		PrintStr("\n")
	}
	if IsDoorOpen(door) {
		PrintStr("\n")
		CloseDoor(door)
		PrintStr("\n")
	}
	if door.state {
		PrintStr("\n")
		CloseDoor(door)
		PrintStr("\n")
	}
}
