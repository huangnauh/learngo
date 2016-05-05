package main

import (
	"fmt"
	//"learngo/mymath"
)

type User struct {
	id int
	name string
}

type Stringer interface {
	String() string
}

type Printer interface {
	Stringer
	Print()
}

type Tester struct {
	s interface{
		String() string
	  }
}

func (self *User) TestPoint() {
	fmt.Printf("Point %p %v\n",self,self)
}

func (self User) TestValue() {
	fmt.Printf("Value %p %v\n",&self,self)
}

func (self *User) String() string {
	return fmt.Sprintf("User id: %d name: %s", self.id, self.name)
}


func (self *User) Print() {
	fmt.Println(self.String())
}

func main() {
	//fmt.Printf("hello,world. Sqrt(2) = %v\n", mymath.Sqrt(2))
	u := User{1, "huang"}
	fmt.Printf("User %p %v\n", &u, u)
	u.TestValue()
	u.TestPoint()
	(&u).TestPoint()
	(&u).TestValue()
	var _ fmt.Stringer = (*User)(nil)
	var t Printer = &User{1, "tom"}
	t.Print()
	x := Tester{&User{2, "huang"}}
	fmt.Println(x.s.String())
	var vi, pi interface{} = u, &u
	fmt.Println(vi.(User).name)
	fmt.Println(pi.(*User).name)

	if i, ok := pi.(fmt.Stringer); ok {
		fmt.Println(ok, " abc:", i)
	}

	switch v := pi.(type) {
	case nil:
		fmt.Println("nil")
	case func() string:
		fmt.Println(v())
	case *User:
		fmt.Println(v.name)
	case fmt.Stringer:
		fmt.Println(1,v)
	default:
		fmt.Println("unkown")
	}
}
