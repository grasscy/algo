package t200_399

type Node struct {
	val    uint8
	childs []*Node
}
type WordDictionary struct {
	root *Node
}

func Constructor() WordDictionary {
	return WordDictionary{&Node{0, []*Node{}}}
}

func (this *WordDictionary) AddWord(word string) {
	cur := this.root
	for i := 0; i < len(word); i++ {
		node := check(cur.childs, word[i])
		if node == nil {
			node = &Node{word[i], []*Node{}}
			cur.childs = append(cur.childs, node)
		}
		cur = node
	}
	cur.childs = append(cur.childs, &Node{0, []*Node{}})
}

func check(ns []*Node, v uint8) *Node {
	for _, vv := range ns {
		if vv.val == v {
			return vv
		}
	}
	return nil
}

func (this *WordDictionary) Search(word string) bool {
	cur := []*Node{this.root}
	for i := 0; i < len(word); i++ {
		t := []*Node{}
		if word[i] == '.' {
			for _, v := range cur {
				t = append(t, v.childs...)
			}
		} else {
			f := false
			for _, v := range cur {
				node := check(v.childs, word[i])
				if node != nil {
					f = true
					t = append(t, node)
				}
			}
			if !f {
				return false
			}
		}
		cur = t
	}
	if len(cur) == 0 {
		return false
	}

	for _, v := range cur {
		for _, vv := range v.childs {
			if vv.val == 0 {
				return true
			}
		}
	}
	return false
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
