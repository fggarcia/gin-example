package handle

import (
	"reflect"
	"testing"
	"unique"
	"unsafe"
)

type HandleStruct struct {
	Text string `json:"text"`
}

func TestHandle(t *testing.T) {
	var str = "string"
	var h = unique.Make(str)
	mapr := make(map[unique.Handle[string]]HandleStruct)
	v := HandleStruct{h.Value()}
	d := h.Value()
	println("d: %p", (*reflect.StringHeader)(unsafe.Pointer(&v.Text)).Data)
	println("str: %p", (*reflect.StringHeader)(unsafe.Pointer(&d)).Data)

	mapr[h] = v

	r := mapr[unique.Make(str)]
	println("r: %p", (*reflect.StringHeader)(unsafe.Pointer(&r.Text)).Data)
	/**
	d: %p 1374390166000
	str: %p 1374390166000
	*/
}
