package t1_199

import "strings"

func fullJustify(words []string, maxWidth int) []string {
	rows := [][]string{}

	nc := 0
	cr := []string{}
	for _, v := range words {
		if len(cr) == 0 {
			cr = append(cr, v)
			nc += len(v)
		} else {
			if nc+1+len(v) <= maxWidth {
				cr = append(cr, v)
				nc += 1 + len(v)
			} else {
				rows = append(rows, append([]string{}, cr...))
				cr = []string{v}
				nc = len(v)
			}
		}
	}
	if len(cr) > 0 {
		rows = append(rows, append([]string{}, cr...))
	}
	ans := []string{}
	for i, v := range rows {
		if i == len(rows)-1 {
			join := strings.Join(v, " ")
			ll := len(join)
			for i := 0; i < maxWidth-ll; i++ {
				join += " "
			}
			ans = append(ans, join)
		} else {
			wll := 0
			for _, w := range v {
				wll += len(w)
			}

			join := v[0]

			if len(v) > 1 {
				kong := (maxWidth - wll) / (len(v) - 1)
				yu := (maxWidth - wll) % (len(v) - 1)

				for i := 1; i < len(v); i++ {
					for j := 0; j < kong; j++ {
						join += " "
					}
					if yu != 0 {
						join += " "
						yu--
					}
					join += v[i]

				}

			} else {
				for i := len(join); i < maxWidth; i++ {
					join += " "
				}
			}

			ans = append(ans, join)

		}
	}
	return ans
}
