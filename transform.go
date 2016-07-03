package transform

import (
	"errors"
	"reflect"
)

var buildin = []process{
	stringHandler,
	int64Handler,
	int32Handler,
	intHandler,
	float64Handler,
	float32Handler,
	boolHandler,
}

type process func(f reflect.Value, t reflect.Value) error

// Go to transform one object to another object
func Go(x, y interface{}, custom ...process) error {
	handlers := make([]process, len(custom)+len(buildin))
	for i, v := range custom {
		handlers[i] = v
	}
	for i, v := range buildin {
		handlers[i+len(custom)] = v
	}
	f := reflect.ValueOf(x)
	t := reflect.ValueOf(y)
	if f.Kind() == reflect.Interface || f.Kind() == reflect.Ptr {
		f = f.Elem()
	}
	if f.Kind() != reflect.Struct {
		return errors.New("object [from] is not struct type object")
	}
	if t.Kind() == reflect.Interface || t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return errors.New("object [target] is not struct type object")
	}
	if f.NumField() != t.NumField() {
		return errors.New("the number of fields of objects mismatch")
	}
	for i := 0; i < f.NumField(); i++ {
		fval := f.Field(i)
		tval := t.FieldByName(f.Type().Field(i).Name)
		if !tval.IsValid() {
			return errors.New("field is invalid")
		}
		if !tval.CanSet() {
			return errors.New("field cannot be set")
		}
		for _, c := range handlers {
			if err := c(fval, tval); err != nil {
				panic(err)
			}
		}
	}
	return nil
}
