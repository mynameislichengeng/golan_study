package main

import "fmt"

func main() {
	//array是值类型，声明之后，如果没有赋值，那么系统会为数组中的每个元素赋值，该值是对应类型的默认值，如int=0 string=''
	var a [3]int
	fmt.Println(a)
}
