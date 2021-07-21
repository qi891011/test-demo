package main

import "fmt"

func main() {
	//testMap 是声明的字典变量名，string 是键的类型，int 则是其中所存放的值类型
	testMap := map[string]int{
		"one":1,
		"two":2,
		"three":3,
	}
	fmt.Println("testMap:",testMap)

	testMap1 := make(map[string]int)
	fmt.Println("testMap1:",testMap1)
	testMap1["one"] = 1
	testMap1["two"] = 2
	testMap1["three"] = 3
	testMap1["four"] = 4
	fmt.Println("testMap1添加键值对:",testMap1)

	//var testMap2 map[string]int  //声明map
	//testMap2["ss"] = 6
	//fmt.Println(testMap2) //报错panic，不能只声明不初始化就赋值，字典初始化之后才能进行赋值操作


	//查找元素
	val, ok := testMap["one"]
	if ok{
		fmt.Println(val)
	}
	//删除元素
	delete(testMap, "three") //键不存在也不会有报错
	fmt.Println("删除元素结果",testMap)

	//遍历字典（键值都要）
	for k, v := range testMap {
		fmt.Println("健",k, "值",v)
	}
	//遍历字典（只要键）
	for k := range testMap {
		fmt.Println("键",k)
	}
	//遍历字典（只要值）
	for _, i2 := range testMap {
		fmt.Println("值",i2)
	}

	//键值对调
	invMap := make(map[int]string)
	fmt.Println("invMap:",invMap)
	for k, v := range testMap {
		invMap[v] = k  //testMap 的值v作为 invMap 的健,testMap 的k作为 invMap 的值
	}
	fmt.Println(invMap)
	for i, s := range invMap {
		fmt.Println(i, s)
	}
}
