package ac

func flipLights(n int, presses int) int {
	if presses == 0 {
		return 1
	}
	if n == 1 {
		return 2
	} else if n == 2 {
		if presses == 1 {
			return 3
		} else {
			return 4

		}
	} else {
		if presses == 1 {
			return 4
		} else {
			if presses == 2 {
				return 7
			}
			return 8
		}
	}
}
