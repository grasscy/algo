package algo

import "testing"

func Test_maxProduct(t *testing.T) {
    type args struct {
        nums []int
    }
    tests := []struct {
        name string
        args args
        want int
    }{
        {"", args{[]int{3, -1, 4}}, 4},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := maxProduct(tt.args.nums); got != tt.want {
                t.Errorf("maxProduct() = %v, want %v", got, tt.want)
            }
        })
    }
}
