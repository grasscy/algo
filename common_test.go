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

func Test_buildTree(t *testing.T) {
	type args struct {
		nums []interface{}
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{"", args{[]interface{}{1, 2, 3, 4, "null", 2, 4, "null", "null", 4}}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildTree(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
