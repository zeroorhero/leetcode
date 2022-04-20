package stack

func removeDuplicates(s string) string {
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		// 注意当出现-1的这种索引的时候  要进行特殊的处理机制的
		if len(stack) > 0 && s[i] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}
