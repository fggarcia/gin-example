package model

import (
	"github.com/amazon-ion/ion-go/ion"
	"github.com/goccy/go-json"
	"strconv"
	"testing"
)

func BenchmarkAlbum_ion(b *testing.B) {
	album := &AlbumION{
		ID:     strconv.Itoa(1),
		Title:  "Blue Train",
		Artist: "John Coltrane",
		Price:  56.99,
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ion.MarshalBinary(album)
	}
}

func BenchmarkAlbum_gojson(b *testing.B) {
	album := &AlbumION{
		ID:     strconv.Itoa(1),
		Title:  "Blue Train",
		Artist: "John Coltrane",
		Price:  56.99,
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		json.Marshal(album)
	}
}
