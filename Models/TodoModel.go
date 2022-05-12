package Models

import (
	"time"
)

type Todo struct{
	ID uint `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	CreatedAt    time.Time      `json:"created_at" form:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at" form:"updated_at"`
 }

 func (b *Todo) TableName() string{
	 return "todo"
 }

 