package pass

import "testing"

func Test_numDecodings(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{"230"}, 0},
		{"", args{"10011"}, 0},
		{"", args{"10"}, 1},
		{"", args{"11106"}, 2},
		{"", args{"12"}, 2},
		{"", args{"226"}, 3},
		{"", args{"0"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numDecodings(tt.args.s); got != tt.want {
				t.Errorf("numDecodings() = %v, want %v", got, tt.want)
			}
		})
	}
}
