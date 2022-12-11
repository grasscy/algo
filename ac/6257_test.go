package ac

import "testing"

func Test_deleteGreatestValue(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{[][]int{{1, 2, 4}, {3, 3, 1}}}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteGreatestValue(tt.args.grid); got != tt.want {
				t.Errorf("deleteGreatestValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
