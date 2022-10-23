package algo

import "testing"

func Test_findTarget(t *testing.T) {
	type args struct {
		root *TreeNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{buildTree([]interface{}{5, 3, 6, 2, 4, "null", 7}), 9}, true},
		{"", args{buildTree([]interface{}{2, 0, 3, -4, 1}), -1}, true},
		{"", args{buildTree([]interface{}{5, 3, 6, 2, 4, "null", 7}), 28}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTarget(tt.args.root, tt.args.k); got != tt.want {
				t.Errorf("findTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}
