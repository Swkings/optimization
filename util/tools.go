package util

import (
	"encoding/json"
	"runtime"
	"strconv"
	"strings"

	"github.com/jinzhu/copier"
)

const (
	intSize         = 32 << (^uint(0) >> 63) // 32 or 64
	IntInf    int   = 1<<(intSize-1) - 1
	IntInfN   int   = -1 << (intSize - 1)
	Int32Inf  int32 = 1<<31 - 1
	Int32InfN int32 = -1 << 31
	Int64Inf  int64 = 1<<63 - 1
	Int64InfN int64 = -1 << 63
	Int8Inf   int8  = 1<<7 - 1
	Int8InfN  int8  = -1 << 7
)

func GetFuncFullName() string {
	pc, _, _, _ := runtime.Caller(1)
	function := runtime.FuncForPC(pc)

	return function.Name()
}

func GetStructName() string {
	pc, _, _, _ := runtime.Caller(1)
	function := runtime.FuncForPC(pc)
	nameArr := strings.Split(function.Name(), ".")
	structName := ""
	if len(nameArr) > 2 {
		structName = strings.Trim(nameArr[1], "()*")
	}

	return structName
}

func GetFmtStructName() string {
	pc, _, _, _ := runtime.Caller(1)
	function := runtime.FuncForPC(pc)
	nameArr := strings.Split(function.Name(), ".")
	structName := ""
	if len(nameArr) > 2 {
		structName = strings.Trim(nameArr[1], "()*")
	}

	return "[" + structName + "]"
}

func GetFuncName(args ...string) string {
	pc, _, _, _ := runtime.Caller(1)
	function := runtime.FuncForPC(pc)
	nameArr := strings.Split(function.Name(), ".")
	funcName := nameArr[len(nameArr)-1]
	for _, arg := range args {
		if arg == "" {
			continue
		}
		funcName += "-" + arg
	}

	return funcName
}

func GetFmtFuncName(args ...string) string {
	pc, _, _, _ := runtime.Caller(1)
	function := runtime.FuncForPC(pc)
	nameArr := strings.Split(function.Name(), ".")
	funcName := nameArr[len(nameArr)-1]
	for _, arg := range args {
		if arg == "" {
			continue
		}
		funcName += "-" + arg
	}

	return "[" + funcName + "]"
}

func GetStructFuncName(args ...string) string {
	pc, _, _, _ := runtime.Caller(1)
	function := runtime.FuncForPC(pc)
	nameArr := strings.Split(function.Name(), ".")
	structName := ""
	if len(nameArr) > 2 {
		structName = strings.Trim(nameArr[1], "()*")
	}
	funcName := nameArr[len(nameArr)-1]
	for _, arg := range args {
		if arg == "" {
			continue
		}
		if funcName == "" {
			funcName = arg
		} else {
			funcName += "-" + arg
		}
	}
	if structName == "" {
		return funcName
	}

	return structName + "." + funcName
}

func GetFmtStructFuncName(args ...string) string {
	pc, _, _, _ := runtime.Caller(1)
	function := runtime.FuncForPC(pc)
	nameArr := strings.Split(function.Name(), ".")
	structName := ""
	if len(nameArr) > 2 {
		structName = strings.Trim(nameArr[1], "()*")
	}
	funcName := nameArr[len(nameArr)-1]
	for _, arg := range args {
		if arg == "" {
			continue
		}
		if funcName == "" {
			funcName = arg
		} else {
			funcName += "-" + arg
		}
	}
	if structName == "" {
		return funcName
	}

	return "[" + structName + "." + funcName + "]"
}

func GetFmtFuncNameInAnonymous(args ...string) string {
	pc, _, _, _ := runtime.Caller(1)
	function := runtime.FuncForPC(pc)
	nameArr := strings.Split(function.Name(), ".")
	funcName := "NULL"
	if len(nameArr) >= 2 {
		funcName = nameArr[len(nameArr)-2]
	}
	for _, arg := range args {
		if arg == "" {
			continue
		}
		if funcName == "" {
			funcName = arg
		} else {
			funcName += "-" + arg
		}

	}

	return "[" + funcName + "]"
}

func GetFmtSelfDefaultFuncName(funcName string, args ...string) string {
	for _, arg := range args {
		if arg == "" {
			continue
		}
		funcName += "-" + arg
	}

	return "[" + funcName + "]"
}

type IntNum interface {
	int | int8 | int32 | int64 | uint | uint8 | uint32 | uint64
}

type Number interface {
	IntNum | float32 | float64
}

//	Range(num) -> 0..num-1
//	Range(startNum, stopNum) -> startNum..stopNum
//	Range(startNum, stopNum, step) -> startNum, startNum+step, startNum+2*step..., startNum+n*step<stopNum
//
// Range general integer sequence
func Range[T IntNum](args ...T) []T {
	list := []T{}
	if len(args) == 0 || len(args) > 3 {
		return list
	}
	if len(args) == 1 {
		num := args[0]
		for i := T(0); i < num; i++ {
			list = append(list, i)
		}
	} else if len(args) == 2 {
		startNum := args[0]
		stopNum := args[1]
		for i := startNum; i <= stopNum; i++ {
			list = append(list, i)
		}
	} else if len(args) == 3 {
		startNum := args[0]
		stopNum := args[1]
		step := args[2]
		for i := startNum; i <= stopNum; i += step {
			list = append(list, i)
		}
	}

	return list
}

// Zero general 0 list
func Zero[T Number](len int) []T {
	return make([]T, len)
}

// ListAdd add a number or list into list
func ListAdd[T Number](src []T, num []T) []T {
	var target []T

	srcL := len(src)
	numL := len(num)
	for i := 0; i < srcL; i++ {
		if numL == 1 {
			target = append(target, src[i]+num[0])
		} else if numL > 1 {
			if i < numL {
				target = append(target, src[i]+num[i])
			} else {
				target = append(target, src[i])
			}
		}
	}

	return target
}

// ListMinus minus a number or list into list
func ListMinus[T Number](src []T, num []T) []T {
	for i := range num {
		num[i] = -num[i]
	}

	return ListAdd(src, num)
}

// ListTimes times a number or list into list
func ListTimes[T Number](src []T, num []T) []T {
	var target []T

	srcL := len(src)
	numL := len(num)
	for i := 0; i < srcL; i++ {
		if numL == 1 {
			target = append(target, src[i]*num[0])
		} else if numL > 1 {
			if i < numL {
				target = append(target, src[i]*num[i])
			} else {
				target = append(target, src[i])
			}
		}
	}

	return target
}

// Max return max number of a and b
func Max[T Number](a, b T) T {
	if a < b {
		return b
	}

	return a
}

// Min return min number of a and b
func Min[T Number](a, b T) T {
	if a < b {
		return a
	}

	return b
}

// MaxList return max number of list
func MaxList[T Number](l []T) T {
	if len(l) == 0 {
		return 0
	}
	maxValue := l[0]
	for _, item := range l {
		if item > maxValue {
			maxValue = item
		}
	}

	return maxValue
}

// MaxMultiElement return max number of list
func MaxMultiElement[T Number](l ...T) T {
	return MaxList(l)
}

// MinList return min number of list
func MinList[T Number](l []T) T {
	if len(l) == 0 {
		return 0
	}
	minValue := l[0]
	for _, item := range l {
		if item < minValue {
			minValue = item
		}
	}

	return minValue
}

// MinMultiElement return min number of list
func MinMultiElement[T Number](l ...T) T {
	return MinList(l)
}

// MaxListOpt return max number and index of list, the compared value cal by f
func MaxListOpt[T any, V Number](l []T, f func(T) V) (V, int) {
	if len(l) == 0 {
		return 0, -1
	}
	maxIndex := 0
	maxValue := f(l[0])
	for i, item := range l {
		if f(item) > maxValue {
			maxValue = f(item)
			maxIndex = i
		}
	}

	return maxValue, maxIndex
}

// MinListOpt return min number and index of list, the compared value cal by f
func MinListOpt[T any, V Number](l []T, f func(T) V) (V, int) {
	if len(l) == 0 {
		return 0, -1
	}
	minIndex := 0
	minValue := f(l[0])
	for i, item := range l {
		if f(item) < minValue {
			minValue = f(item)
			minIndex = i
		}
	}

	return minValue, minIndex
}

// MaxMinListOpt return max and minnindex of list, the compared value cal by f
func MaxMinListOpt[T any, V Number](l []T, f func(T) V) (int, int) {
	if len(l) == 0 {
		return -1, -1
	}
	if len(l) == 1 {
		return 0, 0
	}

	minIndex, maxIndex := 0, 0
	minValue, maxValue := f(l[0]), f(l[0])
	for i, item := range l {
		if f(item) < minValue {
			minValue = f(item)
			minIndex = i
		}
		if f(item) > maxValue {
			maxValue = f(item)
			maxIndex = i
		}
	}

	return minIndex, maxIndex
}

//	if condition is true return ifTrue, else return elseFalse.
//
// Ternary expression
func Ternary[T any](condition bool, ifTrue, elseFalse T) T {
	if condition {
		return ifTrue
	}

	return elseFalse
}

func DeepCopy(toValue interface{}, fromValue interface{}) error {
	return copier.CopyWithOption(toValue, fromValue, copier.Option{DeepCopy: true})
}

func DeepCopyR[T any](src T) T {
	var res T
	err := copier.CopyWithOption(&res, src, copier.Option{DeepCopy: true})
	if err != nil {
		panic("deep copy err")
	}

	return res
}

func Abs[T Number](num T) T {
	if num < 0 {
		return -num
	}

	return num
}

func Dec2BinStr[T IntNum](num T, length int) string {
	bitStr := strconv.FormatInt(int64(num), 2)
	if len(bitStr) >= length {
		return bitStr
	}
	prefixLen := length - len(bitStr)
	prefixStr := ""
	for i := 0; i < prefixLen; i++ {
		prefixStr += "0"
	}
	bitStr = prefixStr + bitStr

	return bitStr
}

func Dec2BinArr[T IntNum](num T, length int) []int {
	bitStr := strconv.FormatInt(int64(num), 2)
	res := []int{}
	if len(bitStr) >= length {
		for i := range bitStr {
			res = append(res, Ternary(string(bitStr[i]) == "1", 1, 0))
		}
		return res
	}
	prefixLen := length - len(bitStr)
	prefixStr := ""
	for i := 0; i < prefixLen; i++ {
		prefixStr += "0"
	}
	bitStr = prefixStr + bitStr
	for i := range bitStr {
		res = append(res, Ternary(string(bitStr[i]) == "1", 1, 0))
	}

	return res
}

func BinArr2Dec[T, V IntNum](binArr []V) T {
	binL := len(binArr)
	res := int64(0)
	j := 0
	for i := binL - 1; i >= 0; i-- {
		res += int64(binArr[i]) * Pow(int64(2), j)
		j++
	}

	return T(res)
}

func BinStr2Dec[T IntNum](str string) T {
	sArr := strings.Split(str, "")
	strL := len(sArr)
	res := int64(0)
	j := 0
	for i := strL - 1; i >= 0; i-- {
		item := Ternary(sArr[i] == "0", 0, Ternary(sArr[i] == "1", 1, 0))
		res += int64(item) * Pow(int64(2), j)
		j++
	}

	return T(res)
}

func Pow[T int64 | float64](num T, n int) T {
	res := T(1)
	for n > 0 {
		if n&1 == 1 {
			res = res * num // res=(res*x)%Max;
		}
		num = num * num // x=(x*x)%Max;
		n >>= 1
	}

	return res
}

func ElementInList[T comparable](elem T, l []T) bool {
	for _, item := range l {
		if elem == item {
			return true
		}
	}
	return false
}

//	map{k1:v1, k2:v2} => slice[k1, v1, k2, v2]
//
// MapUnpack unpack map into kv list
func MapUnpack[K comparable, V interface{}](m map[K]V) []interface{} {
	res := []interface{}{}
	for k, v := range m {
		res = append(res, k, v)
	}

	return res
}

// Unique Array de duplication
func Unique[T interface{}](data []T, compareKey ...func(T) string) []T {
	res := []T{}
	set := make(map[string]int, len(data))
	for i, item := range data {
		var key string
		if len(compareKey) > 0 {
			key = compareKey[0](item)
		} else {
			key = Interface2String(item)
		}
		if _, ok := set[key]; !ok {
			res = append(res, data[i])
			set[key] = i
		}
	}

	return res
}

// Unique Array de duplication
func UniqueIndex[T interface{}](data []T, compareKey ...func(T) string) ([]T, []int) {
	res := []T{}
	resIndex := []int{}
	set := make(map[string]int, len(data))
	for i, item := range data {
		var key string
		if len(compareKey) > 0 {
			key = compareKey[0](item)
		} else {
			key = Interface2String(item)
		}

		if _, ok := set[key]; !ok {
			res = append(res, data[i])
			set[key] = i
			resIndex = append(resIndex, i)
		}
	}

	return res, resIndex
}

// List[V] to Map[getKey]V
func List2Map[T comparable, V any](list []V, GetKey func(item V) T, filter ...func(index int, item V) bool) map[T]V {
	res := map[T]V{}
	for i, item := range list {
		if len(filter) > 0 && filter[0](i, item) {
			continue
		}
		res[GetKey(item)] = item
	}

	return res
}

// List[V] to [getKey]getValue
func List2MapKV[T comparable, V any, W any](list []V, GetKey func(item V) T, GetValue func(item V) W, filter ...func(index int, item V) bool) map[T]W {
	res := map[T]W{}
	for i, item := range list {
		if len(filter) > 0 && filter[0](i, item) {
			continue
		}
		res[GetKey(item)] = GetValue(item)
	}

	return res
}

// List[V] to [getKey]Index
func List2MapKIndex[T comparable, V any](list []V, GetKey func(item V) T, filter ...func(index int, item V) bool) map[T]int {
	res := map[T]int{}
	for i, item := range list {
		if len(filter) > 0 && filter[0](i, item) {
			continue
		}
		res[GetKey(item)] = i
	}

	return res
}

// map[K]V to List[V]
func Map2List[T comparable, V any](m map[T]V, filter ...func(key T, item V) bool) []V {
	res := []V{}
	for key, item := range m {
		if len(filter) > 0 && filter[0](key, item) {
			continue
		}
		res = append(res, item)
	}

	return res
}

// map[K]V to List[GetValue(K,V)]
func Map2ListOpt[T comparable, V any, R any](m map[T]V, GetValue func(T, V) R, filter ...func(key T, item V) bool) []R {
	res := []R{}
	for key, item := range m {
		if len(filter) > 0 && filter[0](key, item) {
			continue
		}
		res = append(res, GetValue(key, item))
	}

	return res
}

// List2List List[T] to List[V]
func List2List[T any, V any](list []T, GetValue func(item T) V, filter ...func(index int, item T) bool) []V {
	res := []V{}
	for i, item := range list {
		if len(filter) > 0 && filter[0](i, item) {
			continue
		}
		res = append(res, GetValue(item))
	}

	return res
}

// IntConvert convert T or []T to []V
func IntConvert[T Number, V Number](items ...T) []V {
	res := []V{}
	for _, item := range items {
		res = append(res, V(item))
	}

	return res
}

// ConvertPtr convert element to *element
func PtrConvert[T any](element T) *T {
	return &element
}

// PtrTrim convert *element to element
func PtrTrim[T any](element *T) T {
	return *element
}

// JSONStringToMap convert json string to map
func JSONStringToMap(payload string) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(payload), &m)
	if err != nil {
		return nil, err
	}

	return m, nil
}

// MapToJSON convert map to json
func MapToJSON(paramMap map[string]interface{}) ([]byte, error) {
	return json.Marshal(paramMap)
}
