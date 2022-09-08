package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	start := 0
	ascBool := true
	desBool := true
	for ind := start; ind < len(a)-1; ind++ {
		if f(a[ind], a[ind+1]) == 0 {
			start += 1
		} else if f(a[ind], a[ind+1]) < 0 {
			start += 1
		} else if f(a[ind], a[ind+1]) > 0 {
			ascBool = false
		}
	}
	for ind := start; ind < len(a)-1; ind++ {
		if f(a[ind], a[ind+1]) == 0 {
			start += 1
		} else if f(a[ind], a[ind+1]) < 0 {
			desBool = false
		} else if f(a[ind], a[ind+1]) > 0 {
			start += 1
		}
	}
	if ascBool || desBool {
		return true
	}
	return false
}
