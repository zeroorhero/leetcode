package Str

import "strings"

func reverseWords(s string) string {
	split := strings.Split(strings.TrimSpace(s), " ")
	l, r := 0, len(split)-1
	for l < r {
		s1 := strings.TrimSpace(split[l])
		s2 := strings.TrimSpace(split[r])
		split[l], split[r] = s2, s1
		l++
		r--
	}
	res := ""
	for _, val := range split {
		if val == "" {
			continue
		} else {
			res += val
			res += " "
		}

	}
	return strings.TrimSpace(res)

}
