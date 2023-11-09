package util

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

var T = func() {
	fullName := GetFuncFullName()
	fmt.Printf("TestFullName T: %v\n", fullName)
}

type TS struct {
	TS2
}
type TS2 struct{}

func (ts *TS) TSFunc() {
	fullName := GetFuncFullName()
	funcName := GetFuncName()
	structName := GetStructName()
	structFuncName := GetStructFuncName()
	fmt.Printf("fullName: %v\n", fullName)
	fmt.Printf("funcName: %v\n", funcName)
	fmt.Printf("structName: %v\n", structName)
	fmt.Printf("structFuncName: %v\n", structFuncName)
}
func (ts2 *TS2) tS2Func() {
	fullName := GetFuncFullName()
	funcName := GetFuncName()
	structName := GetStructName()
	structFuncName := GetStructFuncName()
	fmt.Printf("ts2FullName: %v\n", fullName)
	fmt.Printf("ts2FuncName: %v\n", funcName)
	fmt.Printf("ts2StructName: %v\n", structName)
	fmt.Printf("ts2StructFuncName: %v\n", structFuncName)
}
func TestGetName(t *testing.T) {
	ts := &TS{}
	ts.TSFunc()
	ts.tS2Func()

	fullName := GetFuncFullName()
	funcName := GetFuncName()
	structName := GetStructName()
	structFuncName := GetStructFuncName()
	fmt.Printf("TestFullName: %v\n", fullName)
	fmt.Printf("TestFuncName: %v\n", funcName)
	fmt.Printf("TestStructName: %v\n", structName)
	fmt.Printf("TestStructFuncName: %v\n", structFuncName)

	T()
}

func TestRange(t *testing.T) {
	list1 := Range(10)
	fmt.Printf("list: %v\n", list1)
	list2 := Range(10, 20)
	fmt.Printf("list: %v\n", list2)
	list3 := Range(10, 20, 2)
	fmt.Printf("list: %v\n", list3)
}

func TestZero(t *testing.T) {
	list1 := Zero[int](10)
	fmt.Printf("list: %v\n", list1)
}

func TestListAdd(t *testing.T) {
	list1 := Zero[int](10)
	list2 := ListAdd(list1, []int{1})
	fmt.Printf("list1: %v\n", list1)
	fmt.Printf("list2: %v\n", list2)
}

func TestListMinus(t *testing.T) {
	list1 := Zero[int](10)
	list2 := ListAdd(list1, []int{1, 2, 3})
	list3 := ListMinus(list2, []int{1, 2})
	fmt.Printf("list1: %v\n", list1)
	fmt.Printf("list2: %v\n", list2)
	fmt.Printf("list3: %v\n", list3)
}

func TestListTimes(t *testing.T) {
	list1 := Zero[int](10)
	list2 := ListAdd(list1, []int{1, 2, 3})
	list3 := ListTimes(list2, []int{1, 2})
	fmt.Printf("list1: %v\n", list1)
	fmt.Printf("list2: %v\n", list2)
	fmt.Printf("list3: %v\n", list3)
}

func TestDecBin(t *testing.T) {
	bin := Dec2BinArr(3, 3)
	fmt.Printf("bin %v\n", bin)

	dec := BinArr2Dec[int](bin)
	fmt.Printf("dec %v\n", dec)
}

func TestXOR(t *testing.T) {
	a := 1
	b := 0
	fmt.Printf("a %v\n", a)
	fmt.Printf("b %v\n", b)
	a ^= 1
	b ^= 1
	fmt.Printf("a %v\n", a)
	fmt.Printf("b %v\n", b)
}

type Test struct {
	A int
	B string
	C map[string]int
}

func (a Test) String() string {
	str := strconv.Itoa(a.A) + "_" + a.B
	return str
}

func GetCompareKey(data Test) string {
	str := strconv.Itoa(data.A) + "_" + data.B
	return str
}

func TestUnique(t *testing.T) {
	a := []int{1, 2, 2, 3, 4, 4, 4, 0}
	fmt.Printf("before a: %v\n", a)
	a = Unique(a)
	fmt.Printf("after a: %v\n", a)

	b := [][]int{
		{1, 2},
		{1, 3},
		{1, 2},
		{2, 3},
		{3, 5},
	}
	fmt.Printf("before b: %v\n", b)
	b = Unique(b)
	fmt.Printf("after b: %v\n", b)

	c := []map[string]int{
		{"a": 1, "b": 2},
		{"a": 1, "b": 3},
		{"a": 1, "b": 2},
		{"a": 2, "b": 3},
		{"a": 3, "b": 5},
	}
	fmt.Printf("before c: %v\n", c)
	c = Unique(c)
	fmt.Printf("after c: %v\n", c)

	d := []Test{
		{1, "2", map[string]int{"a": 1, "b": 2}},
		{1, "2", map[string]int{"a": 1}},
		{1, "2", map[string]int{"b": 2, "a": 1}},
		{2, "2", map[string]int{"a": 1}},
		{3, "2", map[string]int{"a": 1}},
	}
	fmt.Printf("type d: %v\n", reflect.TypeOf(d))

	for _, item := range d {
		fmt.Printf("before item: %v\n", item)
	}
	d = Unique(d, GetCompareKey)
	for _, item := range d {
		fmt.Printf("after item: %v\n", item)
	}
}

func TestList2Map(t *testing.T) {
	type SS struct {
		ID  string
		Age int32
	}
	list := []SS{
		{ID: "111", Age: 1},
		{ID: "222", Age: 2},
		{ID: "333", Age: 3},
	}
	res := List2Map(list, func(item SS) string {
		return item.ID
	})

	fmt.Printf("list: %v\n", list)
	fmt.Printf("map: %v\n", res)
}
