package util

import (
	"encoding/json"
	"strconv"
)

type Pair struct {
	First, Second interface{}
}

func (p *Pair) GetFirst() interface{} {
	return p.First
}

func (p *Pair) GetSecond() interface{} {
	return p.Second
}

func (p *Pair) SetFirst(v interface{}) {
	p.First = v
}

func (p *Pair) SetSecond(v interface{}) {
	p.Second = v
}

func (p *Pair) ToString() string {
	return "(" + p.Concat() + ")"
}

func (p *Pair) Concat() string {
	if p.First == "" && p.Second == "" {
		return ""
	}
	if p.First == "" {
		return Interface2String(p.Second)
	}
	if p.Second == "" {
		return Interface2String(p.First)
	}
	return Interface2String(p.First) + "," + Interface2String(p.Second)
}

func (p *Pair) Append(v interface{}) *Pair {
	p.First = p.Concat()
	p.Second = v
	return p
}

func NewTuple(first interface{}, second interface{}) *Pair {
	return &Pair{first, second}
}

func TupleList2StringList(l []*Pair) []string {
	res := []string{}
	for _, v := range l {
		res = append(res, v.ToString())
	}
	return res
}

func Interface2String(v interface{}) string {
	var key string
	if v == nil {
		return key
	}
	switch v.(type) {
	case float64:
		ft := v.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := v.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := v.(int)
		key = strconv.Itoa(it)
	case uint:
		it := v.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := v.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := v.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := v.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := v.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := v.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := v.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := v.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := v.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = v.(string)
	case []byte:
		key = string(v.([]byte))
	default:
		newValue, _ := json.Marshal(v)
		key = string(newValue)
	}
	if len(key) > 0 && key[0] == '"' {
		key = key[1:]
	}
	if len(key) > 0 && key[len(key)-1] == '"' {
		key = key[:len(key)-1]
	}
	return key
}
