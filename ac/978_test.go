package ac

import "testing"

func Test_maxTurbulenceSize(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{[]int{9, 4, 2, 10, 7, 8, 8, 1, 9}}, 5},
		{"", args{[]int{4, 8, 12, 16}}, 2},
		{"", args{[]int{100}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxTurbulenceSize(tt.args.arr); got != tt.want {
				t.Errorf("maxTurbulenceSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
