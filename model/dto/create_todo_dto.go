package dto

// CreateTodoDto is dto to create todo
type CreateTodoDto struct {
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
}
