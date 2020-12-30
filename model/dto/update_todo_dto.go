package dto

type UpdateTodoDto struct {
	Name        string `json:"name" form:"name" validate:"required;string"`
	Description string `json:"description" form:"description" validate:"required;string"`
	Status      string `json:"status" form:"status" validate:"required;string" validate:"required;string"`
}
