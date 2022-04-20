package stack

type MyStack struct {
	qu1 []int
	qu2 []int
}

func Constructor() MyStack {
	return MyStack{
		qu1: make([]int, 0),
		qu2: make([]int, 0),
	}
}

func (this *MyStack) Push(x int) {
	// 每次添加新的元素都是放到qu2的第一个的
	this.qu2 = append(this.qu2, x)
	for len(this.qu1) > 0 {
		this.qu2 = append(this.qu2, this.qu1[0])
		this.qu1 = this.qu1[1:]
	}
	// 进行交换
	this.qu1, this.qu2 = this.qu2, this.qu1
}

func (this *MyStack) Pop() int {
	val := this.qu1[0]
	this.qu1 = this.qu1[1:]
	return val
}

func (this *MyStack) Top() int {
	return this.qu1[0]
}

func (this *MyStack) Empty() bool {
	return len(this.qu1) == 0 && len(this.qu2) == 0
}
