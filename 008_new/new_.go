package main

import "fmt"

func main() {
	//1 new一个对象a,
	// 2 a对象会初始化分配1个指定类型的内存空间，
	// 3 并且对象a指向的内存地址保存的值为对应类型默认值，例如int-0，string-空字符串
	a:=new (int)
	fmt.Println(*a,a)

	b:=new(string)
	fmt.Println(*b,b)
}
