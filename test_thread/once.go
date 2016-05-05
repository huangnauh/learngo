package main

import (
	"sync"
	"time"
)

var a string
var once sync.Once

func setup() {
	println("in once")
	a = "hello world"
}

func doPrint() {
	once.Do(setup)
	println(a)
}
func main() {
	go doPrint()
	go doPrint()
	time.Sleep(time.Second)
}
