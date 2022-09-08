package piscine

func ToUpper(s string) string {
	cast_byt_s := []byte(s)
	screated := []byte{}
	for _, char := range cast_byt_s {
		if char >= 97 && char <= 122 {
			screated = append(screated, char-32)
		} else {
			screated = append(screated, char)
		}
	}
	return string(screated)
}
