package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	file, err := os.OpenFile("app.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	log.SetOutput(file)
	log.SetFlags(log.LstdFlags)

	text := bufio.NewScanner(os.Stdin)
	fmt.Println("Вводите текст,либо `x` для выхода")

	for {
		print("")
		if !text.Scan() {
			break
		}
		t := text.Text()
		if t == "x" {
			break
		}
		log.Println(t)

	}
}
