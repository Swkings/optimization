package matrixop

import "github.com/swkings/optimization/util"

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
