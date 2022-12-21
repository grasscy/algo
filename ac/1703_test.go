package ac

import "testing"

func Test_minMoves(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{[]int{0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 9}, 22},
		{"", args{[]int{1, 0, 0, 1, 0, 1}, 2}, 1},
		{"", args{[]int{1, 0, 0, 0, 0, 0, 1, 1}, 3}, 5},
		{"", args{[]int{1, 1, 0, 1}, 2}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minMoves(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("minMoves() = %v, want %v", got, tt.want)
			}
		})
	}
}
