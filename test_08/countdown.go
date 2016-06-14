package main

import (
	"time"
	"fmt"
	"os"
)

func test1(){
	tick := time.Tick(time.Second)
	for countdown := 10; countdown >0; countdown--{
		fmt.Println(countdown)
		<-tick
	}
	fmt.Println("ok")
}

func test2() {

	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte,1))
		abort <- struct {}{}
	}()

	select {
	case <-time.After(10 * time.Second):
		//
	case <-abort:
		fmt.Println("aborted")
		return
	}
	fmt.Println("ok")
}

func main() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte,1))
		abort <- struct {}{}
	}()

	//tick := time.Tick(time.Second)
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for countdown := 10; countdown >0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-ticker.C:
			//
		case <-abort:
			fmt.Println("aborted")
			return
		}
	}
	fmt.Println("ok")
}