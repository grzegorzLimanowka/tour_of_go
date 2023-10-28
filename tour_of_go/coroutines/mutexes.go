package main

import (
	"fmt"
	"sync"
	"time"
)

// sync.Mutex
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc increments the counter for the given key
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v
	c.v[key]++
	c.mu.Unlock()
}

// Value returns the current value of the counter for the given key
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v
	defer c.mu.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		val := "somekey" + " " + fmt.Sprintf("%v", i%10)

		go c.Inc(val)
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey 1"))
	fmt.Println(c.Value("somekey 3"))
	fmt.Println(c.Value("somekey 6"))
	fmt.Println(c.Value("somekey 8"))
	fmt.Println(c.Value("somekey 9"))
}
