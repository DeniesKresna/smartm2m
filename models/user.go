package models

import "time"

type User struct {
	ID        int64      `json:"id" db:"id"`
	CreatedBy string     `json:"created_by" db:"created_by"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedBy string     `json:"updated_by" db:"updated_by"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
	DeletedBy *string    `json:"deleted_by" db:"deleted_by"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at"`
	FirstName string     `json:"first_name" db:"first_name"`
	LastName  string     `json:"last_name" db:"last_name"`
	Email     string     `json:"email" db:"email"`
	Phone     string     `json:"phone" db:"phone"`
	Password  string     `json:"-" db:"password" sqlq:"userPassword"`
}

func (u User) GetTableName() string {
	return "users"
}

type UserCreatePayload struct {
	FirstName string `json:"first_name" validate:"required" valerr:"First Name should be Filled"`
	LastName  string `json:"last_name" validate:"required" valerr:"Last Name should be Filled"`
	Email     string `json:"email" validate:"required,email" valerr:"Email should be Well Filled"`
	Phone     string `json:"phone" validate:"required" valerr:"Phone should be Filled"`
	Password  string `json:"password" validate:"required,min=8" valerr:"Password should be Filled and minimum 8 chars"`
}
