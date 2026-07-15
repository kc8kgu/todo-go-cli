package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
)

type Todo struct {
	Id    int    `json:"id"`
	Done  bool   `json:"done"`
	Title string `json:"title"`
}

var todos []Todo

const path string = "./todo.json"

func main() {

	if len(os.Args) == 2 && os.Args[1] == "list" {
		listTodos()
	} else if len(os.Args) == 3 && os.Args[1] == "add" {
		addTodos()
	} else if len(os.Args) == 4 && os.Args[1] == "edit" {
		editTodos()
	} else if len(os.Args) == 3 && os.Args[1] == "done" {
		doneTodos()
	} else if len(os.Args) == 3 && os.Args[1] == "delete" {
		deleteTodos()
	} else if len(os.Args) == 2 && os.Args[1] == "clear" {
		clearTodos()
	} else {
		usage()
	}
}

func usage() {

	fmt.Println("usage: td list | add | edit | done| delete | clear <id> <title>")

}

func loadTodos() {

	bytes, err := os.ReadFile(path)
	if err != nil {
		// ignore error since file may not exist yet
		return
	}

	err = json.Unmarshal(bytes, &todos)
	if err != nil {
		log.Fatal(err.Error())
	}

}

func saveTodos() {

	bytes, err := json.Marshal(todos)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = os.WriteFile(path, bytes, 0644)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func listTodos() {

	fmt.Println("listing todos")

	loadTodos()

	if len(todos) == 0 {
		fmt.Println("no todos yet!")
		return
	}

	for _, todo := range todos {
		fmt.Printf("%v, %v, \"%v\"\n", todo.Id, todo.Done, todo.Title)
	}
}

func getNextTodoId() int {

	maxId := 0

	for _, todo := range todos {
		if todo.Id > maxId {
			maxId = todo.Id
		}
	}

	return maxId + 1
}

func addTodos() {

	fmt.Println("adding todo")

	loadTodos()

	nextId := getNextTodoId()

	newTodo := Todo{
		Id:    nextId,
		Done:  false,
		Title: os.Args[2],
	}

	fmt.Printf("id: %v, done: %v, title: \"%v\"\n", newTodo.Id, newTodo.Done, newTodo.Title)

	todos = append(todos, newTodo)
	saveTodos()
}

func editTodos() {

	fmt.Println("editing todo title")

	id, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err.Error())
	}

	title := os.Args[3]

	fmt.Printf("id: %d, title: %s\n", id, title)

	loadTodos()

	for i := 0; i < len(todos); i++ {
		if todos[i].Id == id {
			todos[i].Title = title
			saveTodos()
			return
		}
	}

	fmt.Printf("todo with id %v not found", id)
}

func doneTodos() {

	fmt.Println("toggling todo done field")

	id, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("id: %d\n", id)

	loadTodos()

	for i := 0; i < len(todos); i++ {
		if todos[i].Id == id {
			todos[i].Done = !todos[i].Done
			saveTodos()
			return
		}
	}

	fmt.Printf("todo with id %v not found", id)
}

func deleteTodos() {

	fmt.Println("deleting todos")

	id, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("id: %d\n", id)

	loadTodos()

	for i := 0; i < len(todos); i++ {
		if todos[i].Id == id {
			todos = slices.Delete(todos, i, i+1)
			saveTodos()
			return
		}
	}

	fmt.Printf("todo with id %v not found", id)
}

func clearTodos() {

	fmt.Println("clearing todos")
	os.Remove(path)

}
