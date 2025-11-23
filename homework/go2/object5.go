package main

import "fmt"

/*
题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体
	   实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
*/

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	length float64
	width  float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() float64 {
	return r.length * r.width
}

func (r Circle) Area() float64 {
	return r.radius * r.radius * 3.14
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.length + r.width)
}

func (r Circle) Perimeter() float64 {
	return 2 * r.radius * 3.14
}

func main() {
	var s Shape
	s = Rectangle{length: 2, width: 3}
	fmt.Println("Rectangle Area:", s.Area())
	fmt.Println("Rectangle Perimeter:", s.Perimeter())
	s = Circle{radius: 5}
	fmt.Println("Circle Area:", s.Area())
	fmt.Println("Circle Perimeter:", s.Perimeter())
}
