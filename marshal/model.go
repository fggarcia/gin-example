package marshal

type Person struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
	ID   int    `json:"id"`
}
