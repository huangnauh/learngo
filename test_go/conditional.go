package main

func If(condition bool, trueval, falseval interface{}) interface{} {
	if condition {
		return trueval
	}
	return falseval
}

func main() {
	a, b := 1,2
	max := If(a>b,a,b).(int)
	println(max)
}
