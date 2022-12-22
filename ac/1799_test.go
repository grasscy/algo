package ac

import "testing"

func Test_maxScore(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{[]int{2557, 1997, 23, 599, 37, 1949, 701, 2557, 1997, 23, 599, 37, 1949, 701}}, 44324},
		{"", args{[]int{995853, 58850, 504472, 912272, 382454, 597688, 581332, 209547}}, 81},
		{"", args{[]int{1, 2}}, 1},
		{"", args{[]int{3, 4, 6, 8}}, 11},
		{"", args{[]int{1, 2, 3, 4, 5, 6}}, 14},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxScore(tt.args.nums); got != tt.want {
				t.Errorf("maxScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
