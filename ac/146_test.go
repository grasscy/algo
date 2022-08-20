package ac

import (
	"fmt"
	"testing"
)

func TestLRUCache_Get(t *testing.T) {
	lRUCache := Constructor(2)
	lRUCache.Put(2, 1)     // 缓存是 {1=1}
	lRUCache.Put(2, 2)     // 缓存是 {1=1, 2=2}
	get := lRUCache.Get(2) // 返回 1
	fmt.Println(get)
	lRUCache.Put(1, 1)    // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
	lRUCache.Put(4, 1)    // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
	get = lRUCache.Get(2) // 返回 -1 (未找到)
	fmt.Println(get)
}
