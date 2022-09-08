package piscine

func ToLower(s string) string {
	cast_byt_s := []byte(s)
	screated := []byte{}
	for _, char := range cast_byt_s {
		if char >= 65 && char <= 90 {
			screated = append(screated, char+32)
		} else {
			screated = append(screated, char)
		}
	}
	return string(screated)
}
