package t1_199

import (
	"reflect"
	"testing"
)

func Test_mergeKLists(t *testing.T) {
	type args struct {
		lists []*ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"", args{[]*ListNode{buildList([]int{1, 4, 5}), buildList([]int{1, 3, 4}), buildList([]int{2, 6})}}, nil},
		{"", args{[]*ListNode{}}, nil},
		{"", args{[]*ListNode{{}}}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeKLists(tt.args.lists); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeKLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
