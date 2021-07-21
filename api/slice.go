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

	//2种遍历切片
	for i := 0; i < len(summer); i++ {
		fmt.Println("summer[", i, "]", summer[i])
	}
	fmt.Println("/////////////////////")
	for i, v := range summer {
		fmt.Println("summer[", i, "]", v)
	}

	//动态增加元素
	oldSlice := make([]int, 5, 10)
	newSlice := append(oldSlice, 1,2,3)
	fmt.Println("oldSlice:", oldSlice)
	fmt.Println("newSlice:", newSlice)
	appendSlice := []int{1,2,3,4,5}
	newSlice1 := append(oldSlice, appendSlice...)
	fmt.Println("newSlice1:", newSlice1) //[0 0 0 0 0 1 2 3 4 5]

	slice1 := []int{1,2,3,4,5}
	slice2 := []int{5,4,3}
	// 复制 slice1 到 slice 2
	//copy(slice2, slice1)
	//fmt.Println(slice2)
	copy(slice1, slice2)
	fmt.Println("copy_slice1:", slice1) //[5 4 3 4 5]

	slice3 := []int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println("slice3原始长度:", len(slice3))
	fmt.Println("slice3原始容量:", cap(slice3))
	fmt.Println(slice3)
	slice3 = slice3[:len(slice3) - 5]
	fmt.Println(slice3)
	slice3 = slice3[5:]
	fmt.Println(slice3)
	fmt.Println("slice3长度变为", len(slice3))
	fmt.Println("slice3的容量变为",cap(slice3))
}







