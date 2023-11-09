package mop

import (
	"sort"

	"github.com/swkings/optimization/util"
)

//	@input: data [][]T
//	@return []int, first rank index list of data
// NDSet cal non-dominated rank
func NDSet[T Number](data [][]T) []int {
	objNum := len(data[0])
	if objNum < 2 {
		_, minIndex := util.MinListOpt(data, func(item []T) int {
			return int(item[0])
		})
		res := []int{}
		for i, item := range data {
			if item[0] == data[minIndex][0] {
				res = append(res, i)
			}
		}
		return res
	}
	if objNum == 2 {
		return NDSetLD(data)
	} else {
		return NDSetHD(data)
	}
}

// f := func [V Number](item []V) V {
// 	return item[0]
// }

//	@input: data [][2]T
//	@return []int, first rank index list of data
// NDSetLD cal non-dominated rank for 2-dimension
func NDSetLD[T Number](data [][]T) []int {
	if len(data[0]) != 2 {
		return []int{}
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
	// first value must be first rank
	firstRankSet := []int{index[0]}
	// record first rank min value
	minRankValue := data[index[0]]
	for i, posI := range index {
		if i == 0 {
			continue
		}
		compareValue := data[posI]
		if compareValue[1] < minRankValue[1] { // if value less minRankValue, the index of compareValue must be first rank
			firstRankSet = append(firstRankSet, posI)
			minRankValue = compareValue
		} else if compareValue[1] == minRankValue[1] && compareValue[0] == minRankValue[0] { // compareValue equal minRankValue, add into first rank
			firstRankSet = append(firstRankSet, posI)
		}
	}

	return firstRankSet
}

//	@input: data [][>=3]T
//	@return []int, first rank index list of data
// NDSetHD cal non-dominated rank for >= 3-dimension
func NDSetHD[T Number](data [][]T) []int {
	if len(data[0]) < 2 {
		return []int{}
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
	// first value must be first rank
	firstRankSet := []int{index[0]}
	// compare array, try to reduce the number of unrelated elements in the array and the number of comparisons
	tempSet := [][]T{data[index[0]]}
	// combination element, consisting of the minimum values of each dimension
	combineElem := data[index[0]]
	prevValue := data[index[0]]
	for i, posI := range index {
		if i == 0 {
			continue
		}
		compareValue := data[posI]
		// first compare with composite elements
		relevantValue := Dominate(combineElem[1:], compareValue[1:])
		preComRelevant := Dominate(prevValue[1:], compareValue[1:])
		if relevantValue != ADominateB && relevantValue != AEqualB { // if the combination element cannot dominate the comparison element, the comparison element must be the element in the first level
			firstRankSet = append(firstRankSet, posI)
			prevValue = compareValue
			// update combined element minimum
			combineElemCopy := []T{}
			for j := range combineElem {
				if compareValue[j] < combineElem[j] {
					combineElemCopy = append(combineElemCopy, compareValue[j])
				} else {
					combineElemCopy = append(combineElemCopy, combineElem[j])
				}
			}
			combineElem = combineElemCopy
			if relevantValue == ANonDominatedB { // if they do not dominate each other, the comparison elements are added to the comparison set
				tempSet = append(tempSet, compareValue)
			} else { // if the comparison element dominates the combination element, clear the comparison set, because the comparison element is optimal
				tempSet = [][]T{compareValue}
			}
		} else if AllEqual(compareValue, prevValue) { // when the comparison element is equal to the last element of the first level, it must be the first level
			firstRankSet = append(firstRankSet, posI)
		} else if preComRelevant == ADominateB || preComRelevant == AEqualB { // the last element of the first level dominates the comparison element, so the comparison element must not be the first level
			continue
		} else if preComRelevant == BDominateA { // if the last element of the first level is dominated by the comparison element, the comparison element must be the first level
			firstRankSet = append(firstRankSet, posI)
			prevValue = compareValue
		} else { // if none of the preceding conditions are met, it will be compared with the elements in the comparison set
			isNonDominated := true
			for j := len(tempSet) - 1; j >= 0; j-- { // from the last comparison to the first, keep order. Once the elements dominate the comparison, they must dominate with the previous ones or not dominate each other
				relevantValue = Dominate(tempSet[j][1:], compareValue[1:])
				if relevantValue == ADominateB || relevantValue == AEqualB { // once a comparison element is dominated by an element in the comparison set, it must not be the first level in the comparison element and exit directly
					isNonDominated = false
					break
				} else if relevantValue == BDominateA { // once a comparison element dominates an element in the comparison set, it must be the first level in the comparison element and exit directly
					tempSet = append(tempSet[:j], tempSet[j+1:]...)
					break
				}
			}
			if isNonDominated {
				tempSet = append(tempSet, prevValue, compareValue)
				firstRankSet = append(firstRankSet, posI)
				prevValue = compareValue
			}
		}
	}

	return firstRankSet
}
