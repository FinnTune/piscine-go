package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	num := n
	crtslc := []rune{}
	for num > 0 {
		crtslc = append(crtslc, rune(num%10))
		num /= 10
	}
	for i := 0; i < len(crtslc); i++ {
		for j := i + 1; j < len(crtslc); j++ {
			if crtslc[i] > crtslc[j] {
				temp := crtslc[i]
				crtslc[i] = crtslc[j]
				crtslc[j] = temp
			}
		}
	}
	if len(crtslc) == 0 {
		z01.PrintRune(48)
	}
	for i := 0; i < len(crtslc); i++ {
		z01.PrintRune(crtslc[i] + 48)
	}
}
