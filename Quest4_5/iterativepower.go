package piscine

func IterativePower(nb int, power int) int {
	result := 1
	if nb > 63 || power < 0 || nb == 0 && power == 1 {
		return 0
	} else if nb == 1 && power == 1 {
		return 1
	} else if power == 0 {
		return 1
	} else if power >= 1 {
		for i := 1; i <= power; i++ {
			result = result * nb
		}
	}
	return result
}
