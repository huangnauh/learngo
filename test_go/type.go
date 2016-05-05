package main

import (
	"fmt"
	"unsafe"
	"bytes"
	"encoding/gob"
)

func deep_copy(dsc, src interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil{
		return err
	}
	fmt.Println(buf)
	return gob.NewDecoder(&buf).Decode(dsc)
	//return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dsc)
}

func test() {
	var str string = "huang"
	strp := (*struct {str uintptr
					len int})(unsafe.Pointer(&str))
	fmt.Printf("str: %+v\n",strp)

	var slice []int = make([]int,5,10)
	p := (*struct {
			array uintptr
			len int
			cap int})(unsafe.Pointer(&slice))
	fmt.Printf("slice: %+v\n",p)

	var array = [...]int{1,2,3,4,5}
	var sli = array[0:3]
	fmt.Printf("改变slice之前: array=%+v, slice=%+v\n", array, sli)
	var sl1 = append(sli,1)
	fmt.Printf("改变slice之前: array=%p, slice=%p, sl1=%p\n", &array, &sli,&sl1)
	var sl = append(sli, 6,7,8)
	fmt.Printf("改变slice之后: array=%+v, slice=%+v, sl=%+v\n", array, sli, sl)

	var m = make(map[string]int32, 10)
    m["hello"] = 123
    pp := (*struct {
        count      int
        flags      uint32
        hash0      uint32
        B          uint8
        keysize    uint8
        valuesize  uint8
        bucketsize uint16

        buckets    uintptr
        oldbuckets uintptr
        nevacuate  uintptr
    })(unsafe.Pointer(&m))

    fmt.Printf("output map: %+v\n", pp)
}

type Stringer interface {
	String() string
}

type E struct {
	X,Y,Z int
	Name interface {
			String() string}
}

type F struct {
	X,Y,Z int
	Name T
}

type T struct {
	Test int
}

func (this *T) String() string {
	return fmt.Sprintf("test:%d",this.Test)
}

func main() {
	a := [...]int{1,2,3,4,5}
	b := [5]int{}
	fmt.Println(b)
	deep_copy(&b, a)
	fmt.Printf("%v %p %p\n",b, &a,&b)

	var network bytes.Buffer
	enc := gob.NewEncoder(&network)
	dec := gob.NewDecoder(&network)

	gob.Register(T{})

	t := T{1}

	err := enc.Encode(E{1,2,3, &t})
	if err != nil {
		fmt.Println(err)
	}
	var x F
	err = dec.Decode(&x)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v: {%d,%d}\n", x.Name, x.X, x.Y)
}