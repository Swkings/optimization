package mop

import (
	"fmt"
	"sort"

	"github.com/swkings/optimization/util"
)

//	@input: data [][]T
//	@return []int, first rank index list of data
// NDSet cal non-dominated rank
func NDSetRank[T Number](data [][]T) [][]int {
	objNum := len(data[0])
	if objNum < 2 {
		return [][]int{}
	}
	if objNum == 2 {
		return NDSetRankLD(data)
	} else {
		return NDSetRankHD(data)
	}
}

//	@input: data [][2]T
//	@return []int, first rank index list of data
// NDSetLD cal non-dominated rank for 2-dimension
func NDSetRankLD[T Number](data [][]T) [][]int {
	if len(data[0]) != 2 {
		return [][]int{}
	}
	// arg sort for data
	index := util.Range(len(data))
	sort.Slice(index, func(i, j int) bool {
		posI, posJ := index[i], index[j]
		if data[posI][0] == data[posJ][0] {
			return data[posI][1] < data[posJ][1]
		}
		return data[posI][0] < data[posJ][0]
	})
	// result array
	rankSet := [][]int{
		{index[0]},
	}
	minRankValueSet := [][]T{
		data[index[0]],
	}
	prevValueRank := 0
	for i, posI := range index {
		if i == 0 {
			continue
		}
		compareValue := data[posI]
		minRank, maxRank := 0, prevValueRank
		if compareValue[1] < minRankValueSet[prevValueRank][1] {
			minRank, maxRank = 0, prevValueRank
		} else if AllEqual(compareValue, minRankValueSet[prevValueRank]) {
			minRank, maxRank = prevValueRank, prevValueRank
		} else {
			minRank, maxRank = prevValueRank+1, len(rankSet)
		}
		iRank := findRankLD(rankSet, minRankValueSet, compareValue, minRank, maxRank)
		rankSet, minRankValueSet = addToRankSetLD(rankSet, minRankValueSet, posI, compareValue, iRank)
		prevValueRank = iRank
	}

	return rankSet
}

// findRankLD binary search compareValue rank
func findRankLD[T Number](rankSet [][]int, minRankValueSet [][]T, compareValue []T, minRank int, maxRank int) int {
	for minRank < maxRank {
		midRank := (minRank + maxRank) / 2
		midRankMinValue := minRankValueSet[midRank]
		if compareValue[1] < midRankMinValue[1] {
			maxRank = midRank
		} else if compareValue[1] > midRankMinValue[1] {
			minRank = midRank + 1
		} else {
			return midRank + 1
		}
	}
	return minRank
}

func addToRankSetLD[T Number](rankSet [][]int, minRankValueSet [][]T, index int, compareValue []T, iRank int) ([][]int, [][]T) {
	if iRank >= len(rankSet) {
		rankSet = append(rankSet, []int{index})
		minRankValueSet = append(minRankValueSet, compareValue)
	} else {
		rankSet[iRank] = append(rankSet[iRank], index)
		if compareValue[1] < minRankValueSet[iRank][1] {
			minRankValueSet[iRank] = compareValue
		}
	}

	return rankSet, minRankValueSet
}

//	@input: data [][>=3]T
//	@return []int, first rank index list of data
// NDSetHD cal non-dominated rank for >= 3-dimension
func NDSetRankHD[T Number](data [][]T) [][]int {
	if len(data[0]) < 2 {
		return [][]int{}
	}
	// arg sort for data
	index := util.Range(len(data))
	sort.Slice(index, func(i, j int) bool {
		posI, posJ := index[i], index[j]
		for inI := range data[posI] {
			if data[posI][inI] == data[posJ][inI] {
				continue
			}
			return data[posI][inI] < data[posJ][inI]
		}

		return data[posI][0] < data[posJ][0]
	})
	fmt.Printf("index: %v\n", index)
	// result array
	rankSet := [][]int{
		{index[0]},
	}
	minRankValueSet := [][][]T{
		{data[index[0]]},
	}
	prevValueRank := 0
	for i, posI := range index {
		if i == 0 {
			continue
		}
		prevValue := data[index[i-1]]
		nextValue := data[posI]
		relevantValue := Dominate(prevValue[1:], nextValue[1:])
		minRank, maxRank := prevValueRank, prevValueRank
		if prevValue[0] == nextValue[0] && relevantValue == AEqualB {
			minRank, maxRank = prevValueRank, prevValueRank
		} else if relevantValue == AEqualB {
			minRank, maxRank = prevValueRank+1, prevValueRank+1
		} else if relevantValue == ADominateB {
			minRank, maxRank = prevValueRank+1, len(rankSet)
		} else if relevantValue == ANonDominatedB {
			minRank, maxRank = 0, len(rankSet)
		} else if relevantValue == BDominateA {
			minRank, maxRank = 0, prevValueRank
		}
		iRank := findRankHD(rankSet, minRankValueSet, nextValue, minRank, maxRank)
		rankSet, minRankValueSet = addToRankSetHD(rankSet, minRankValueSet, posI, nextValue, iRank)
		prevValueRank = iRank
	}
	return rankSet
}

// findRankLD binary search compareValue rank
func findRankHD[T Number](rankSet [][]int, minRankValueSet [][][]T, compareValue []T, minRank int, maxRank int) int {
	for minRank < maxRank {
		midRank := (minRank + maxRank) / 2

		midRankMinValueSet := minRankValueSet[midRank]
		isNonDominated := true
		for i := len(midRankMinValueSet) - 1; i >= 0; i-- {
			if Dominate(midRankMinValueSet[i], compareValue) != ADominateB {
				continue
			} else {
				minRank = midRank + 1
				isNonDominated = false
				break
			}
		}
		if isNonDominated {
			maxRank = midRank
		}
	}

	return minRank
}

func addToRankSetHD[T Number](rankSet [][]int, minRankValueSet [][][]T, index int, compareValue []T, iRank int) ([][]int, [][][]T) {
	if iRank >= len(rankSet) {
		rankSet = append(rankSet, []int{index})
		minRankValueSet = append(minRankValueSet, [][]T{compareValue})
	} else {
		rankSet[iRank] = append(rankSet[iRank], index)
		minRankValueSet[iRank] = append(minRankValueSet[iRank], compareValue)
	}

	return rankSet, minRankValueSet
}
