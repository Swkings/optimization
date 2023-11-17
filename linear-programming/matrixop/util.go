package matrixop

import (
	"github.com/swkings/optimization/util"
)

func I[T util.Number](n int) [][]T {
	var iMatrix [][]T = make([][]T, n)
	for i := 0; i < n; i++ {
		iMatrix[i] = make([]T, n)
		iMatrix[i][i] = 1
	}

	return iMatrix
}

func Swap[T any](arr []T, i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func NewEmptyMatrix[T util.Number](n, m int) [][]T {
	if n == 0 || m == 0 {
		return [][]T{}
	}
	matrix := make([][]T, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]T, m)
	}

	return matrix
}

// T matrix transpose
func T[S ~[][]T, T util.Number](matrix S) S {
	n, m := len(matrix), 0
	if n == 0 {
		return matrix
	}
	m = len(matrix[0])
	if m == 0 {
		return matrix
	}

	tMatrix := NewEmptyMatrix[T](m, n)
	for i := 0; i < n; i++ {
		iVector := matrix[i]
		for j := 0; j < m; j++ {
			if iVector[j] == 0 {
				continue
			}
			tMatrix[j][i] = iVector[j]
		}
	}

	return tMatrix
}

// DotProductVector obtain scalar product for vector
func DotProductVector[S ~[]T, T util.Number](vector1, vector2 S) T {
	var res T
	n1, n2 := len(vector1), len(vector2)
	if n1 != n2 || n1 == 0 {
		return res
	}

	for i := 0; i < n1; i++ {
		res += vector1[i] * vector2[i]
	}

	return res
}

// DotProductMatrix obtain scalar product for matrix
func DotProductMatrix[S ~[][]T, T util.Number](matrix1, matrix2 S) S {
	var res S

	n1, m1 := len(matrix1), 0
	n2, m2 := len(matrix2), 0
	if n1 == 0 || n2 == 0 {
		return res
	}
	m1, m2 = len(matrix1[0]), len(matrix2[0])

	if n1 != n2 || m1 != m2 {
		return res
	}

	res = NewEmptyMatrix[T](n1, m1)
	for i := 0; i < n1; i++ {
		for j := 0; j < m1; j++ {
			res[i][j] = matrix1[i][j] * matrix2[i][j]
		}
	}

	return res
}

// obtain vector product for matrix
func CrossProduct[S ~[][]T, T util.Number](matrix1, matrix2 S) S {
	var matrix S

	n1, m1 := len(matrix1), 0
	n2, m2 := len(matrix2), 0
	if n1 == 0 || n2 == 0 {
		return matrix
	}
	m1, m2 = len(matrix1[0]), len(matrix2[0])

	if m1 != n2 {
		return matrix
	}

	matrix = NewEmptyMatrix[T](n1, m2)
	for i := 0; i < n1; i++ {
		for j := 0; j < m2; j++ {
			for k := 0; k < m1; k++ {
				matrix[i][j] += matrix1[i][k] * matrix2[k][j]
			}
		}
	}

	return matrix
}
