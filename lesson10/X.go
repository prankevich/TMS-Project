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
		fmt.Println("Ошибка чтения файла")
		return
	}
	var v Library
	if err := xml.Unmarshal(data, &v); err != nil {
		fmt.Println("Ошибка чтения XML", err)
		return
	}
	for i := range v.Books {
		v.Books[i].Year++
		fmt.Printf("Год: %d, Наименование: %s\n", v.Books[i].Year, v.Books[i].Title)

	}

}
