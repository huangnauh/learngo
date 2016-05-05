package main

import (
	"time"
	"fmt"
)

func main() {
	ticker := time.NewTicker(time.Second)
	go func(){
		for t := range ticker.C {
			fmt.Println(t)
		}
	}()

	time := time.NewTimer(10 * time.Second)
	<- time.C
	ticker.Stop()
	fmt.Println("expired")
}
