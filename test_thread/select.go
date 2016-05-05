package main

import (
	"fmt"
	"os"
)

func main() {
	a, b := make(chan int, 3), make(chan int)

	go func(){
		v, ok, s := 0, false, ""

		for {
			select {
			case v, ok = <-a:
				s = "a"
			case v, ok = <-b:
				s = "b"
			}

			if ok {
				fmt.Println("out", s, v)
			} else {
				fmt.Println("end 1")
				os.Exit(0)
			}
		}
	}()

	for i := 0; i < 5; i++{
		select {
		case a <- i:
			fmt.Println("in a",i)
		case b <- i:
			fmt.Println("in b",i)
		}
	}

	close(a)
	fmt.Println("end 2")
	select {

	}
}
