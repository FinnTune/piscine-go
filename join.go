package piscine

func Join(strs []string, sep string) string {
	newStr := ""
	for i := 0; i < len(strs); i++ {
		if i+1 == len(strs) {
			newStr += strs[i]
		} else {
			newStr += strs[i] + sep
		}
	}
	return newStr
}
