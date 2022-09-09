package ac

import "testing"

func Test_shipWithinDays(t *testing.T) {
	type args struct {
		weights []int
		days    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{[]int{1, 2, 3, 1, 1}, 4}, 3},
		{"", args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5}, 15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shipWithinDays(tt.args.weights, tt.args.days); got != tt.want {
				t.Errorf("shipWithinDays() = %v, want %v", got, tt.want)
			}
		})
	}
}
