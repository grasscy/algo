package pass

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{[]int{4, 5, 6, 7, 8, 1, 2, 3}, 8}, 4},
		{"", args{[]int{4, 5, 6, 7, 0, 1, 2}, 6}, 2},
		{"", args{[]int{4, 5, 6, 7, 0, 1, 2}, 0}, 4},
		{"", args{[]int{4, 5, 6, 7, 0, 1, 2}, 3}, -1},
		{"", args{[]int{1}, 0}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
