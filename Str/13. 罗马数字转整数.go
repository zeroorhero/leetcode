package Str

func romanToInt1(s string) int {
	// map在分割的时候使用的是， 而不是；
	hashMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	res := 0
	// 在对字符串进行便利的时候，每一个是字节的是byte形式的
	for i, val := range s {
		res += hashMap[string(val)]
		if i < len(s)-1 && hashMap[string(val)] < hashMap[string(s[i+1])] {
			res -= 2 * hashMap[string(val)]
		}
	}
	return res
}

func romanToInt(s string) int {
	hashMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	res := 0
	for i := range s {
		res += hashMap[s[i]]
		if i < len(s)-1 && hashMap[s[i]] < hashMap[s[i+1]] {
			res -= 2 * hashMap[s[i]]
		}
	}
	return res

}
