package main

import "fmt"

type Employee struct {
	Name     string
	Age      int
	Position string
	Salary   float64
}

func main() {
	emp := []Employee{
		{Name: "Alex", Age: 20, Salary: 1000, Position: "developer"}}
	emp = addingEmployee(emp, "Sveta", 25, 1500, "manager")
	fmt.Println(emp)
	sal := make(map[string]float64)
	for _, e := range emp {
		sal[e.Position] = e.Salary
		fmt.Println(sal["manager"])
	}
}

func (e Employee) Info() string {
	return fmt.Sprint(e.Name, ", ", e.Age, ", ", e.Position, ", ", e.Salary)

}

func addingEmployee(emp []Employee, n string, a int, s float64, p string) []Employee {
	newEmp := Employee{
		Name:     n,
		Age:      a,
		Salary:   s,
		Position: p,
	}
	return append(emp, newEmp)

}
