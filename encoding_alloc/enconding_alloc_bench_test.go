package encoding_alloc

import (
	"gin-example/model"
	"github.com/goccy/go-json"
	"testing"
)

func NewFromStack(data []byte, ptr *model.Album) *model.Album {
	json.Unmarshal(data, ptr)
	return ptr
}

func NewFromHeap(data []byte) *model.Album {
	var album model.Album
	json.Unmarshal(data, &album)
	return &album
}
func BenchmarkEncodingAlloc(b *testing.B) {
	album := model.Album{
		ID:     "1",
		Title:  "title",
		Artist: "artist",
		Price:  9.99,
	}

	encoded, _ := json.Marshal(album)

	b.Run("pre_stack", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			albumInstance := model.Album{}
			_ = NewFromStack(encoded, &albumInstance)
			doSomething(&albumInstance)
		}
	})
	b.Run("pre_stack_2", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			albumPtr := new(model.Album)
			_ = NewFromStack(encoded, albumPtr)
			doSomething(albumPtr)
		}
	})
	b.Run("from_heap", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			albumPtr := NewFromHeap(encoded)
			doSomething(albumPtr)
		}
	})
}

func doSomething(ptr *model.Album) {
}
