package Models

import (
	"first_api/tuts/Config"

	_ "github.com/go-sql-driver/mysql"
)


func GetAllTodos(todo *[]Todo) (err error){
	
	if err = Config.DB.Find(todo).Error;err!=nil{
		return err
	}
	return nil
}

func CreateTodo(todo *Todo) (err error){
	
	if err = Config.DB.Create(todo).Error;err!=nil{
		return err
	}
	return nil
}