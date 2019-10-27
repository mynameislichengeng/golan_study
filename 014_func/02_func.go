package main

import "fmt"

func main() {
	//可以不用变量 接收，有返回值的函数
	//show()
	//如果接收有返回值的函数,接收变量必须与函数一一对应
	name, age, address := show()
	fmt.Println(name, age, address)
	fmt.Printf("%p,%p,%p\n", &name, &age, &address)

	//如果不想接收某个值，则可以使用占位符来代替
	name1, age1, _ := show1 ()
	fmt.Println(name1, age1)
	fmt.Printf("%p,%p\n", &name1, &age1)
}

//多返回值函数
//func 函数名([参数名 参数类型]) (返回值1的类型，返回值2的类型,返回值3的类型（以此类推）){}
func show() (string, int, string) {
	fmt.Println("执行show()")
	return "lc", 30, "wuhan"
}

func show1() (name string, age int,address string) {
	fmt.Println("执行show()")
	name= "lc"
	age= 18
	address="wh"
	return
}