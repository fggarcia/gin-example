package common

import (
	"github.com/mitchellh/copystructure"
	//"github.com/jinzhu/copier"
	//"log"
)

func DeepCopy(v interface{}) (interface{}, error) {
	return copystructure.Copy(v)

	/*
		var result interface{}
		err := copier.CopyWithOption(&result, v, copier.Option{DeepCopy: true})
		if err != nil {
			return nil, err
		}
		return result, nil
	*/
}
