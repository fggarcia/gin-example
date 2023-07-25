package stack

import (
	"testing"
)

func BenchmarkStack(b *testing.B) {
	var (
		name     = "name"
		likes    = []string{"likes"}
		dislikes = []string{"dislikes"}
		age      = 1
	)
	stack := NewStack(name, likes, dislikes, age)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		stack.WithName("name_2")
	}
}

func BenchmarkStackPointer(b *testing.B) {
	var (
		name     = "name"
		likes    = []string{"likes"}
		dislikes = []string{"dislikes"}
		age      = 1
	)
	stack := NewStackPointer(name, likes, dislikes, age)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		stack = WithNamePointer(stack, "name_2")
	}
}

func BenchmarkStackCallPointer(b *testing.B) {
	var (
		name     = "name"
		likes    = []string{"likes"}
		dislikes = []string{"dislikes"}
		age      = 1
	)
	stack := NewStack(name, likes, dislikes, age)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		WithNamePointer(&stack, "name_2")
	}
}
