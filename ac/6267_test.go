package ac

import "testing"

func Test_isPossible(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{6, [][]int{{1, 2}, {1, 3}, {1, 4}, {4, 5}, {5, 6}}}, true},
		{"", args{4, [][]int{{4, 1}, {3, 2}, {2, 4}, {1, 3}}}, true},
		{"", args{4, [][]int{{1, 2}, {3, 4}}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPossible(tt.args.n, tt.args.edges); got != tt.want {
				t.Errorf("isPossible() = %v, want %v", got, tt.want)
			}
		})
	}
}
