package utils

import (
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
