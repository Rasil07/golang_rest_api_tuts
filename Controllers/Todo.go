package Controllers

import (
	"first_api/tuts/Models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	var todos []Models.Todo
	err := Models.GetAllTodos(&todos)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todos)
	}
}

func CreateTodo(c *gin.Context) {
	var todo Models.Todo
	c.BindJSON(&todo)

	err := Models.CreateTodo(&todo)

	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func GetById(c *gin.Context) {
	var todo Models.Todo
	id := c.Param("id")

	err := Models.GetTodoById(&todo, id)

	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func UpdateById(c *gin.Context) {
	var todo Models.Todo
	id := c.Param("id")

	err := Models.GetTodoById(&todo, id)
	if err != nil {
		c.JSON(http.StatusNotFound, todo)
	}
	c.BindJSON(&todo)
	err = Models.UpdateTodoById(&todo)

	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func DeleteTodoById(c *gin.Context) {
	var todo Models.Todo
	id := c.Param("id")

	err := Models.GetTodoById(&todo, id)

	if err != nil {
		c.JSON(http.StatusNotFound, todo)
	}
	err = Models.DeleteTodosById(&todo)

	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}
