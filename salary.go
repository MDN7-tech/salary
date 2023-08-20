package main

import (
	"fmt"
)

func main() {
	var name string
	var salary int32
	var insurance float32 = 0.1
	var tax float32 = 0.09
	var employees [20]string
	var salary_employee [20]int32

	for i := 0; i < 20; i++ {
		fmt.Print("enter employee name: ")
		fmt.Scanln(&name)
		employees[i] = name
		fmt.Println("enter employee", name, "salary: ")
		fmt.Scanln(&salary)
		salary_employee[i] = salary

	}

}
