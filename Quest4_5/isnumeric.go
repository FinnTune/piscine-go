package piscine

func IsNumeric(s string) bool {
	var defaultbool bool

	if s == "" {
		return true
	}
	for _, char := range s {
		if char >= 48 && char <= 57 {
			defaultbool = true
		} else {
			return false
		}
	}
	return defaultbool
}
