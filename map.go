package piscine

func Map(f func(int) bool, a []int) []bool {
	slc := []bool{}
	for _, val := range a {
		slc = append(slc, f(val))
	}
	return slc
}
