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
	v := HandleStruct{h.Value()}
	d := h.Value()
	println("d: %p", (*reflect.StringHeader)(unsafe.Pointer(&v.Text)).Data)
	println("str: %p", (*reflect.StringHeader)(unsafe.Pointer(&d)).Data)

	/**
	d: %p 1374390166000
	str: %p 1374390166000
	*/
}
