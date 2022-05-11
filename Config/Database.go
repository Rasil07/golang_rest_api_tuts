package Config

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

type DBConfig struct{
	Host string
	Port int
	User string
	DBName string
	Password string
}

func BuildConfig() *DBConfig{

// 	DB_USERNAME = "root"
// DB_PASSWORD = "Coder@123"
// DB_NAME = "todo_db"
// DB_HOST = "127.0.0.1"
// DB_PORT = 3306
	dbConfig:= DBConfig{
		Host: "localhost",
		Port: 3306,
		User: "root",
		Password: "Coder@123",
		DBName: "todo_db",
	}
	return &dbConfig
}

func DbURL(dbConfig *DBConfig) string{
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	   )
}