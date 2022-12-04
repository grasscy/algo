package ac

import "testing"

func Test_closestCost(t *testing.T) {
	type args struct {
		baseCosts    []int
		toppingCosts []int
		target       int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{[]int{3, 10}, []int{2, 5}, 9}, 8},
		{"", args{[]int{3008, 3801, 3589, 808}, []int{680, 350, 9676, 23, 62, 8624, 2298, 1279, 364}, 7064}, 7064},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := closestCost(tt.args.baseCosts, tt.args.toppingCosts, tt.args.target); got != tt.want {
				t.Errorf("closestCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
