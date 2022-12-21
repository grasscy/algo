package t1_199

import (
	"reflect"
	"testing"
)

func Test_wordBreak(t *testing.T) {
	type args struct {
		s        string
		wordDict []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"", args{"a", []string{"a"}}, []string{"a"}},
		{"", args{"catsanddog", []string{"cat", "cats", "and", "sand", "dog"}}, []string{"cats and dog", "cat sand dog"}},
		{"", args{"pineapplepenapple", []string{"apple", "pen", "applepen", "pine", "pineapple"}}, []string{"pine apple pen apple", "pineapple pen apple", "pine applepen apple"}},
		{"", args{"catsandog", []string{"cats", "dog", "sand", "and", "cat"}}, []string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordBreak(tt.args.s, tt.args.wordDict); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("wordBreak() = %v, want %v", got, tt.want)
			}
		})
	}
}
