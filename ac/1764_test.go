package ac

import "testing"

func Test_canChoose(t *testing.T) {
	type args struct {
		groups [][]int
		nums   []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{[][]int{{1, 2, 3}, {3, 4}}, []int{1, 2, 3, 3, 4}}, true},
		{"", args{[][]int{{9099312, -7882487, -1441304, 6624042, -6043305}}, []int{-1441304, 9099312, -7882487, -1441304, 6624042, -6043305, -1441304}}, true},
		{"", args{[][]int{{1, 2}}, []int{1, 3, 2}}, false},
		{"", args{[][]int{{21, 22, 21, 22, 21, 30}}, []int{21, 22, 21, 22, 21, 22, 21, 30}}, true},
		{"", args{[][]int{{352529, -4284030, 6431387, 2432677, -3305342, -4342915, -9007629, 3195451, -9587228, 5747770, 8106556, -2385247, 3207013, -8540809, 9400364, 6852329}, {9903943}},
			[]int{-8540809, 5747770, 2432677, -9587228, 3195451, 352529, -4284030, 6431387, 2432677, -3305342, -4342915, -9007629, 3195451, -9587228, 5747770, 8106556, -2385247, 3207013, -8540809, -8540809, 6852329, -9007629, 352529, 9903943}},
			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canChoose(tt.args.groups, tt.args.nums); got != tt.want {
				t.Errorf("canChoose() = %v, want %v", got, tt.want)
			}
		})
	}
}
