package main

import "fmt"

/*
 函数,可以作为函数的形参
 */
func main() {
	myDo(func(name string) {
		fmt.Println(name)
	})
}


func myDo(args func(name string)){
	fmt.Println("执行myDo函数")
	args("smaling")
}