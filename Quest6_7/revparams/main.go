package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	for i := len(args) - 1; i >= 1; i-- {
		for _, value := range args[i] {
			z01.PrintRune(value)
		}
		z01.PrintRune('\n')
	}
	// picto := []rune{}
	/* for index, arg := range arguments {
		if index >= 1 {
			for j := 0; j < len(arg); j++ {
				slc_rns := []rune(arg)
				z01.PrintRune(slc_rns[j])
				// picto = append(picto, rune(value)) - wrong because string cannot become rune
				for k := len(arguments) - 1; k >= 0; k-- {
					z01.PrintRune(slc_rns[j])
				}
			}
			z01.PrintRune('\n')
		}
	} */
}
