package algo

import (
	"reflect"
	"testing"
)

func Test_buildList(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"", args{[]int{1, 2, 3, 4, 5}}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildList(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildList() = %v, want %v", got, tt.want)
			}
		})
	}
}
