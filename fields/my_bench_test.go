package fields

import "testing"

func BenchmarkItemOmitEmptyTag(b *testing.B) {
	s := Person{}
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		ItemStructField(s)
	}
}

func BenchmarkCustomOmitEmptyTag(b *testing.B) {
	s := Person{}
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		CustomStructField(s)
	}
}
