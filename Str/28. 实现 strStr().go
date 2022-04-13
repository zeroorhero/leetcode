package Str

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	m, n := len(haystack), len(needle)
outer:
	for i := 0; i+n <= m; i++ {
		for j := 0; j < n; j++ {
			// 注意这个是i + j的
			if haystack[i+j] != needle[j] {
				continue outer
			}
		}
		return i
	}
	return -1
}
