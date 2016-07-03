package transform

import (
	"reflect"
	"strconv"
	"time"
)

var int32Handler = func(o reflect.Value, t reflect.Value) error {
	if o.Kind() != reflect.Int32 {
		return nil
	}
	switch {
	case t.Kind() == reflect.String:
		v := strconv.Itoa(int(o.Interface().(int32)))
		t.SetString(v)
	case t.Kind() == reflect.Int:
		v := int(o.Interface().(int32))
		t.Set(reflect.ValueOf(v))
	case t.Kind() == reflect.Int8:
		v := int8(o.Interface().(int32))
		t.Set(reflect.ValueOf(v))
	case t.Kind() == reflect.Int16:
		v := int16(o.Interface().(int32))
		t.Set(reflect.ValueOf(v))
	case t.Kind() == reflect.Int32:
		v := int32(o.Interface().(int32))
		t.Set(reflect.ValueOf(v))
	case t.Kind() == reflect.Int64:
		v := int64(o.Interface().(int32))
		t.SetInt(v)
	case t.Kind() == reflect.Uint:
		v := uint(o.Interface().(int32))
		t.Set(reflect.ValueOf(v))
	case t.Kind() == reflect.Uint8:
		v := uint8(o.Interface().(int32))
		t.Set(reflect.ValueOf(v))
	case t.Kind() == reflect.Uint16:
		v := uint16(o.Interface().(int32))
		t.Set(reflect.ValueOf(v))
	case t.Kind() == reflect.Uint32:
		v := uint32(o.Interface().(int32))
		t.Set(reflect.ValueOf(v))
	case t.Kind() == reflect.Uint64:
		v := uint64(o.Interface().(int32))
		t.SetUint(v)
	case t.Kind() == reflect.Float32:
		v := float32(o.Interface().(int32))
		t.Set(reflect.ValueOf(v))
	case t.Kind() == reflect.Float64:
		v := float64(o.Interface().(int32))
		t.SetFloat(v)
	case t.Kind() == reflect.Bool:
		t.SetBool(o.Interface().(int32) == 1)
	case t.Kind() == reflect.Interface:
		t.Set(reflect.ValueOf(o.Interface()))
	case t.Type().String() == "time.Time":
		v := time.Unix(int64(o.Interface().(int32)), 0)
		t.Set(reflect.ValueOf(v))
	}
	return nil
}
