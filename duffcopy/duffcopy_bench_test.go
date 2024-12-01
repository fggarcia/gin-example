package duffcopy

import (
	"testing"
)

// LargeStruct is a struct with many fields to make the difference more noticeable
type LargeStruct struct {
	Field1, Field2, Field3, Field4, Field5   int
	Field6, Field7, Field8, Field9, Field10  int
	Field11, Field12, Field13, Field14, Field15 int
	Field16, Field17, Field18, Field19, Field20 int
}

func BenchmarkSliceIteration(b *testing.B) {
	// Create a slice with 1024 LargeStruct elements
	slice := make([]LargeStruct, 1024)
	for i := range slice {
		slice[i] = LargeStruct{
			Field1: i, Field2: i, Field3: i, Field4: i, Field5: i,
			Field6: i, Field7: i, Field8: i, Field9: i, Field10: i,
			Field11: i, Field12: i, Field13: i, Field14: i, Field15: i,
			Field16: i, Field17: i, Field18: i, Field19: i, Field20: i,
		}
	}

	b.Run("Range", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			sum := 0
			for _, v := range slice {
				sum += v.Field1 + v.Field10 + v.Field20
			}
		}
	})

	b.Run("For with len", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			sum := 0
			for j := 0; j < len(slice); j++ {
				sum += slice[j].Field1 + slice[j].Field10 + slice[j].Field20
			}
		}
	})

	b.Run("For with constant", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			sum := 0
			for j := 0; j < 1024; j++ {
				sum += slice[j].Field1 + slice[j].Field10 + slice[j].Field20
			}
		}
	})
}
