package ac

import "testing"

func Test_minFlipsMonoIncr(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{"00110"}, 1},
		{"", args{"010110"}, 2},
		{"", args{"00011000"}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minFlipsMonoIncr(tt.args.s); got != tt.want {
				t.Errorf("minFlipsMonoIncr() = %v, want %v", got, tt.want)
			}
		})
	}
}
