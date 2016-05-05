package main

import (
	"point/p"
	"unsafe"
	"fmt"
)

type Stringer interface {
    ToString() string
}

type String struct {
	s interface{     // 匿名接⼝口
		ToString() string
	}
}

type Resource struct {
	id int
	name string
}
type Classify struct {
    id int
}

type File struct {
	name string
	size int
	attr struct {
		perm int
		owner int
	     }
}

type Foo struct {
	name string "username"
	size int
}

type User struct {
	Resource
	Classify
	name string  // 遮蔽 Resource.name
}

type Manager struct {
	User
	title string
}

func (self *User) ToString() string{    //匿名字段,可以像字段成员那样访问匿名字段方法,编译器负责查找。
	return fmt.Sprintf("User: %p,%v", self,self)
}

//func (self Manager) ToString() string{    //method value 会复制 receiver。
//	x  := fmt.Sprintf("manager: %p,%v", &self,self)
//	self.name = "hagha"
//	return x
//}

func main(){
	foo := Foo{name:"foo"}
	foo.size = 10
	fmt.Println(foo.name)
	sl1 := []int{0,1,2,8:100}
	fmt.Println(sl1)
	f := File{
		name: "test.txt",
		size: 100,
		attr: struct {perm int
				owner int}{2, 0555},
	}
	fmt.Println(f.attr.perm, f.attr.owner)
	f.attr.perm = 0755
	f.attr.owner = 1
	var ma Manager
	ma.Resource.id = 1
	ma.Classify.id = 10
	//可以像普通字段那样访问匿名字段成员
	//编译器从外向内逐级查找所有层次的匿名字段, 直到发现⺫⽬目标或出错。
	ma.name = "huang"
	ma.title = "admin"

	//stri := String{ma}
	//strin := Stringer(ma)
	//pstrin := Stringer(&ma)
	//ma.name = "yu"   //数据指针持有的是⺫⽬目标对象的只读复制品,复制完整对象或指针。
	////strin.(Manager).name = "abc"  cannot assign to strin.(Manager).User.name
	//pstrin.(*Manager).name = "Jack"
	//fmt.Println(stri.s.ToString())
	//fmt.Println(strin.ToString(),"&&&&",strin.(Manager))
	//fmt.Println(pstrin.ToString(),"&&&&",pstrin.(*Manager))

	fmt.Printf("%p\n", &ma)
	fmt.Println(ma.ToString())
	//内存布局和 C struct 相同,没有任何附加的 object 信息。
	println(unsafe.Sizeof(ma.Resource))
	println(unsafe.Offsetof(ma.Classify))
	println(unsafe.Sizeof(ma.User))
	println(unsafe.Offsetof(ma.Resource.id))
	println(unsafe.Offsetof(ma.Resource.name))
	println(unsafe.Sizeof(ma.Resource.name))
	println("Classify id:",(*Classify)(unsafe.Pointer(uintptr(unsafe.Pointer(&ma)) + unsafe.Offsetof(ma.Classify))).id)
	println(unsafe.Offsetof(ma.name))
	//type File struct {
	//	name string
	//	size int
	//	attr struct{
	//		perm int
	//		owner int
	//	     }
	//}
	//
	ff := File{
		name: "test.a",
	}
	ff.size = 100
	ff.attr.owner = 1
	ff.attr.perm = 0755
	println(ff.attr.perm,ff.size)
	for i :=0; i < 1; i++{
		m := map[int]string{
			0: "a", 1:  "a", 2:  "a", 3:  "a", 4:  "a",
			5: "a", 6:  "a", 7:  "a", 8:  "a", 9:  "a",
		}
		for k := range m {
			fmt.Println(k)
			delete(m, k+k)
			//在迭代时安全删除键值
			//如果先遍历到1，则存在4，如果先遍历到2，则存在8
		}
		fmt.Println(m)
	}

	m := map[string] struct {
		name string
		age int
	}{
		"huang": {"user1", 10},
		"li": {"user2", 20},
	}
	if v,ok := m["huang"]; ok{  //判断key是否存在
		println(ok, v.name)
	}
	//从 map 中取回的是一个 value 临时复制品,对其成员的修改是没有意义的。
	//当 map 因扩张⽽而重新哈希时,各键值项存储位置都会发⽣生改变。 因此,map 被设计成 not addressable。
	//m["huang"].name = "user4"   cannot assign to m["huang"].name
	println(m["123"].name)   // 对于不存在的 key,直接返回 \0,不会出错
	m["yu"] = struct {
		name string
		age int
	}{"user3", 40}
	println(len(m))
	for k,v := range m {    // 随机顺序返回,每次不相同。
		println(k,v.name,v.age)
	}
	s := []int{0,1,2,3,4,5,6,7,8,11:9}
	s1 := s[2:4]
	s2 := s1[2:5]
	s3 := s2[2:5]
	s4 := append(s3, 100,200)
	fmt.Println(s,s1,s2,s3,s4)
	for i := range s3 {
		fmt.Println(s3[i])
	}
	fmt.Printf("地址： %p %p %p", &s[6], &s3[0], &s4[0])
	data := [][]int{
		[]int{1,2,3},
		[]int{4,5,6},
		[]int{7,8,9},
	}
	st := [5]struct {
		x int
	}{}
	stp := st[:]
	st[1].x = 10
	stp[2].x = 20
	fmt.Printf("%v,%p,%p\n",st, &st,&st[0])
	q := &s[2]
	*q +=200
	fmt.Println(s,data)

	t1 := make([]int, 1, 5)
	t2 := append(t1, 1)
	fmt.Printf("t1 %v &t1 %p t2 %v &t2 %p ", t1, &t1, t2, &t2)

	var v *p.V = new(p.V)
	fmt.Printf("size=%d\n", unsafe.Sizeof(*v))
	var i *int32 = (*int32)(unsafe.Pointer(v))
	*i = int32(48)
	var j *int64 = (*int64)(unsafe.Pointer(uintptr(unsafe.Pointer(v))+8))   //注意字节对齐
	*j = int64(123)
	v.PutI()
	v.PutJ()
}


