package main

import (
	"sync"
	"fmt"
	"runtime"
	"math"
	"time"
)

var counter int = 0

func count(lock *sync.Mutex){
	lock.Lock()
	counter++
	fmt.Println(counter)
	lock.Unlock()
}

func sum(id int) {
	var x int64
	for i:=0; i < math.MaxUint32; i++ {
		x += int64(i)
	}

	println(id, x)
}

func main() {
	lock := &sync.Mutex{}

	for i := 0;i < 10;i++{
		go count(lock)
	}
	for {
		lock.Lock()
		c := counter
		lock.Unlock()
		runtime.Gosched()
		if c >= 10 {
			break
		}
	}
	fmt.Println("again")
	wg := new(sync.WaitGroup)
	//wg.Add(2)
	//
	//for i:=0; i < 2;i++ {
	//	go func(id int) {
	//		defer wg.Done()
	//		sum(id)
	//	}(i)
	//}
	//
	//wg.Wait()

	wg.Add(2)

	go func() {
		defer wg.Done()

		for i := 0; i < 6; i++ {
			println(i)
			if i == 3 {
				time.Sleep(time.Second)
				runtime.Gosched()
			}
		}
	}()

	go func() {
		defer wg.Done()
		time.Sleep(2 * time.Second)
		println("haha")
	}()

	wg.Wait()

}
