package ac

import "testing"

func Test_maximalSquare(t *testing.T) {
    type args struct {
        matrix [][]byte
    }
    tests := []struct {
        name string
        args args
        want int
    }{
        {"", args{[][]byte{{'0', '1'}}}, 1},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := maximalSquare(tt.args.matrix); got != tt.want {
                t.Errorf("maximalSquare() = %v, want %v", got, tt.want)
            }
        })
    }
}
