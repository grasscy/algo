package ac

import "testing"

func Test_knightProbability(t *testing.T) {
	type args struct {
		n      int
		k      int
		row    int
		column int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"", args{3, 2, 0, 0}, 0.0625},
		{"", args{10, 13, 5, 3}, 0.11734},
		{"", args{1, 0, 0, 0}, 1.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := knightProbability(tt.args.n, tt.args.k, tt.args.row, tt.args.column); got != tt.want {
				t.Errorf("knightProbability() = %v, want %v", got, tt.want)
			}
		})
	}
}
