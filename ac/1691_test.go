package ac

import (
	"testing"
)

func Test_maxHeight(t *testing.T) {
	type args struct {
		cuboids [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{[][]int{{74, 66, 55}, {97, 97, 38}, {21, 69, 43}, {88, 83, 87}, {11, 4, 96}, {17, 98, 21}, {45, 43, 12}, {19, 67, 24}, {72, 5, 81}, {30, 53, 100}, {38, 30, 29}, {81, 53, 42}, {78, 80, 9}, {3, 80, 66}, {74, 16, 92}, {18, 17, 70}, {66, 88, 56}, {7, 51, 49}, {9, 59, 13}, {44, 67, 21}, {9, 95, 14}, {88, 100, 37}, {23, 76, 24}, {15, 38, 41}, {47, 98, 66}, {25, 39, 23}, {74, 49, 28}, {100, 82, 62}, {96, 73, 52}, {9, 22, 5}, {83, 99, 28}, {9, 35, 5}, {26, 53, 33}, {53, 98, 93}}}, 605},
		{"", args{[][]int{{86, 17, 6}, {96, 59, 97}, {78, 72, 34}, {17, 79, 41}, {51, 64, 43}, {94, 77, 19}, {28, 44, 42}, {72, 77, 7}, {31, 90, 43}, {50, 90, 93}, {50, 55, 35}, {22, 63, 96}, {36, 41, 26}, {50, 39, 64}, {93, 27, 2}, {26, 52, 80}, {32, 64, 96}, {95, 61, 44}, {16, 64, 81}, {50, 85, 24}, {78, 80, 3}, {87, 61, 82}, {97, 75, 92}, {45, 2, 75}, {41, 18, 1}}}, 501},
		{"", args{[][]int{{36, 46, 41}, {15, 100, 100}, {75, 91, 59}, {13, 82, 64}}}, 182},
		{"", args{[][]int{{38, 25, 45}, {76, 35, 3}}}, 76},
		{"", args{[][]int{{50, 26, 84}, {2, 55, 62}, {64, 63, 72}}}, 134},
		{"", args{[][]int{{50, 45, 20}, {95, 37, 53}, {45, 23, 12}}}, 190},
		{"", args{[][]int{{7, 11, 17}, {7, 17, 11}, {11, 7, 17}, {11, 17, 7}, {17, 7, 11}, {17, 11, 7}}}, 102},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxHeight(tt.args.cuboids); got != tt.want {
				t.Errorf("maxHeight() = %v, want %v", got, tt.want)
			}
		})
	}
}
