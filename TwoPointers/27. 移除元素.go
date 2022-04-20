package TwoPointers

func removeElement1(nums []int, val int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	l, r := 0, length
	for l < r {
		if nums[l] == val {
			nums[l] = nums[r-1]
			r--
		} else {
			l++
		}
	}
	return l
}

func removeElement(nums []int, val int) int {
	// 使用fast指定前面是否存在val相等的值的
	// 使用slow用于指定
	length := len(nums)
	slow, fast := 0, 0
	for fast < length {
		if nums[fast] == val {
			fast++
		} else {
			nums[slow] = nums[fast]
			slow++
			fast++
		}
	}
	return slow
}
