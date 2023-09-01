package encoder

import (
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
