package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//字符串切片,左含右不含(左闭右开)
	str := "Hello, world"
	str1 := str[:5] //获取索引5（不含）之前的子串
	str2 := str[7:] //获取索引7（含）之后的子串
	str3 := str[0:5] //获取从索引0（含）到索引6（不含）之间的子串
	str4 := str[:]
	//字符串函数
	str5 := strings.Replace(str4, "o", "k", -1) //func Replace(s, old, new string, n int) string 在s字符串中，把old字符串替换为new字符串，n表示替换的次数，小于0表示全部替换
	fmt.Println("str1:", str1)
	fmt.Println("str2:", str2)
	fmt.Println("str3:", str3)
	fmt.Println("str4:", str4)
	fmt.Println("str5:", str5)
	//字符串遍历
	str_ := "Hello, niho"
	n := len(str_)
	for i := 0; i < n; i++ {
		ch := str_[i]
		fmt.Println(i, ch)
		fmt.Println(i, string(ch))
	}
	//for true{
	//	fmt.Printf("这是无限循环。\n")
	//}
	v1 := 65
	fmt.Println(string(v1))
	v2 := 30028
	fmt.Println(string(v2))
	v3 := "100"
	v4, _ := strconv.Atoi(v3)
	fmt.Println(v4)
	v5 := 100
	v6 := strconv.Itoa(v5)
	fmt.Println(v6)
}
