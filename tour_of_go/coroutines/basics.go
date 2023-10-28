package coroutines

import (
	"fmt"
	"time"
)

var x int = 0

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		x += 1

		fmt.Println(s)
	}
}

func calc_numbers(x []int, ch chan int) {
	sum := 0

	for _, v := range x {
		sum += v
	}

	ch <- sum // send
}

func main() {
	go say("world")
	go say("obama")

	for i := 0; i < 10; i++ {
		go say(fmt.Sprintf("loopz %v", i))
	}

	say("hello")

	nums := []int{2, 3, 5, 6, 2}
	n1 := nums[:len(nums)/2]
	n2 := nums[len(nums)/2:]

	ch := make(chan int)

	go calc_numbers(n1, ch)
	go calc_numbers(n2, ch)

	sum1, sum2 := <-ch, <-ch // recv

	fmt.Println(sum1 + sum2)
}
