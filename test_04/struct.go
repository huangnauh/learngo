package main

import (
	"fmt"
)

type Emp struct {
	id int
	name string
}

type Mana struct {
	Emp
	age int
}

func EmpById(emps []Emp, id int) *Emp {
	return &emps[id]
}

func main() {
	emps := []Emp{
		1: Emp{1, "huang"},
	}
	EmpById(emps,1).name = "j"
	fmt.Println(EmpById(emps,1))

	seen := make(map[string]struct{})
	// 模拟 set 数据结构,强调 key 的重要性 struct{} 大小为0
	if _, ok := seen["1"]; !ok {
		seen["1"] = struct{}{}
	}
	fmt.Println(seen)

}
