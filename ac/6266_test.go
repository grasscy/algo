package ac

import "testing"

func Test_smallestValue(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{15}, 5},
		{"", args{3}, 3},
		{"", args{99952}, 5},
		{"", args{4}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestValue(tt.args.n); got != tt.want {
				t.Errorf("smallestValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
