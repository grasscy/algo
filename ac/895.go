package ac

type FreqStack struct {
	mmax int
	cs   map[int]int
	head *Node
}

//type FreqStack struct {
//    freq    map[int]int
//    group   map[int][]int
//    maxFreq int
//}

type Node struct {
	v            int
	before, next *Node
}

func Constructor() FreqStack {
	return FreqStack{
		mmax: 0,
		cs:   map[int]int{},
		head: nil,
	}
}

func (this *FreqStack) Push(val int) {
	this.cs[val]++
	if this.cs[val] > this.mmax {
		this.mmax = this.cs[val]
	}
	n := &Node{val, nil, nil}
	if this.head == nil {
		this.head = n
	} else {
		this.head.before = n
		n.next = this.head
		this.head = n
	}
}

func (this *FreqStack) Pop() int {
	ans := -1
	cur := this.head
	for cur != nil {
		if this.cs[cur.v] == this.mmax {
			ans = cur.v
			this.cs[cur.v]--

			if cur == this.head && cur.next != nil {
				cur.next.before = nil
				this.head = cur.next
			} else {
				before := cur.before
				next := cur.next
				if before != nil {
					before.next = next
				}
				if next != nil {
					next.before = before
				}
			}
			break
		}
		cur = cur.next
	}
	if ans == -1 {
		this.mmax--
		return this.Pop()
	}
	return ans
}
