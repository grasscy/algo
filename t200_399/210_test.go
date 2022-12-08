package t200_399

import (
	"reflect"
	"testing"
)

func Test_findOrder(t *testing.T) {
	type args struct {
		numCourses    int
		prerequisites [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"", args{2, [][]int{{0, 1}}}, []int{1, 0}},
		{"", args{2, [][]int{{1, 0}, {0, 1}}}, []int{}},
		{"", args{4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}}, []int{0, 2, 1, 3}},
		{"", args{2, [][]int{{1, 0}}}, []int{0, 1}},
		{"", args{1, [][]int{}}, []int{0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findOrder(tt.args.numCourses, tt.args.prerequisites); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
