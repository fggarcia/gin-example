package fields

import (
	"fmt"
	"reflect"
	"strings"
	"sync"
)

type MyStruct struct {
	Field1 int
	Field2 bool
	Field3 string
	Field4 bool
	Field5 float64
	Field6 bool
}

var fieldCache = sync.Map{}

func createSliceFor(t reflect.Type) []byte {
	byteSlice := make([]byte, (t.NumField()+7)/8)

	for i := 0; i < t.NumField(); i++ {
		fieldType := t.Field(i)

		tag := fieldType.Tag.Get("json")
		// Si el tag contiene omitempty, establecer el bit correspondiente a 1
		if strings.Contains(tag, "omitempty") {
			byteIndex := i / 8
			bitIndex := i % 8
			byteSlice[byteIndex] |= 1 << bitIndex
		}
	}
	return byteSlice
}

func isOmitEmpty(byteSlice []byte, fieldNumber int) bool {
	byteIndex := fieldNumber / 8
	bitIndex := fieldNumber % 8
	return (byteSlice[byteIndex] & (1 << bitIndex)) != 0
}

func CustomStructField(s interface{}) {
	var slice []byte
	t := reflect.TypeOf(s)
	//fullName := util.ToBytes(t.String())
	fullName := t.String()

	if value, ok := fieldCache.Load(fullName); !ok {
		//fmt.Printf("creating slice for %s\n", fullName)
		slice = createSliceFor(t)
		fieldCache.Store(fullName, slice)
	} else {
		slice = value.([]byte)
	}

	for i := 0; i < t.NumField(); i++ {
		//field := t.Field(i)
		//fmt.Printf("custom field %d name: %s\n", i, field.Name)
		if isOmitEmpty(slice, i) {
			//fmt.Printf("omitting field %s\n", field.Name)
			continue
		}
	}
}

func myFunction() {
	// Crear una instancia de la estructura
	s := MyStruct{}

	t := reflect.TypeOf(s)

	// Crear un slice de bytes con suficiente capacidad para almacenar un bit por cada campo
	byteSlice := make([]byte, (t.NumField()+7)/8)

	for i := 0; i < t.NumField(); i++ {
		fieldType := t.Field(i).Type

		// Si el campo es bool, establecer el bit correspondiente a 1
		if fieldType.Kind() == reflect.Bool {
			byteIndex := i / 8
			bitIndex := i % 8
			byteSlice[byteIndex] |= 1 << bitIndex
		}
	}

	// Obtener el nombre del paquete y la estructura
	fullName := t.String()
	//splitName := strings.Split(fullName, ".")
	//packageName := splitName[0]
	//structName := splitName[1]

	// Crear la key del mapa
	mapKey := []byte(fullName)

	// Crear el mapa
	myMap := make(map[string][]byte)
	myMap[string(mapKey)] = byteSlice

	// Imprimir el mapa
	for k, v := range myMap {
		fmt.Printf("Key: %s, Value: ", k)
		for _, b := range v {
			fmt.Printf("%08b ", b)
		}
		fmt.Println()
	}
}
