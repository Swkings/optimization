package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"reflect"
	"strings"

	"golang.org/x/text/width"
)

func PrettyMapStruct(v interface{}, indent bool) interface{} {
	jsonV, err := json.Marshal(v)
	if err != nil {
		return v
	}

	var out bytes.Buffer
	if indent {
		err = json.Indent(&out, jsonV, "", "  ")
	} else {
		out = *bytes.NewBuffer(jsonV)
	}

	if err != nil {
		return v
	}
	return out.String()
}

func PrettyStructTable[T any](elementList []T, rightAlign ...bool) string {
	if elementList == nil || (elementList != nil && len(elementList) == 0) {
		return ""
	}

	elementListType := reflect.TypeOf(elementList)
	if elementListType.Kind() != reflect.Slice {
		return ""
	}

	keyList := []string{}
	valueArray := [][]string{}
	for i, element := range elementList {
		et, ev := reflect.TypeOf(element), reflect.ValueOf(element)
		switch et.Kind() {
		case reflect.Ptr:
			if ev.IsNil() {
				// fmt.Printf("ev is nil %v\n", element)
				continue
			}

			zeroElem := reflect.New(et.Elem()).Elem().Interface()
			evElem := ev.Elem().Interface()
			if reflect.DeepEqual(zeroElem, evElem) {
				fmt.Printf("ev is empty %v\n", element)
				continue
			}

			et, ev = et.Elem(), ev.Elem()
		case reflect.Struct:
			if ev.IsZero() {
				continue
			}
		}

		if i == 0 {
			keyList = GetKey(et)
		}
		valueArray = append(valueArray, GetValue(ev))
	}

	return PrettyArrayTable(keyList, valueArray, rightAlign...)
}

func GetFieldValue(t reflect.Type, v reflect.Value) (keyList []string, valueList []string) {
	for k := 0; k < t.NumField(); k++ {
		field := t.Field(k).Name
		value := v.Field(k)
		keyList = append(keyList, fmt.Sprint(field))
		valueList = append(valueList, fmt.Sprint(value))
	}

	return keyList, valueList
}
func GetKey(t reflect.Type) (keyList []string) {
	for k := 0; k < t.NumField(); k++ {
		field := t.Field(k).Name
		keyList = append(keyList, fmt.Sprint(field))
	}

	return keyList
}
func GetValue(v reflect.Value) (valueList []string) {
	for k := 0; k < v.NumField(); k++ {
		value := v.Field(k)
		valueList = append(valueList, fmt.Sprint(value))
	}

	return valueList
}

func PrettyArrayTable[T any](title []string, elementArray [][]T, rightAlign ...bool) string {
	if title == nil && (elementArray == nil || (elementArray != nil && len(elementArray) == 0)) {
		return ""
	}

	isRightAlign := false
	if len(rightAlign) > 0 {
		isRightAlign = rightAlign[0]
	}

	table := []string{}

	colNum := 0
	if title != nil {
		colNum = len(title)
	}
	if colNum == 0 && len(elementArray) > 0 {
		colNum = len(elementArray[0])
	}

	if colNum == 0 {
		return ""
	}

	minColumWidth := ListAdd(Zero[int](colNum), []int{math.MaxInt})
	maxColumWidth := Zero[int](colNum)
	for j, element := range title {
		eW := Width(fmt.Sprint(element))
		if eW < minColumWidth[j] {
			minColumWidth[j] = eW
		}
		if eW > maxColumWidth[j] {
			maxColumWidth[j] = eW
		}
	}
	for _, elementList := range elementArray {
		if elementList == nil {
			continue
		}
		for j, element := range elementList {
			eW := Width(fmt.Sprint(element))
			if eW < minColumWidth[j] {
				minColumWidth[j] = eW
			}
			if eW > maxColumWidth[j] {
				maxColumWidth[j] = eW
			}
		}
	}

	fmtStr := ""
	for j := range title {
		if isRightAlign {
			fmtStr += "%" + fmt.Sprint(maxColumWidth[j]) + "v "
		} else {
			fmtStr += "%-" + fmt.Sprint(maxColumWidth[j]) + "v "
		}
	}
	table = append(table, fmt.Sprintf(fmtStr, func() []any {
		res := []any{}
		for _, element := range title {
			res = append(res, element)
		}
		return res
	}()...))

	for _, elementList := range elementArray {
		fmtStr := ""
		for j := range elementList {
			if isRightAlign {
				fmtStr += "%" + fmt.Sprint(maxColumWidth[j]) + "v "
			} else {
				fmtStr += "%-" + fmt.Sprint(maxColumWidth[j]) + "v "
			}
		}

		table = append(table, fmt.Sprintf(fmtStr, func() []any {
			res := []any{}
			for _, element := range elementList {
				res = append(res, element)
			}
			return res
		}()...))
	}

	return strings.Join(table, "\n")
}

func Width(s string) (w int) {
	for _, r := range s {
		switch width.LookupRune(r).Kind() {
		case width.EastAsianFullwidth, width.EastAsianWide:
			w += 2
		case width.EastAsianHalfwidth, width.EastAsianNarrow,
			width.Neutral, width.EastAsianAmbiguous:
			w += 1
		}
	}

	return w
}
