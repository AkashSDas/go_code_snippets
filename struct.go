package main

import "fmt"

type Employee struct {
	firstName string
	lastName  string
}

type Manager struct {
	department string
	emp        *Employee
}

func getFirstName(emp *Employee) string {
	return emp.firstName
}

func setFirstName(emp *Employee, newName string) {
	emp.firstName = newName
}

func main() {
	var emp1 Employee = Employee{"James", "Bond"}
	emp2 := Employee{firstName: "Ethan", lastName: "Hunt"}

	fmt.Println(emp1.firstName, emp2.lastName)

	mng1 := Manager{
		department: "Hell",
		emp:        &Employee{"Tony", "Ferguson"},
	}
	fmt.Println(mng1)
	fmt.Println(getFirstName(mng1.emp))
	setFirstName(mng1.emp, "Sherlok")
	fmt.Println(getFirstName(mng1.emp))
}
