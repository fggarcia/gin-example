package stack_heap

import "testing"

const (
	name = "John"
	age  = 30
)

func BenchmarkStack(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		createPersonStack(name, age)
	}
}

func BenchmarkHeap(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		createPersonHeap(name, age)
	}
}

func BenchmarkPointer(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var person Person
		createPersonPointer(&person, name, age)
	}
}
