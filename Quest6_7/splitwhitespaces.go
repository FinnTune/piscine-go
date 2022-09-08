package piscine

func SplitWhiteSpaces(s string) []string {
	var slc []string
	var str string
	start := 0
	for i := start; i < len(s); i++ {
		if string(s[i]) == " " || string(s[i]) == "\t" || string(s[i]) == "\n" {
			str = s[start:i]
			start = i + 1
			if str != "" {
				slc = append(slc, str)
			}
			continue
		}
	}
	slc = append(slc, s[start:])
	return slc
}
