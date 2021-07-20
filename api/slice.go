package main

import "fmt"

func main() {
	//基于数组
	// 先定义一个数组
	months := [...]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}

	// 基于数组创建切片
	q2 := months[3:6]    // 第二季度
	summer := months[5:8]  // 夏季

	fmt.Println(q2)
	fmt.Println(len(q2))
	fmt.Println(cap(q2))
	fmt.Println(summer)
	fmt.Println("//////////")
	//基于切片
	firsthalf := months[:6]
	fmt.Println(cap(firsthalf))
	q1 := firsthalf[:3]
	fmt.Println(q1)
	//直接创建切片
	mySlice1 := make([]int, 5)
	mySlice2 := make([]int, 5, 10)
	mySlice3 := []int{1,2,3,4,5,6}
	fmt.Println(mySlice1)
	fmt.Println("mySlice1的容量：", cap(mySlice1))
	fmt.Println(mySlice2)
	fmt.Println("mySlice2的容量:", cap(mySlice2))
	fmt.Println("mySlice2的长度:", len(mySlice2))
	fmt.Println(mySlice3)
	fmt.Println("mySlice3的容量", cap(mySlice3))

	for i := 0; i < len(summer); i++ {
		fmt.Println("summer[", i, "]", summer[i])
	}
	fmt.Println("/////////////////////")
	for i, v := range summer {
		fmt.Println("summer[", i, "]", v)
	}
}
