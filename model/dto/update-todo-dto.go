package dto

type UpdateTodoDto struct {
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	Status      string `json:"status" form:"status"`
}
