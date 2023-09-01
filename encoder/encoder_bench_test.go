package encoder

import (
	"gin-example/model"
	"testing"
)

func BenchmarkEncoder(b *testing.B) {
	encoder := JsonEncoder[model.Album]{}
	album := model.Album{
		ID:     "1",
		Title:  "title",
		Artist: "artist",
		Price:  9.99,
	}
	b.ResetTimer()
	b.Run("generic", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			encoded, _ := encoder.Encode(album)
			doSomethingEncoded(encoded)
		}
	})
	b.ResetTimer()
	b.Run("old", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			encoded, _ := oldEncoder(album)
			doSomethingEncoded(encoded)
		}
	})
}

func BenchmarkDecoder(b *testing.B) {
	encoder := JsonEncoder[model.Album]{}
	album := model.Album{
		ID:     "1",
		Title:  "title",
		Artist: "artist",
		Price:  9.99,
	}
	encoded, _ := encoder.Encode(album)
	b.ResetTimer()
	b.Run("generic", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			decoded, _ := encoder.Decode(encoded)
			doSomethingDecoded(decoded)
		}
	})
	b.ResetTimer()
	b.Run("old", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			decoded, _ := oldDecoder(encoded)
			doSomethingDecoded(decoded)
		}
	})
}

func doSomethingEncoded(encoded []byte) {

}

func doSomethingDecoded(decoded *model.Album) {

}
