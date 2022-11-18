package pass

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	getMin := minStack.GetMin()
	fmt.Println(getMin)
	minStack.Pop()
	minStack.Top()
	getMin = minStack.GetMin()
	fmt.Println(getMin)
}
