package pass

import (
    "reflect"
    "testing"
)

func Test_sortList(t *testing.T) {
    type args struct {
        head *ListNode
    }
    tests := []struct {
        name string
        args args
        want *ListNode
    }{
        {"", args{buildList([]int{4, 2, 1, 3})}, nil},
        {"", args{buildList([]int{-1, 5, 3, 4, 0})}, nil},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := sortList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("sortList() = %v, want %v", got, tt.want)
            }
        })
    }
}
