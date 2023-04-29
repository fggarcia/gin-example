package common

import "github.com/mitchellh/copystructure"

func DeepCopy(v interface{}) (interface{}, error) {
	return copystructure.Copy(v)
	/*
		srcType := reflect.TypeOf(v)
		srcValue := reflect.ValueOf(v)

		if srcType.Kind() == reflect.Ptr {
			srcType = srcType.Elem()
			srcValue = srcValue.Elem()
		}

		dest := reflect.New(srcType).Interface()
		err := copier.Copy(dest, v)
		if err != nil {
			return nil, err
		}

		return dest, nil
	*/
}
