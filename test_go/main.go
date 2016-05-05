package main

import (
	"time"
	"fmt"
)

var c chan int

func ready(w string, sec int){
	time.Sleep(time.Duration(sec) * time.Second)
	fmt.Println(w, "is ready!")
	c <- 1
}

func main() {
	c = make(chan int)
	i := 0
	go ready("Tea", 2)
	go ready("Coffee", 1)
	fmt.Println("I'm waiting")
L:	for {
		select {
		case <- c:
			i++
			if i == 2 {
				break L
			}

		}
	}
}


