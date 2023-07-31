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

func BenchmarkWithoutPoolSmall(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		unzipValue(gzipSmallPayload)
	}
}

func BenchmarkWithPoolSmall(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		unzipPoolValue(gzipSmallPayload)
	}
}

func BenchmarkWithoutPoolMedium(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		unzipValue(gzipMediumPayload)
	}
}

func BenchmarkWithPoolMedium(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		unzipPoolValue(gzipMediumPayload)
	}
}

func BenchmarkWithoutPoolLarge(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		unzipValue(gzipLargePayload)
	}
}

func BenchmarkWithPoolLarge(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		unzipPoolValue(gzipLargePayload)
	}
}
