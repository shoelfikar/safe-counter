
package main

import (
	"fmt"
	"sync"
)

// Counter struct dengan mutex agar thread-safe
type Counter struct {
	mu    sync.Mutex
	value int
}

// Inc menambahkan counter secara thread-safe
func (c *Counter) Inc() {
	c.mu.Lock()
	c.value++
	c.mu.Unlock()
}

// Value mengembalikan nilai counter secara thread-safe
func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {
	var wg sync.WaitGroup
	counter := &Counter{}

	// jalankan 100 goroutine
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// tiap goroutine menambah 1000 kali
			for j := 0; j < 1000; j++ {
				counter.Inc()
			}
		}()
	}

	// tunggu semua goroutine selesai
	wg.Wait()

	// cetak hasil akhir
	fmt.Println("Final Counter Value:", counter.Value())
}
