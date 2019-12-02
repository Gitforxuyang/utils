package utils

import "testing"

func TestSortIdToLongId(t *testing.T) {
	if id := SortIdToLongId(1); id != "000000000000000000000001" {
		t.Error("转换错误")
	} else {
		t.Log("转换通过")
	}
}
