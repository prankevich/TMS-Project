package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Message struct {
	Name string `json:"name"`
}

func main() {
	http.HandleFunc("/echo", echoHandler)
	http.HandleFunc("/greet", messageApi)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Ошибка запуска")
	}

}

func echoHandler(w http.ResponseWriter, reg *http.Request) {
	body, err := io.ReadAll(reg.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = w.Write(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func messageApi(w http.ResponseWriter, reg *http.Request) {
	var msg Message
	err := json.NewDecoder(reg.Body).Decode(&msg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	msg2 := "Привет " + msg.Name
	err = json.NewEncoder(w).Encode(msg2)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
