package ac

import "testing"

func Test_reorderSpaces(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{"  this   is  a sentence "}, "this   is   a   sentence"},
		{"", args{" practice   makes   perfect"}, "practice   makes   perfect "},
		{"", args{"hello   world"}, "hello   world"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reorderSpaces(tt.args.text); got != tt.want {
				t.Errorf("reorderSpaces() = %v, want %v", got, tt.want)
			}
		})
	}
}
