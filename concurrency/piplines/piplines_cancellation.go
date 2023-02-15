package piplines

import (
	"fmt"
	"sync"
)

func Run() {

	//done := make(chan struct{}, 2)
	done := make(chan struct{})
	defer close(done)

	in := gen(2, 2)
	in2 := gen(4, 4)

	// Distribute the sq work across two goroutines that both read from in.
	c1 := sq(done, in)
	c2 := sq(done, in2)
	c3 := sq(done, in2)
	//c3 := sq(in)
	//c4 := sq(in)

	// Consume the merged output from c1 and c2.
	for n := range merge(done, c1, c2, c3) {
		fmt.Println(n) // 4 then 9, or 9 then 4
	}

	// Tell the remaining senders we're leaving.
	//done <- struct{}{}
	//done <- struct{}{}

}

func gen(nums ...int) <-chan int {
	//out := make(chan int)
	//go func() {
	//	for _, n := range nums {
	//		out <- n
	//	}
	//	close(out)
	//}()
	//return out

	//BEST

	out := make(chan int, len(nums))
	for _, n := range nums {
		out <- n
	}
	close(out)
	return out
}

func sq(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case out <- n * n:
			case <-done:
				return
			}
		}
	}()
	return out
}

func merge(done <-chan struct{}, cs ...<-chan int) <-chan int {
	out := make(chan int, 1)
	var wg sync.WaitGroup

	output := func(c <-chan int) {
		defer wg.Done()
		for n := range c {
			select {
			case out <- n:
			case <-done:
				return
			}
		}
	}

	wg.Add(len(cs))

	for _, c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
