package t200_399

import "testing"

func Test_canFinish(t *testing.T) {
	type args struct {
		numCourses    int
		prerequisites [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{4, [][]int{{0, 1}, {3, 1}, {1, 3}, {3, 2}}}, false},
		{"", args{8, [][]int{{1, 0}, {2, 6}, {1, 7}, {6, 4}, {7, 0}, {0, 5}}}, true},
		{"", args{5, [][]int{{1, 4}, {2, 4}, {3, 1}, {3, 2}}}, true},
		{"", args{2, [][]int{{1, 0}}}, true},
		{"", args{2, [][]int{{1, 0}, {0, 1}}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canFinish(tt.args.numCourses, tt.args.prerequisites); got != tt.want {
				t.Errorf("canFinish() = %v, want %v", got, tt.want)
			}
		})
	}
}
