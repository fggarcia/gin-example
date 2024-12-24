package split

import (
	"testing"
)

func BenchmarkSplitBy(b *testing.B) {
	b.Run("split", func(b *testing.B) {
		b.ReportAllocs()
		testString := "12-34"
		for i := 0; i < b.N; i++ {
			splitByDashUsingSplit(testString, '-')
		}
	})
	b.Run("bytes", func(b *testing.B) {
		b.ReportAllocs()
		testString := "12-34"
		for i := 0; i < b.N; i++ {
			splitByDashUsingBytes(testString, '-')
		}
	})
}
