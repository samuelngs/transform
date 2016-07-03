package transform

import (
	"reflect"
	"strconv"
	"time"
)

var intHandler = func(o reflect.Value, t reflect.Value) error {
	if o.Kind() != reflect.Int {
		return nil
	}
	switch {
	case t.Kind() == reflect.String:
		v := strconv.FormatInt(o.Int(), 10)
		t.SetString(v)
	case t.Kind() == reflect.Int:
		t.SetInt(o.Int())
	case t.Kind() == reflect.Int8:
		t.SetInt(o.Int())
	case t.Kind() == reflect.Int16:
		t.SetInt(o.Int())
	case t.Kind() == reflect.Int32:
		t.SetInt(o.Int())
	case t.Kind() == reflect.Int64:
		t.SetInt(o.Int())
	case t.Kind() == reflect.Uint:
		t.SetUint(o.Uint())
	case t.Kind() == reflect.Uint8:
		t.SetUint(o.Uint())
	case t.Kind() == reflect.Uint16:
		t.SetUint(o.Uint())
	case t.Kind() == reflect.Uint32:
		t.SetUint(o.Uint())
	case t.Kind() == reflect.Uint64:
		t.SetUint(o.Uint())
	case t.Kind() == reflect.Float32:
		v := float32(o.Int())
		t.Set(reflect.ValueOf(v))
	case t.Kind() == reflect.Float64:
		t.SetFloat(float64(o.Int()))
	case t.Kind() == reflect.Bool:
		t.SetBool(o.Int() == 1)
	case t.Kind() == reflect.Interface:
		t.Set(reflect.ValueOf(o.Interface()))
	case t.Type().String() == "time.Time":
		v := time.Unix(o.Int(), 0)
		t.Set(reflect.ValueOf(v))
	}
	return nil
}
