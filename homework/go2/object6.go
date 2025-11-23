package main

/*
题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体,
       组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
*/
import (
	"fmt"
	"strconv"
)

type Person struct {
	name string
	age  int
}

type Employee struct {
	Person
	EmployeeID int
}

func (r Employee) PrintInfo() {
	fmt.Println("name: " + r.name + ", age: " + strconv.Itoa(r.age) + ", EmployeeID: " + strconv.Itoa(r.EmployeeID))
}

func main() {
	emp := Employee{
		Person:     Person{name: "Alice", age: 30},
		EmployeeID: 12345,
	}
	emp.PrintInfo()
}
