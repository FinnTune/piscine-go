package piscine

func ActiveBits(n int) int {
	num := n
	crtslc := []int{}
	count := 0
	for num > 0 {
		crtslc = append(crtslc, num%2)
		// fmt.Printf("Slice:%v\n", crtslc)
		num /= 2
	}
	for i := 0; i < len(crtslc); i++ {
		if crtslc[i] == 1 {
			count++
		}
	}
	// fmt.Printf("Slice:%v\n", crtslc)
	return count
}
