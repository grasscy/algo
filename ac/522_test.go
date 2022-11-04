package ac

import "testing"

func Test_findLUSlength(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{[]string{"aba", "cdc", "eae"}}, 3},
		{"", args{[]string{"aaa", "aaa", "aa"}}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLUSlength(tt.args.strs); got != tt.want {
				t.Errorf("findLUSlength() = %v, want %v", got, tt.want)
			}
		})
	}
}
