package ac

import "testing"

func Test_maxPathSum(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{buildTree([]interface{}{-2, -1})}, -1},
		{"", args{buildTree([]interface{}{-3})}, -3},
		{"", args{buildTree([]interface{}{-10, 9, 20, "null", "null", 15, 7})}, 42},
		{"", args{buildTree([]interface{}{1, 2, 3})}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPathSum(tt.args.root); got != tt.want {
				t.Errorf("maxPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
