package ac

import "testing"

func Test_reformatNumber(t *testing.T) {
	type args struct {
		number string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{"1-23-45 6"}, "123-456"},
		{"", args{"123 4-567"}, "123-45-67"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reformatNumber(tt.args.number); got != tt.want {
				t.Errorf("reformatNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
