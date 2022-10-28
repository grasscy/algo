package algo

import "testing"

func Test_jump(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{[]int{1, 1, 1, 1}}, 3},
		{"", args{[]int{1, 2}}, 1},
		{"", args{[]int{2, 3, 1, 1, 4}}, 2},
		{"", args{[]int{2, 0, 2, 0, 1}}, 2},
		{"", args{[]int{2, 3, 0, 1, 4}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := jump(tt.args.nums); got != tt.want {
				t.Errorf("jump() = %v, want %v", got, tt.want)
			}
		})
	}
}
