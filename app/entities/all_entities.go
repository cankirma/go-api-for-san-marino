package entities

import (
	"github.com/google/uuid"
	"time"
)

type SignUp struct {
	Email    string `json:"email" validate:"lte=255"`
	Password string `json:"password" validate:"required,lte=255"`
	UserRole string `json:"user_role" validate:"required,lte=25"`
}

type Product struct {
	ProductId  int `json:"product_id" db:"product_id"`
	ProductName  string `json:"product_name" db:"product_name" validate:"required ,lte=40"`
	SupplierId  int `json:"supplier_id" db:"supplier_id" validate:"required"`
	CategoryId  int `json:"category_id" db:"category_id"`
	QuantityPerUnit  string `json:"quantity_per_unit" db:"quantity_per_unit"`
	UnitPrice  float64 `json:"unit_price" db:"unit_price"`
	UnitsInStock int `json:"units_in_stock" db:"units_in_stock"`
	UnitsOnOrder int `json:"units_on_order" db:"units_on_order"`
	ReOrderLevel int `json:"reorder_level" db:"reorder_level"`
	Discontinued int `json:"discontinued" db:"discontinued"`
}
type SignIn struct {
	Email    string `json:"email" validate:"required,email,lte=255"`
	Password string `json:"password" validate:"required,lte=255"`
}
type Category struct {
	CategoryId int `json:"category_id" db:"category_id"`
	CategoryName string `json:"category_name" db:"category_name"`
	Description string `json:"description" db:"description"`
	Picture string  `json:"picture" db:"picture"`
}


type Renew struct {
	RefreshToken string `json:"refresh_token"`
}

type User struct {
	ID           uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
	Email        string    `db:"email" json:"email" validate:"required,email,lte=255"`
	PasswordHash string    `db:"password_hash" json:"password_hash,omitempty" validate:"required,lte=255"`
	UserStatus   int       `db:"user_status" json:"user_status" validate:"required,len=1"`
	UserRole     string    `db:"user_role" json:"user_role" validate:"required,lte=25"`
}

