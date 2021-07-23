package main

import (
	"fmt"
	"unsafe"
)

//指针
func main() {
	var ptr *int
	fmt.Println(ptr)

	a := 100
	ptr = &a
	fmt.Println(ptr)
	fmt.Println(*ptr)

	ptr1 := new(int)
	*ptr1 = 100
	fmt.Println("ptr1:", ptr1)

	i := 10
	var p *int = &i
	fmt.Println("p:", p)
	fmt.Println("p:", *p)

	arr := [3]int{1,2,3}
	ap := &arr

	sp := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(ap)) + unsafe.Sizeof(arr[0])))
	//fmt.Println(unsafe.Sizeof(arr[0]))
	//fmt.Println("ap:",ap)
	*sp += 3
	fmt.Println(arr)
}

//func swap(a, b *int){
//	*a, *b = *b, *a
//	fmt.Println(*a, *b)
//}