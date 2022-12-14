package t1_199

import (
	"reflect"
	"testing"
)

func Test_threeSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"", args{[]int{0, 0, 0, 0}}, [][]int{{0, 0, 0}}},
		{"", args{[]int{-1, 0, 1, 2, -1, -4}}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{"", args{[]int{-2, 0, 0, 2, 2}}, [][]int{{-2, 0, 2}}},
		{"", args{[]int{-2, 0, 1, 1, 2}}, [][]int{{-2, 0, 2}, {-2, 1, 1}}},
		{"", args{[]int{}}, [][]int{}},
		{"", args{[]int{0}}, [][]int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSum(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
