package matrixop

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/swkings/optimization/util"
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

func TestT(t *testing.T) {
	matrix := [][]float64{
		{2, 1, 1, 5},
		{3, 4, 5, 6},
		{5, 2, 1, 7},
	}
	fmt.Printf("origin matrix: %v\n\n", util.PrettyArrayTable(nil, matrix))

	tMatrix := T(matrix)
	fmt.Printf("transpose matrix: %v\n\n", util.PrettyArrayTable(nil, tMatrix))
}

func TestProduct(t *testing.T) {
	matrix := [][]float64{
		{2, 1, 1},
		{3, 4, 5},
		{5, 2, 1},
	}
	fmt.Printf("origin matrix: %v\n\n", util.PrettyArrayTable(nil, matrix))

	dotProduct := DotProductMatrix(matrix, matrix)
	fmt.Printf("DotProduct: %v\n\n", util.PrettyArrayTable(nil, dotProduct))

	rMatrix := GetInverseMatrixByGaussElimination(matrix)
	crossProduct := CrossProduct(matrix, rMatrix)
	fmt.Printf("CrossProduct: %v\n\n", util.PrettyArrayTable(nil, crossProduct))

	UnitMatrix(crossProduct)
	fmt.Printf("CrossProduct: %v\n\n", util.PrettyArrayTable(nil, crossProduct))
}
