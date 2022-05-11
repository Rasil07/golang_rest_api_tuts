package Controllers

import (
	"first_api/tuts/Models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)


func GetTodos( c *gin.Context){
	var todos []Models.Todo
	err:=Models.GetAllTodos(&todos)

	if err!=nil{
		c.AbortWithStatus(http.StatusNotFound)
	}else{
		c.JSON(http.StatusOK,todos)
	}
}

func CreateTodo ( c *gin.Context){
	var todo Models.Todo
	c.BindJSON(&todo)

	err:= Models.CreateTodo(&todo)

	if err!=nil{
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	}else{
		c.JSON(http.StatusOK,todo)
	}
}