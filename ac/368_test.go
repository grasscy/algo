package ac

import (
	"reflect"
	"testing"
)

func Test_largestDivisibleSubset(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"", args{[]int{3, 4, 6, 8, 12, 16, 32}}, []int{4, 8, 16, 32}},
		{"", args{[]int{2, 3, 4, 9, 8}}, []int{2, 4, 8}},
		{"", args{[]int{4, 8, 10, 240}}, []int{4, 8, 240}},
		{"", args{[]int{2, 4, 6, 9, 19, 81, 729}}, []int{9, 81, 729}},
		{"", args{[]int{3, 4, 16, 8}}, []int{4, 8, 16}},
		{"", args{[]int{1, 2, 3}}, []int{1, 2}},
		{"", args{[]int{1, 2, 4, 8}}, []int{1, 2, 4, 8}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestDivisibleSubset(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("largestDivisibleSubset() = %v, want %v", got, tt.want)
			}
		})
	}
}
