package matrixop

import (
	"fmt"
	"slices"
	"testing"

	"github.com/swkings/optimization/util"
)

func TestGaussElimination(t *testing.T) {
	matrix := [][]float64{
		{2, 1, 1, 5},
		{3, 4, 5, 6},
		{5, 2, 1, 7},
	}
	fmt.Printf("origin matrix: %v\n\n", util.PrettyArrayTable(nil, matrix))

	matrix1 := slices.Clone(matrix)
	matrix11 := slices.Clone(matrix)
	UpperTriangularByGaussElimination(matrix1, matrix11)
	fmt.Printf("Upper:%v\n", util.PrettyArrayTable(nil, matrix1))
	fmt.Printf("Upper 11:%v\n", util.PrettyArrayTable(nil, matrix11))
	UnitMatrixByUpperTriangular(matrix1, matrix11)
	fmt.Printf("Transfer:%v\n\n", util.PrettyArrayTable(nil, matrix1))
	fmt.Printf("Transfer 11:%v\n\n", util.PrettyArrayTable(nil, matrix11))

	matrix2 := slices.Clone(matrix)
	matrix22 := slices.Clone(matrix)
	LowerTriangularByGaussElimination(matrix2, matrix22)
	fmt.Printf("Lower:%v\n", util.PrettyArrayTable(nil, matrix2))
	fmt.Printf("Lower 22:%v\n", util.PrettyArrayTable(nil, matrix22))
	UnitMatrixByLowerTriangular(matrix2, matrix22)
	fmt.Printf("Transfer:%v\n\n", util.PrettyArrayTable(nil, matrix2))
	fmt.Printf("Transfer 22:%v\n\n", util.PrettyArrayTable(nil, matrix22))

	fmt.Printf("origin matrix:%v\n\n", util.PrettyArrayTable(nil, matrix))
}
