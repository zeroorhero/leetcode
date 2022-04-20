package stack

// 栈每次只能操作切片的最后一个数据

type MyQueue struct {
	in  []int
	out []int
}

func Constructor1() MyQueue {
	return MyQueue{
		in:  make([]int, 0),
		out: make([]int, 0),
	}
}

func (this *MyQueue) Push(x int) {
	this.in = append(this.in, x)
}

// 该函数的作用是将in中的数据全部传输到out中去的
func (this *MyQueue) intoout() {
	for len(this.in) != 0 {
		this.out = append(this.out, this.in[len(this.in)-1])
		this.in = this.in[:len(this.in)-1]
	}
}

func (this *MyQueue) Pop() int {
	if len(this.out) == 0 {
		this.intoout()
	}
	// 只能操作最后一个数
	val := this.out[len(this.out)-1]
	this.out = this.out[:len(this.out)-1]
	return val
}

func (this *MyQueue) Peek() int {
	if len(this.out) == 0 {
		this.intoout()
	}
	return this.out[len(this.out)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.in) == 0 && len(this.out) == 0
}
