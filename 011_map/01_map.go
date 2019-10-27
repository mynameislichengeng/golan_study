package main

func main() {
	//1 声明map，不会分配内存地址
	//var m map[string]string
	//fmt.Println(m==nil)
	//fmt.Printf("%T: %p",m,m)

	//2 使用make来实例化一个没有初始值的map
	//var m map[string]string =make(map[string]string)
	//fmt.Println(m==nil)
	//fmt.Printf("%T: %p",m,m)

	//3 在map声明时，就给初始化赋值
	//var m map[string]string = map[string]string{"name":"lc","age":"18"}
	//fmt.Println(m)

}
