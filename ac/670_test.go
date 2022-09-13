package ac

import "testing"

func Test_maximumSwap(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{1993}, 9913},
		{"", args{2736}, 7236},
		{"", args{9973}, 9973},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumSwap(tt.args.num); got != tt.want {
				t.Errorf("maximumSwap() = %v, want %v", got, tt.want)
			}
		})
	}
}
