package gzip

import (
	"testing"
)

func BenchmarkWithoutPool(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		unzipValue(gzipAlbum)
	}
}

func BenchmarkWithPool(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		unzipPoolValue(gzipAlbum)
	}
}
