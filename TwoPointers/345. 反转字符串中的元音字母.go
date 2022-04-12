package TwoPointers

import "strings"

func reverseVowels1(s string) string {
	tem := "aeiouAEIOU"
	chars := []byte(s)
	l, r := 0, len(chars)-1
	for l < r {
		if strings.Contains(tem, string(chars[l])) && strings.Contains(tem, string(chars[r])) {
			chars[l], chars[r] = chars[r], chars[l]
			// 交换完毕后要进行更新操作
			l = l + 1
			r--
		} else if !strings.Contains(tem, string(chars[l])) {
			l++
		} else {
			r--
		}
	}
	return string(chars)
}

func reverseVowels(s string) string {
	tem := "aeiouAEIOU"
	chars := []byte(s)
	l, r := 0, len(chars)-1
	for l < r {
		// 注意后面的l和r大小的判断的
		for !strings.Contains(tem, string(chars[l])) && l < r {
			l++
		}
		for !strings.Contains(tem, string(chars[r])) && r > 0 {
			r--
		}
		if strings.Contains(tem, string(chars[l])) && strings.Contains(tem, string(chars[r])) {
			chars[l], chars[r] = chars[r], chars[l]
			l++
			r--
		}
	}
	return string(chars)
}
