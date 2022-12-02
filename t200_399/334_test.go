package t200_399

import "testing"

func Test_increasingTriplet(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{[]int{1, 2, 1, 3}}, true},
		{"", args{[]int{1, 6, 2, 5, 1}}, true},
		{"", args{[]int{2, 1, 5, 0, 4, 6}}, true},
		{"", args{[]int{1, 2, 3, 4, 5}}, true},
		{"", args{[]int{5, 4, 3, 2, 1}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := increasingTriplet(tt.args.nums); got != tt.want {
				t.Errorf("increasingTriplet() = %v, want %v", got, tt.want)
			}
		})
	}
}
