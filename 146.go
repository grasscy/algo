package algo

type LRUCache struct {
    m          map[int]*Node
    size       int
    capacity   int
    head, tail *Node
}

type Node struct {
    key        int
    val        int
    pre, after *Node
}

func Constructor(capacity int) LRUCache {
    lru := LRUCache{
        m:        map[int]*Node{},
        size:     0,
        capacity: capacity,
        head:     &Node{0, 0, nil, nil},
        tail:     &Node{0, 0, nil, nil},
    }
    lru.head.after = lru.tail
    lru.tail.pre = lru.head

    return lru
}

func (this *LRUCache) Get(key int) int {
    if v, ok := this.m[key]; ok {
        this.mvToHead(v)
        return v.val
    } else {
        return -1
    }
}

func (this *LRUCache) Put(key int, value int) {
    n := &Node{key, value, nil, nil}

    if v, ok := this.m[key]; ok {
        this.mvToHead(v)
        this.m[key].val=value
    } else {
        if this.size == this.capacity {
            tail := this.rmTail()
            delete(this.m, tail.key)
            this.size--
        }
        this.addHead(n)
        this.m[key] = n
        this.size++
    }
}

func (this *LRUCache) addHead(n *Node) {
    n.pre = this.head
    n.after = this.head.after
    this.head.after.pre = n
    this.head.after = n
}

func (this *LRUCache) rmTail() *Node {
    n := this.tail.pre
    this.rmNode(n)
    return n
}

func (this *LRUCache) mvToHead(n *Node) {
    this.rmNode(n)
    this.addHead(n)
}

func (this *LRUCache) rmNode(n *Node) {
    n.after.pre = n.pre
    n.pre.after = n.after
}