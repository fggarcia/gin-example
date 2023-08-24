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
