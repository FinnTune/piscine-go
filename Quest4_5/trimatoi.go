package piscine

func TrimAtoi(s string) int {
	num := []byte(s)
	var number int
	negative := false
	for i := 0; i <= len(s)-1; i++ {
		if num[i] >= 48 && num[i] <= 57 {
			number = number*10 + int(num[i]-48)
		} else if number == 0 {
			if num[i] == 45 {
				negative = true
			}
		}
	}
	if negative {
		return number * -1
	}
	return number
}
