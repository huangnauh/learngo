package main

import (
	"crypto/sha256"
	"fmt"
	"bytes"
)

func reverse(s []int) {
	for i, j := 0, len(s) - 1; i < j; i, j = i + 1, j - 1 {
		s[i], s[j] = s[j], s[i]
	}
}

func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i, v := range x {
		if v != y[i] {
			return false
		}
	}
	return true
}

func equal1(x,y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}

//输入的 slice 和输出的 slice 共享一个底层数组
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func nonempty2(strings []string) []string {
	out := strings[:0]
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func main() {
	//stringer.NewStringer()
	//
	fmt.Println(equal1(map[string]int{"A":0}, map[string]int{"B":0}))
	var f []int
	f = append(f, 10)
	fmt.Println(f)
	var ages map[string]int
	fmt.Println(len(ages))
	//ages["a"] = 10  panic: assignment to entry in nil map
	ages = map[string]int{}
	ages["a"] = 10
	fmt.Println(ages)
	e := []string{"1", "", "2"}
	fmt.Printf("%v\n", nonempty2(e))
	fmt.Printf("%v\n", e)
	d2 := [...]string{"123", "456"}
	d1 := [...]string{"123", "456"}
	d3 := d2[0:1]
	d4 := d1[0:1]
	fmt.Printf("%p %p %v\n", &d3, &d4, equal(d3, d4))
	c := [...]byte{1, 2, 3, 4}
	cc := [...]byte{1, 2, 3, 4}
	c1 := c[0:1]
	c2 := cc[0:1]
	fmt.Printf("%p %p %v\n", &c1, &c2, bytes.Equal(c1, c2))
	b1 := sha256.Sum256([]byte("x"))
	b2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", b1, b2, b1 == b2, b1)

	a := [...]int{1, 2, 3}
	reverse(a[:])
	fmt.Println(":", a)
	const (
		USD int = iota
		EUR
	)
	symbol := [...]string{USD: "$", EUR: "€"}
	fmt.Println(EUR, symbol[EUR])
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}
}
