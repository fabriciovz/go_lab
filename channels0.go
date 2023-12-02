package go_lab

import (
	"fmt"
)

func MainChain0() {
	dataChan := make(chan int, 1)

	dataChan <- 789 //Add data to the channel

	n := <-dataChan //Get data to the channel
	fmt.Printf("n =%d\n", n)
}

func MainChain1() {
	dataChan := make(chan int, 2)

	dataChan <- 789 //Add data to the channel
	dataChan <- 123 //Add data to the channel

	n := <-dataChan //Get data to the channel
	fmt.Printf("n =%d\n", n)

	n = <-dataChan //Get data to the channel
	fmt.Printf("n =%d\n", n)
}

func MainChain2() {
	dataChan := make(chan int)

	go func() {
		dataChan <- 789 //Add data to the channel
	}()

	n := <-dataChan //Get data to the channel
	fmt.Printf("n =%d\n", n)
}
