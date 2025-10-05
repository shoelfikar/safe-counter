
package main

import (
	"fmt"
	"sync"
)

// Counter struct dengan mutex
type Counter struct {
	mu    sync.Mutex
	value int
}

// Inc function untuk menambahkan counter
func (c *Counter) Inc() {
	c.mu.Lock()
	c.value++
	c.mu.Unlock()
}

// Value function mengembalikan nilai counter secara thread-safe
func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {
	var wg sync.WaitGroup
	counter := &Counter{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter.Inc()
			}
		}()
	}

	wg.Wait()

	fmt.Println("Total Counter Value:", counter.Value())
}
