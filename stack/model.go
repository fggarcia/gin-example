package stack

type Stack struct {
	name     string
	likes    []string
	dislikes []string
	age      int
}

func NewStack(name string, likes []string, dislikes []string, age int) Stack {
	return Stack{
		name:     name,
		likes:    likes,
		dislikes: dislikes,
		age:      age,
	}
}

//go:noinline
func NewStackPointer(name string, likes []string, dislikes []string, age int) *Stack {
	return &Stack{
		name:     name,
		likes:    likes,
		dislikes: dislikes,
		age:      age,
	}
}

func (s *Stack) WithName(name string) {
	s.name = name
}

//go:noinline
func WithNamePointer(s *Stack, name string) *Stack {
	s.name = name
	return s
}
