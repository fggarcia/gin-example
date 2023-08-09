package oreilly

import (
	"fmt"
	"testing"
)

func TestRange(t *testing.T) {
	friends := [5]string{"Annie", "Betty", "Charley", "Doug", "Edward"}
	for i, v := range friends {
		fmt.Printf("Value [%s] \tAddress [%p] \tIndexAddr[%p]\n", v, &v, &friends[i])
	}
}
