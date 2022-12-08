package t1_199

import "testing"

func Test_isNumber(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{"+"}, false},
		{"", args{"4."}, true},
		{"", args{"+."}, false},
		{"", args{"4e+"}, false},
		{"", args{".1."}, false},
		{"", args{"e9"}, false},
		{"", args{"+6e-1"}, true},
		{"", args{"0"}, true},
		{"", args{"0089"}, true},
		{"", args{"-0.1"}, true},
		{"", args{"+3.14"}, true},
		{"", args{"-.9"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isNumber(tt.args.s); got != tt.want {
				t.Errorf("isNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
