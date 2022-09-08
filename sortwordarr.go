package piscine

/*
func SortWordArr1(a []string) {
	byteSlc := []byte{}
	strSlc := []string{}
	for i := 0; i <= len(a)-1; i++ {
		for j := 0; j <= len(a[i])-1; j++ {
			byteSlc = append(byteSlc, byte(a[i][j]))
		}
	}
	// fmt.Printf("byteSlc: %v\n", byteSlc)
	for i := 0; i < len(byteSlc); i++ {
		for j := i + 1; j < len(byteSlc); j++ {
			if byteSlc[i] > byteSlc[j] {
				temp := byteSlc[i]
				byteSlc[i] = byteSlc[j]
				byteSlc[j] = temp
			}
		}
	}
	for _, val := range byteSlc {
		strSlc = append(strSlc, string(rune(val)))
	}
	// fmt.Printf("byteSlc: %v\n", strSlc)
} */

func SortWordArr(a []string) {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}
