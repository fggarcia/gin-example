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

	b.ReportAllocs()

	for b.Loop() {
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

	b.ReportAllocs()

	for b.Loop() {
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

	b.ReportAllocs()

	for b.Loop() {
		WithNamePointer(&stack, "name_2")
	}
}
