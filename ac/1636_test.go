package ac

import (
	"reflect"
	"testing"
)

func Test_frequencySort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"", args{[]int{1, 1, 2, 2, 2, 3}}, []int{3, 1, 1, 2, 2, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := frequencySort(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("frequencySort() = %v, want %v", got, tt.want)
			}
		})
	}
}
