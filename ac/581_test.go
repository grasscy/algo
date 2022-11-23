package ac

import "testing"

func Test_findUnsortedSubarray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{[]int{2, 6, 4, 8, 10, 9, 15}}, 5},
		{"", args{[]int{1, 3, 2, 2, 2}}, 4},
		{"", args{[]int{1, 2, 3, 4}}, 0},
		{"", args{[]int{1}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findUnsortedSubarray(tt.args.nums); got != tt.want {
				t.Errorf("findUnsortedSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
