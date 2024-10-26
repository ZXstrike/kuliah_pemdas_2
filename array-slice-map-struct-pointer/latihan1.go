package arrayslicemapstructpointer

import (
	"fmt"
)

type Employee struct {
	Name   string
	Age    int
	Salary float64
}

var employees []Employee
var employeeMap map[string]*Employee

func init() {
	employeeMap = make(map[string]*Employee)
}

func addEmployee(name string, age int, salary float64) {
	employee := Employee{Name: name, Age: age, Salary: salary}
	employees = append(employees, employee)
	employeeMap[name] = &employee
}

func displayEmployees() {
	for _, employee := range employees {
		fmt.Printf("Name: %s, Age: %d, Salary: %.2f\n", employee.Name, employee.Age, employee.Salary)
	}
}

func findEmployeeByName(name string) *Employee {
	if employee, exists := employeeMap[name]; exists {
		return employee
	}
	return nil
}

func updateEmployee(name string, age int, salary float64) {
	if employee, exists := employeeMap[name]; exists {
		employee.Age = age
		employee.Salary = salary
	}
}

func deleteEmployee(name string) {
	for i, employee := range employees {
		if employee.Name == name {
			employees = append(employees[:i], employees[i+1:]...)
			delete(employeeMap, name)
			break
		}
	}
}

func Latihan1() {
	addEmployee("Alice", 30, 50000)
	addEmployee("Bob", 25, 45000)
	displayEmployees()

	fmt.Println("Finding Alice:")
	if emp := findEmployeeByName("Alice"); emp != nil {
		fmt.Printf("Found: Name: %s, Age: %d, Salary: %.2f\n", emp.Name, emp.Age, emp.Salary)
	}

	fmt.Println("Updating Bob's salary:")
	updateEmployee("Bob", 25, 48000)
	displayEmployees()

	fmt.Println("Deleting Alice:")
	deleteEmployee("Alice")
	displayEmployees()
}
