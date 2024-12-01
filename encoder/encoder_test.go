package encoder

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"gin-example/model"
	"testing"
)

func TestEncoder(t *testing.T) {
	encoder := &JsonEncoder[model.Album]{}
	album := model.Album{
		ID:     "1",
		Title:  "title",
		Artist: "artist",
		Price:  9.99,
	}
	encoded, err := encoder.Encode(album)
	if err != nil {
		t.Fatal("failed to encode album")
	}
	decodedAlbum, err := encoder.Decode(encoded)

	if album != *decodedAlbum {
		t.Fatal("failed to decode album")
	}
}

func TestEmptySlice(t *testing.T) {
	var slice []model.Album
	encoded, err:= encode(slice)
	if err!= nil {
        t.Fatal("failed to encode empty slice")
    }

	fmt.Println(encoded)
}

func encode(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(v)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
