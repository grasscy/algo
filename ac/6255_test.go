package ac

import "testing"

func Test_minScore(t *testing.T) {
	type args struct {
		n     int
		roads [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{16, [][]int{{6, 10, 2373}, {2, 12, 9953}, {12, 3, 8785}, {3, 2, 1055}, {11, 1, 4921}, {6, 11, 5874}, {11, 3, 1369}, {3, 16, 1669}, {16, 11, 9097}, {1, 6, 9547}, {12, 11, 2554}, {10, 11, 5446}, {4, 11, 4889}, {10, 4, 8651}, {4, 12, 4138}}}, 1055},
		{"", args{4, [][]int{{1, 2, 9}, {2, 3, 6}, {2, 4, 5}, {1, 4, 7}}}, 5},
		{"", args{4, [][]int{{1, 2, 2}, {1, 3, 4}, {3, 4, 7}}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minScore(tt.args.n, tt.args.roads); got != tt.want {
				t.Errorf("minScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
