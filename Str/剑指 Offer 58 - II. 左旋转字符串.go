package Str

func reverseLeftWords1(s string, n int) string {
	// 字符串的切片还是字符串的类型的
	//res :=  append(s[n:len(s)], s[0:n]...)
	//tem := ""
	//for _, val := range res{
	//	tem += val
	//}
	//return tem
}

func reverseLeftWords(s string, n int) string {
	return s[n:len(s)] + s[0:n]
}
