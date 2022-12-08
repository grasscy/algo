package t1_199

func reverseBits(num uint32) uint32 {
	var ans uint32
	for i := 0; i < 32; i++ {
		u := num % 2
		ans = ans*2 + u
		num /= 2
	}
	return ans
}
