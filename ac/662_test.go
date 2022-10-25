package ac

import "testing"

func Test_widthOfBinaryTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{buildTree([]interface{}{1, 3, 2, 5, "null", "null", 9, 6, "null", 7})}, 7},
		{"", args{buildTree([]interface{}{0, 0, 0, "null", 0, 0, "null", "null", 0, 0, "null", "null", 0, 0, "null", "null", 0, 0, "null", "null", 0, 0, "null", "null", 0, 0, "null", "null", 0, 0, "null", "null", 0, 0, "null", "null", 0, 0, "null", "null", 0, 0, "null", "null", 0, 0, "null", "null", 0, 0, "null", "null", 0, 0, "null", "null", 0, 0, "null", "null", 0, 0, "null", "null", 0, 0, "null", "null", 0, 0, "null", "null", 0, 0, "null", "null", 0, 0, "null", "null", 0, 0, "null", "null", 0, 0, "null", "null", 0, 0, "null", "null", 0, 0, "null", "null", 0, 0, "null", "null", 0, 0, "null", "null", 0, 0, "null", "null", 0, 0, "null", "null", 0, 0, "null", "null", 0, 0, "null", "null", 0, 0, "null", "null", 0, 0, "null", "null", 0, 0, "null", "null", 0, 0, "null", "null", 0, 0, "null", "null", 0, 0, "null", "null", 0, 0, "null", "null", 0, 0, "null", "null", 0, 0, "null", "null", 0, 0, "null", "null", 0, 0, "null", "null", 0, 0, "null", "null", 0, 0, "null", "null", 0, 0, "null", "null", 0, 0, "null", "null", 0, 0, "null", "null", 0, 0, "null", "null", 0, 0, "null", "null", 0, 0, "null", "null", 0, 0, "null", "null", 0, 0, "null", "null", 0, 0, "null", "null", 0, 0, "null", "null", 0, 0, "null", "null", 0, 0, "null", "null", 0, 0, "null", "null", 0, 0, "null", "null", 0, 0, "null", "null", 0, 0, "null", "null", 0, 0, "null", "null", 0, 0, "null", "null", 0, 0, "null", "null", 0, 0, "null", "null", 0, 0, "null", "null", 0, 0, "null", "null", 0, 0, "null", "null", 0, 0, "null", "null", 0, 0, "null"})}, 2},
		{"", args{buildTree([]interface{}{1, 1, 1, 1, 1, 1, 1, "null", "null", "null", 1, "null", "null", "null", "null", 2, 2, 2, 2, 2, 2, 2, "null", 2, "null", "null", 2, "null", 2})}, 8},
		{"", args{buildTree([]interface{}{1, 3, 2, 5, 3, "null", 9})}, 4},
		{"", args{buildTree([]interface{}{1, 3, 2, 5})}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := widthOfBinaryTree(tt.args.root); got != tt.want {
				t.Errorf("widthOfBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
