package _map

import (
"testing"
)

const (
	numItems = 32 // Number of items to add to the map
)

func BenchmarkMapAllocation(b *testing.B) {
	b.Run("DynamicGrowth", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			m := make(map[int]int)
			for j := 0; j < numItems; j++ {
				m[j] = j
			}
		}
	})

	b.Run("PrecalculatedSize", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			m := make(map[int]int, numItems)
			for j := 0; j < numItems; j++ {
				m[j] = j
			}
		}
	})
}
