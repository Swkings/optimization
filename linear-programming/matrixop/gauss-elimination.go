package matrixop

import (
	"github.com/swkings/optimization/util"
)

/**
 * 高斯消元法: 初等行变换
 */

func GetUpperTriangularByGaussElimination[T util.Number](matrix [][]T) {
	n, m := len(matrix), 0
	if n > 0 {
		m = len(matrix[0])
	}
	for i := 0; i < n; i++ {
		if i >= m {
			continue
		}
		minElementIndex, maxElementIndex := util.MaxMinListOpt(matrix[i:], func(item []T) T {
			return item[i]
		})
		minElementIndex += i
		maxElementIndex += i

		if matrix[i][maxElementIndex] == 0 {
			if i != minElementIndex {
				Swap(matrix, i, minElementIndex)
			}
		} else {
			if i != maxElementIndex {
				Swap(matrix, i, maxElementIndex)
			}
		}

		rate := matrix[i][i]
		if rate == 0 {
			continue
		}

		matrix[i] = util.ListTimes(matrix[i], []T{1 / rate})
		for nextIndex := i + 1; nextIndex < n; nextIndex++ {
			matrix[nextIndex] = util.ListMinus(matrix[nextIndex], util.ListTimes(append([]T{}, matrix[i]...), []T{matrix[nextIndex][i]}))
		}
	}
}

func GetLowerTriangularByGaussElimination[T util.Number](matrix [][]T) {
	n, m := len(matrix), 0
	if n > 0 {
		m = len(matrix[0])
	}
	for i := n - 1; i >= 0; i-- {
		if i < n-m {
			continue
		}
		minElementIndex, maxElementIndex := util.MaxMinListOpt(matrix[:i+1], func(item []T) T {
			return item[i]
		})

		if matrix[i][maxElementIndex] == 0 {
			if i != minElementIndex {
				Swap(matrix, i, minElementIndex)
			}
		} else {
			if i != maxElementIndex {
				Swap(matrix, i, maxElementIndex)
			}
		}

		rate := matrix[i][i]
		if rate == 0 {
			continue
		}

		matrix[i] = util.ListTimes(matrix[i], []T{1 / rate})
		for nextIndex := i - 1; nextIndex >= 0; nextIndex-- {
			matrix[nextIndex] = util.ListMinus(matrix[nextIndex], util.ListTimes(append([]T{}, matrix[i]...), []T{matrix[nextIndex][i]}))
		}
	}
}

func GetUnitMatrixByUpperTriangular[T util.Number](matrix [][]T) {
	n, m := len(matrix), 0
	if n > 0 {
		m = len(matrix[0])
	}
	if n < 2 || m < 2 {
		return
	}
	for i := 0; i < n-1; i++ {
		for k := i + 1; k < n; k++ {
			rate := matrix[i][k]
			if rate == 0 {
				continue
			}
			for j := k; j < m; j++ {
				matrix[i][j] -= rate * matrix[k][j]
			}
		}
	}
}

func GetUnitMatrixByLowerTriangular[T util.Number](matrix [][]T) {
	n, m := len(matrix), 0
	if n > 0 {
		m = len(matrix[0])
	}
	if n < 2 || m < 2 {
		return
	}
	for i := n - 1; i > 0; i-- {
		for k := i - 1; k >= 0; k-- {
			rate := matrix[i][k]
			if rate == 0 {
				continue
			}
			for j := k; j >= 0; j-- {
				matrix[i][j] -= rate * matrix[k][j]
			}
			matrix[i][m-1] -= rate * matrix[k][m-1]
		}
	}
}
