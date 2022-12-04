package ac

import "sort"

func dividePlayers(skill []int) int64 {
	if len(skill)%2 == 1 {
		return -1
	}
	sum := 0
	for _, v := range skill {
		sum += v
	}
	if sum%(len(skill)/2) != 0 {
		return -1
	}
	sum /= len(skill) / 2

	sort.Ints(skill)

	var ans int

	for i := 0; i < len(skill)/2; i++ {
		if skill[i]+skill[len(skill)-1-i] != sum {
			return -1
		}
		ans += skill[i] * skill[len(skill)-1-i]
	}
	return int64(ans)
}
