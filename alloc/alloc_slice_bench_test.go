package alloc

import (
	"testing"
)

// go test -gcflags="-m -m" -run=^$ -bench=BenchmarkAlloc -cpuprofile=cpu.pprof -memprofile=mem.pprof -count=3 | tee new.txt
var allocNumber = 512

const (
	allocNumberConst = 512
)

func createSliceConstNumber() []byte {
	return make([]byte, 512)
}
func createSliceVar() []byte {
	return make([]byte, allocNumber)
}

func createSliceConst() []byte {
	return make([]byte, allocNumberConst)
}

func BenchmarkAlloc(b *testing.B) {
	b.Run("const_number", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = make([]byte, 512)
		}
	})
	b.Run("const_var", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = make([]byte, allocNumberConst)
		}
	})
	b.Run("var", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = make([]byte, allocNumber)
		}
	})
	b.Run("func_const_number", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = createSliceConstNumber()
		}
	})
	b.Run("func_const_var", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = createSliceConst()
		}
	})
	b.Run("func_var", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = createSliceVar()
		}
	})
}
