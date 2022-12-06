package t1_199

import "testing"

func Test_findMin(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{[]int{2, 1}}, 1},
		{"", args{[]int{3, 4, 5, 1, 2}}, 1},
		{"", args{[]int{4, 5, 6, 7, 0, 1, 2}}, 0},
		{"", args{[]int{11, 13, 15, 17}}, 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMin(tt.args.nums); got != tt.want {
				t.Errorf("findMin() = %v, want %v", got, tt.want)
			}
		})
	}
}
