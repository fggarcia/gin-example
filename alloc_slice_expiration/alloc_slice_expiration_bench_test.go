package alloc

import (
	"encoding/binary"
	"testing"
	"time"
)

// go test -gcflags="-m -m" -run=^$ -bench=BenchmarkExpirationToSlice -cpuprofile=cpu.pprof -memprofile=mem.pprof -count=3 | tee new.txt
func expirationToSlicePtr(expirationTime time.Time) *[]byte {
	var (
		expirationVarInt = expirationTime.Unix()
		bytes            = make([]byte, binary.MaxVarintLen64)
	)

	binary.PutVarint(bytes, expirationVarInt)

	return &bytes
}

func expirationToSlice(expirationTime time.Time) []byte {
	var (
		expirationVarInt = expirationTime.Unix()
		bytes            = make([]byte, binary.MaxVarintLen64)
	)

	binary.PutVarint(bytes, expirationVarInt)

	return bytes
}

func expirationToSliceRef(buf *[]byte, expirationTime time.Time) bool {
	var (
		expirationVarInt = expirationTime.Unix()
	)

	binary.PutVarint(*buf, expirationVarInt)

	return true
}

func BenchmarkExpirationToSlice(b *testing.B) {
	var timeNow = time.Now()
	b.Run("heap", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = expirationToSlicePtr(timeNow)
		}
	})
	b.Run("stack", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = expirationToSlice(timeNow)
		}
	})
	b.Run("ref", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			var buf = make([]byte, binary.MaxVarintLen64)
			_ = expirationToSliceRef(&buf, timeNow)
		}
	})
}
