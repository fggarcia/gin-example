package stack_heap

import "testing"

const (
	name = "John"
	age  = 30
)

func BenchmarkStack(b *testing.B) {
	b.ReportAllocs()

	for b.Loop() {
		createPersonStack(name, age)
	}
}

func BenchmarkHeap(b *testing.B) {
	b.ReportAllocs()

	for b.Loop() {
		createPersonHeap(name, age)
	}
}

func BenchmarkPointer(b *testing.B) {
	b.ReportAllocs()

	for b.Loop() {
		var person Person
		createPersonPointer(&person, name, age)
	}
}
