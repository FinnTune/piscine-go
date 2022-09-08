package piscine

func Index(s string, toFind string) int {
	/* 	txt := []byte(s)
	   	totxt := []byte(toFind) */
	sliceoftoFind := []rune(toFind)
	startindex := 0
	toFindIndex := 0
	for index, char := range s {
		if sliceoftoFind[toFindIndex] == char {
			if toFindIndex == 0 {
				startindex = index
			}
			toFindIndex++
		} else if toFindIndex > 0 && sliceoftoFind[0] == char {
			startindex = index
			toFindIndex++
		}
		if toFindIndex == len(toFind) {
			return startindex
		}
		/* if totxt[i] == txt[i] && totxt[i+len(toFind)] == txt[i+len(toFind)] {
			return i
		} */

		/* 	if a == b {
			return 0
		} else if a < b {
			return -1
		} else if a > b {
			return 1
		} */
	}
	return -1
}
