package ac

import "testing"

func Test_subArrayRanges(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"", args{[]int{1, 2, 3}}, 4},
		{"", args{[]int{1, 3, 3}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subArrayRanges(tt.args.nums); got != tt.want {
				t.Errorf("subArrayRanges() = %v, want %v", got, tt.want)
			}
		})
	}
}
