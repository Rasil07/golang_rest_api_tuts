package Middlewares

import (
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	// return func(c *gin.Context) {
	// 	method := c.Request.Method

	// 	//You can replace * with the specified domain name

	// 	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	// 	// c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
	// 	// c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
	// 	// c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
	// 	// c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

	// 	if method == "OPTIONS" {
	// 		c.AbortWithStatus(http.StatusNoContent)
	// 	}

	// 	c.Next()
	// }

}
