package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	for index, value := range arguments {
		if index >= 1 {
			for j := 0; j < len(value); j++ {
				picto := []rune(value)
				z01.PrintRune(picto[j])
			}
			z01.PrintRune('\n')
		}
	}
}
