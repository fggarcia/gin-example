package gzip

import (
	"testing"
)

func BenchmarkGzipAlbum(b *testing.B) {
	b.Run("without_pool", func(b *testing.B) {
		b.ReportAllocs()

		for b.Loop() {
			unzipValue(gzipAlbum)
		}
	})

	b.Run("with_pool", func(b *testing.B) {
		b.ReportAllocs()

		for b.Loop() {
			unzipPoolValue(gzipAlbum)
		}
	})
}

func BenchmarkGzipItem(b *testing.B) {
	b.Run("without_pool", func(b *testing.B) {
		b.ReportAllocs()

		for b.Loop() {
			unzipValue(gzipItem)
		}
	})

	b.Run("with_pool", func(b *testing.B) {
		b.ReportAllocs()

		for b.Loop() {
			unzipPoolValue(gzipItem)
		}
	})

}

func BenchmarkGzipSearch(b *testing.B) {
	b.Run("without_pool", func(b *testing.B) {
		b.ReportAllocs()

		for b.Loop() {
			unzipValue(gzipSearchApiGOJSON)
		}
	})

	b.Run("with_pool", func(b *testing.B) {
		b.ReportAllocs()

		for b.Loop() {
			unzipPoolValue(gzipSearchApiGOJSON)
		}
	})

}

func BenchmarkItemKVSWithoutPool(b *testing.B) {
	b.ReportAllocs()

	for b.Loop() {
		unzipValue(gzipItemKVSGOJSON)
	}
}

func BenchmarkItemKVSWithPool(b *testing.B) {
	b.ReportAllocs()

	for b.Loop() {
		unzipPoolValue(gzipItemKVSGOJSON)
	}
}

func Benchmark64bWithoutPool(b *testing.B) {
	b.ReportAllocs()

	for b.Loop() {
		unzipValue(gzipPayload64b)
	}
}

func Benchmark64bWithPool(b *testing.B) {
	b.ReportAllocs()

	for b.Loop() {
		unzipPoolValue(gzipPayload64b)
	}
}

func Benchmark1kWithoutPool(b *testing.B) {
	b.ReportAllocs()

	for b.Loop() {
		unzipValue(gzipPayload1k)
	}
}

func Benchmark1kWithPool(b *testing.B) {
	b.ReportAllocs()

	for b.Loop() {
		unzipPoolValue(gzipPayload1k)
	}
}

func Benchmark5kWithoutPool(b *testing.B) {
	b.ReportAllocs()

	for b.Loop() {
		unzipValue(gzipPayload5k)
	}
}

func Benchmark5kWithPool(b *testing.B) {
	b.ReportAllocs()

	for b.Loop() {
		unzipPoolValue(gzipPayload5k)
	}
}

func Benchmark1MWithoutPool(b *testing.B) {
	b.ReportAllocs()

	for b.Loop() {
		unzipValue(gzipPayload1M)
	}
}

func Benchmark1MWithPool(b *testing.B) {
	b.ReportAllocs()

	for b.Loop() {
		unzipPoolValue(gzipPayload1M)
	}
}
