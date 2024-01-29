package marshal

import (
	"fmt"
	"gin-example/util"
	jsoniter "github.com/json-iterator/go"
	"testing"
)

func TestConvertToMap(t *testing.T) {
	var person Person
	person.Name = "John"
	person.Age = 30
	person.ID = 1

	json := jsoniter.ConfigCompatibleWithStandardLibrary

	bytes, _ := json.Marshal(person)

	var entityMap map[string]interface{}

	json.Unmarshal(bytes, &entityMap)

	println(fmt.Sprintf("map: %v", entityMap))
	println(fmt.Sprintf("string: %s", util.ToString(bytes)))

	var p1 Person
	var p2 Person

	entityMap2, _ := json.Marshal(entityMap)
	json.Unmarshal(entityMap2, &p1)
	json.Unmarshal(bytes, &p2)

	println(fmt.Sprintf("p1: %v", p1))
	println(fmt.Sprintf("p2: %v", p2))
}


func TestMapStruct(t *testing.T) {
	someMap := make(map[string]interface{})

	someMap["key"] = struct{}{}

	val, ok := someMap["key"]
	if _, okCast := val.(struct{}); ok && okCast {
		println("ok")
	} else {
		println("not ok")
	}
}
