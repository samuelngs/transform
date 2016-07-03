package transform

import (
	"reflect"
	"strconv"
	"strings"
	"time"
)

var stringHandler = func(o reflect.Value, t reflect.Value) error {
	if o.Kind() != reflect.String {
		return nil
	}
	switch {
	case t.Kind() == reflect.String:
		t.Set(o)
	case t.Kind() == reflect.Int:
		i, err := strconv.Atoi(strings.Split(o.String(), ".")[0])
		if err != nil {
			return err
		}
		t.Set(reflect.ValueOf(i))
	case t.Kind() == reflect.Int8:
		i, err := strconv.Atoi(strings.Split(o.String(), ".")[0])
		if err != nil {
			return err
		}
		t.Set(reflect.ValueOf(int8(i)))
	case t.Kind() == reflect.Int16:
		i, err := strconv.Atoi(strings.Split(o.String(), ".")[0])
		if err != nil {
			return err
		}
		t.Set(reflect.ValueOf(int16(i)))
	case t.Kind() == reflect.Int32:
		i, err := strconv.Atoi(strings.Split(o.String(), ".")[0])
		if err != nil {
			return err
		}
		t.Set(reflect.ValueOf(int32(i)))
	case t.Kind() == reflect.Int64:
		i, err := strconv.ParseInt(strings.Split(o.String(), ".")[0], 10, 64)
		if err != nil {
			return err
		}
		t.SetInt(i)
	case t.Kind() == reflect.Uint:
		i, err := strconv.Atoi(strings.Split(o.String(), ".")[0])
		if err != nil {
			return err
		}
		t.Set(reflect.ValueOf(uint(i)))
	case t.Kind() == reflect.Uint8:
		i, err := strconv.Atoi(strings.Split(o.String(), ".")[0])
		if err != nil {
			return err
		}
		t.Set(reflect.ValueOf(uint8(i)))
	case t.Kind() == reflect.Uint16:
		i, err := strconv.Atoi(strings.Split(o.String(), ".")[0])
		if err != nil {
			return err
		}
		t.Set(reflect.ValueOf(uint16(i)))
	case t.Kind() == reflect.Uint32:
		i, err := strconv.Atoi(strings.Split(o.String(), ".")[0])
		if err != nil {
			return err
		}
		t.Set(reflect.ValueOf(uint32(i)))
	case t.Kind() == reflect.Uint64:
		i, err := strconv.ParseUint(strings.Split(o.String(), ".")[0], 10, 64)
		if err != nil {
			return err
		}
		t.Set(reflect.ValueOf(uint64(i)))
	case t.Kind() == reflect.Float32:
		f, err := strconv.ParseFloat(o.String(), 32)
		if err != nil {
			return err
		}
		t.Set(reflect.ValueOf(float32(f)))
	case t.Kind() == reflect.Float64:
		f, err := strconv.ParseFloat(o.String(), 64)
		if err != nil {
			return err
		}
		t.SetFloat(f)
	case t.Kind() == reflect.Bool:
		if s := strings.ToLower(o.String()); s == "true" {
			t.SetBool(true)
		} else {
			t.SetBool(false)
		}
	case t.Kind() == reflect.Interface:
		t.Set(reflect.ValueOf(o.Interface()))
	case t.Type().String() == "time.Time":
		v, err := time.Parse(time.RFC3339, o.String())
		if err != nil {
			return err
		}
		t.Set(reflect.ValueOf(v))
	}
	return nil
}
