package ac

import "testing"

func Test_boxDelivering(t *testing.T) {
	type args struct {
		boxes      [][]int
		portsCount int
		maxBoxes   int
		maxWeight  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{[][]int{{2, 4}, {2, 5}, {3, 1}, {3, 2}, {3, 7}, {3, 1}, {4, 4}, {1, 3}, {5, 2}}, 5, 5, 7}, 14},
		{"", args{[][]int{{1, 1}, {2, 1}, {1, 1}}, 2, 3, 3}, 4},
		{"", args{[][]int{{1, 2}, {3, 3}, {3, 1}, {3, 1}, {2, 4}}, 3, 3, 6}, 6},
		{"", args{[][]int{{1, 4}, {1, 2}, {2, 1}, {2, 1}, {3, 2}, {3, 4}}, 3, 6, 7}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := boxDelivering(tt.args.boxes, tt.args.portsCount, tt.args.maxBoxes, tt.args.maxWeight); got != tt.want {
				t.Errorf("boxDelivering() = %v, want %v", got, tt.want)
			}
		})
	}
}
