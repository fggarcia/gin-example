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
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			encoded, _ := encoder.Encode(album)
			doSomethingEncoded(encoded)
		}
	})

	b.Run("old", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
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
	b.Run("generic", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			decoded, _ := encoder.Decode(encoded)
			doSomethingDecoded(decoded)
		}
	})
	b.Run("generic_stack", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			var instance model.Album
			encoder.DecodeTo(encoded, &instance)
			doSomethingDecoded(&instance)
		}
	})
	b.Run("old", func(b *testing.B) {
		b.ResetTimer()
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
