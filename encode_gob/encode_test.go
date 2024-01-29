package encode_gob

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"regexp"
	"sync"
	"testing"
)

type MyStruct struct {
	Field1 int `json:"field_1"`
	Field2 string `json:"field_2"`
}

type MyStructRegex struct {
	Field1 int `json:"field_1"`
	Field2 string `json:"field_2"`
	RegexString string `json:"regex"`
	regex *regexp.Regexp
	lock sync.RWMutex
}

func (m *MyStructRegex) GetRegex() *regexp.Regexp {
	m.lock.Lock()
	defer m.lock.Unlock()
	if m.regex == nil {
		m.regex = regexp.MustCompile(m.RegexString)
	}
	return m.regex
}

func TestGob(T *testing.T) {
	// Crear una instancia de la estructura
	myStruct := MyStruct{Field1: 42, Field2: "Hello World", }

	// Buffer para almacenar la codificación
	var buf,buf2 bytes.Buffer

	// Crear un encoder que escriba en el buffer
	enc := gob.NewEncoder(&buf)

	// Codificar el puntero a la estructura
	if err := enc.Encode(&myStruct); err != nil {
		log.Fatal("Encode error:", err)
	}

	myStructRegex := MyStructRegex{Field1: 42, Field2: "Hello World", RegexString: "^\\d+$", }
	enc = gob.NewEncoder(&buf2)
	if err := enc.Encode(&myStructRegex); err != nil {
		log.Fatal("Encode error:", err)
	}

	// Imprimir el slice de bytes resultante
	fmt.Printf("%x\n", buf.Bytes())
	fmt.Printf("%x\n", buf2.Bytes())
}
