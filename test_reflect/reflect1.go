package main

import (
	"reflect"
	"fmt"
)

type User struct {
	Username string `field:"username" type:"nvarchar(20)"`
	age int  `field:"age" type:"tiny int"`
}

type Admin struct {
	User
	title string
}

func (*User) String(){

}

func (Admin) test() {}


func main() {
	u := new(Admin)
	t := reflect.TypeOf(u)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	f, _ := t.FieldByName("User")
	fmt.Println(f.Name)
	f, _ = t.FieldByName("Username")
	fmt.Println(f.Tag)
	fmt.Println(f.Tag.Get("field"))
	fmt.Println(f.Tag.Get("type"))
	f = t.FieldByIndex([]int{0,1})
	fmt.Println("2", f.Name,f.Type)
	for i, n :=0, t.NumField();i < n;i++{
		f := t.Field(i)
		fmt.Println(f.Name, f.Type)
	}

	var uu Admin
	var x reflect.Type = reflect.TypeOf(&uu)
	for i, n := 0, x.NumMethod(); i<n; i++{
		fmt.Println(i,n)
		m := x.Method(i)
		fmt.Println(m.Name)
	}

	c := reflect.ChanOf(reflect.SendDir, reflect.TypeOf(""))
	fmt.Println(c.Elem())

	m := reflect.MapOf(reflect.TypeOf(""), reflect.TypeOf(0))
	fmt.Println(m)

	s := reflect.SliceOf(reflect.TypeOf(0))
	fmt.Println(s)

	tt := struct {Name string}{}
	p := reflect.PtrTo(reflect.TypeOf(tt))
	fmt.Println(p)


}
