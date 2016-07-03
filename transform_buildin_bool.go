package transform

import "reflect"

var boolHandler = func(o reflect.Value, t reflect.Value) error {
	if o.Kind() != reflect.Bool {
		return nil
	}
	switch {
	case t.Kind() == reflect.String:
		if o.Bool() {
			t.SetString("true")
		} else {
			t.SetString("false")
		}
	case t.Kind() == reflect.Int:
		if o.Bool() {
			t.SetInt(1)
		} else {
			t.SetInt(0)
		}
	case t.Kind() == reflect.Int8:
		if o.Bool() {
			t.SetInt(1)
		} else {
			t.SetInt(0)
		}
	case t.Kind() == reflect.Int16:
		if o.Bool() {
			t.SetInt(1)
		} else {
			t.SetInt(0)
		}
	case t.Kind() == reflect.Int32:
		if o.Bool() {
			t.SetInt(1)
		} else {
			t.SetInt(0)
		}
	case t.Kind() == reflect.Int64:
		if o.Bool() {
			t.SetInt(1)
		} else {
			t.SetInt(0)
		}
	case t.Kind() == reflect.Uint:
		if o.Bool() {
			t.SetUint(1)
		} else {
			t.SetUint(0)
		}
	case t.Kind() == reflect.Uint8:
		if o.Bool() {
			t.SetUint(1)
		} else {
			t.SetUint(0)
		}
	case t.Kind() == reflect.Uint16:
		if o.Bool() {
			t.SetUint(1)
		} else {
			t.SetUint(0)
		}
	case t.Kind() == reflect.Uint32:
		if o.Bool() {
			t.SetUint(1)
		} else {
			t.SetUint(0)
		}
	case t.Kind() == reflect.Uint64:
		if o.Bool() {
			t.SetUint(1)
		} else {
			t.SetUint(0)
		}
	case t.Kind() == reflect.Float32:
		if o.Bool() {
			t.SetFloat(1.0)
		} else {
			t.SetUint(0.0)
		}
	case t.Kind() == reflect.Float64:
		if o.Bool() {
			t.SetFloat(1.0)
		} else {
			t.SetUint(0.0)
		}
	case t.Kind() == reflect.Bool:
		t.SetBool(o.Bool())
	case t.Kind() == reflect.Interface:
		t.Set(reflect.ValueOf(o.Interface()))
	}
	return nil
}
