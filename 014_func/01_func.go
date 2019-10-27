package main

import "fmt"

func main() {
	//1 无参数无返回值调用
	//demo();

	//2 有参数无返回值调用
	//exam("lc",20)

	//3 有参数有返回值调用
	z:=sample1(1, 3)
	fmt.Printf("%d",z);
}


//1 无参数无返回值函数
//func 函数名(){}
func demo() {
	fmt.Println("执行无参数无返回值函数");
}

//2 有参数无返回值函数，
// 这里要注意，一般声明函数中的参数，叫做形参，调用函数中的参数，叫做实参
//func 函数名(参数名 参数类型){}
func exam(name string,age int)  {
	fmt.Printf("我的名字:%s,年龄:%d",name,age)
}

//3 有参数有返回值函数
//func 函数名(参数名 参数类型) 返回值类型
func sample(a,c int) int{
	return a+c;
}
//这是 有参数有返回值函数的另一种写法，需要注意
func sample1(a,c int) (s int){
	s =a+c
	return
}
