package Models

type Todo struct{
	Id uint `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
 }

 func (b *Todo) TableName() string{
	 return "todo"
 }

 