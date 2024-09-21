package main

import (
	"fmt"
	"go-lang-basic/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/todos", controller.GetTodos)
	router.POST("/addtodo", controller.AddTodo)
	router.Run("localhost:8080")
	fmt.Print("Servidor rodando")

}
