package main

import "fmt"

func main() {
	//1标准for循环
	//for i:=0;i<5;i++{
	//	fmt.Println(i)
	//}
	//2 死循环
	//for ; ;  {
	//	fmt.Println("执行")
	//}

	//3 省略--for循环
	i:=0
	for ;i<5;i++{
		fmt.Println(i)
	}

	j:=0
	for j<5{
		fmt.Println(j)
		j++
	}
}
