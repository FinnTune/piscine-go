package piscine

func BasicAtoi(s string) int {
	num := []byte(s)
	var number int
	for i := 0; i <= len(s)-1; i++ {
		number = number*10 + int(num[i]-48)
	}
	return number
}
