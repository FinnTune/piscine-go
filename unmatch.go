package piscine

func Unmatch(a []int) int {
	for _, val := range a {
		count := 0
		for _, val2 := range a {
			if val == val2 {
				count++
			}
		}
		if count == 1 || count%2 == 1 {
			return val
		}
	}
	return -1
}

/* slc := []int(a)
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
for j := 0; j < len(slc); j++ {

}
return 0
}
*/
