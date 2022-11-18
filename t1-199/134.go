package pass

func canCompleteCircuit(gas []int, cost []int) int {
	sum := 0
	for i := 0; i < len(gas); i++ {
		sum += gas[i] - cost[i]
	}
	if sum < 0 {
		return -1
	}
	tank, start := 0, 0
	for i := 0; i < len(gas); i++ {
		tank += gas[i] - cost[i]
		if tank < 0 {
			tank = 0
			start = i + 1
		}
	}
	if start == len(gas) {
		return -1
	}
	return start
}
