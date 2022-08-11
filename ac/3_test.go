package ac

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
    type args struct {
        s string
    }
    tests := []struct {
        name string
        args args
        want int
    }{
        {"", args{"abba"}, 2},
        {"", args{""}, 0},
        {"", args{"au"}, 2},
        {"", args{"abcabcbb"}, 3},
        {"", args{"bbbbb"}, 1},
        {"", args{"pwwkew"}, 3},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
                t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
            }
        })
    }
}
