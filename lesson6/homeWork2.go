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

	person1 := Employee{"Gleb", 27, "baker", 2700}
	person2 := Employee{"Sveta", 27, "baker", 1700}
	person3 := Employee{"Alex", 20, "developer", 1000}

	var emp []Employee
	emp = addingEmployee(emp, person1)
	emp = addingEmployee(emp, person2)
	emp = addingEmployee(emp, person3)

	avgmap := averageSalary(emp)
	fmt.Println("Вывод обновленных данных о средней З/П по должности  :", avgmap["baker"])
	fmt.Printf("Средний возраст  всех сотрудников: %2.f \n", averageAge(emp))
	result, err := searchEmployee(emp, "Alex")
	if err != nil {
		fmt.Println(err)
	}
	for _, e := range result {
		var p Printable = e
		p.Print()
	}

}

func (e Employee) Info() string {
	return fmt.Sprintf("Имя:%-10s Возраст:%-3d Должность:%-10s Зарплата:%.2f", e.Name, e.Age, e.Position, e.Salary)

}
func addingEmployee(emp []Employee, person Employee) []Employee {
	newEmp := append(emp, person)
	return newEmp
}
func (e Employee) Print() {
	fmt.Println(e.Info())
}

func searchEmployee(emp []Employee, name string) ([]Employee, error) {
	resultSearch := make([]Employee, 0)
	for _, e := range emp {
		if e.Name == name {
			resultSearch = append(resultSearch, e)
		}
	}
	if len(resultSearch) == 0 {
		return nil, errors.New("Нет такого имени")
	}
	return resultSearch, nil
}
func averageAge(emp []Employee) float64 {
	var sum float64
	for _, e := range emp {
		sum += float64(e.Age)
	}
	return sum / float64(len(emp))

}
func averageSalary(emp []Employee) map[string]float64 {
	countSum := make(map[string]float64)
	countPers := make(map[string]int)
	for _, e := range emp {
		countSum[e.Position] += e.Salary
		countPers[e.Position]++
	}

	average := make(map[string]float64)
	for p, s := range countSum {
		average[p] = s / float64(countPers[p])

	}
	return average
}
