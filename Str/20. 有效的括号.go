package Str

func isValid(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}
	stack := []byte{}
	for i := 0; i < n; i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack = append(stack, s[i])
		} else {
			// 注意这个一定要进行判断的
			// 当是括号的后面一个的时候
			// stack中一定不是为空的
			if len(stack) == 0 {
				return false
			} else if s[i] == ')' && stack[len(stack)-1] != '(' {
				return false
			} else if s[i] == ']' && stack[len(stack)-1] != '[' {
				return false
			} else if s[i] == '}' && stack[len(stack)-1] != '{' {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		}

	}
	return len(stack) == 0
}
