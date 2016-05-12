package main

import (
	"time"
	"log"
	"fmt"
	"os"
	"runtime"
)

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}

func main() {
	defer trace("main")()
	defer printStack()
	time.Sleep(2 * time.Second)
	f(3)
}

func f(x int) {
	if x == 0 {
		return
	}
	fmt.Printf("f(%d)\n", x + 0 / x)
	defer fmt.Println("defer", x)
	f(x - 1)
}

func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stderr.Write(buf[:n])
}
