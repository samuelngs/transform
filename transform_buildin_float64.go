package transform

import (
	"reflect"
	"strconv"
	"time"
)

var float64Handler = func(o reflect.Value, t reflect.Value) error {
	if o.Kind() != reflect.Float64 {
		return nil
	}
	switch {
	case t.Kind() == reflect.String:
		v := strconv.FormatFloat(o.Float(), 'f', -1, 64)
		t.SetString(v)
	case t.Kind() == reflect.Int:
		t.SetInt(int64(o.Float()))
	case t.Kind() == reflect.Int8:
		t.SetInt(int64(o.Float()))
	case t.Kind() == reflect.Int16:
		t.SetInt(int64(o.Float()))
	case t.Kind() == reflect.Int32:
		t.SetInt(int64(o.Float()))
	case t.Kind() == reflect.Int64:
		t.SetInt(int64(o.Float()))
	case t.Kind() == reflect.Uint:
		t.SetUint(uint64(o.Float()))
	case t.Kind() == reflect.Uint8:
		t.SetUint(uint64(o.Float()))
	case t.Kind() == reflect.Uint16:
		t.SetUint(uint64(o.Float()))
	case t.Kind() == reflect.Uint32:
		t.SetUint(uint64(o.Float()))
	case t.Kind() == reflect.Uint64:
		t.SetUint(uint64(o.Float()))
	case t.Kind() == reflect.Float32:
		t.SetFloat(o.Float())
	case t.Kind() == reflect.Float64:
		t.SetFloat(o.Float())
	case t.Kind() == reflect.Bool:
		t.SetBool(o.Float() == 1.0)
	case t.Kind() == reflect.Interface:
		t.Set(reflect.ValueOf(o.Interface()))
	case t.Type().String() == "time.Time":
		v := time.Unix(int64(o.Float()), 0)
		t.Set(reflect.ValueOf(v))
	}
	return nil
}
