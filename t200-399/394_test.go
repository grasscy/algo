package t200_399

import "testing"

func Test_decodeString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{"3[a2[c]]"}, "accaccacc"},
		{"", args{"3[a]2[bc]"}, "aaabcbc"},
		{"", args{"leetcode"}, "leetcode"},
		{"", args{"2[abc]3[cd]ef"}, "abcabccdcdcdef"},
		{"", args{"abc3[cd]xyz"}, "abccdcdcdxyz"},
		{"", args{"3[z]2[2[y]pq4[2[jk]e1[f]]]ef"}, "zzzyypqjkjkefjkjkefjkjkefjkjkefyypqjkjkefjkjkefjkjkefjkjkefef"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decodeString(tt.args.s); got != tt.want {
				t.Errorf("decodeString() = %v, want %v", got, tt.want)
			}
		})
	}
}
