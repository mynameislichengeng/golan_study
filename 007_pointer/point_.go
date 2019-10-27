package main

import "fmt"

func main()  {
	i:=1
	//fmt.Println(i,&i)
	//
	k:=i
	//fmt.Println(k,&k)
	//fmt.Printf("%T\n",k)
	//
	//c:=&k
	//fmt.Println(c)
	//fmt.Printf("%T\n",c)
	//fmt.Printf("%d",*c)

	//声明一个指针
	var d *int
	//给声明的指针赋值
	d = &k
	//方法1：修改指针指向内存地址的值
	//*d = 10
	//方法2：修改指针指向内存地址的值
	*&k = 12
	fmt.Println(d,&k)
	fmt.Println(*d,k,i)


	//声明一个指针时，是不会为他开辟一个内存地址的,会初始化它为nil
	var z *int
	fmt.Println(z)
	fmt.Println(z==nil)



}
