package piscine

func Rot14(s string) string {
	ansString := ""

	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			char += 14
			if char > 'z' {
				char -= 26
			}
		} else if char >= 'A' && char <= 'Z' {
			char += 14
			if char > 'Z' {
				char -= 26
			}
		}

		ansString += string(char)
	}

	return ansString
}
