package t1_199

import "testing"

func Test_multiply(t *testing.T) {
	type args struct {
		num1 string
		num2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{"3876620623801494171", "6529364523802684779"}, "25311869173291309803201639921518923209"},
		{"", args{"3896737933784656127", "925101087184894"}, "3604876499018802870538090258945538"},
		{"", args{"925101087184894", "3896737933784656127"}, "3604876499018802870538090258945538"},
		{"", args{"498828660196", "840477629533"}, "419254329864656431168468"},
		{"", args{"123", "456"}, "56088"},
		{"", args{"2", "3"}, "6"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multiply(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}
