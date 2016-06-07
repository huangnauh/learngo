package main

import "fmt"

func sqlQuote(x interface{}) string {
	// 类型断言
	if x == nil {
		return "NULL"
	} else if _, ok := x.(int); ok {
		return fmt.Sprintf("%d", x)
	} else if _, ok := x.(uint); ok {
		return fmt.Sprintf("%d", x)
	} else if b, ok := x.(bool); ok{
		if b {
			return "TRUE"
		} else {
			return "FALSE"
		}
	} else {
		panic("unexpected")
	}
}

func sqlQ(x interface{}) string {
	// 类型开关, 隐式创建了一个语言块
	print("type %T: %v", x, x)
	switch x := x.(type) {
	case nil:
		return "NULL"
	case int:
		return fmt.Sprintf("type %T: %v", x, x)
	case bool:
		if x {
			return "TRUE"
		} else {
			return "FALSE"
		}
	default:
		 panic(fmt.Sprintf("unexpected type %T: %v", x, x))
	}

}

func main() {
	print(sqlQ("1"))
}