package print

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"testing"
	"unsafe"
)

var result string

func Benchmark(b *testing.B) {
	b.Run("Sprintf", func(b *testing.B) {
		var r string
		s1 := "hello"
		s2 := "world"
		f := 3.14159265359
		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			r = fmt.Sprintf("%s_%s_%f", s1, s2, f)
		}
		result = r
	})
	b.Run("Fprintf", func(b *testing.B) {
		var r string
		s1 := "hello"
		s2 := "world"
		f := 3.14159265359
		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			var buf bytes.Buffer
			fmt.Fprintf(&buf, "%s_%s_%f", s1, s2, f)
			r = buf.String()
		}
		result = r
	})
	b.Run("NoAlloc", func(b *testing.B) {
		var r string
		s1 := "hello"
		s2 := "world"
		f := 3.14159265359
		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			buf := make([]byte, 0, 64)
			buf = append(buf, s1...)
			buf = append(buf, '_')
			buf = append(buf, s2...)
			buf = append(buf, '_')
			buf = strconv.AppendFloat(buf, f, 'f', -1, 64)
			r = string(buf) // esta conversión aloca
		}
		result = r
	})
}

func BenchmarkStringConcatenation(b *testing.B) {
	s1 := "hello"
	s2 := "world"
	f := 3.14159265359

	b.Run("WithMake", func(b *testing.B) {
		var r string
		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			buf := make([]byte, 0, 64)
			buf = append(buf, s1...)
			buf = append(buf, '_')
			buf = append(buf, s2...)
			buf = append(buf, '_')
			buf = strconv.AppendFloat(buf, f, 'f', -1, 64)
			r = *(*string)(unsafe.Pointer(&buf))
		}
		result = r
	})

	b.Run("WithStackArray", func(b *testing.B) {
		var r string
		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			var arr [64]byte  // Array en la pila (stack)
			buf := arr[:0]    // Slice que apunta al array pero con longitud 0
			buf = append(buf, s1...)
			buf = append(buf, '_')
			buf = append(buf, s2...)
			buf = append(buf, '_')
			buf = strconv.AppendFloat(buf, f, 'f', -1, 64)
			r = *(*string)(unsafe.Pointer(&buf))
		}
		result = r
	})

	b.Run("WithStackArrayOutsideLoop", func(b *testing.B) {
		var r string
		var arr [64]byte  // Array fuera del bucle
		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			buf := arr[:0]
			buf = append(buf, s1...)
			buf = append(buf, '_')
			buf = append(buf, s2...)
			buf = append(buf, '_')
			buf = strconv.AppendFloat(buf, f, 'f', -1, 64)
			r = *(*string)(unsafe.Pointer(&buf))
		}
		result = r
	})

	// Comparemos también con una implementación usando strings.Builder
	b.Run("WithStringsBuilder", func(b *testing.B) {
		var r string
		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			var sb strings.Builder
			sb.Grow(64)
			sb.WriteString(s1)
			sb.WriteByte('_')
			sb.WriteString(s2)
			sb.WriteByte('_')
			sb.WriteString(strconv.FormatFloat(f, 'f', -1, 64))
			r = sb.String()
		}
		result = r
	})

	// Verifiquemos que todas las soluciones producen el mismo resultado
	b.Run("VerifyResults", func(b *testing.B) {
		expected := fmt.Sprintf("%s_%s_%f", s1, s2, f)

		buf1 := make([]byte, 0, 64)
		buf1 = append(buf1, s1...)
		buf1 = append(buf1, '_')
		buf1 = append(buf1, s2...)
		buf1 = append(buf1, '_')
		buf1 = strconv.AppendFloat(buf1, f, 'f', 6, 64)
		result1 := *(*string)(unsafe.Pointer(&buf1))

		var arr2 [64]byte
		buf2 := arr2[:0]
		buf2 = append(buf2, s1...)
		buf2 = append(buf2, '_')
		buf2 = append(buf2, s2...)
		buf2 = append(buf2, '_')
		buf2 = strconv.AppendFloat(buf2, f, 'f', 6, 64)
		result2 := *(*string)(unsafe.Pointer(&buf2))

		if expected != result1 || expected != result2 {
			b.Fatalf("Results don't match: expected=%s, result1=%s, result2=%s",
				expected, result1, result2)
		}

		b.Logf("All results correct: %s", expected)
	})
}
