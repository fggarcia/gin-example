package duffcopy

import (
	"testing"
)

const _heavyStructSize = 1_000

// Estructura mucho más pesada
type HeavyStruct struct {
	A int64
	B int64
	C int64
	D int64
	E [1024*1024]byte // Incrementamos el tamaño del array
}

// Benchmark para for range con copia directa
func BenchmarkForRangeCopy(b *testing.B) {
	// Creamos un slice grande de estructuras
	slice := make([]HeavyStruct, _heavyStructSize)
	for i := range slice {
		slice[i] = HeavyStruct{
			A: int64(i),
			B: int64(i * 2),
			C: int64(i * 3),
			D: int64(i * 4),
		}
	}

	var result HeavyStruct // Usamos esto para asegurarnos de que las copias no sean optimizadas
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, item := range slice {
			result = item // Copia explícita para evitar optimizaciones
		}
	}
	_ = result // Evita optimización del compilador
}
func BenchmarkForRangeIdx(b *testing.B) {
	// Creamos un slice grande de estructuras
	slice := make([]HeavyStruct, _heavyStructSize)
	for i := range slice {
		slice[i] = HeavyStruct{
			A: int64(i),
			B: int64(i * 2),
			C: int64(i * 3),
			D: int64(i * 4),
		}
	}

	var result HeavyStruct // Usamos esto para asegurarnos de que las copias no sean optimizadas
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for idx := range slice {
			result = slice[idx] // Copia explícita para evitar optimizaciones
		}
	}
	_ = result // Evita optimización del compilador
}

// Benchmark para for range con punteros
func BenchmarkForRangePointer(b *testing.B) {
	// Creamos un slice grande de punteros a estructuras
	slice := make([]*HeavyStruct, _heavyStructSize)
	for i := range slice {
		slice[i] = &HeavyStruct{
			A: int64(i),
			B: int64(i * 2),
			C: int64(i * 3),
			D: int64(i * 4),
		}
	}

	var result *HeavyStruct // Usamos esto para asegurarnos de que los accesos no sean optimizados
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, item := range slice {
			result = item // Acceso directo a punteros
		}
	}
	_ = result // Evita optimización del compilador
}

func BenchmarkForRangeByIndex(b *testing.B) {
	// Creamos un slice grande de estructuras
	slice := make([]HeavyStruct, _heavyStructSize)
	for i := range slice {
		slice[i] = HeavyStruct{
			A: int64(i),
			B: int64(i * 2),
			C: int64(i * 3),
			D: int64(i * 4),
		}
	}

	var result HeavyStruct
	// Benchmark
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for idx := 0; idx < len(slice); idx++ {
			result = slice[idx] // Simula el uso del valor (evita que el compilador lo optimice fuera)
		}
		_ = result // Evita optimización del compilador
	}
}
