package main

import "fmt"

func main() {
	var a [8]byte // 长度为8的数组，每个元素为一个字节
	var b [3][3]int // 二维数组（9宫格）
	var c [3][3][3]float64 // 三维数组（立体的9宫格）
	var d = [3]int{1, 2, 35}  // 声明时初始化
	var e = [3]int{35}   // 通过 new 初始化
	var f = new([3]string) // 声明时初始化
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
	fmt.Println("d:", d)
	fmt.Println("e:", e)
	fmt.Println("f:", f)

	v := [...]int{1:2, 3:7} //[0 2 0 7]
	v1 := [5]int{1:3, 4:7} //[0 3 0 0 7]
	fmt.Println(v)
	fmt.Println(v1)
	fmt.Println(len(v))
	fmt.Println(len(v1))
	fmt.Println(len(v1)-1) //len(arr)-1 表示最后一个元素的下标
	a1, a2 := v[1], len(v1)-1
	fmt.Println(a1, a2)
	v[0] = 100 //通过下标设置数组v对应索引位置的元素值
	fmt.Println(v)

	//遍历数组
	for i := 0; i < len(v1); i++ {
		fmt.Println("数组下标为",i,"的值为", v1[i])
	}
	fmt.Println("---------------------")
	for i, i2 := range v1 {
		fmt.Println("数组下标为",i,"的值为", i2)
	}
	fmt.Println("//////////////////////")
	for _, i2 := range v1 {
		fmt.Println("不想要数组下标，值为", i2)
	}
	fmt.Println("//////////////////////")
	for i := range v1 {
		fmt.Println("只要数组下标i:", i)
	}
}