package ac

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {

	trie := Constructor()
	trie.Insert("apple")
	r := trie.Search("apple") // 返回 True
	fmt.Println(r)
	r = trie.Search("app") // 返回 False
	fmt.Println(r)

	r = trie.StartsWith("app") // 返回 True
	fmt.Println(r)

	trie.Insert("app")
	r = trie.Search("app") // 返回 True
	fmt.Println(r)

}
