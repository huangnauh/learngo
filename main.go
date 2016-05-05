package main

import (
	"errors"
	"fmt"
	"unsafe"
)

var (
	is_active       bool          // 声明变量的一般方法
	enable, disable bool = true, false // 声明变量，并初始化
)

const (
	i  = 100
	x  = iota
	Pi = 3.14
	y
	z = iota
	w
)

type Color int

const (
	Black Color = iota
	Red
	Blue
)

func testColor(c Color){
	fmt.Println(c)
}

func test() {
	var available bool
	vaild := false   // 简短声明
	available = true // 常规赋值
	data, j := [3]int{0, 0, 0}, 1
	//j, data[j] = 0, 100
	data1 := make([]int, 4) // make 被编译器翻译成具体的创建函数,由其分配内存和初始化成员结构,返回对象
	data2 := new([]int)  // new 计算类型⼤大⼩小,为其分配零值内存,返回指针
	s := "abc"
	println("&s", &s)
	s, t := "hello", 20  //s重新赋值，地址不变
	s = `a
	b\r\n\x00
	c`
	s = "电脑app"
	us := []rune(s)
	us[2] = '哎'
	println(string(us))
	str := "abc"
	println("&s", &s,s)
	println("str",str[0]=='\x61', str[1]==0x62, str[2]==99)
	println(&s)
	{
		s := 100
		fmt.Printf("&s:%v t:%v\n", &s, t)  //定义新变量，不在同一层次代码块
	}
	var c complex64 = 5+5i
	var hello string = "hello"
	fmt.Println(vaild)
	fmt.Println(available)
	fmt.Println(hello)
	fmt.Printf("Value is:%v\n", c)
	fmt.Printf("data is:%v data1:%v, data2:%v i is:%v\n", data, data1,data2,j)
	fmt.Printf("const is:%v, %v, %v, %v, %v, %v\n", i,x,Pi,y,z,w)

	d := Black
	testColor(d)
	//e := 10   cannot use e (type int) as type Color in argument to testColor
	//testColor(e)
}

type person struct {
	name string
	age  int
}

type student struct {
	person
	id string
}

func Older(p1, p2 student) (student, int) {
	if p1.age > p2.age {
		return p1, p1.age - p2.age
	}
	return p2, p2.age - p1.age
}

func testin() *int {
	d := struct {
		s string
		x int
	}{"abc", 100}

	p := uintptr(unsafe.Pointer(&d))
	p += unsafe.Offsetof(d.x)
	p2 := unsafe.Pointer(p)
	px := (*int)(p2)
	*px = 200
	fmt.Printf("%v\n",d)
	tt := 100
	fmt.Println("tt:",tt)
	return &tt  //返回局部变量指针是安全的,编译器会根据需要将其分配在 GC Heap 上。
}

func test1(){
	type data struct {
		a int
	}
	pp := testin()
	fmt.Println(*pp)

	var d = data{1234}
	var p *data
	p = &d
	fmt.Printf("%v,%v\n", p, p.a)

	x := 0x98765432
	q := unsafe.Pointer(&x)
	n := (*[4]byte)(q)
	for i :=0;i < len(n); i++ {
		fmt.Printf("%v ", n[i])
	}



}

func main() {
	test1()
	tom := student{person{"Tom", 22}, "B06021730"}
	var jack student
	jack.name, jack.age, jack.id = "Jack", 25, "B06021731"
	old_person, age_diff := Older(tom, jack)
	fmt.Print(old_person, age_diff, "\n")
	const Pi float32 = 3.1415926
	fmt.Print(x, "\n")
	var arr [10]int
	arr[0] = 42
	//a := [...]int{10, 20, 30}
	d := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}}
	e := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
	fmt.Print("arr:", d[0], e[1])
	var numbers map[string]int
	numbers = make(map[string]int)
	fmt.Print("\nmap:", numbers["a"])
	aa, aaa := numbers["a"]
	fmt.Print("\naa:", aa, "\naaa:", aaa)
	numbers["one"] = 1
	if xx := 5; xx < 10 {
		fmt.Print("\nx is less then 10", xx)
	}
	// fmt.Print("\n", xx)
	s := "hello"
	s1 := `c
	dsds`
	fmt.Printf("%s\n", s1)
	c := []byte(s)
	c[0] = 'c'
	s2 := string(c)
	fmt.Printf("%s\n", s2)
	err := errors.New("huangnauh")
	if err != nil {
		fmt.Print(err)
	}
}
