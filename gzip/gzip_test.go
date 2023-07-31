package gzip

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEqual(t *testing.T) {
	var err error
	withoutPool, err := unzipValue(gzipAlbum)

	if err != nil {
		t.Error(err)
	}

	withPool, err := unzipPoolValue(gzipAlbum)

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, withoutPool, withPool)
}
