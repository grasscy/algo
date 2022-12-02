package t1_199

import "testing"

func Test_canJump(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{[]int{1, 0, 1, 0}}, false},
		{"", args{[]int{1, 1, 2, 2, 0, 1, 1}}, true},
		{"", args{[]int{0}}, true},
		{"", args{[]int{0, 2, 3}}, false},
		{"", args{[]int{2, 3, 1, 1, 4}}, true},
		{"", args{[]int{3, 2, 1, 0, 4}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canJump(tt.args.nums); got != tt.want {
				t.Errorf("canJump() = %v, want %v", got, tt.want)
			}
		})
	}
}
