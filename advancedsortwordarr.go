package piscine

func AdvancedSortWordArr1(a []string, f func(a, b string) int) {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			// if f(a[i], a[j]) {
			a[i], a[j] = a[j], a[i]
		}
	}
}
