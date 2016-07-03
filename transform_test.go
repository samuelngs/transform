package transform

import (
	"reflect"
	"testing"
	"time"
)

type date struct {
	Year   int32
	Month  int32
	Day    int32
	Hour   int32
	Minute int32
	Second int32
}

type object1 struct {
	Field1 int
	Field2 int64
	Field3 float32
	Field4 string
	Field5 bool
	Field6 *date
	Field7 string
	Field8 string
}

type object2 struct {
	Field1 int64
	Field2 float64
	Field3 int
	Field4 string
	Field5 bool
	Field6 time.Time
	Field7 int
	Field8 time.Time
}

var sample = &object1{
	Field1: 50,
	Field2: 60,
	Field3: 70.5,
	Field4: "value-4",
	Field5: true,
	Field6: &date{
		Year:   2016,
		Month:  6,
		Day:    1,
		Hour:   12,
		Minute: 5,
		Second: 10,
	},
	Field7: "30.5",
	Field8: "2014-11-12T11:45:26.371Z",
}

func dateHandler(o reflect.Value, t reflect.Value) error {
	if o.Type().String() != "*transform.date" {
		return nil
	}
	switch {
	case t.Type().String() == "time.Time":
		s := o.Interface().(*date)
		v := time.Date(
			int(s.Year),
			time.Month(s.Month),
			int(s.Day),
			int(s.Hour),
			int(s.Minute),
			int(s.Second),
			0, time.UTC)
		t.Set(reflect.ValueOf(v))
	}
	return nil
}

func TestGo(t *testing.T) {
	modified := &object2{}
	if err := Go(sample, modified, dateHandler); err != nil {
		t.Error(err)
	}
	t.Logf("\nwas: %v\nnow: %v", sample, modified)
}
