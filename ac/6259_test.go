package ac

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	var r int
	loc := Constructor(5) // 初始化一个大小为 10 的内存数组，所有内存单元都是空闲的。

	r = loc.Free(4) // 释放 mID 为 2 的所有内存单元。内存数组变为 [1, ,3, , , , , , , ] 。返回 1 ，因为只有 1 个 mID 为 2 的内存单元。
	fmt.Println(r)
	r = loc.Allocate(9, 5) // 最左侧的块的第一个下标是 1 。内存数组变为 [1,2, , , , , , , , ]。返回 1 。
	fmt.Println(r)
	r = loc.Allocate(5, 8) // 最左侧的块的第一个下标是 2 。内存数组变为 [1,2,3, , , , , , , ]。返回 2 。
	fmt.Println(r)
	r = loc.Allocate(4, 8) // 最左侧的块的第一个下标是 3 。内存数组变为 [1, ,3,4,4,4, , , , ]。返回 3 。
	fmt.Println(r)
	r = loc.Allocate(1, 1) // 最左侧的块的第一个下标是 1 。内存数组变为 [1,1,3,4,4,4, , , , ]。返回 1 。
	fmt.Println(r)
	r = loc.Allocate(1, 1) // 最左侧的块的第一个下标是 6 。内存数组变为 [1,1,3,4,4,4,1, , , ]。返回 6 。
	fmt.Println(r)
	r = loc.Free(1) // 释放 mID 为 1 的所有内存单元。内存数组变为 [ , ,3,4,4,4, , , , ] 。返回 3 ，因为有 3 个 mID 为 1 的内存单元。
	fmt.Println(r)
	r = loc.Allocate(10, 2) // 无法找出长度为 10 个连续空闲内存单元的空闲块，所有返回 -1 。
	fmt.Println(r)
	r = loc.Free(7) // 释放 mID 为 7 的所有内存单元。内存数组保持原状，因为不存在 mID 为 7 的内存单元。返回 0 。
	fmt.Println(r)
}
