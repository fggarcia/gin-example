package fields

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLink(T *testing.T) {
	var person Person
	person.Name = "John"
	person.Age = 30
	person.ID = 1

	StructField(person)
}

func TestSameOrderName(T *testing.T) {
	var person Person
	for i := 0; i < 3; i++ {
		structValue := reflect.Indirect(reflect.ValueOf(person))
		for i := 0; i < structValue.NumField(); i++ {
			field := structValue.Type().Field(i)
			name := getFieldName(field)
			fmt.Printf("field %d name: %s\n", i, name)
		}
	}
}
