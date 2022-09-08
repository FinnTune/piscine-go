package piscine

func IsAlpha(s string) bool {
	var defaultbool bool

	if s == "" {
		return true
	}
	for _, char := range s {
		if char >= 48 && char <= 57 || char >= 65 && char <= 90 || char >= 97 && char <= 122 {
			defaultbool = true
		} else {
			return false
		}
	}
	return defaultbool
}
