package zlib

import (
	"fmt"
	"testing"
)

func BenchmarkZlibEncoder(b *testing.B) {
	b.Run("flush", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			bytes, err := zlibEncodeFlush(itemKVSJsonBytes)
			if err != nil {
				fmt.Printf("err: %v\n", err)
			}
			something(bytes)
		}
	})
	b.Run("close", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			bytes, err := zlibEncodeClose(itemKVSJsonBytes)
			if err != nil {
				fmt.Printf("err: %v\n", err)
			}
			something(bytes)
		}
	})
	b.Run("old", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			bytes, err := zlibEncode2(itemKVSJsonBytes)
			if err != nil {
				fmt.Printf("err: %v\n", err)
			}
			something(bytes)
		}
	})
}
func something(b []byte) {

}
