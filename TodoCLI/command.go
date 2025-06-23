package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add  string
	Del  int
	Edit string
	Tog  int
	List bool
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add a new todo specify title")
	flag.StringVar(&cf.Edit, "edit", "", "Edit the todo by specify index & specify a new title. id:new_title")
	flag.IntVar(&cf.Del, "del", -1, "Specify the index of the todo to delete")
	flag.IntVar(&cf.Tog, "toggle", -1, "Specify the index of the todo to toggle")
	flag.BoolVar(&cf.List, "list", false, "List all todos")

	flag.Parse()

	return &cf
}

func (cf *CmdFlags) Execute(todos *Todos) {
	switch {
	case cf.List:
		todos.print()
	case cf.Add != "":
		todos.add(cf.Add)
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Invalid edit command. Use id:new_title")
			os.Exit(1)
		}
		index, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Invalid index:", err)
			os.Exit(1)
		}
		todos.edit(index, parts[1])
	case cf.Tog != -1:
		todos.toggle(cf.Tog)

	case cf.Del != -1:
		todos.delete(cf.Del)

	default:
		fmt.Println("Invalid command. Use -help for usage.")
	}
}
