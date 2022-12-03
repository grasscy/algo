package ac

import "testing"

func Test_secondHighest(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{"abc1111"}, -1},
		{"", args{"dfa12321afd"}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := secondHighest(tt.args.s); got != tt.want {
				t.Errorf("secondHighest() = %v, want %v", got, tt.want)
			}
		})
	}
}
