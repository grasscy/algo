package t200_399

type Trie struct {
	root *Node
	// opt
	//children [26]*Trie
	//isEnd    bool
}

type Node struct {
	val   int32
	nexts []*Node
	isEnd bool
}

func Constructor() Trie {
	t := Trie{}
	root := &Node{0, []*Node{}, false}
	t.root = root
	return t
}

func (this *Trie) Insert(word string) {
	cur := this.root
	for _, v := range word {
		exsit := false
		for _, n := range cur.nexts {
			if n.val == v {
				cur = n
				exsit = true
				break
			}
		}
		if !exsit {
			nn := &Node{v, []*Node{}, false}
			cur.nexts = append(cur.nexts, nn)
			cur = nn
		}
	}
	cur.isEnd = true
}

func (this *Trie) Search(word string) bool {
	cur := this.root
	for _, v := range word {
		exsit := false
		for _, n := range cur.nexts {
			if n.val == v {
				cur = n
				exsit = true
				break
			}
		}
		if !exsit {
			return false
		}
	}
	return cur.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	cur := this.root
	for _, v := range prefix {
		exsit := false
		for _, n := range cur.nexts {
			if n.val == v {
				cur = n
				exsit = true
				break
			}
		}
		if !exsit {
			return false
		}
	}
	return true
}
