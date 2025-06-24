package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	todos   Todos
	storage *Storage[Todos]
)

func main() {
	todos = Todos{}
	storage = CreateNewStorage[Todos]("todos.json")
	storage.Load(&todos)

	r := gin.Default()
	r.GET("/", handleHome)
	r.GET("/home", handleWelcomeHome)

	//todo
	r.GET("/getAll", handleGetAllTodos)
	r.POST("/add", HandleAddTodo)
	r.DELETE("/del/:id", HandleDeleteTodo)
	r.PUT("/toggle/:id", HandleToggleTodo)
	r.PUT("/edit/:id", HandleEditTodo)
	r.Run(":8080")
}

// todo API's
func handleGetAllTodos(c *gin.Context) {
	if len(todos) == 0 {
		c.JSON(http.StatusAccepted, gin.H{"message": "No Todos Available"})
		return
	}
	c.JSON(http.StatusOK, todos)
}

// to add new todo
func HandleAddTodo(c *gin.Context) {
	NewTitle := c.Query("title")
	if NewTitle == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title is Required"})
	}
	todos.add(NewTitle)
	storage.save(todos)
	c.JSON(http.StatusOK, gin.H{"message": "Todo Added Successsfully"})
}

//--------*********--------------

func HandleDeleteTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if err := todos.validateIndex(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Index"})
		return
	}
	if err := todos.delete(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	storage.save(todos)
	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
}

//--------*********--------------

func HandleToggleTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": "Invalid ID"})
		return
	}
	if err := todos.validateIndex(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Index"})
		return
	}
	if err := todos.toggle(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Index"})
		return
	}
	storage.save(todos)
	c.JSON(http.StatusOK, gin.H{"message": "Todo Toggled"})
}
func HandleEditTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	title := c.Query("title")
	if title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "New title is required"})
		return
	}
	if err := todos.validateIndex(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Index"})
		return
	}
	if err := todos.edit(id, title); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}
	storage.save(todos)
	c.JSON(http.StatusOK, gin.H{"message": "Todo Updated"})
}

func handleHome(c *gin.Context) {
	c.String(http.StatusOK, "Server running in port :8080")
}

func handleWelcomeHome(c *gin.Context) {
	c.String(http.StatusOK, "Welcome to Todo App")
}
