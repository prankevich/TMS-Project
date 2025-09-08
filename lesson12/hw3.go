package main

import (
	"encoding/json"
	"fmt"
	"github.com/xuri/excelize/v2"
	"os"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	data, err := os.ReadFile("lesson12/users.json")
	if err != nil {
		return
	}
	var users []User
	err = json.Unmarshal(data, &users)
	if err != nil {
		return
	}
	f := excelize.NewFile()
	sheet := "Sheet1"
	err = f.SetCellValue(sheet, "A1", "Name")
	if err != nil {
		return
	}
	err = f.SetCellValue(sheet, "B1", "Age")
	if err != nil {
		return
	}
	for i, u := range users {
		row := i + 2
		err = f.SetCellValue(sheet, fmt.Sprintf("A%d", row), u.Name)
		if err != nil {
			return
		}
		err = f.SetCellValue(sheet, fmt.Sprintf("B%d", row), u.Age)
		if err != nil {
			return
		}
	}
	err = f.SaveAs("report.xlsx")
	if err != nil {
		return

	}
}
