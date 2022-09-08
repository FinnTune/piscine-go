package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:][0]
	for _, v := range args {
		for i := v - 'a' + 1; i >= 0; i-- {
			z01.PrintRune(v)
		}
	}
}
