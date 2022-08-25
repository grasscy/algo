package ac

import (
	"reflect"
	"testing"
)

func Test_findClosestElements(t *testing.T) {
	type args struct {
		arr []int
		k   int
		x   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"", args{[]int{1, 1, 1, 10, 10, 10}, 1, 9}, []int{10}},
		{"", args{[]int{1, 2, 3, 4, 5}, 4, 3}, []int{1, 2, 3, 4}},
		{"", args{[]int{1, 2, 3, 4, 5}, 4, -1}, []int{1, 2, 3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findClosestElements(tt.args.arr, tt.args.k, tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findClosestElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
