package algo

import (
	"reflect"
	"testing"
)

func Test_trimBST(t *testing.T) {
	type args struct {
		root *TreeNode
		low  int
		high int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{"", args{buildTree([]interface{}{1, 0, 2}), 1, 2}, nil},
		{"", args{buildTree([]interface{}{3, 0, 4, "null", 2, "null", "null", 1}), 1, 3}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trimBST(tt.args.root, tt.args.low, tt.args.high); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("trimBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
