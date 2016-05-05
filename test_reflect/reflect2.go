package main

import (
	"reflect"
	"fmt"
)

type Data struct {

}

func (*Data) String() string {
	return ""
}

func main() {
	var d *Data
	t := reflect.TypeOf(d)

	it := reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
	fmt.Println(t.Implements(it))
}
