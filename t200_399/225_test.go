package t200_399

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	obj := Constructor()
	obj.Push(1) // q1 ,q2 1,
	obj.Push(2) // q1 2 ,q2 1,;q1 2,1 ,q2 ;q1, q2 2 1
	obj.Push(3) // q1 3,q2 2,1;q1 3 2 1,q2;q1, q2 321

	var r int
	r = obj.Pop()
	fmt.Println(r)
	r = obj.Top()
	fmt.Println(r)
}
