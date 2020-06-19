package slice

import (
	"reflect"
)

const PanicTheSecondParameterMustBeSlice = "The Second Parameter Must Be Slice"

func In(item interface{}, slice interface{}) bool {
	result := false
	value := reflect.ValueOf(slice)
	if value.Kind() == reflect.Slice {
		for i := 0; i < value.Len(); i++ {
			if item == value.Index(i).Interface() {
				result = true
				break
			}
		}
	} else {
		panic(PanicTheSecondParameterMustBeSlice)
	}
	return result
}
