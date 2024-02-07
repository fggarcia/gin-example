package gzip

import (
	"testing"
)

func BenchmarkGzipAlbum(b *testing.B) {
	b.Run("without_pool", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			unzipValue(gzipAlbum)
		}
	})

	b.Run("with_pool", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			unzipPoolValue(gzipAlbum)
		}
	})
}

func BenchmarkGzipItem(b *testing.B) {
	b.Run("without_pool", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			unzipValue(gzipItem)
		}
	})

	b.Run("with_pool", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			unzipPoolValue(gzipItem)
		}
	})

}

func BenchmarkGzipSearch(b *testing.B) {
	b.Run("without_pool", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			unzipValue(gzipSearchApiGOJSON)
		}
	})

	b.Run("with_pool", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			unzipPoolValue(gzipSearchApiGOJSON)
		}
	})

}

func BenchmarkItemKVSWithoutPool(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		unzipValue(gzipItemKVSGOJSON)
	}
}

func BenchmarkItemKVSWithPool(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		unzipPoolValue(gzipItemKVSGOJSON)
	}
}

func Benchmark64bWithoutPool(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		unzipValue(gzipPayload64b)
	}
}

func Benchmark64bWithPool(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		unzipPoolValue(gzipPayload64b)
	}
}

func Benchmark1kWithoutPool(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		unzipValue(gzipPayload1k)
	}
}

func Benchmark1kWithPool(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		unzipPoolValue(gzipPayload1k)
	}
}

func Benchmark5kWithoutPool(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		unzipValue(gzipPayload5k)
	}
}

func Benchmark5kWithPool(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		unzipPoolValue(gzipPayload5k)
	}
}

func Benchmark1MWithoutPool(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		unzipValue(gzipPayload1M)
	}
}

func Benchmark1MWithPool(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		unzipPoolValue(gzipPayload1M)
	}
}
