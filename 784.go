package algo

var ans []string

func letterCasePermutation(s string) []string {
	ans = []string{}
	d(s, 0, "")
	return ans
}

func d(s string, index int, tmp string) {
	if index == len(s) {
		ans = append(ans, tmp)
		return
	}
	i := index
	if s[i] >= 'A' && s[i] <= 'Z' {
		d(s, index+1, tmp+string(s[i]))
		d(s, index+1, tmp+string(s[i]+'a'-'A'))
	} else if s[i] >= 'a' && s[i] <= 'z' {
		d(s, index+1, tmp+string(s[i]))
		d(s, index+1, tmp+string(s[i]-'a'+'A'))
	} else {
		d(s, index+1, tmp+string(s[i]))
	}
}
