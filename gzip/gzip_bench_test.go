package gzip

import (
	"testing"
)

func BenchmarkWithoutPoolAlbum(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		unzipValue(gzipAlbum)
	}
}

func BenchmarkWithPoolAlbum(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		unzipPoolValue(gzipAlbum)
	}
}

func Benchmark_WithoutPool_64b(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		unzipValue(gzipPayload64b)
	}
}

func Benchmark_WithPool_64b(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		unzipPoolValue(gzipPayload64b)
	}
}

func Benchmark_WithoutPool_1k(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		unzipValue(gzipPayload1k)
	}
}

func Benchmark_WithPool_1k(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		unzipPoolValue(gzipPayload1k)
	}
}

func Benchmark_WithoutPool_5k(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		unzipValue(gzipPayload5k)
	}
}

func Benchmark_WithPool_5k(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		unzipPoolValue(gzipPayload5k)
	}
}

func Benchmark_WithoutPool_1M(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		unzipValue(gzipPayload1M)
	}
}

func Benchmark_WithPool_1M(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		unzipPoolValue(gzipPayload1M)
	}
}
