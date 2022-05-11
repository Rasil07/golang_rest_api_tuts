package Routes

import (
	"first_api/tuts/Controllers"

	"github.com/gin-gonic/gin"
)


func SetupRouter() *gin.Engine{

	r:=gin.Default()
	grp1 := r.Group("/api")
	{
		grp1.GET("/todos",Controllers.GetTodos)
		grp1.POST("/todos",Controllers.CreateTodo)
		// grp1.GET("/todos",Controllers.GetTodoById)
		// grp1.PUT("/todos",Controllers.UpdateTodos)
		// grp1.DELETE("/todos",Controllers.DeleteTodos)

	}
	return r
}