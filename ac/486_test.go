package ac

import "testing"

func TestPredictTheWinner(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{[]int{1, 5, 2}}, false},
		{"", args{[]int{1, 5, 233, 7}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PredictTheWinner(tt.args.nums); got != tt.want {
				t.Errorf("PredictTheWinner() = %v, want %v", got, tt.want)
			}
		})
	}
}
