package algo

import "testing"

func Test_pathSum(t *testing.T) {
	type args struct {
		root      *TreeNode
		targetSum int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{buildTree([]interface{}{10, 5, -3, 3, 2, "null", 11, 3, -2, "null", 1}), 8}, 3},
		{"", args{buildTree([]interface{}{-2, "null", -3}), -5}, 1},
		{"", args{buildTree([]interface{}{5, 4, 8, 11, "null", 13, 4, 7, 2, "null", "null", 5, 1}), 22}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pathSum(tt.args.root, tt.args.targetSum); got != tt.want {
				t.Errorf("pathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
