package main

import "fmt"

func main()  {
	arr1:=[2]string{"1","samling"}
	fmt.Println(arr1)

	//第1种，遍历循环
	//for i:=0;i<len(arr1);i++{
	//	fmt.Println(arr1[i])
	//}

	//第2种，遍历循环
	//for j,n:= range arr1{
	//	fmt.Println(j,n)
	//}
	//第3种，遍历循环
	//_ 下划线表示的是，一个占位符
	for _,n:= range arr1{
		fmt.Println(n)
	}
}
