package algo

import "testing"

func Test_scoreOfParentheses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{"(()()())"}, 6},
		{"", args{"(()(()))"}, 6},
		{"", args{"()()"}, 2},
		{"", args{"()"}, 1},
		{"", args{"(())"}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := scoreOfParentheses(tt.args.s); got != tt.want {
				t.Errorf("scoreOfParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
