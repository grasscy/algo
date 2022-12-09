package t200_399

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	wordDictionary := Constructor()
	wordDictionary.AddWord("a")
	wordDictionary.AddWord("ab")
	wordDictionary.AddWord("an")
	wordDictionary.AddWord("add")

	var r bool
	r = wordDictionary.Search("a") // 返回 False
	fmt.Println(r)
	r = wordDictionary.Search("a.") // 返回 False
	fmt.Println(r)
	r = wordDictionary.Search(".") // 返回 True
	fmt.Println(r)
	//wordDictionary := Constructor()
	//wordDictionary.AddWord("a")
	//var r bool
	//r = wordDictionary.Search("a.") // 返回 False
	//fmt.Println(r)
	//wordDictionary := Constructor()
	//wordDictionary.AddWord("bad")
	//wordDictionary.AddWord("dad")
	//wordDictionary.AddWord("mad")
	//
	//var r bool
	//r = wordDictionary.Search("pad") // 返回 False
	//fmt.Println(r)
	//r = wordDictionary.Search("bad") // 返回 True
	//fmt.Println(r)
	//r = wordDictionary.Search(".ad") // 返回 True
	//fmt.Println(r)
	//r = wordDictionary.Search("b..") // 返回 True
	//fmt.Println(r)
}
