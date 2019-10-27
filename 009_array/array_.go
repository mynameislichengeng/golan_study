package main

import "fmt"

func main() {
	////1 方式1
	//var a [3]int=[3]int{1,2,3}
	//fmt.Println(a)

	////2 方式2
	//b:=[3]int{5,6,7}
	//fmt.Println(b)
	////如果设置的数组的长度是7，但是初始化只指定了3个，那么没有指定的，初始化，就使用对应类型的默认值
	//c:=[7]int{5,6,7}
	//fmt.Println(c)
	//
	////3 方式3,{}里面给定了多少个值，那么就设定这个数组的长度为多少
	//d:=[...]int{11,12,13,14,15}
	//fmt.Println(d)



	//注意重点
	//go 语言中的数组，是值类型，而不是像java中的引用类型
	arr1:=[3]int{1,2,3};
	//这里相当于,arr2重新分配了1个内存地址，然后赋值arr1数组的值
	arr2:=arr1;
	fmt.Println(arr1,arr2)
	//arr1和arr2不是同1个地址,因为 arr2是重新分配的内存地址
	fmt.Printf("%p %p\n",&arr1,&arr2)

	//arr2的值变化，不会引起arr1的值变化
	arr2[0] = 10
	fmt.Println(arr1,arr2)

}
