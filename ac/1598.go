package ac

func minOperations(logs []string) int {
	st := []string{}
	for _, l := range logs {
		switch l {
		case "./":
		case "../":
			if len(st) > 0 {
				st = st[:len(st)-1]
			}
		default:
			st = append(st, l)
		}
	}
	return len(st)
}
