package singleflight

import (
	"encoding/json"
	"golang.org/x/sync/singleflight"
	"strconv"
	"sync"
	"testing"
)

var totalTasks = 10

func Benchmark(b *testing.B) {
	album := Album{
		ID:     strconv.Itoa(1),
		Title:  "Blue Train",
		Artist: "John Coltrane",
		Price:  56.99,
	}
	bytes, _ := json.Marshal(&album)

	b.ReportAllocs()
	b.ResetTimer()
	b.Run("default", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			var wg sync.WaitGroup
			wg.Add(totalTasks)
			for i := 0; i < totalTasks; i++ {
				go func() {
					var item Album
					toJsonEncode(&item, bytes)
					doSomething(&item)
					wg.Done()
				}()
			}
			wg.Wait()
		}
	})
	b.Run("work-pattern", func(b *testing.B) {
		b.ReportAllocs()
		var workerCount = 5 // number of workers
		for i := 0; i < b.N; i++ {
			var wg sync.WaitGroup
			tasks := make(chan int, totalTasks)
			for i := 0; i < workerCount; i++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					for range tasks {
						var item Album
						err := toJsonEncode(&item, bytes)
						if err != nil {
							b.Error(err)
							return
						}
						doSomething(&item)
					}
				}()
			}
			for i := 0; i < totalTasks; i++ {
				tasks <- i
			}
			close(tasks)
			wg.Wait()
		}
	})
	b.Run("Singleflight", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			var g singleflight.Group
			var wg sync.WaitGroup
			wg.Add(totalTasks)
			for i := 0; i < totalTasks; i++ {
				go func() {
					g.Do("key", func() (interface{}, error) {
						var item Album
						err := toJsonEncode(&item, bytes)
						doSomething(&item)
						return item, err
					})
					wg.Done()
				}()
			}
			wg.Wait()
		}
	})
}

func doSomething(entity *Album) {

}
