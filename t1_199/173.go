package t1_199

type BSTIterator struct {
	ns    []*TreeNode
	index int
}

func Constructor(root *TreeNode) BSTIterator {
	ret := &BSTIterator{
		ns:    []*TreeNode{},
		index: 0,
	}
	ret.d(root)
	return *ret
}

func (this *BSTIterator) d(r *TreeNode) {
	if r == nil {
		return
	}
	this.d(r.Left)
	this.ns = append(this.ns, r)
	this.d(r.Right)
}

func (this *BSTIterator) Next() int {
	n := this.ns[this.index]
	this.index++
	return n.Val
}

func (this *BSTIterator) HasNext() bool {
	if this.index == len(this.ns) {
		return false
	}
	return true
}
