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
	GetUpperTriangularByGaussElimination(matrix1)
	fmt.Printf("Upper:%v\n", util.PrettyArrayTable(nil, matrix1))
	GetUnitMatrixByUpperTriangular(matrix1)
	fmt.Printf("Transfer:%v\n\n", util.PrettyArrayTable(nil, matrix1))

	matrix2 := slices.Clone(matrix)
	GetLowerTriangularByGaussElimination(matrix2)
	fmt.Printf("Lower:%v\n", util.PrettyArrayTable(nil, matrix2))
	GetUnitMatrixByLowerTriangular(matrix2)
	fmt.Printf("Transfer:%v\n\n", util.PrettyArrayTable(nil, matrix2))

	fmt.Printf("origin matrix:%v\n\n", util.PrettyArrayTable(nil, matrix))
}
