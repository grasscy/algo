package t1_199

import "testing"

func Test_recoverTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
	}{
		{"", args{buildTree([]interface{}{3, 1, 4, "null", "null", 2})}},
		{"", args{buildTree([]interface{}{1, 3, "null", "null", 2})}},
		{"", args{buildTree([]interface{}{146, 71, -13, 55, "null", 231, 399, 321, "null", "null", "null", "null", "null", -33})}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recoverTree(tt.args.root)
		})
	}
}
