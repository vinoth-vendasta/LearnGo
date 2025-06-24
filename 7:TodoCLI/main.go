package main

import "fmt"

func main() {
	fmt.Println("Inside main in Todo App")

	todos := Todos{}

	storage := NewStorage[Todos]("todos.json")
	storage.Load(&todos)

	cmdFlage := NewCmdFlags()
	cmdFlage.Execute(&todos)

	storage.Save(todos)
}
