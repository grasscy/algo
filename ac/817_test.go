package ac

import "testing"

func Test_numComponents(t *testing.T) {
	type args struct {
		head *ListNode
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{buildList([]int{0, 1, 2}), []int{1, 0}}, 1},
		{"", args{buildList([]int{0, 1, 2, 3}), []int{0, 1, 3}}, 2},
		{"", args{buildList([]int{0, 1, 2, 3, 4}), []int{0, 3, 1, 4}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numComponents(tt.args.head, tt.args.nums); got != tt.want {
				t.Errorf("numComponents() = %v, want %v", got, tt.want)
			}
		})
	}
}
