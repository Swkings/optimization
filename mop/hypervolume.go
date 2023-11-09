package mop

import (
	"sort"

	"github.com/swkings/optimization/util"
)

// pareto solution set
type ParetoSolutionSet[T Number] struct {
	Points    [][]T
	Dimension int
	Size      int
}

func NewParetoSolutionSet[T Number](data [][]T) ParetoSolutionSet[T] {
	size := len(data)
	dimension := 0
	if size > 0 {
		dimension = len(data[0])
	}
	return ParetoSolutionSet[T]{
		Points:    data,
		Dimension: dimension,
		Size:      size,
	}
}

func (h ParetoSolutionSet[T]) NadirPoint() []T {
	return NadirPoint(h.Points)
}

func (h ParetoSolutionSet[T]) IdealPoint() []T {
	return IdealPoint(h.Points)
}

func (h ParetoSolutionSet[T]) RefPoint() []T {
	return h.NadirPoint()
}

func (h ParetoSolutionSet[T]) NDSet(points [][]T) [][]T {
	index := NDSet(points)
	ndSet := [][]T{}
	for _, item := range index {
		ndSet = append(ndSet, points[item])
	}

	return ndSet
}

func (h ParetoSolutionSet[T]) VolumeBetween(a []T, b []T) float64 {
	return VolumeBetween(a, b)
}

func (h ParetoSolutionSet[T]) HyperVolume2D() float64 {
	return HyperVolume2D(h.Points)
}

func (h ParetoSolutionSet[T]) HyperVolume(refPointArg ...[]T) float64 {
	h.Points = util.Unique(h.Points)
	var refPoint []T
	if len(refPointArg) > 0 {
		refPoint = refPointArg[0]
	} else {
		refPoint = h.RefPoint()
	}
	return h.hyperVolume(h.NDSet(h.Points), refPoint)
}

func (h ParetoSolutionSet[T]) hyperVolume(point [][]T, refPoint []T) float64 {
	size := len(point)
	if size == 0 {
		return 0
	}
	dim := len(refPoint)

	if dim == 1 {
		return float64(refPoint[0] - point[0][0])
	}
	var sortPoint [][]T = make([][]T, len(point))
	copy(sortPoint, point)
	sort.Slice(sortPoint, func(i, j int) bool {
		return sortPoint[i][0] < sortPoint[j][0]
	})

	prevPoint := refPoint
	var hv float64 = 0

	for i := size - 1; i >= 0; i-- {
		currentPoint := sortPoint[i]
		width := prevPoint[0] - currentPoint[0]

		tmpPoints := append([][]T{}, sortPoint[:i+1]...)
		for idx := range tmpPoints {
			tmpPoints[idx] = append([]T{}, tmpPoints[idx][1:]...)
		}

		ndSet := h.NDSet(tmpPoints)
		restRef := append([]T{}, refPoint[1:]...)

		subHV := h.hyperVolume(ndSet, restRef)

		hv += float64(width) * subHV
		prevPoint = currentPoint
	}

	return hv
}

// Calculate a nadir point
func NadirPoint[T Number](data [][]T) []T {
	dataL := len(data)
	if dataL == 0 {
		return []T{}
	}

	dim := len(data[0])

	var refPoint []T = make([]T, len(data[0]))
	copy(refPoint, data[0])

	for i := 0; i < dim; i++ {
		for _, item := range data {
			refPoint[i] = util.Max(refPoint[i], item[i])
		}
	}

	return refPoint
}

// Calculate a ideal point
func IdealPoint[T Number](data [][]T) []T {
	dataL := len(data)
	if dataL == 0 {
		return []T{}
	}

	dim := len(data[0])

	var refPoint []T = make([]T, len(data[0]))
	copy(refPoint, data[0])

	for i := 0; i < dim; i++ {
		for _, item := range data {
			refPoint[i] = util.Min(refPoint[i], item[i])
		}
	}

	return refPoint
}

// Calculate a reference point
func RefPoint[T Number](data [][]T) []T {
	return NadirPoint(data)
}

func VolumeBetween[T Number](a []T, b []T) float64 {
	aL, bL := len(a), len(b)
	if aL != bL {
		return 0
	}

	var volume float64 = 1
	for i := 0; i < aL; i++ {
		volume *= float64(a[i] - b[i])
	}

	return util.Ternary(volume < 0, -volume, volume)
}

// Compute HyperVolume for 2-D data
func HyperVolume2D[T Number](data [][]T, refPointArg ...[]T) float64 {
	var hv float64 = 0
	dataL := len(data)
	if dataL == 0 {
		return hv
	}

	var refPoint []T
	if len(refPointArg) > 0 {
		refPoint = refPointArg[0]
	} else {
		refPoint = RefPoint(data)
	}

	if dataL == 1 {
		return VolumeBetween(data[0], refPoint)
	}

	var points [][]T = make([][]T, len(data))
	copy(points, data)

	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	w := refPoint[0] - points[0][0]
	for i := 0; i < dataL-1; i++ {
		hv += float64((points[i+1][1] - points[i][1]) * w)
		w = util.Max(w, refPoint[0]-points[i+1][0])
	}
	hv += float64((refPoint[1] - points[dataL-1][1]) * w)

	return hv
}
