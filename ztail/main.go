package main

import (
	"fmt"
	"os"
)

var (
	errBool bool = false
	flPrint bool = false
)

func main() {
	max := 0
	if os.Args[1][0] == '-' && len(os.Args[1]) > 1 {
		if os.Args[1] == "-c" {
			max = BasicAtoi(os.Args[2])
			// fmt.Println(max)
		}
	}
	files := os.Args[1:]
	if max > 0 {
		files = os.Args[3:]
	}
	for _, oneFile := range files {
		concat(oneFile, max)
	}
	if errBool {
		os.Exit(1)
	}
}

func concat(oneFile string, max int) {
	file, err := os.Open(oneFile)
	if err != nil {
		fmt.Printf("%v\n", err.Error())
		errBool = true
		flPrint = true
		return
	}
	dat, _ := file.Stat()
	fileByteSlc := make([]byte, dat.Size())
	file.Read(fileByteSlc)
	if flPrint {
		fmt.Println("")
	}
	flPrint = true

	fmt.Printf("==> %v <==\n", oneFile)

	ansStr := string(fileByteSlc)
	srt := 0
	if max > 0 && max < len(fileByteSlc) {
		srt = len(fileByteSlc) - max
		ansStr = ansStr[srt:]
	}
	fmt.Print(ansStr)
}

func BasicAtoi(s string) int {
	num := []byte(s)
	var number int
	for i := 0; i <= len(s)-1; i++ {
		number = number*10 + int(num[i]-48)
	}
	return number
}
