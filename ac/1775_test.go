package ac

import "testing"

func Test_minOperations(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{[]int{6}, []int{5, 4, 1, 2, 5, 3, 2, 5, 1, 5, 6, 6, 3, 6, 1, 6}}, -1},
		{"", args{[]int{6, 6}, []int{1}}, 3},
		{"", args{[]int{1, 2, 3, 4, 5, 6}, []int{1, 1, 2, 2, 2, 2}}, 3},
		{"", args{[]int{1, 1, 1, 1, 1, 1, 1}, []int{6}}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minOperations(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("minOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
