package t1_199

import (
	"reflect"
	"testing"
)

func Test_fullJustify(t *testing.T) {
	type args struct {
		words    []string
		maxWidth int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"", args{[]string{"This", "is", "an", "example", "of", "text", "justification."}, 16}, []string{"This    is    an",
			"example  of text",
			"justification.  "}},
		{"", args{[]string{"What", "must", "be", "acknowledgment", "shall", "be"}, 16}, []string{"What   must   be",
			"acknowledgment  ",
			"shall be        "}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fullJustify(tt.args.words, tt.args.maxWidth); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fullJustify() = %v, want %v", got, tt.want)
			}
		})
	}
}
