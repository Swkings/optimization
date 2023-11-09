package util

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"reflect"
	"strconv"
	"time"
)

// RandSE random general integer number [start, end)
func RandSE[T Number](start T, end T) T {
	if end < start {
		return 0
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	switch reflect.TypeOf(start).Kind() {
	case reflect.Int:
		return T(r.Intn(int(end-start)) + int(start))
	case reflect.Int32:
		return T(r.Int31n(int32(end-start)) + int32(start))
	case reflect.Int64:
		return T(r.Int63n(int64(end-start)) + int64(start))
	default:
		return T(r.Intn(int(end-start)) + int(start))
	}
}

// RandN random general integer number [0, n)
func RandN[T Number](n T) T {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	switch reflect.TypeOf(n).Kind() {
	case reflect.Int:
		return T(r.Intn(int(n)))
	case reflect.Int32:
		return T(r.Int31n(int32(n)))
	case reflect.Int64:
		return T(r.Int63n(int64(n)))
	default:
		return T(r.Int())
	}
}

// RandN random general number
func Rand[T Number]() T {
	var t T
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	switch reflect.TypeOf(t).Kind() {
	case reflect.Int:
		return T(r.Int())
	case reflect.Int32:
		return T(r.Int31())
	case reflect.Int64:
		return T(r.Int63())
	case reflect.Float32:
		return T(r.Float32())
	case reflect.Float64:
		return T(r.Float64())
	default:
		return T(r.Int())
	}
}

func RandListV2[T Number](size int32, args ...T) []T {
	res := []T{}
	argsL := len(args)
	var f T
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := int32(0); i < size; i++ {
		if argsL == 1 {
			n := args[0]
			switch reflect.TypeOf(n).Kind() {
			case reflect.Int:
				f = T(r.Intn(int(n)))
			case reflect.Int32:
				f = T(r.Int31n(int32(n)))
			case reflect.Int64:
				f = T(r.Int63n(int64(n)))
			default:
				f = T(r.Int())
			}
		} else if argsL == 2 {
			start, end := args[0], args[1]
			switch reflect.TypeOf(start).Kind() {
			case reflect.Int:
				f = T(r.Intn(int(end-start)) + int(start))
			case reflect.Int32:
				f = T(r.Int31n(int32(end-start)) + int32(start))
			case reflect.Int64:
				f = T(r.Int63n(int64(end-start)) + int64(start))
			default:
				f = T(r.Intn(int(end-start)) + int(start))
			}
		} else {
			switch reflect.TypeOf(f).Kind() {
			case reflect.Int:
				f = T(r.Int())
			case reflect.Int32:
				f = T(r.Int31())
			case reflect.Int64:
				f = T(r.Int63())
			case reflect.Float32:
				f = T(r.Float32())
			case reflect.Float64:
				f = T(r.Float64())
			default:
				f = T(r.Int())
			}
		}
		res = append(res, f)
	}

	return res
}

//	RandList(size): list[...].
//	RandList(size, n): list[0, n).
//	RandList(size, s, e): list[s, e).
// RandList random general list.
func RandList[T Number](size int32, args ...T) []T {
	res := []T{}
	argsL := len(args)

	for i := int32(0); i < size; i++ {
		f := Rand[T]()
		if argsL == 1 {
			f = RandN(args[0])
		} else if argsL == 2 {
			f = RandSE(args[0], args[1])
		}
		res = append(res, f)
	}

	return res
}

// RandElement random general element from arr
func RandElement[T any](arr []T) T {
	arrL := len(arr)
	if arrL == 0 {
		panic("input arr empty")
	}
	if arrL == 1 {
		return arr[0]
	}

	return arr[RandN(arrL)]
}

// RandElementList random general n element from arr
func RandElementListV2[T any](arr []T, n int, diff bool) []T {
	arrL := len(arr)
	if arrL == 0 {
		panic("input arr empty")
	}
	if arrL < n {
		panic("n is out arr index")
	}
	if arrL == 1 {
		return arr
	}
	indexList := make([]int, 0)
	if diff {
		for len(indexList) < n {
			i := RandN(arrL)
			// check repeat
			exist := false
			for _, v := range indexList {
				if v == i {
					exist = true
					break
				}
			}
			if !exist {
				indexList = append(indexList, i)
			}
		}
	} else {
		indexList = RandList(int32(n), arrL)
	}

	res := []T{}
	for _, i := range indexList {
		res = append(res, arr[i])
	}
	return res
}

// RandElementList random general n element from arr
func RandElementList[T any](arr []T, n int, diff bool) []T {
	arrL := len(arr)
	if arrL == 0 {
		// panic("input array empty")
		return arr
	}
	if diff && arrL < n {
		// panic("n is out array index")
		return arr
	}
	if diff && arrL == 1 {
		return arr
	}
	if diff {
		RandShuffle(arr)
		return append([]T{}, arr[:n]...)
	}

	indexList := RandList(int32(n), arrL)
	res := []T{}
	for _, i := range indexList {
		res = append(res, arr[i])
	}

	return res
}

// RandShuffle shuffle slice
func RandShuffle[T any](slice []T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}

// RandomClientID returns random client id
func RandomClientID() string {
	nano := time.Now().UnixNano()
	nanoStr := strconv.FormatInt(nano, 10)
	sum := md5.Sum([]byte(nanoStr))

	return fmt.Sprintf("%x", sum[12:])
}
