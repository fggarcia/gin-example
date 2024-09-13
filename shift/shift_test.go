package shift

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShift(t *testing.T) {
	assert.Equal(t, 1<<1, 2)
	assert.Equal(t, 1<<2, 4)
	assert.Equal(t, 2<<1, 4)
}
