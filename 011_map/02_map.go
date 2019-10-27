package main

import "fmt"


func main() {

	m:=map[string]string{"name":"lc","age":"19"}
	fmt.Println(m)

	//添加
	//1当key不存在时，表示新增
	//m["school"]="college"
	//fmt.Println(m)
	////2当key存在时，表示覆盖
	//m["name"]="kg"
	//fmt.Println(m)

	//删除
	//3 当key存在时，正常删除
	//delete(m,"name")
	//fmt.Println(m)
	////4 当key不存在时，不会有任何报错
	//delete(m,"meimei")
	//fmt.Println(m)

	// 查找
	//5 当key存在时，就正常返回key对应的value
	//fmt.Println(m["name"])
	//6 当key不存在时，就返回该value对应类型的默认值，比如int 返回0,string返回''
	//fmt.Println(m["meimei"])
	//fmt.Println("hello")
	//多返回值
	//第1个值，表示key对应的value值
	//第2个值，表示该key是否存在，boolean类型
	//_value,ok := m["name"]
	//fmt.Println(_value,ok)
	//_value1,ok1 := m["meimei"]
	//fmt.Println(_value1,ok1)

	//循环遍历
	for key,value :=range m{
		fmt.Println(key,value)
	}
	for _,value :=range m{
		fmt.Println(value)
	}
}
