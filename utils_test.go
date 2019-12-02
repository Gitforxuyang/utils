package utils

import (
	"testing"
)

func TestSortIdToLongId(t *testing.T) {
	if id := SortIdToLongId(1); id != "000000000000000000000001" {
		t.Error("err")
	} else {
	}
}

func TestLongIdToSortId(t *testing.T) {
	if id, err := LongIdToSortId("000000000000000000000001"); err != nil || id != 1 {
		t.Error("err")
	}
}

func TestStringToObjectId(t *testing.T) {
	if id, err := StringToObjectId("1"); err != nil || id.Hex() != "000000000000000000000001" {
		t.Error("err")
	} else {
	}

	if id, err := StringToObjectId("ged"); err != nil && id == nil {
	} else {
		t.Error("err")
	}
}

type E struct {
	Z int `default:"1"`
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
	O int64   `default:",NOW_SECOND()"`
	P int64   `default:",NOW_MILLI()"`
	Z E
	Y *E
}

type DemoErr struct {
	A *uint8 `default:"1"`
	B *uint8 `default:"a"`
}

func TestDefault(t *testing.T) {
	d := new(Demo)
	d.Y = new(E)
	if err := Default(d); err != nil {
		t.Error("err")
	}
	if err := Default(*d); err == nil {
		t.Error("err")
	}
	if err := Default(new(DemoErr)); err == nil {
		t.Error("err")
	}
}
