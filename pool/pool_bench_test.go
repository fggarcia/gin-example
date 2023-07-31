package pool

import (
	"gin-example/model"
	"gin-example/util"
	"github.com/segmentio/encoding/json"
	"sync"
	"testing"
)

const (
	albumStr = `{"id":"1","title":"The Dark Side of the Moon","artist":"Pink Floyd","price":10.99}`
	iter     = 12_500
)

var (
	entityPool = &sync.Pool{
		New: func() interface{} {
			return &model.Album{}
		},
	}
	album2Bytes = util.ToBytes(albumStr)
)

func BenchmarkPool(b *testing.B) {
	var album *model.Album
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for j := 0; j < iter; j++ {
			album = entityPool.Get().(*model.Album)
			json.Unmarshal(album2Bytes, album)
			entityPool.Put(album)
		}
	}
}

func BenchmarkNoPool(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for j := 0; j < iter; j++ {
			album := new(model.Album)
			json.Unmarshal(album2Bytes, album)
		}
	}
}
