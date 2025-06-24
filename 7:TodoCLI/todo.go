package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (todos *Todos) validateIndex(i int) error {
	if i < 0 || i >= len(*todos) {
		err := errors.New("index out of bound")
		return err
	}
	return nil
}

func (todos *Todos) add(title string) {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}
	*todos = append(*todos, todo)
	fmt.Println("Todo Added Successfully")
}

func (todos *Todos) delete(i int) error {
	if err := todos.validateIndex(i); err != nil {
		return err
	}
	*todos = append((*todos)[:i], (*todos)[i+1:]...)
	return nil
}

func (todos *Todos) edit(i int, con string) error {
	t := *todos
	if err := todos.validateIndex(i); err != nil {
		return err
	}
	t[i].Title = con
	t[i].Completed = false
	t[i].CompletedAt = nil
	return nil
}

func (todos *Todos) toggle(i int) error {
	t := *todos
	if err := todos.validateIndex(i); err != nil {
		return err
	}
	isCompleted := t[i].Completed
	if !isCompleted {
		completedTime := time.Now()
		t[i].CompletedAt = &completedTime
	}
	t[i].Completed = !isCompleted
	return nil
}

func (todos *Todos) print() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Completed", "Created At", "Completed At")

	if len(*todos) == 0 {
		fmt.Println("No todos found. Add some with -add!")
		return
	}

	for index, t := range *todos {
		completed := "❌"
		CompletedAt := ""
		if t.Completed {
			completed = "✅"
			if t.CompletedAt != nil {
				CompletedAt = t.CompletedAt.Format("2006-01-02 15:04:05")
			}
		}
		table.AddRow(strconv.Itoa(index), t.Title, completed, t.CreatedAt.Format("2006-01-02 15:04:05"), CompletedAt)
	}
	table.Render()
}
