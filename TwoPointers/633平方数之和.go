package TwoPointers

import "math"

func judgeSquareSum(c int) bool {
	// 双指针中右边的数字转化为平方根
	// 使用math这个包
	l, r := 0, int(math.Sqrt(float64(c)))
	for l <= r {
		if l*l+r*r == c {
			return true
		} else if l*l+r*r < c {
			l++
		} else {
			r--
		}
	}
	return false
}
