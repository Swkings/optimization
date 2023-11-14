package matrixop

import "github.com/swkings/optimization/util"

func I[T util.Number](n int) [][]T {
	var iMaxtrix [][]T = make([][]T, n)
	for i := 0; i < n; i++ {
		iMaxtrix[i] = make([]T, n)
		iMaxtrix[i][i] = 1
	}

	return iMaxtrix
}

func Swap[T any](arr []T, i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
