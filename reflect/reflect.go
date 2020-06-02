package reflect

import (
	"reflect"
)

const (
	PanicArgumentMustBePointerOrReference = "the argument must be pointer or reference"
	PanicArgumentMustNotBePointerAndReference = "the argument must'nt be pointer and reference"
)

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
