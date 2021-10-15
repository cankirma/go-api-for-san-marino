package dtos

type CreateCategoryModel struct {
	CategoryId string `json:"category_id" db:"category_id" validate:"required,int"`
	CategoryName string `json:"category_name" db:"category_name" validate:"required,string"`
	Description string `json:"description" db:"description" validate:"required,string"`
}


type UpdateCategoryModel struct {
	CategoryId int `json:"category_id" db:"category_id"`
	CategoryName string `json:"category_name" db:"category_name"`
	Description string `json:"description" db:"description"`
}