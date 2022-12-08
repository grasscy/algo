package t200_399

func isIsomorphic(s string, t string) bool {
	m := map[int32]uint8{}

	mm := map[uint8]bool{}
	if len(s) != len(t) {
		return false
	}
	for i, v := range s {
		if vv, ok := m[v]; ok {
			if vv != t[i] {
				return false
			}
		} else {
			if mm[t[i]] {
				return false
			}
			m[v] = t[i]
			mm[t[i]] = true
		}
	}
	return true
}
