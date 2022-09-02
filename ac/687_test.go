package ac

import "testing"

func Test_longestUnivaluePath(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{buildTree([]interface{}{5, 4, 5, 1, 1, 5})}, 2},
		{"", args{buildTree([]interface{}{1, 4, 5, 4, 4, 5})}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestUnivaluePath(tt.args.root); got != tt.want {
				t.Errorf("longestUnivaluePath() = %v, want %v", got, tt.want)
			}
		})
	}
}
