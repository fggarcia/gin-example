package zlib

import (
	"fmt"
	"testing"
)

func BenchmarkZlibEncoder(b *testing.B) {
	b.Run("flush", func(b *testing.B) {
		b.ReportAllocs()

		for b.Loop() {
			bytes, err := zlibEncodeFlush(itemKVSJsonBytes)
			if err != nil {
				fmt.Printf("err: %v\n", err)
			}
			something(bytes)
		}
	})
	b.Run("close", func(b *testing.B) {
		b.ReportAllocs()

		for b.Loop() {
			bytes, err := zlibEncodeClose(itemKVSJsonBytes)
			if err != nil {
				fmt.Printf("err: %v\n", err)
			}
			something(bytes)
		}
	})
	b.Run("old", func(b *testing.B) {
		b.ReportAllocs()

		for b.Loop() {
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
