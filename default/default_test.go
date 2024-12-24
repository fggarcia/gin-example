package _default

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDefaultValues(t *testing.T) {
	var slice []int

	if slice != nil {
		t.Error("Expected slice to be nil")
	}

	assert.Nil(t, slice)
	assert.Equal(t, 0, cap(slice))
	assert.Equal(t, 0, len(slice))
}
