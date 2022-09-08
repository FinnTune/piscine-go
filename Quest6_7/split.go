package piscine

func Split(s, sep string) []string {
	var slc []string
	var str string
	start := 0
	// fmt.Printf("Here: %v", s[2:5])
	if len(s) == 0 {
		return nil
	} else if len(sep) == 0 {
		return nil
	}
	for i := start; i < len(s); i++ {
		if len(sep) > len(s[i:]) {
			break
		} else if string(s[i:i+len(sep)]) == sep {
			str = s[start:i]
			start = i + len(sep)
			if str != "" {
				slc = append(slc, str)
			}
			continue
		}
	}
	slc = append(slc, s[start:])
	return slc
	/* sent := []string{}
	for i := 0; i < len(s); i++ {
		if s[i:i+len(sep)] == sep {
			sent = append(sent, s[:i-1])
		}
	}
	return sent */
}
