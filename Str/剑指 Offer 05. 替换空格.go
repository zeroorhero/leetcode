package Str

import "strings"

func replaceSpace(s string) string {
	s = strings.Replace(s, " ", "%20", -1)
	return s
}

func replaceSpace1(s string) string {
	res := []byte{}
	for i := range s {
		// 字符串直接使用+ 也是可以进行链接操作的
		if string(s[i]) == " " {
			res = append(res, '%', '2', '0')
		} else {
			res = append(res, []byte(s[i:i+1])...)
		}
	}
	return string(res)
}

func replaceSpace2(s string) string {
	var res string
	for _, v := range s {
		switch v {
		case ' ': //遇到空格则加上%20
			res += "%20"
		default: //默认加上v值，但是要转为string
			res += string(v)
		}
	}
	return res
}
