package piscine

func Max(a []int) int {
	slc := []int(a)
	if a == nil {
		return 0
	}
	for i := 0; i < len(slc); i++ {
		for j := i + 1; j < len(slc); j++ {
			if slc[i] < slc[j] {
				slc[i], slc[j] = slc[j], slc[i]
			}
		}
	}
	// fmt.Printf("Here: %v\n", slc)
	return slc[0]
}
