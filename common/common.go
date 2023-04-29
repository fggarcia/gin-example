package common

import (
	"github.com/mitchellh/copystructure"
)

func DeepCopy(v interface{}) (interface{}, error) {
	return copystructure.Copy(v)
}
