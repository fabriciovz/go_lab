package go_lab

import (
	"fmt"
	"sync"
)

func MainChain30() {
	var wait sync.WaitGroup
	wait.Add(1)
	go func() {
		fmt.Println("Hello World!")
		wait.Done()
	}()
	wait.Wait()
}

func MainChain31() {
	var wait sync.WaitGroup
	wait.Add(1)
	go func() {
		fmt.Println("Hello World!")
		wait.Add(-1)
	}()
	wait.Wait()
}

func MainChain32() {
	var wait sync.WaitGroup
	goRoutines := 5
	wait.Add(goRoutines)
	for i := 0; i < goRoutines; i++ {
		go func(goRoutineID int) {
			fmt.Printf("ID:%d: Hello goroutines!\n", goRoutineID)
			wait.Done()
		}(i)
	}
	wait.Wait()
}
