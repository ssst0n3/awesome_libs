package awesome_reflect

import (
	"github.com/ssst0n3/awesome_libs/test_data"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestIsPointer(t *testing.T) {
	t.Run("pointer", func(t *testing.T) {
		assert.Equal(t, true, IsPointer(&test_data.Challenge{}))
	})
	t.Run("not pointer", func(t *testing.T) {
		assert.Equal(t, false, IsPointer(test_data.Challenge{}))
	})
}

func TestMustPointer(t *testing.T) {
	t.Run("not pointer", func(t *testing.T) {
		assert.PanicsWithValue(t, PanicArgumentMustBePointerOrReference, func() {
			MustPointer(test_data.Challenge{})
		})
	})
	t.Run("pointer", func(t *testing.T) {
		assert.NotPanics(t, func() {
			MustPointer(&test_data.Challenge{})
		})
	})
}

func TestMustNotPointer(t *testing.T) {
	t.Run("not pointer", func(t *testing.T) {
		assert.NotPanics(t, func() {
			MustNotPointer(test_data.Challenge{})
		})
	})
	t.Run("pointer", func(t *testing.T) {
		assert.PanicsWithValue(t, PanicArgumentMustNotBePointerAndReference, func() {
			MustNotPointer(&test_data.Challenge{})
		})
	})
}

func TestReflect(t *testing.T) {
	expect := reflect.ValueOf(test_data.Challenge{})
	t.Run("not pointer", func(t *testing.T) {
		value := Value(test_data.Challenge{})
		assert.Equal(t, expect.Interface(), value.Interface())
	})
	t.Run("pointer", func(t *testing.T) {
		value := Value(&test_data.Challenge{})
		assert.Equal(t, expect.Interface(), value.Interface())
	})
}

func TestValueByModel(t *testing.T) {
	t.Run("not pointer", func(t *testing.T) {
		assert.Equal(t, Value(test_data.Challenge{}).Interface(), ValueByModel(test_data.Challenge{}).Interface())
	})
	t.Run("pointer", func(t *testing.T) {
		assert.PanicsWithValue(t, PanicArgumentMustNotBePointerAndReference, func() {
			ValueByModel(&test_data.Challenge{})
		})
	})
}

func TestValueByPtr(t *testing.T) {
	t.Run("pointer", func(t *testing.T) {
		assert.Equal(t, Value(test_data.Challenge{}).Interface(), ValueByPtr(&test_data.Challenge{}).Interface())
	})
	t.Run("not pointer", func(t *testing.T) {
		assert.PanicsWithValue(t, PanicArgumentMustBePointerOrReference, func() {
			ValueByPtr(test_data.Challenge{})
		})
	})
}

func TestFieldByJsonTag(t *testing.T) {
	t.Run("struct", func(t *testing.T) {
		field, find := FieldByJsonTag(
			Value(struct {
				Name string `json:"name"`
			}{Name: "john"}),
			"name",
		)
		assert.Equal(t, true, find)
		assert.Equal(t, Value("john").Interface(), field.Interface())
	})

	// TODO: is there need to add support for type slice?
	t.Run("slice", func(t *testing.T) {
		//field, find := FieldByJsonTag(
		//	Value([]struct {Name string `json:"name"`}{{Name: "john"}}),
		//	"name",
		//)
		//log.Logger.Info(field)
		//assert.Equal(t, true, find)
		//assert.Equal(t, Value("john").Interface(), field.Interface())
	})
}

func TestEmptyPointerModel(t *testing.T) {
	type model struct {
		A string
	}
	m := model{}
	assert.Equal(t, reflect.TypeOf(model{}), reflect.ValueOf(EmptyPointerOfModel(m)).Elem().Type())
}


func TestEmptyPointerSliceModel(t *testing.T) {
	type model struct {
		A string
	}
	m := model{}
	sm := EmptyPointerOfSliceModel(m)
	assert.Equal(t, reflect.TypeOf([]model{}), reflect.ValueOf(sm).Elem().Type())
}
