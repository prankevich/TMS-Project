package main

import (
	"errors"
	"fmt"
)

type Employee struct {
	Name     string
	Age      int
	Position string
	Salary   float64
}
type Printable interface {
	Print()
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
	var p Printable = emp[1]
	p.Print()
	fmt.Println(searchEmployee(emp, "Sveta"))
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
func (e Employee) Print() {
	fmt.Println(e.Info())
}

func searchEmployee(emp []Employee, name string) ([]Employee, error) {
	resultSearch := make([]Employee, 0)
	for _, e := range emp {
		if e.Name == name {
			resultSearch = append(resultSearch, e)
		} else {
			return nil, errors.New("Нет имени ")
		}
	}
	return resultSearch, nil

}
