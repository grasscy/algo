package ac

type Allocator struct {
	m []int
}

func Constructor(n int) Allocator {
	a := make([]int, n)
	return Allocator{a}
}

func (this *Allocator) Allocate(size int, mID int) int {
	var start int
	f := false
	for i := 0; i <= len(this.m)-size; i++ {
		if this.m[i] != 0 {
			continue
		}
		j := 0
		for ; j < size && this.m[i+j] == 0; j++ {
		}
		if j >= size {
			for k := i; k < i+size; k++ {
				this.m[k] = mID
			}
			start = i
			f = true
			break
		}
	}
	if !f {
		return -1
	}
	return start
}

func (this *Allocator) Free(mID int) int {
	ret := 0
	for i, v := range this.m {
		if v == mID {
			this.m[i] = 0
			ret++
		}
	}
	return ret
}

/**
 * Your Allocator object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Allocate(size,mID);
 * param_2 := obj.Free(mID);
 */
