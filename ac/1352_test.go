package ac

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	productOfNumbers := Constructor()
	productOfNumbers.Add(3)        // [3]
	productOfNumbers.Add(0)        // [3,0]
	productOfNumbers.Add(2)        // [3,0,2]
	productOfNumbers.Add(5)        // [3,0,2,5]
	productOfNumbers.Add(4)        // [3,0,2,5,4]
	productOfNumbers.GetProduct(2) // 返回 20 。最后 2 个数字的乘积是 5 * 4 = 20
	productOfNumbers.GetProduct(3) // 返回 40 。最后 3 个数字的乘积是 2 * 5 * 4 = 40
	productOfNumbers.GetProduct(4) // 返回  0 。最后 4 个数字的乘积是 0 * 2 * 5 * 4 = 0
	productOfNumbers.Add(8)        // [3,0,2,5,4,8]
	productOfNumbers.GetProduct(2) // 返回 32 。最后 2 个数字的乘积是 4 * 8 = 32

}
