package main
import "fmt"
import (
	"sync"
	"time"
)

func gen(done <-chan struct{}, nums ... int) <-chan int {
	out := make(chan int, len(nums))

	go func() {
		defer close(out)
		for k, n := range nums {
			select{
			case out <- n:
				fmt.Println("gen", k)
			case <-done:
				fmt.Println("gen done", k)
				return
			}
		}
	}()

	return out
}

func sq(done <-chan struct{}, in <-chan int) <-chan int{
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case out <- n *n:
				fmt.Println("sq", n*n)
			case <-done:
				return
			}

		}
	}()
	return out
}

func merge(done <-chan struct{}, cs ... <-chan int) <-chan int{
	var wg sync.WaitGroup //在调用close之前确认所有的发送动作都执行完毕
	out := make(chan int)

	output := func(k int, c <-chan int){
		defer wg.Done()
		for n := range(c){
			select {
			case out <-n:
				fmt.Println("merge", k, n)
			case <-done:
				fmt.Println("done")
				return
			}
		}
	}

	wg.Add(len(cs))
	for k, c := range cs {
		go output(k, c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	done := make(chan struct{})
	defer close(done)

	c := gen(done, 2, 3)


	o1 := sq(done, c)
	o2 := sq(done, c)

	n := merge(done,o1,o2)
	fmt.Println("main", <-n)
	time.Sleep(time.Second)
}
