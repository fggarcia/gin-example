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

	b.Run("generic", func(b *testing.B) {
		b.ReportAllocs()

		for b.Loop() {
			encoded, _ := encoder.Encode(album)
			doSomethingEncoded(encoded)
		}
	})

	b.Run("old", func(b *testing.B) {
		b.ReportAllocs()

		for b.Loop() {
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
	b.Run("generic", func(b *testing.B) {
		b.ReportAllocs()

		for b.Loop() {
			decoded, _ := encoder.Decode(encoded)
			doSomethingDecoded(decoded)
		}
	})
	b.Run("generic_stack", func(b *testing.B) {
		b.ReportAllocs()

		for b.Loop() {
			var instance model.Album
			encoder.DecodeTo(encoded, &instance)
			doSomethingDecoded(&instance)
		}
	})
	b.Run("old", func(b *testing.B) {
		b.ReportAllocs()

		for b.Loop() {
			decoded, _ := oldDecoder(encoded)
			doSomethingDecoded(decoded)
		}
	})
}

func doSomethingEncoded(encoded []byte) {

}

func doSomethingDecoded(decoded *model.Album) {

}
