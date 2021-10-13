package restful

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ToDo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func Demo1() {
	response, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(response.Body)

	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
	var todoResponse ToDo
	json.Unmarshal(bodyBytes, &todoResponse) // burada json dan todoya dönüştürüyoruz
	fmt.Println(todoResponse)
}

func Demo2() {
	todo := ToDo{1, 2, "alışveriş yapılacak", false}
	jsonTodo, err := json.Marshal(todo) // burada todo yu json formatına çevirdik

	response, err := http.Post("https://jsonplaceholder.typicode.com/todos", "application/json;charset=utf-8", bytes.NewBuffer(jsonTodo))
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	// body okuma işlemleri
	bodyBytes, _ := ioutil.ReadAll(response.Body)

	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
	var todoResponse ToDo
	json.Unmarshal(bodyBytes, &todoResponse) // burada json dan todoya dönüştürüyoruz
	fmt.Println(todoResponse)
}
