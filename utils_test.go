package utils

import (
	"fmt"
	"testing"
)

func TestSortIdToLongId(t *testing.T) {
	if id := SortIdToLongId(1); id != "000000000000000000000001" {
		t.Error("转换错误")
	} else {
		t.Log("转换通过")
	}
}

func TestLongIdToSortId(t *testing.T) {
	if id, err := LongIdToSortId("000000000000000000000001"); err != nil || id != 1 {
		t.Error("转换错误")
	} else {
		t.Log("转换通过")
	}
}

func TestStringToObjectId(t *testing.T) {
	if id, err := StringToObjectId("1"); err != nil || id.Hex() != "000000000000000000000001" {
		t.Error("转换错误")
	} else {
		t.Log("success")
	}
}

func TestStringToObjectIdError(t *testing.T) {
	if id, err := StringToObjectId("ged"); err != nil && id == nil {
		t.Error("success")
	} else {
		t.Log("err")
	}
}

type Demo struct {
	A uint8   `default:"1"`
	B uint16  `default:"2"`
	C uint32  `default:"3"`
	D uint    `default:"4"`
	E uint64  `default:"5"`
	F int8    `default:"6"`
	G int16   `default:"7"`
	H int32   `default:"8"`
	I int     `default:"9"`
	J int64   `default:"10"`
	K float32 `default:"0.1"`
	L float64 `default:"0.2"`
	M string  `default:"str"`
	N bool    `default:"true"`
}

func TestDefault(t *testing.T) {
	d := new(Demo)
	if err := Default(d); err == nil {
		t.Log("success")
	} else {
		fmt.Println(err)
		t.Error("err")
	}
}
