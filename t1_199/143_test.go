package t1_199

import "testing"

func Test_reorderList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
	}{
		{"", args{buildList([]int{1, 2, 3, 4})}},
		{"", args{buildList([]int{1, 2, 3, 4, 5})}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reorderList(tt.args.head)
		})
	}
}
