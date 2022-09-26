package ac

import "testing"

func Test_rotate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
	}{
		{"", args{[]int{1, 2, 3, 4, 5, 6, 7}, 1}},
		{"", args{[]int{1, 2, 3, 4, 5, 6, 7}, 2}},
		{"", args{[]int{1, 2, 3, 4, 5, 6, 7}, 3}},

		{"", args{[]int{-1, -100, 3, 99}, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.args.nums, tt.args.k)
		})
	}
}
