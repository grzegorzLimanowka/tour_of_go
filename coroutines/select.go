package coroutines

import "fmt"

// select statement lets a goroutine wait on mulptiple communication operations

// A select blocks until one of its cases can run, then it executes that case. It chooses at random if multiple are ready

func fibonaci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("Quit")
			return
		}

	}

}

func main() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	fibonaci(c, quit)
}
