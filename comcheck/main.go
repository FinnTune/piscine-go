package main

import (
	"os"
)

func main() {
	slc := []string{"01", "galaxy", "galaxy 01"}
	args := os.Args[1:]
	for _, val1 := range slc {
		// fmt.Printf("slc: %v\n", val1)
		for _, val2 := range args {
			// fmt.Printf("args: %v\n", val2)
			if val1 == val2 {
				// fmt.Println("Alert!!!")
				os.Stdout.WriteString("Alert!!!\n")
				return
			}
		}
	}
}
