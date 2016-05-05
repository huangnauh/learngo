package main

import (
	"reflect"
	"fmt"
)

func main() {
	var x float32 = 3.4
	//t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)
	fmt.Println("type: ", v.Kind() == reflect.Float32, "value: ", v.Float())

	p := reflect.ValueOf(&x)
	fmt.Println(p.Type())
	fmt.Println(p.CanSet())

	pv := p.Elem()
	fmt.Println(pv.CanSet())
	pv.SetFloat(7.1)
	fmt.Println(pv.Interface())
	fmt.Println(x)

	type T struct {
		A int
		B string
	}

	t := T{100, "huang"}
	s := reflect.ValueOf(&t).Elem()
	typeoft := s.Type()
	for i :=0;i< s.NumField();i++{
		f := s.Field(i)
		fmt.Println(i, typeoft.Field(i).Name,
		f.Type(),f.Interface())
	}

}
