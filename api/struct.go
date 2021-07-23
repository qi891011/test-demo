package main

import "fmt"

//结构体
type Books struct {
	title string
	author string
	subject string
	book_id int
}
func main() {
	//创建新的结构体
	struct1 := Books{"Go语言", "RobertGriesemer", "Go语言教程", 10001}
	fmt.Println(struct1)
	//key=>value格式
	struct2 := Books{title: "Go语言", author: "RobertGriesemer", subject: "Go语言教程", book_id: 10001}
	fmt.Println(struct2)
	//也可以忽略字段， 忽略的字段为 0 或 空,不按顺序也可以，但是显示是按上文的结构体的顺序
	struct3 := Books{book_id: 10001, author: "RobertGriesemer", title: "Go语言"}
	fmt.Println(struct3)

	//访问结构体成员
	fmt.Println(struct3.title)
}
