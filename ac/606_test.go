package ac

import "testing"

func Test_tree2str(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{buildTree([]interface{}{1, 2, 3, 4})}, "1(2(4))(3)"},
		{"", args{buildTree([]interface{}{1, 2, 3, "null", 4})}, "1(2()(4))(3)"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tree2str(tt.args.root); got != tt.want {
				t.Errorf("tree2str() = %v, want %v", got, tt.want)
			}
		})
	}
}
