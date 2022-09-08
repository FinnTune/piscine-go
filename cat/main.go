package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := len(os.Args)
	if args < 2 {
		inputString, _ := ioutil.ReadAll(os.Stdin)
		PrintRunes(string(inputString))
		// fmt.Println("Hello\nHello\n^C")
	}
	for i := 1; i < args; i++ {
		data, err := ioutil.ReadFile(os.Args[i])
		if err != nil {
			errorstring := "ERROR: "
			errorstring += err.Error()
			errorstring += "\n"
			// fmt.Printf("Error: %v\n", err.Error())
			PrintRunes(errorstring)
			os.Exit(1)
		}
		PrintRunes(string(data))
		// fmt.Println(string(data))
	}
}

func PrintRunes(str string) {
	for _, char := range str {
		z01.PrintRune(char)
	}
}
