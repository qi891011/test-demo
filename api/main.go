package main

import (
	"fmt"
	"math"
)
const (
	c0 = iota
	c1
	c2
)
const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	numberOfDays
)
func GetName()(userName, nickName string)  {
	return "nonfu", "学院君"
}
//Go 语言的 main() 函数不能带参数，也不能定义返回值。
func main() {
	const Tuesday = iota
	//_, nickName := GetName()
	_, nickName := GetName()
	fmt.Println("长度：", len(nickName))
	fmt.Printf("昵称:%s\n",nickName)
	fmt.Println(c1)
	fmt.Println(numberOfDays)
	fmt.Println(Tuesday)
	//floatValue1 := 102
	var floatValue1 float32
	floatValue1 = 10
	fmt.Println(floatValue1)
	floatValue2 := 10.0 // 如果不加小数点，floatValue2 会被推导为整型而不是浮点型
	//floatValue3 := 1.1E-10
	floatValue4 := 0.1
	floatValue5 := 0.7
	floatValue6 := floatValue4 + floatValue5
	fmt.Println(floatValue6)

	p := 0.00001
	// 判断 floatValue1 与 floatValue2 是否相等
	if math.Dim(float64(floatValue1), floatValue2) < p {
		fmt.Println("floatValue1 和 floatValue2 相等")
	}else {
		fmt.Println("floatValue1 和 floatValue2 不相等")
	}
	var str string
	str = "hello world"
	ch := str[0]
	fmt.Printf("The length of \"%s\" is %d \n", str, len(str))
	fmt.Printf("The first character of \"%s\" is %c.\n", str, ch)
	result := `Search result for "Golang"
	-GO
	- Golang
	Golang Programming`
	fmt.Printf("%s\n", result)

	str1 := "Hello"
	str1 = str1 +
		", 学院君"
	fmt.Println(str1)
}