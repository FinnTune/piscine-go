package piscine

func Compact(ptr *[]string) int {
	slc := []string{}
	for i := 0; i < len(*ptr); i++ {
		if (*ptr)[i] != "" {
			slc = append(slc, (*ptr)[i])
		}
	}
	*ptr = make([]string, len(slc))
	copy(*ptr, slc)
	return len(slc)
}
