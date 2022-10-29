package ac

import "testing"

func Test_canIWin(t *testing.T) {
	type args struct {
		maxChoosableInteger int
		desiredTotal        int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{5, 50}, false},
		{"", args{4, 6}, true},
		{"", args{10, 11}, false},
		{"", args{10, 0}, true},
		{"", args{10, 1}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canIWin(tt.args.maxChoosableInteger, tt.args.desiredTotal); got != tt.want {
				t.Errorf("canIWin() = %v, want %v", got, tt.want)
			}
		})
	}
}
