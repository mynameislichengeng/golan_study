package main

import "fmt"

/**
函数变量，作为正常的变量
 */
/**
1 函数类型的变量，简称函数变量
2 函数变量声明后，系统不会马上为其分配内存空间
3 可以为函数变量赋值,赋值可以为如下:
     ---匿名函数
     ---普通函数
4 函数变量的使用，跟正常函数一样，只不过，函数名就是变量名
5 函数是引用类型
*/
func main() {
	a_2()
}


func a_1(){
	fmt.Println("")
	var a func()
	fmt.Println(a)
	a = func(){
		fmt.Println("我是一个函数类型的变量")
	}
	a()
}

func a_2(){
	fmt.Println("")
	var a func()
	fmt.Println(a)
	a = func(){
		fmt.Println("我是一个函数类型的变量2")
	}
	a()
	a=a_1
	a()

}


