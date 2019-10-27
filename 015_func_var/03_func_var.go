package main

import "fmt"

/**
 函数变量，作为返回值
 */
func main() {

	fmt.Println("excute")
	result := a();
	r2 := result();
	fmt.Println(r2)

}

func a() func() int {
	return func() int {
		return 10
	}
}
