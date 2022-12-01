package ac

import (
	"fmt"
	"testing"
)

func Test_FreqStack(t *testing.T) {
	freqStack := Constructor()
	freqStack.Push(5) //堆栈为 [5]
	freqStack.Push(7) //堆栈是 [5,7]
	freqStack.Push(5) //堆栈是 [5,7,5]
	freqStack.Push(7) //堆栈是 [5,7,5,7]
	freqStack.Push(4) //堆栈是 [5,7,5,7,4]
	freqStack.Push(5) //堆栈是 [5,7,5,7,4,5]

	var ret int
	ret = freqStack.Pop() //返回 5 ，因为 5 出现频率最高。堆栈变成 [5,7,5,7,4]。
	fmt.Println(ret)
	ret = freqStack.Pop() //返回 7 ，因为 5 和 7 出现频率最高，但7最接近顶部。堆栈变成 [5,7,5,4]。
	fmt.Println(ret)
	ret = freqStack.Pop() //返回 5 ，因为 5 出现频率最高。堆栈变成 [5,7,4]。
	fmt.Println(ret)
	ret = freqStack.Pop() //返回 4 ，因为 4, 5 和 7 出现频率最高，但 4 是最接近顶部的。堆栈变成 [5,7]。
	fmt.Println(ret)
}
