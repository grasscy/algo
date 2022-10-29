package ac

var memo map[int]bool

func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	if maxChoosableInteger*(maxChoosableInteger+1) < (desiredTotal * 2) {
		return false
	}
	memo = map[int]bool{}
	return d(0, desiredTotal, 0, maxChoosableInteger)
}

func d(nums int, target int, sum int, maxChoosableInteger int) bool {
	if memo[nums] {
		return memo[nums]
	}
	ans := false
	for i := 1; i <= maxChoosableInteger; i++ {
		if (nums>>i)&1 == 0 {
			if sum+i >= target {
				ans = true
				break
			}
			if !d(nums|(1<<i), target, sum+i, maxChoosableInteger) {
				ans = true
				break
			}
		}
	}
	memo[nums] = ans
	return ans
}
