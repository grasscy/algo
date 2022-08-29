package algo

import (
	"reflect"
	"testing"
)

func Test_buildTree(t *testing.T) {
	type args struct {
		root []interface{}
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{"", args{[]interface{}{10, 5, -3, 3, 2, "null", 11, 3, -2, "null", 1}}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildTree(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
