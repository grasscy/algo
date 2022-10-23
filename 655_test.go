package algo

import (
	"reflect"
	"testing"
)

func Test_printTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{"", args{buildTree([]interface{}{1, 2})}, [][]string{
			{"", "1", ""},
			{"2", "", ""}}},
		{"", args{buildTree([]interface{}{1, 2, 3, "null", 4})}, [][]string{
			{"", "", "", "1", "", "", ""},
			{"", "2", "", "", "", "3", ""},
			{"", "", "4", "", "", "", ""}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := printTree(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("printTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
