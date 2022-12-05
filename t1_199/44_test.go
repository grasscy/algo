package t1_199

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
		{"3", args{"a", ""}, false},
		{"3", args{"a", "?*"}, true},
		{"4", args{"a", "a*"}, true},
		{"1", args{"adceb", "*a*b"}, true},
		{"6", args{"", "******"}, true},
		{"2", args{"a", "aa"}, false},
		{"5", args{"abcabczzzde", "*abc???de*"}, true},
		{"8", args{"acdcb", "a*c?b"}, false},
		{"9", args{"aa", "a"}, false},
		{"10", args{"aa", "*"}, true},
		{"11", args{"cb", "?a"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
