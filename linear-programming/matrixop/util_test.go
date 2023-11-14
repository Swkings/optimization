package matrixop

import (
	"reflect"
	"testing"
)

func TestI(t *testing.T) {
	type args struct {
		n int
	}
	type testCase struct {
		name string
		args args
		want [][]float64
	}
	tests := []testCase{
		{
			name: "i-n2",
			args: args{n: 2},
			want: [][]float64{
				{1, 0},
				{0, 1},
			},
		},
		{
			name: "i-n3",
			args: args{n: 3},
			want: [][]float64{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := I[float64](tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("I() = %v, want %v", got, tt.want)
			}
		})
	}
}
