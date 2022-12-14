package ac

import "testing"

func Test_leastInterval(t *testing.T) {
	type args struct {
		tasks []byte
		n     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{[]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 50}, 104},
		{"", args{[]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2}, 8},
		{"", args{[]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 0}, 6},
		{"", args{[]byte{'A', 'A', 'A', 'A', 'A', 'A', 'B', 'C', 'D', 'E', 'F', 'G'}, 2}, 16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := leastInterval(tt.args.tasks, tt.args.n); got != tt.want {
				t.Errorf("leastInterval() = %v, want %v", got, tt.want)
			}
		})
	}
}
