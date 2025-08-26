package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	user, err := load("C:\\Users\\dell\\GolandProjects\\TMS-Project1\\lesson10\\users.json")
	if err != nil || user == nil {
		fmt.Println("Ошибка загрузки файла :", err)
		return
	}
	uppAge(user)
	for _, user := range user {
		fmt.Println(user)
	}
}

func load(path string) ([]User, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg []User
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

func uppAge(users []User) {
	for i, user := range users {
		users[i].Age = user.Age + 1
	}
}
