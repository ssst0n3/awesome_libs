package reflect

import (
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
