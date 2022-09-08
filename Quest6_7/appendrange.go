package piscine

func AppendRange(min, max int) []int {
	slc := make([]int, 0)
	if min >= max {
		return []int(nil)
	} else {
		for i := min; i < max; i++ {
			slc = append(slc, i)
		}
	}
	return slc
}
