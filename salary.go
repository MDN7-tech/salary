package main

import (
	"fmt"
	"os"
)

func main() {
	var name string
	var salary float32
	var insurance float32 = 0.1
	var tax float32 = 0.09
	var employees [20]string
	var salary_employee [20]float32
	var netـsalary [20]float32
	var employee_num int
	var reoprt_mod string

// Open the file for writing
    file, errs := os.Create("salary_report.csv")
     if errs != nil {
      fmt.Println("Failed to create file:", errs)
      return
     }
     defer file.Close()
	 _, errs = file.WriteString(fmt.Sprintf("id,employees,salary_employee,netـsalary\n"))
	 if errs != nil {
		fmt.Println("Failed to write to file:", errs) //print the failed message
		return
	 }

//get How many employees do you have?
	fmt.Println(" How many employees do you have?(1-20)")
	fmt.Scanln(&employee_num)
//get data
	for i := 0; i < employee_num; i++ {
//get nameemployee
		fmt.Println("Name of employee number",i+1 ,":")
		fmt.Scanln(&name)
		employees[i] = name
//get salary
		fmt.Println("Salary of employee number",i+1 ,":")
		fmt.Scanln(&salary)
		salary_employee[i] = salary
//calculate netsalary
		netـsalary [i]= salary - salary*(tax+insurance)
	}
//get report mode
	fmt.Println("how do you want export report? (screen,csv,both)")
	fmt.Scanln(&reoprt_mod)
//report
	if reoprt_mod=="screen" {
      for i := 0; i < employee_num; i++ {
	    fmt.Println(i+1,employees[i],salary_employee[i],netـsalary [i])
        }
	} else if reoprt_mod=="csv"{
		for i := 0; i < employee_num; i++ {
			_, errs = file.WriteString(fmt.Sprintf("%d,%s,%g,%g \n", i+1, employees[i], salary_employee[i],netـsalary [i]))
			if errs != nil {
			   fmt.Println("Failed to write to file:", errs) //print the failed message
			   return
			}
			
		}
	} else{
		for i := 0; i < employee_num; i++ {
			fmt.Println(i+1,employees[i],salary_employee[i],netـsalary [i])
			_, errs = file.WriteString(fmt.Sprintf("%d,%s,%g,%g \n", i+1, employees[i], salary_employee[i],netـsalary [i]))
			if errs != nil {
			   fmt.Println("Failed to write to file:", errs) //print the failed message
			   return			
			}    
	}
}
}