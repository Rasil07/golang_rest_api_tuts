package main

import (
	"first_api/tuts/Config"
	"first_api/tuts/Models"
	"first_api/tuts/Routes"
	"fmt"

	"github.com/jinzhu/gorm"
)

var err error


func main(){

	Config.DB,err = gorm.Open("mysql", Config.DbURL(Config.BuildConfig()) )

	if err!=nil{
		fmt.Println("Status :",err)
	}

	defer Config.DB.Close()
		Config.DB.AutoMigrate(&Models.Todo{})
	r := Routes.SetupRouter()


	r.Run()	
}