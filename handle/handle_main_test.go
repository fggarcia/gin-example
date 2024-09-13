package handle

import (
	"fmt"
	"testing"
	"unique"
	"unsafe"
)

type KeyMapStruct struct {
	Part1, Part2 string
}

func TestSizeOf(t *testing.T) {
	var str = "string"
	var part1 = "part1"
	var part2 = "part2"
	var handle = unique.Make(str)
	var handle2 = unique.Make(str)

	key := KeyMapStruct{
		Part1: part1,
		Part2: part2,
	}

	//get memory size of handle
	println("size of handle: ", unsafe.Sizeof(handle))
	println("size of string: ", unsafe.Sizeof("string"))
	fmt.Println(fmt.Sprintf("string pointer handle: %v", unsafe.StringData(handle.Value())))
	fmt.Println(fmt.Sprintf("string pointer handle2: %v", unsafe.StringData(handle2.Value())))
	fmt.Println(fmt.Sprintf("string pointer: %v", unsafe.StringData(str)))
	fmt.Println(fmt.Sprintf("key/Part1: %v", unsafe.StringData(key.Part1)))
	fmt.Println(fmt.Sprintf("key/Part2: %v", unsafe.StringData(key.Part2)))
	fmt.Println(fmt.Sprintf("part1 %v", unsafe.StringData(part1)))
	fmt.Println(fmt.Sprintf("part2 %v", unsafe.StringData(part2)))
}
