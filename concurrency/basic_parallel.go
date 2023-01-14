package concurrency

import (
	"fmt"
	"math/rand"
	"time"
)

type Result string
type Search func(query string) Result

var (
	Web1   = fakeSearch("Web1")
	Web2   = fakeSearch("Web2")
	Image1 = fakeSearch("Image1")
	Image2 = fakeSearch("Image2")
	Video1 = fakeSearch("Video1")
	Video2 = fakeSearch("Video2")
)

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}

func First(query string, replicas ...Search) Result {
	c := make(chan Result)
	for i := range replicas {
		go func(idx int) {
			c <- replicas[idx](query)
		}(i)
	}
	// the magic is here. First function always waits for 1 time after receiving the result
	return <-c
}

func Google(query string) []Result {
	c := make(chan Result)

	//each search platforms in a goroutine
	go func() {
		c <- First(query, Web1, Web2)
	}()

	go func() {
		c <- First(query, Image1, Image2)
	}()

	go func() {
		c <- First(query, Video1, Video2)
	}()

	var result []Result
	// the global timeout for 3 queries
	// it means after 50ms, it ignores the result from the server that taking response greater than 50ms

	timeout := time.After(50 * time.Millisecond)
	for i := 0; i < 3; i++ {
		select {
		case r := <-c:
			result = append(result, r)
		case <-timeout:
			fmt.Println("timeout")
			return result
		}
	}

	return result
}

func Run() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	result := Google("golang")
	elapsed := time.Since(start)
	fmt.Println(result)
	fmt.Println(elapsed)
}
