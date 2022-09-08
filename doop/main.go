package main

import (
	"os"
)

func main() {
	args := os.Args
	for i := 1; i <= len(args)-1; i++ {
		if !IsNumeric(args[i][1:]) {
			return
		}
	}
	if len(args[1]) >= 19 || len(args[3]) >= 19 {
		return
	}
	if len(args) > 4 {
		return
	} else if args[2] == "/" && args[3] == "0" {
		os.Stdout.WriteString("No division by 0\n")
		// fmt.Println("No division by 0")
	} else if args[2] == "%" && args[3] == "0" {
		os.Stdout.WriteString("No modulo by 0\n")
		// fmt.Println("No modulo by 0")
	} else if args[2] != "+" && args[2] != "-" && args[2] != "*" && args[2] != "/" && args[2] != "%" {
		return
	} else {
		operand1 := BasicAtoi(args[1])
		// fmt.Printf("%v\n", operand1)
		operand2 := BasicAtoi(args[3])
		// fmt.Printf("%v\n", operand2)
		operator := args[2]
		// fmt.Printf("%v\n", operator)
		ans := 0
		if operator == "+" {
			ans = operand1 + operand2
		} else if operator == "-" {
			ans = operand1 - operand2
		} else if operator == "*" {
			ans = operand1 * operand2
		} else if operator == "/" {
			ans = operand1 / operand2
		} else if operator == "%" {
			ans = operand1 % operand2
		}
		intToStr(ans)
	}
}

// os.Stderr.WriteString("String Written")

func BasicAtoi(s string) int {
	num := []byte(s)
	var number int
	boolNeg := false
	// fmt.Printf("boolNeg: %v\n", boolNeg)
	for i := 0; i < len(s); i++ {
		if s[0] == '-' {
			boolNeg = true
			// fmt.Printf("boolNeg: %v\n", boolNeg)
		}
		if s[i] == '-' {
			i++
		}
		number = number*10 + int(num[i]-48)
		// fmt.Printf("BasicAtoI: %v\n", number)
	}
	if boolNeg {
		number *= -1
		// fmt.Printf("BasicAtoIneg:%v\n", number)
		return number
	} else {
		return number
	}
}

func intToStr(n int) {
	num := n
	crtslc := []rune{}
	negNum := false
	if num < 0 {
		negNum = true
		num *= -1
	} else if num == 0 {
		os.Stdout.WriteString("0")
	}
	for num > 0 {
		crtslc = append(crtslc, rune(num%10))
		num /= 10
	}
	if negNum {
		crtslc = append(crtslc, '-')
		// fmt.Println(crtslc)
	}
	if len(crtslc) > 19 {
		return
	}
	for i := len(crtslc) - 1; i >= 0; i-- {
		if negNum {
			os.Stdout.WriteString(string(crtslc[i]))
			negNum = false
			continue
		}
		os.Stdout.WriteString(string(crtslc[i] + 48))
	}
	os.Stdout.WriteString("\n")
}

func IsNumeric(s string) bool {
	var defaultbool bool

	if s == "" {
		return true
	}
	for _, char := range s {
		if char >= 48 && char <= 57 {
			defaultbool = true
		} else {
			return false
		}
	}
	return defaultbool
}
