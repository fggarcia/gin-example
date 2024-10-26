package panic

import (
	"fmt"
	"testing"
)

func TestPain(t *testing.T) {
	fmt.Println("Inicio del programa")
	safeFunction()
	fmt.Println("Fin del programa")
}

func safeFunction() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Se atrapó un panic:", r)
		}
	}()

	defer func() {
		fmt.Println("Segundo defer: Esto se ejecuta antes del recover")
	}()

	causePanic()
}

func causePanic() {
	panic("¡Algo salió mal!")
}
