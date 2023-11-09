package mop

type DominateRes int

const (
	ADominateB     DominateRes = 1
	ANonDominatedB DominateRes = 0
	BDominateA     DominateRes = -1
	AEqualB        DominateRes = 2
)

type Number interface {
	int | int8 | int32 | int64 | uint | uint8 | uint32 | uint64 | float32 | float64
}
