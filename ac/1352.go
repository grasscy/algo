package ac

type ProductOfNumbers struct {
	pre []int
}

func Constructor() ProductOfNumbers {
	p := ProductOfNumbers{pre: []int{1}}
	return p
}

func (this *ProductOfNumbers) Add(num int) {
	if num == 0 {
		this.pre = []int{1}
	} else {
		pre := this.pre
		this.pre = append(pre, pre[len(pre)-1]*num)
	}
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	if len(this.pre) <= k {
		return 0
	}
	pre := this.pre
	return pre[len(pre)-1] / pre[len(pre)-1-k]
}
