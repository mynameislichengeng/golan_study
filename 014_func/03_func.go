package main

import "fmt"

func main() {
	//show11("lc", "kg")
	show22(19,"lc", "kg")
}

func show11(names ...string) {
	for i, name := range names {
		fmt.Println(i, name)
	}
}

//对于这种既有固定参数，又有可变参数的，必须把可变参数放在最后
func show22(age int,names ...string) {
	fmt.Println(age)
	for i, name := range names {
		fmt.Println(i, name)
	}
}