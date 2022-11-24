package pass

import "testing"

func Test_isMatch(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{"a", ".*"}, true},
		{"", args{"ab", ".*c"}, false},
		{"", args{"bbbba", ".*a*a"}, true},
		{"", args{"a", "ab*"}, true},
		{"", args{"ac", ".*c"}, true},
		{"", args{"aaa", "ab*a*c*a"}, true},
		{"", args{"aa", "a*"}, true},
		{"", args{"ab", ".*"}, true},
		{"", args{"aab", "a*c*b"}, true},
		{"", args{"aab", "c*a*b"}, true},
		{"", args{"aa", "a"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
