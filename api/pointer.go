package main

import "fmt"

//指针
func main() {
	//var ptr *int
	//fmt.Println(ptr)
	//a := 100
	//ptr = &a
	//fmt.Println(ptr)
	//fmt.Println(*ptr)


	//a := 100
	//ptr := &a
	//fmt.Printf("%p\n", ptr)
	//fmt.Printf("%d\n", *ptr)

	a := 1
	b := 2
	swap(&a, &b)
	fmt.Println(a, b)
}
func swap(a, b *int){
	*a, *b = *b, *a
	fmt.Println(*a, *b)
}