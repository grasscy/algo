package ac

import "testing"

func Test_countSubIslands(t *testing.T) {
	type args struct {
		grid1 [][]int
		grid2 [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{[][]int{{1, 1, 1, 0, 0}, {0, 1, 1, 1, 1}, {0, 0, 0, 0, 0}, {1, 0, 0, 0, 0}, {1, 1, 0, 1, 1}},
			[][]int{{1, 1, 1, 0, 0}, {0, 0, 1, 1, 1}, {0, 1, 0, 0, 0}, {1, 0, 1, 1, 0}, {0, 1, 0, 1, 0}}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSubIslands(tt.args.grid1, tt.args.grid2); got != tt.want {
				t.Errorf("countSubIslands() = %v, want %v", got, tt.want)
			}
		})
	}
}
