package transform

import (
	"reflect"
	"strconv"
	"time"
)

var int64Handler = func(o reflect.Value, t reflect.Value) error {
	if o.Kind() != reflect.Int64 {
		return nil
	}
	switch {
	case t.Kind() == reflect.String:
		v := strconv.FormatInt(o.Int(), 10)
		t.SetString(v)
	case t.Kind() == reflect.Int:
		v := int(o.Int())
		t.Set(reflect.ValueOf(v))
	case t.Kind() == reflect.Int8:
		v := int8(o.Int())
		t.Set(reflect.ValueOf(v))
	case t.Kind() == reflect.Int16:
		v := int16(o.Int())
		t.Set(reflect.ValueOf(v))
	case t.Kind() == reflect.Int32:
		v := int32(o.Int())
		t.Set(reflect.ValueOf(v))
	case t.Kind() == reflect.Int64:
		t.SetInt(o.Int())
	case t.Kind() == reflect.Uint:
		v := uint(o.Int())
		t.Set(reflect.ValueOf(v))
	case t.Kind() == reflect.Uint8:
		v := uint8(o.Int())
		t.Set(reflect.ValueOf(v))
	case t.Kind() == reflect.Uint16:
		v := uint16(o.Int())
		t.Set(reflect.ValueOf(v))
	case t.Kind() == reflect.Uint32:
		v := uint32(o.Int())
		t.Set(reflect.ValueOf(v))
	case t.Kind() == reflect.Uint64:
		v := uint64(o.Int())
		t.Set(reflect.ValueOf(v))
	case t.Kind() == reflect.Float32:
		v := float32(o.Int())
		t.Set(reflect.ValueOf(v))
	case t.Kind() == reflect.Float64:
		v := float64(o.Int())
		t.SetFloat(v)
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
