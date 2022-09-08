package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func PrintNbrInOrder(n int) {
	num := n
	crtslc := []rune{}
	for num > 0 {
		crtslc = append(crtslc, rune(num%10))
		num /= 10
	}
	if len(crtslc) == 0 {
		z01.PrintRune(48)
	}
	for i := len(crtslc) - 1; i >= 0; i-- {
		z01.PrintRune(crtslc[i] + 48)
	}
}

func PrintStr(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
}

func main() {
	points := &point{}
	s1 := "x = "
	s2 := ", y = "
	setPoint(points)

	PrintStr(s1)
	PrintNbrInOrder(points.x)
	PrintStr(s2)
	PrintNbrInOrder(points.y)
	z01.PrintRune('\n')
}
