//Go Object-Oriented programming example

package main

import (
    "fmt"
)

type Employee struct {
    Name string
    Id int
    Department string
}

// *Employee is the pointer type of Employee struct
// By using the pointer reference to update the employee details
func(e *Employee) UpdateEmp(){
	e.Name = "Ram"
}

//pass by value, instance method
func(e Employee) IsItDepartment() bool{
	if e.Department == "IT"{
		return true
	} else {
		return false
	}
}

func CreateEmp(name string, id int, dep string) Employee {
    emp := Employee{name, id, dep}
    return emp
}

func main(){
	emp := CreateEmp("viki",123,"IT")
	fmt.Println("New Employee:", emp)
	emp.UpdateEmp()
	fmt.Println("Updated Employee:", emp)
	fmt.Println("IT Department?:", emp.IsItDepartment())
}
