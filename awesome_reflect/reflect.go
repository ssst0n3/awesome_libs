package awesome_reflect

import (
	"fmt"
	"reflect"
)

const (
	PanicArgumentMustBePointerOrReference     = "the argument must be pointer or reference"
	PanicArgumentMustNotBePointerAndReference = "the argument must'nt be pointer and reference"
)

/*
!!!reflect attention, may cause panic!!!
*/
func IsPointer(model interface{}) bool {
	return reflect.ValueOf(model).Kind() == reflect.Ptr
}

func MustPointer(model interface{}) {
	if !IsPointer(model) {
		// It's a Programming Error, so use panic
		panic(PanicArgumentMustBePointerOrReference)
	}
}

func MustNotPointer(model interface{}) {
	if IsPointer(model) {
		// It's a Programming Error, so use panic
		panic(PanicArgumentMustNotBePointerAndReference)
	}
}

/*
!!!reflect attention, may cause panic!!!
*/
func Value(model interface{}) reflect.Value {
	switch reflect.ValueOf(model).Kind() {
	case reflect.Ptr:
		return reflect.ValueOf(model).Elem()
	default:
		return reflect.ValueOf(model)
	}
}

/*
!!!reflect attention, may cause panic!!!
*/
func ValueByModel(model interface{}) reflect.Value {
	MustNotPointer(model)
	return reflect.ValueOf(model)
}

/*
!!!reflect attention, may cause panic!!!
*/
func ValueByPtr(modelPtr interface{}) reflect.Value {
	MustPointer(modelPtr)
	return reflect.ValueOf(modelPtr).Elem()
}

/*
!!!reflect attention, may cause panic!!!
*/
func FieldByJsonTag(v reflect.Value, jsonTag string) (reflect.Value, bool) {
	switch v.Kind() {
	case reflect.Struct:
		for i := 0; i < v.Type().NumField(); i++ {
			if v.Type().Field(i).Tag.Get("json") == jsonTag {
				return v.Field(i), true
			}
			if v.Field(i).Kind() == reflect.Struct {
				if value, find := FieldByJsonTag(v.Field(i), jsonTag); find {
					return value, find
				} else {
					continue
				}
			}
		}
	default:
		panic(fmt.Sprintf("Not support type: %s", v.Kind().String()))
	}
	return reflect.Value{}, false
}
