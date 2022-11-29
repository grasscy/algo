package pass

import (
	"reflect"
	"testing"
)

func Test_findSubstring(t *testing.T) {
	type args struct {
		s     string
		words []string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"", args{"foobarfoobar", []string{"foo", "bar"}}, []int{0, 3, 6}},
		{"", args{"barfoothefoobarman", []string{"foo", "bar"}}, []int{0, 9}},
		{"", args{"wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSubstring(tt.args.s, tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
