package model

import (
	json "encoding/json"
	"fmt"
	"strconv"
	"testing"

	"github.com/amazon-ion/ion-go/ion"
	"github.com/apache/fory/go/fory"
	"github.com/bytedance/sonic"
	jsonv2 "github.com/go-json-experiment/json"
	gojson "github.com/goccy/go-json"
	jsoniter "github.com/json-iterator/go"
	segment "github.com/segmentio/encoding/json"
	sonnet "github.com/sugawarayuuta/sonnet"
	jettison "github.com/wI2L/jettison"
)

func marshalJson(name string, f func(interface{}) ([]byte, error), v interface{}) []byte {
	var bytes []byte
	bytes, _ = f(v)
	fmt.Println(fmt.Sprintf("Marshaling %s %d bytes", name, len(bytes)))
	return bytes
}

func TestAlbumMarshal(t *testing.T) {
	var bytes []byte
	var err error

	album := &AlbumION{
		ID:     strconv.Itoa(1),
		Title:  "Blue Train",
		Artist: "John Coltrane",
		Price:  56.99,
	}

	fory := fory.NewFory(true)
	if err = fory.RegisterTagType("album_bench_test.AlbumION", AlbumION{}); err != nil {
		panic(err)
	}

	jsoniterV := jsoniter.ConfigCompatibleWithStandardLibrary

	marshalJson("fory", fory.Marshal, album)
	marshalJson("standard", json.Marshal, album)
	marshalJson("gojson", gojson.Marshal, album)
	marshalJson("jsoniter", jsoniterV.Marshal, album)
	marshalJson("segment", segment.Marshal, album)
	marshalJson("jettison", jettison.Marshal, album)
	marshalJson("sonnet", sonnet.Marshal, album)
	marshalJson("sonic", sonic.Marshal, album)

	bytes, _ = jsonv2.Marshal(album)
	fmt.Println(fmt.Sprintf("Marshaling %s %d bytes", "jsonv2", len(bytes)))
	bytes, _ = ion.MarshalBinary(album)
	fmt.Println(fmt.Sprintf("Marshaling %s %d bytes", "ion", len(bytes)))
}
