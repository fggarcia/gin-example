package model

import (
	json "encoding/json"
	"github.com/amazon-ion/ion-go/ion"
	fury "github.com/apache/fury/go/fury"
	"github.com/bytedance/sonic"
	jsonv2 "github.com/go-json-experiment/json"
	gojson "github.com/goccy/go-json"
	jsoniter "github.com/json-iterator/go"
	segment "github.com/segmentio/encoding/json"
	sonnet "github.com/sugawarayuuta/sonnet"
	jettison "github.com/wI2L/jettison"
	"reflect"
	"strconv"
	"testing"
)

func BenchmarkAlbumMarshal(b *testing.B) {
	album := &AlbumION{
		ID:     strconv.Itoa(1),
		Title:  "Blue Train",
		Artist: "John Coltrane",
		Price:  56.99,
	}

	b.Run("standard", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			json.Marshal(album)
		}
	})

	b.Run("ion", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			ion.MarshalBinary(album)
		}
	})

	b.Run("gojson", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			gojson.Marshal(album)
		}
	})

	b.Run("jsoniter", func(b *testing.B) {
		var json = jsoniter.ConfigCompatibleWithStandardLibrary
		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			json.Marshal(album)
		}
	})

	b.Run("sonic", func(b *testing.B) {
		sonic.Pretouch(reflect.TypeOf(album))
		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			sonic.Marshal(album)
		}
	})

	b.Run("segmentio", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			segment.Marshal(album)
		}
	})

	b.Run("jettison", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			jettison.Marshal(album)
		}
	})

	b.Run("sonnet", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			sonnet.Marshal(album)
		}
	})
	b.Run("jsonV2", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			jsonv2.Marshal(album)
		}
	})
	b.Run("fury", func(b *testing.B) {
		fury := fury.NewFury(true)
		if err := fury.RegisterTagType("album_bench_test.AlbumION", AlbumION{}); err != nil {
			panic(err)
		}

		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			fury.Marshal(album)
		}
	})
}
