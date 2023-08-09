package oreilly

import "fmt"

var array = [3]int{1, 2, 3}

// MODULO 3.2 ARRAY
func iterateArray() {
	//aca accedo a la posicion del array posta
	for i := 0; i < len(array); i++ {
		_ = array[i]
	}

	//aca itero sobre una copia del array
	for _, v := range array {
		_ = v
	}

	//prueba
	friends := [5]string{"Annie", "Betty", "Charley", "Doug", "Edward"}
	for i, v := range friends {
		fmt.Printf("Value [%s] \tAddress [%p] \tIndexAddr[%p]\n", v, &v, &friends[i])
	}
}

//ESTRUCTURA DEL SLICE PTR al array, len, cap
//struct {} NO GENERA ALOCACION
//var slice []string crea algo que apunta a nil, 0, 0
//var slice = []string{} crea algo que apunta a un array struct{}, 0, 0
//mas eficiente for con el length que hacer el append
//buff[2:4] es un slice que apunta al array buff, desde el indice 2 hasta el 4 (no inclusive) con len 2 y cap 3 (suponiendo buff un slice de 5)
//buff[2:4:4] es un slice que apunta al array buff, desde el indice 2 hasta el 4 (no inclusive) con len 2 y cap 2 (suponiendo buff un slice de 5)

//ARRAY TO SLICE 3.3 slices - part5
//buf[:] crear un slice que apunta al array buf, desde el indice 0 hasta el len del array, con len igual al len del array y cap igual al len del array
//por ejemplo para cuando se hace un append o un copy

//Value and pointer semantics
//cuando paso un slice a un método, mejor dejar que haga un copy, total lo que unico que hace es un copy del puntero, len y cap y no una copia del slice en heap
//to do lo que sea built-in types, maps, slices, channels, interfaces, son reference types -> pass by value, excepto si hay metodos de marshal/unmarshal o encode/decode

// 4.1 methods
type data struct {
	name string
	age  int
}

func (d data) displayName() {
	fmt.Println("My Name Is", d.name)
}

func (d *data) setAge(age int) {
	d.age = age
}

func main() {
	d := data{
		name: "bill",
	}

	//this is syntactic sugar for (&d).setAge(45)
	d.displayName()
	d.setAge(45)
	data.displayName(d)
	(*data).setAge(&d, 45)
}

//cast
//var ml MoverLocker
// var m Mover
// ml = bike{}
// m = ml //copy
// b := m.(bike) //cast can thought panic
// b, ok := m.(bike) not panic
