package main

import (
	"first_api/tuts/Config"
	"first_api/tuts/Models"
	"first_api/tuts/Routes"
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/jinzhu/gorm"
)

var err error

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}

func main(){

	Config.DB,err = gorm.Open("mysql", Config.DbURL(Config.BuildConfig()) )

	if err!=nil{
		fmt.Println("Status :",err)
	}

	defer Config.DB.Close()
		Config.DB.AutoMigrate(&Models.Todo{})
	r := Routes.SetupRouter()

	r.Use(CORSMiddleware())

	r.Run()	
}