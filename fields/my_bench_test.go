package fields

import "testing"

func BenchmarkItemOmitEmptyTag(b *testing.B) {
	b.ReportAllocs()

	s := Person{}

	for b.Loop() {
		ItemStructField(s)
	}
}

func BenchmarkCustomOmitEmptyTag(b *testing.B) {
	b.ReportAllocs()

	s := Person{}

	for b.Loop() {
		CustomStructField(s)
	}
}
