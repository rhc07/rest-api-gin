package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []Todo{
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Make Bed", Completed: true},
	{ID: "3", Item: "Do Homework", Completed: false},
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func addTodo(context *gin.Context) {
	var newTodo Todo

	if err := context.BindJSON(&newTodo); err != nil {
		return
	}
	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodoById(id string) (*Todo, error) {
	for i, td := range todos {
		if td.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("todos not found")
}

func getTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
	}
	context.IndentedJSON(http.StatusOK, todo)
}

func toggleTodoStatus(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
	}
	todo.Completed = !todo.Completed
	context.IndentedJSON(http.StatusOK, todo)
}

func main() {
	fmt.Printf("Start server")

	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.PATCH("/todos/:id", toggleTodoStatus)
	router.POST("/todos", addTodo)

	router.Run("localhost:9090")
}
