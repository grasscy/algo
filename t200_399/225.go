package t200_399

type MyStack struct {
	q1 []int
	q2 []int
}

func Constructor() MyStack {
	return MyStack{[]int{}, []int{}}
}

func (this *MyStack) Push(x int) {
	this.q1 = append(this.q1, x)
	for len(this.q2) != 0 {
		this.q1 = append(this.q1, this.q2[0])
		this.q2 = this.q2[1:]
	}
	this.q1, this.q2 = this.q2, this.q1
}

func (this *MyStack) Pop() int {
	i := this.q2[0]
	this.q2 = this.q2[1:]
	return i
}

func (this *MyStack) Top() int {
	return this.q2[0]
}

func (this *MyStack) Empty() bool {
	return len(this.q2) == 0
}
