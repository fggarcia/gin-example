package handle

import (
	"fmt"
	"sync"
	"testing"
)

func TestSome(t *testing.T) {
	var wg sync.WaitGroup
	var slice []int = make([]int, 10)

	wg.Add(len(slice) * 2)

	for i := range slice {
		myInt := i
		go myGoroutine(myInt, &wg)
		go func(i int) {
			defer wg.Done()
			println(myInt)
		}(i)
	}

	wg.Done()
}

func myGoroutine(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	println(i + 10)
	fmt.Sprintf("%d %v", i, &i)
}
