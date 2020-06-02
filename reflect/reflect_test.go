package reflect

import (
	"github.com/ssst0n3/awesome_libs/test_data"
	"github.com/stretchr/testify/assert"
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
