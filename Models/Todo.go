package Models

import (
	"first_api/tuts/Config"

	_ "github.com/go-sql-driver/mysql"
)

func GetAllTodos(todo *[]Todo) (err error) {

	if err = Config.DB.Find(todo).Error; err != nil {
		return err
	}
	return nil
}

func CreateTodo(todo *Todo) (err error) {

	if err = Config.DB.Create(todo).Error; err != nil {
		return err
	}
	return nil
}

func GetTodoById(todo *Todo, id string) (err error) {
	if err = Config.DB.Where("id=?", id).First(todo).Error; err != nil {
		return err
	}
	return nil
}

func UpdateTodoById(todo *Todo) (err error) {
	if err = Config.DB.Save(todo).Error; err != nil {
		return err
	} else {
		return nil
	}
}

func DeleteTodosById(todo *Todo) (err error) {
	if err = Config.DB.Delete(todo).Error; err != nil {
		return err
	} else {
		return nil
	}
}
