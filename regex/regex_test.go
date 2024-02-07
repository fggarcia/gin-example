package regex

import (
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

const regexPattern = "^(?!(453998|426398|462437|451212|456188|435087|404025|409280|406176|478507|430360|451302|410349|432958|417401)).*"

func TestRegex(t *testing.T) {
	_, err := regexp.Compile(regexPattern)
	assert.NotEqual(t, nil, err)
	ok, err2 := regexp.Match(regexPattern, []byte("453998"))
	assert.Equal(t, false, ok)
	assert.NotEqual(t, nil, err2)
}
