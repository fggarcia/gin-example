package string

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConcatenate(t *testing.T) {
	value := "value"
	concatFmt := fmtSprintf(value)
	concatSlice := sliceCompose(value)
	concatSB := concatStringBuilder(value)
	assert.Equal(t, concatFmt, "key_value")
	assert.Equal(t, concatFmt, concatSlice)
	assert.Equal(t, concatFmt, concatSB)
}
