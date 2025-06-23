package main

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", handleHome)
	r.GET("/hello", helloHandlr)
	r.GET("/users", handleGetUsers)
	r.GET("/items", handleGetItems)
	r.GET("/file", handleGetDataFromFile)
	r.GET("/download/:filename", handleDownloadFile)
	r.GET("/search", Handlquery)
	r.Run(":8080")
}

func helloHandlr(c *gin.Context) {
	c.String(http.StatusOK, "Hello from go server")
}
func handleHome(c *gin.Context) {
	c.String(http.StatusOK, "Server is running on port :8080")
}

func handleGetUsers(c *gin.Context) {
	users := gin.H{
		"name": "vinoth",
		"age":  21,
	}
	c.JSON(http.StatusOK, users)
}
func handleGetItems(c *gin.Context) {
	users := []gin.H{
		{"id": 1, "name": "apple", "price": 12},
		{"id": 2, "name": "orannge", "price": 15},
		{"id": 3, "name": ",mango", "price": 19},
	}
	c.JSON(http.StatusOK, users)
}

func handleGetDataFromFile(c *gin.Context) {
	c.File("./text.txt")
}

func handleDownloadFile(c *gin.Context) {
	fname := c.Param("filename")
	basDir := "./files"
	filepath := filepath.Join(basDir, fname)

	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		c.String(http.StatusNotFound, "File not found")
		return
	}
	c.Header("Content-Disposition", "attachment; filename="+fname)
	c.File(filepath)
}

func Handlquery(c *gin.Context) {
	name := c.DefaultQuery("name", "Guest")
	c.String(http.StatusOK, "Hello! %s", name)
}
