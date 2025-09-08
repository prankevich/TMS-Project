package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("lesson12/divide.txt")
	if err != nil {
		fmt.Println("Ошибка открытия", err)
		return
	}
	defer file.Close()

	var lines strings.Builder
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		cleaned := workString(line)
		lines.WriteString(cleaned + "\n")
	}
	if err = scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	words := strings.Fields(lines.String())
	wordCount := make(map[string]int)
	for _, word := range words {
		wordCount[word]++
	}

	for word, count := range wordCount {
		fmt.Printf("%s: %d\n", word, count)
	}

	file, err = os.Create("lesson12/result.csv")
	if err != nil {
		fmt.Println("Ошибка в создании файла", err)
		return
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, word := range words {
		writer.Write([]string{word, fmt.Sprint(wordCount[word])})
	}

}

func workString(s string) string {
	s = strings.ToLower(s)
	s = strings.TrimSpace(s)
	return s
}
