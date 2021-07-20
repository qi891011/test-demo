package main

import (
	"fmt"
	"strconv"
)

func main() {
	v1 := "100"
	v2, _ := strconv.Atoi(v1)  // 将字符串转化为整型，v2 = 100
	fmt.Println(v2)
	v3 := 100
	v4 := strconv.Itoa(v3)   // 将整型转化为字符串, v4 = "100"
	fmt.Println(v4)
	v5 := "true"
	v6, _ := strconv.ParseBool(v5)  // 将字符串转化为布尔型
	v5 = strconv.FormatBool(v6)  // 将布尔值转化为字符串
	fmt.Println(v5)
	fmt.Println(v6)
	v7 := "100"
	v8, _ := strconv.ParseInt(v7, 10, 64)   // 将字符串转化为整型，第二个参数表示进制，第三个参数表示最大位数
	v7 = strconv.FormatInt(v8, 10)   // 将整型转化为字符串，第二个参数表示进制
	fmt.Println("v7:", v7)
	fmt.Println("v8:", v8)
}
