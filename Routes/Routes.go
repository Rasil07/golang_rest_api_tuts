package Routes

import (
	"first_api/tuts/Controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	grp1 := r.Group("/api")
	{
		grp1.GET("/todos", Controllers.GetTodos)
		grp1.POST("/todos", Controllers.CreateTodo)
		grp1.GET("/todos/:id", Controllers.GetById)
		grp1.PATCH("/todos/:id", Controllers.UpdateById)
		grp1.DELETE("/todos/:id", Controllers.DeleteTodoById)

	}
	return r
}
