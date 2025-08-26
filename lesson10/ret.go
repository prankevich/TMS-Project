package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Book struct {
	Title string `xml:"title"`
	Year  int    `xml:"year"`
}
type Library struct {
	Books []Book `xml:"book"`
}

func main() {
	data, err := os.ReadFile("C:\\Users\\dell\\GolandProjects\\TMS-Project1\\lesson10\\books.xml")
	if err != nil {
		fmt.Println(err)
	}
	var v Library
	xml.Unmarshal(data, &v)

	for _, v := range v.Books {
		v.Year = v.Year + 1
		fmt.Println("Год:", v.Year, v.Title)
	}

}
