package alloc

import (
	"gin-example/model"
	"gin-example/util"
	json "github.com/goccy/go-json"
	segmentio "github.com/segmentio/encoding/json"
	"sync"
	"testing"
)

const (
	album2 = `{"id":"1","title":"The Dark Side of the Moon","artist":"Pink Floyd","price":10.99}`
)

var (
	entityPool = &sync.Pool{
		New: func() interface{} {
			return &model.Album{}
		},
	}
	album2Bytes = util.ToBytes(album2)
)

func BenchmarkUnmarshalAlbum(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		album := &model.Album{}
		err := json.Unmarshal(album2Bytes, &album)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkUnmarshalSyncAlbum(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		albumPtr := entityPool.Get().(*model.Album)
		err := json.Unmarshal(album2Bytes, albumPtr)
		if err != nil {
			b.Fatal(err)
		}
		entityPool.Put(albumPtr)
	}
}

func BenchmarkUnmarshalSyncAlbum2(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		albumPtr := entityPool.Get().(*model.Album)
		err := segmentio.Unmarshal(album2Bytes, albumPtr)
		if err != nil {
			b.Fatal(err)
		}
		entityPool.Put(albumPtr)
	}
}
