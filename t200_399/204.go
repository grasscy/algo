package t200_399

func countPrimes(n int) int {
	isprime := make([]bool, n)
	for i := 0; i < n; i++ {
		isprime[i] = true
	}
	for i := 2; i*i < n; i++ {
		if isprime[i] {
			for j := i * i; j < n; j += i {
				isprime[j] = false
			}
		}
	}
	var ans int
	for i := 2; i < n; i++ {
		if isprime[i] {
			ans++
		}
	}
	return ans
}
