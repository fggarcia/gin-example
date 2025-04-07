package stack

import (
	"testing"
)

func BenchmarkStack(b *testing.B) {
	b.Run("stack", func(b *testing.B) {
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
	})
	b.Run("stack-pointer", func(b *testing.B) {
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
	})
	b.Run("stack-call-pointer", func(b *testing.B) {
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
	})
}
