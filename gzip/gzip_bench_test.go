package gzip

import (
	"testing"
)

func Benchmark_Album_WithoutPool(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		unzipValue(gzipAlbum)
	}
}

func Benchmark_Album_WithPool(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		unzipPoolValue(gzipAlbum)
	}
}

func Benchmark_Item_WithoutPool(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		unzipValue(gzipItem)
	}
}

func Benchmark_Item_WithPool(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		unzipPoolValue(gzipItem)
	}
}

func Benchmark_Search_WithoutPool(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		unzipValue(gzipSearchApiGOJSON)
	}
}

func Benchmark_Search_WithPool(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		unzipPoolValue(gzipSearchApiGOJSON)
	}
}

func Benchmark_ItemKVS_WithoutPool(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		unzipValue(gzipItemKVSGOJSON)
	}
}

func Benchmark_ItemKVS_WithPool(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		unzipPoolValue(gzipItemKVSGOJSON)
	}
}

func Benchmark_64b_WithoutPool(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		unzipValue(gzipPayload64b)
	}
}

func Benchmark_64b_WithPool(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		unzipPoolValue(gzipPayload64b)
	}
}

func Benchmark_1k_WithoutPool(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		unzipValue(gzipPayload1k)
	}
}

func Benchmark_1k_WithPool(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		unzipPoolValue(gzipPayload1k)
	}
}

func Benchmark_5k_WithoutPool(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		unzipValue(gzipPayload5k)
	}
}

func Benchmark_5k_WithPool(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		unzipPoolValue(gzipPayload5k)
	}
}

func Benchmark_1M_WithoutPool(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		unzipValue(gzipPayload1M)
	}
}

func Benchmark_1M_WithPool(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		unzipPoolValue(gzipPayload1M)
	}
}
