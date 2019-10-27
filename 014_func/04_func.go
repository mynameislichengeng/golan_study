package main

import "fmt"

func main() {
	demo4()
}

//匿名函数
//1 匿名函数就是没有函数名的函数,形如func(){}
//2 匿名函数必须在函数内部声明，比如下面的func(){}必须在函数demo4(){}中声明
//3 匿名函数声明后，就必须立即调用，调用的方式就是在匿名函数结尾()
func demo4() {

	//1 无参数无返回值匿名函数
	func() {
		fmt.Println("happy birthday")
	}()

	//2 有参数无返回值匿名函数
	func(sprot string) {
		fmt.Println(sprot)
	}("play basketball")

	//3 有参数有返回值匿名函数
	k := func(sprot string) (age int) {
		age = 20
		return
	}("play basketball")
	fmt.Println(k)
}
