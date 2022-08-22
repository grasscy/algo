package ac

import "math"

// opt: min stack
type MinStack struct {
    stack      []*Node
    head, tail *Node
    size       int
}

type Node struct {
    val          int
    before, next *Node
}

func Constructor() MinStack {
    s := MinStack{
        stack: []*Node{},
        head:  &Node{math.MaxInt64, nil, nil},
        tail:  &Node{math.MinInt64, nil, nil},
        size:  0,
    }
    s.head.next = s.tail
    s.tail.before = s.head

    return s
}

func (this *MinStack) Push(val int) {
    newn := &Node{val, nil, nil}

    this.stack = append(this.stack, newn)
    this.size++

    this.insertNode(newn)

    // opt
    //this.stack = append(this.stack, x)
    //top := this.minStack[len(this.minStack)-1]
    //this.minStack = append(this.minStack, min(x, top))

}

func (this *MinStack) Pop() {
    node := this.stack[this.size-1]
    this.stack = this.stack[:this.size-1]
    this.size--
    this.rmNode(node)
}

func (this *MinStack) Top() int {
    return this.stack[this.size-1].val
}

func (this *MinStack) GetMin() int {
    return this.tail.before.val
}

func (this *MinStack) insertNode(node *Node) {
    cur := this.head
    for cur.val > node.val {
        cur = cur.next
    }
    before := cur.before
    before.next = node
    node.before = before
    node.next = cur
    cur.before = node
}

func (this *MinStack) rmNode(node *Node) {
    before := node.before
    next := node.next
    before.next = next
    next.before = before
}
