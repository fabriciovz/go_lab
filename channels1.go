package go_lab

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func MainChain3() {
	dataChan := make(chan int)

	go func() {
		for i := 0; i < 1000; i++ {
			dataChan <- i

		}
		close(dataChan)
	}()

	for n := range dataChan {
		fmt.Printf("n =%d\n", n)
	}
}

func DoWork() int {
	time.Sleep(time.Second)
	return rand.Intn(100)
}

func MainChain4() {
	dataChan := make(chan int)

	go func() {
		wg := sync.WaitGroup{}
		for i := 0; i < 1000; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				result := DoWork()
				dataChan <- result
			}()
		}
		wg.Wait()
		close(dataChan)
	}()

	for n := range dataChan {
		fmt.Printf("n =%d\n", n)
	}
}

func Worker(done chan bool) {
	fmt.Println("Working...")
	time.Sleep(time.Second)
	fmt.Println("Done")

	done <- true
}

func MainChain5() {
	done := make(chan bool, 1)
	go Worker(done)

	<-done
}
