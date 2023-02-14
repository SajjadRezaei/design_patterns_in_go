package concurrency

import (
	"fmt"
	"time"
)

func Up() {
	transferCh := make(chan struct{})

	go handler(transferCh)

	time.Sleep(5 * time.Second)
	transferCh <- struct{}{}

	time.Sleep(15 * time.Second)
	transferCh <- struct{}{}

	time.Sleep(5 * time.Second)
	transferCh <- struct{}{}

	time.Sleep(100 * time.Second)
}

func handler(ch chan struct{}) {
	for {
		select {
		case <-ch:
			startTransfer()
		default:
			scanBlock()
		}
	}
}

func startTransfer() {
	fmt.Println("start transfer")
	for i := 0; i < 10; i++ {
		fmt.Println(fmt.Sprintf("transferd request %v", i))
		time.Sleep(time.Second * 1)
	}
	fmt.Println("transfer end......")
}

func scanBlock() {
	for i := 0; i < 10; i++ {
		fmt.Println("scan block ", i)
		time.Sleep(1 * time.Second)
	}
}
