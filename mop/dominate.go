package mop

import "github.com/swkings/optimization/util"

//	@params
//		- A: list
//		- B: list
//		- minProblemArg: optional bool
//	@return
//		-  1: A \succ B
//		- -1: B \succ A
//		-  0: A \not \succ B, B \not \succ A
//		-  2: A == B
// Dominate judge A dominate B
func Dominate[T Number](A []T, B []T, minProblemArg ...bool) DominateRes {
	minProblem := true
	if len(minProblemArg) > 0 {
		minProblem = minProblemArg[0]
	}

	ABetterCount, BBetterCount, equalCount := 0, 0, 0

	for i := 0; i < len(A); i++ {
		if A[i] < B[i] {
			ABetterCount += 1
		} else if B[i] < A[i] {
			BBetterCount += 1
		} else {
			equalCount += 1
		}
		if ABetterCount > 0 && BBetterCount > 0 {
			return 0
		}
	}

	if (ABetterCount+equalCount == len(A)) && ABetterCount > 0 {
		return util.Ternary(minProblem, ADominateB, BDominateA)
	} else if (BBetterCount+equalCount) == len(A) && BBetterCount > 0 {
		return util.Ternary(minProblem, BDominateA, ADominateB)
	} else if equalCount == len(A) {
		return AEqualB
	} else {
		return ANonDominatedB
	}
}

//	@params
//		- A: list
//		- B: list
//	@return
//		-  true: all(A == B)
//		-  false: A != B, A[i] != B[i]
// AllEqual judge all(A == B)
func AllEqual[T Number](A []T, B []T) bool {
	if len(A) != len(B) {
		return false
	}
	for i := range A {
		if A[i] != B[i] {
			return false
		}
	}

	return true
}

//	@params
//		- A: list
//		- B: list
//	@return
//		-  true: all(A < B)
//		-  false: A != B, A[i] >= B[i]
// AllLt judge all(A < B)
func AllLt[T Number](A []T, B []T) bool {
	if len(A) != len(B) {
		return false
	}
	minLen := len(A)
	BL := len(B)
	if BL < minLen {
		minLen = BL
	}
	for i := 0; i < BL; i++ {
		if A[i] >= B[i] {
			return false
		}
	}

	return true
}

//	@params
//		- A: list
//		- B: list
//	@return
//		-  true: all(A <= B)
//		-  false: A != B, A[i] > B[i]
// AllLeq judge all(A <= B)
func AllLeq[T Number](A []T, B []T) bool {
	if len(A) != len(B) {
		return false
	}
	minLen := len(A)
	BL := len(B)
	if BL < minLen {
		minLen = BL
	}
	for i := 0; i < BL; i++ {
		if A[i] > B[i] {
			return false
		}
	}

	return true
}

//	@params
//		- A: list
//		- B: list
//	@return
//		-  true: all(A > B)
//		-  false: A != B, A[i] <= B[i]
// AllGt judge all(A > B)
func AllGt[T Number](A []T, B []T) bool {
	if len(A) != len(B) {
		return false
	}
	minLen := len(A)
	BL := len(B)
	if BL < minLen {
		minLen = BL
	}
	for i := 0; i < BL; i++ {
		if A[i] <= B[i] {
			return false
		}
	}

	return true
}

//	@params
//		- A: list
//		- B: list
//	@return
//		-  true: all(A >= B)
//		-  false: A != B, A[i] < B[i]
// AllGeq judge all(A >= B)
func AllGeq[T Number](A []T, B []T) bool {
	if len(A) != len(B) {
		return false
	}
	minLen := len(A)
	BL := len(B)
	if BL < minLen {
		minLen = BL
	}
	for i := 0; i < BL; i++ {
		if A[i] < B[i] {
			return false
		}
	}

	return true
}
