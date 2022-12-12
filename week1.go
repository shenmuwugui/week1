package main

import "fmt"

func main() {
	var a int = 10
	var b int = 2
	var num1, num2, num3, num4 int
	num1 = a + b
	num2 = a - b
	num3 = a * b
	num4 = a / b
	fmt.Println("加法：", num1)
	fmt.Println("减法：", num2)
	fmt.Println("乘法：", num3)
	fmt.Println("除法：", num4)

	var i float32 = 5
	_ = i
	fmt.Printf("%.2f\n", i)

	var sum int = 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	var school string = "八维教育"
	fmt.Printf("我爱%s\n", school)
	myschool(school)
}

func myschool(school_name string) {
	fmt.Printf("%s", school_name)
}
