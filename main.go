package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Todo struct {
	Id    int    `json:"id"`
	Done  bool   `json:"done"`
	Title string `json:"title"`
}

var todos []Todo

const PATH string = "./todo.json"

func main() {

	if len(os.Args) == 2 && os.Args[1] == "list" {
		listTodos()
	} else if len(os.Args) == 3 && os.Args[1] == "add" {
		addTodos()
	} else if len(os.Args) == 4 && os.Args[1] == "edit" {
		editTodos()
	} else {
		usage()
	}
}

func usage() {
	fmt.Println("usage: td command [id] [title]")
}

func loadTodos() {

	//jsonString := `[{"id": 1, "done": false, "title": "git init"}]`

	jsonBytes, err := os.ReadFile(PATH)
	if err != nil {
		log.Fatal("error reading file")
	}

	err = json.Unmarshal(jsonBytes, &todos)
	if err != nil {
		log.Fatal("error parsing json")
	}
}

func listTodos() {
	fmt.Println("list todos")

	loadTodos()

	for _, todo := range todos {
		fmt.Printf("%v, %v, %v, ", todo.Id, todo.Done, todo.Title)
	}
}

func addTodos() {
	fmt.Println("add todos")

	fmt.Printf("title: %s\n", os.Args[2])
}

func editTodos() {
	fmt.Println("edit todos")

	var id, title string = os.Args[2], os.Args[3]
	fmt.Printf("id: %s, title: %s\n", id, title)
}
