package singleflight

import (
	"encoding/json"
	"golang.org/x/sync/singleflight"
	"strconv"
	"sync"
	"testing"
	"time"
)

type internalTask struct {
	wg *sync.WaitGroup
}

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

		for b.Loop() {
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
		tasks := make(chan internalTask, totalTasks)
		for k := 0; k < workerCount; k++ {
			go func() {
				for it := range tasks {
					var item Album
					err := toJsonEncode(&item, bytes)
					if err != nil {
						b.Error(err)
						return
					}
					doSomething(&item)
					it.wg.Done()
				}
			}()
		}

		for b.Loop() {
			var wg sync.WaitGroup
			wg.Add(totalTasks)
			for j := 0; j < totalTasks; j++ {
				tasks <- internalTask{&wg}
			}
			wg.Wait()
		}
		close(tasks)
	})
	b.Run("Singleflight", func(b *testing.B) {
		b.ReportAllocs()

		for b.Loop() {
			var g singleflight.Group
			var wg sync.WaitGroup
			wg.Add(totalTasks)
			for j := 0; j < totalTasks; j++ {
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
	time.Sleep(10 * time.Millisecond)
}
