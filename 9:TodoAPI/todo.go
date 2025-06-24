package main

import (
	"errors"
	"fmt"
	"time"
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
