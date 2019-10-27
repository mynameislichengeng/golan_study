package main

import (
	"fmt"
	"strconv"
)

func main()  {
	s:="11"
	//第1个参数表示需要转换的字符串
	//第2个参数表示需要转换的字符串，是2进制度，还是8进制。。
	//第3个参数表示的是转换后的数是按 int=0,int8=8,int16=16,int32=32,int64=64
	k,_:=strconv.ParseInt(s,2,8)
	fmt.Println(k)
	fmt.Printf("%b",k)

}
