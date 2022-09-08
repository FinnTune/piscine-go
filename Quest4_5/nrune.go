package piscine

func NRune(s string, n int) rune {
	element := []rune(s)
	if n <= 0 || n > len(s) {
		return 0
	} else {
		return element[n-1]
	}
}
