package Str

import (
	"fmt"
	"sort"
)

// 注意在对字符串进行遍历的时候
// 如果使用的是for range 得到每一个元素的类型是rune
// 如果使用的是for i进行遍历 得到的类型是byte
func isAnagram(s string, t string) bool {
	map1 := map[byte]int{}
	for i := 0; i < len(s); i++ {
		if _, ok := map1[s[i]]; ok {
			map1[s[i]] = map1[s[i]] + 1
		} else {
			map1[s[i]] = 1
		}
	}

	for i := 0; i < len(t); i++ {
		if _, ok := map1[t[i]]; !ok {
			return false
		} else {
			map1[t[i]] = map1[t[i]] - 1
		}
	}

	for _, v := range map1 {
		if v != 0 {
			return false
		}
	}
	return true
}

func isAnagram1(s string, t string) bool {
	s1, s2 := []byte(s), []byte(t)
	sort.Slice(s1, func(i, j int) bool {
		return s1[i] < s1[j]
	})
	for _, val := range s1 {
		fmt.Println(val)
	}
	sort.Slice(s2, func(i, j int) bool {
		return s2[i] < s2[j]
	})
	return string(s1) == string(s2)
}

func isAnagram2(s, t string) bool {
	// 数组之间可以进行比较  而切片之间是不能进行比较的
	var c1, c2 [26]int
	for _, ch := range s {
		c1[ch-'a']++
	}
	for _, ch := range t {
		c2[ch-'a']++
	}
	return c1 == c2
}
