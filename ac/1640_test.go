package ac

import "testing"

func Test_canFormArray(t *testing.T) {
	type args struct {
		arr    []int
		pieces [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{[]int{15, 88}, [][]int{{88}, {15}}}, true},
		{"", args{[]int{49, 18, 16}, [][]int{{16, 18, 49}}}, false},
		{"", args{[]int{91, 4, 64, 78}, [][]int{{78}, {4, 64}, {91}}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canFormArray(tt.args.arr, tt.args.pieces); got != tt.want {
				t.Errorf("canFormArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
