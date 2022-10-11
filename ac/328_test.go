package ac

import (
	"reflect"
	"testing"
)

func Test_oddEvenList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"", args{buildList([]int{1, 2, 3, 4, 5, 6, 7, 8})}, nil},
		{"", args{buildList([]int{1, 2, 3, 4, 5})}, nil},
		{"", args{buildList([]int{2, 1, 3, 5, 6, 4, 7})}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := oddEvenList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("oddEvenList() = %v, want %v", got, tt.want)
			}
		})
	}
}
