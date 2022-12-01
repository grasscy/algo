package pass

import "testing"

func Test_isInterleave(t *testing.T) {
	type args struct {
		s1 string
		s2 string
		s3 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{"a", "", "c"}, false},
		{"", args{"aabc", "abad", "aabadabc"}, true},
		{"", args{"aabcc", "dbbca", "aadbbcbcac"}, true},
		{"", args{"aabcc", "dbbca", "aadbbbaccc"}, false},
		{"", args{"", "", ""}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isInterleave(tt.args.s1, tt.args.s2, tt.args.s3); got != tt.want {
				t.Errorf("isInterleave() = %v, want %v", got, tt.want)
			}
		})
	}
}
