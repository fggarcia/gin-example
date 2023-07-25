package fields

import (
	"fmt"
	"github.com/goccy/go-json"
	"github.com/google/go-cmp/cmp"
	"reflect"
	"strings"
)

type Person struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
	ID   int    `json:"id"`
}

func HasValueStruct(i interface{}) bool {
	if i == nil {
		return false
	}
	iType := reflect.TypeOf(i)
	zero := reflect.Zero(iType)
	return !cmp.Equal(i, zero.Interface())
}

func getFieldName(field reflect.StructField) string {
	tag := field.Tag.Get("json")
	if i := strings.Index(tag, ","); i != -1 {
		return tag[:i]
	}
	return tag
}

func shouldOmit(field reflect.StructField, value interface{}) bool {
	tag := field.Tag.Get("json")
	return strings.Contains(tag, "omitempty") && !HasValueStruct(value)
}

func shouldOmitEmpty(field reflect.StructField) bool {
	tag := field.Tag.Get("json")
	return strings.Contains(tag, "omitempty")
}

func getFieldJson(field reflect.StructField, value interface{}) (fieldJson string, err error) {
	name := getFieldName(field)
	bytes, err := json.Marshal(value)
	if err != nil {
		return
	}

	fieldJson = `"` + name + `":` + string(bytes)

	return
}

func StructField(i interface{}) {
	structValue := reflect.Indirect(reflect.ValueOf(i))
	pkgName := structValue.Type().PkgPath()
	structName := structValue.Type().Name()
	fmt.Printf("pkg: %v name: %v\n", pkgName, structName)
	for i := 0; i < structValue.NumField(); i++ {
		field := structValue.Type().Field(i)
		value := structValue.Field(i).Interface()

		if shouldOmit(field, value) {
			continue
		}

		fieldJson, _ := getFieldJson(field, value)

		fmt.Println(fieldJson)
	}
}

func ItemStructField(i interface{}) {
	structValue := reflect.Indirect(reflect.ValueOf(i))
	for i := 0; i < structValue.NumField(); i++ {
		field := structValue.Type().Field(i)
		//fmt.Printf("item field %d name: %s\n", i, field.Name)
		if shouldOmitEmpty(field) {
			//fmt.Printf("omitting field %s\n", field.Name)
			continue
		}
	}
}
