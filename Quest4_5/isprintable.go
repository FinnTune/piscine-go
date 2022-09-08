package piscine

func IsPrintable(s string) bool {
	var defaultbool bool

	if s == "" {
		return true
	}
	for _, char := range s {
		if char > 31 {
			defaultbool = true
		} else {
			return false
		}
	}
	return defaultbool
}
