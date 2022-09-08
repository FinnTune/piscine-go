package piscine

func Capitalize(s string) string {
	ans := ""
	cap := false
	for i, let := range s {
		str := s[i]
		if IsAlpha(string(str)) {
			if cap {
				if IsLower(string(str)) {
					ans += (string(str - 32))
					cap = false
				} else {
					ans += string(let)
					cap = false
				}
			} else if i == 0 {
				if IsLower(string(str)) {
					ans += string(let - 32)
				} else {
					ans += string(let)
				}
			} else if IsUpper(string(str)) {
				ans += string(str + 32)
			} else {
				ans += string(let)
			}
		} else {
			ans += string(let)
			cap = true
		}
	}
	return ans
	/* str := []byte(s)
	cr_str := []rune{}
	for i := 0; i < len(s); i++ {
		if str[i] == 32 { // space
			cr_str = append(cr_str, rune(str[i]))
		} else if i == 0 { // first letter
			cr_str = append(cr_str, rune(str[0]))
		} else if str[i] >= 65 && str[i] <= 90 && str[i-1] == 32 {
			cr_str = append(cr_str, rune(str[i]-32))
		} else if str[i] >= 33 && str[i] <= 64 { // printable
			cr_str = append(cr_str, rune(str[i]))
		} else if str[i] >= 65 && str[i] <= 90 { // capitals
			cr_str = append(cr_str, rune(str[i]+32))
		} else if str[i] >= 123 && str[i] <= 126 { // printable
			cr_str = append(cr_str, rune(str[i]))
		} else if str[i] >= 97 && str[i] <= 122 { // lower
			cr_str = append(cr_str, rune(str[i]))
		}
	}
	fmt.Println(string(cr_str))
	return string(cr_str) */
}
