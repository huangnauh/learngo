package p

import (
	"fmt"
	"unsafe"
)

type V struct {
	i int32
	j int64
}

func (this V) PutI() {
	fmt.Printf("i=%d\n", this.i)
	fmt.Printf("i=%d\n", unsafe.Offsetof(this.j))

}

func (this V) PutJ() {
	fmt.Printf("J=%d\n", this.j)
}