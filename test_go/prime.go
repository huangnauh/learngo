package main

import (
	"runtime"
	"runtime/debug"
	"fmt"
)

func generate() chan int{
	ch := make(chan int)
	go func() {
		for i:=2;;i++{
			ch <- i
		}
	}()
	return ch
}

func filter(in <-chan int, prime int) chan int{
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}

func main() {
	runtime.GOMAXPROCS(1)
	debug.SetMaxThreads(5)

	ch := generate()
	for i:=0; i<100;i++{
		prime := <-ch
		fmt.Printf("%v,%v\n", i+1, prime)
		ch = filter(ch, prime)
	}
}
