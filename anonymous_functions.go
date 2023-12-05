package go_lab

import "time"

func MainChain10() {
	go func() {
		println("Hello World")
	}()
	time.Sleep(time.Second)
}

func MainChain11() {
	go func(msg string) {
		println(msg)
	}("Hello World")
	time.Sleep(time.Second)
}

func MainChain12() {
	messagePrinter := func(msg string) {
		println(msg)
	}
	go messagePrinter("Hello World")
	go messagePrinter("Hello goroutine")
	time.Sleep(time.Second)
}
