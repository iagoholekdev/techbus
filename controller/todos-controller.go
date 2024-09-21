package controller

import (
	"fmt"
	todos "go-lang-basic/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTodos(context *gin.Context) {
	context.JSON(200, todos.Todos)
}

func AddTodo(context *gin.Context) {
	var newTodo todos.Todo

	if err := context.BindJSON(&newTodo); err != nil {
		fmt.Print(err)
		return
	}

	todos.Todos = append(todos.Todos, newTodo)

	context.JSON(http.StatusCreated, newTodo)
}
