package Str

func longestPalindrome(s string) int {
	hashmap := map[byte]int{}
	for i := 0; i < len(s); i++ {
		if _, ok := hashmap[s[i]]; ok {
			hashmap[s[i]] += 1
		} else {
			hashmap[s[i]] = 1
		}
	}
	res := 0
	flag := false
	for _, v := range hashmap {
		if v%2 == 0 {
			res += v
		} else {
			flag = true
			res += (v / 2) * 2
		}
	}
	if flag {
		return res + 1
	} else {
		return res
	}
}
