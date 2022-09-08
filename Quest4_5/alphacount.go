package piscine

func AlphaCount(s string) int {
	size := []byte(s)
	count := 0
	for i := 0; i <= len(size)-1; i++ {
		if size[i] > 64 && size[i] < 91 || size[i] > 96 && size[i] < 123 {
			count++
		}
	}
	return count
}
