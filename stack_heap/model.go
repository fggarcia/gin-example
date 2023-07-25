package stack_heap

type Person struct {
	name string
	age  int
}

//go:noinline
func createPersonHeap(name string, age int) *Person {
	return &Person{name: name, age: age}
}

func createPersonStack(name string, age int) Person {
	return Person{name: name, age: age}
}

func createPersonPointer(person *Person, name string, age int) {
	person.name = name
	person.age = age
	return
}
