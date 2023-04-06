package models

type Reputation struct {
	Badge string `json:"badge" db:"badge" jsondb:"data>$.reputation.badge>CHAR(20)"`
	Value int    `json:"value" db:"value" jsondb:"data>$.reputation.value>UNSIGNED" validate:"required,min=0,max=1000" valerr:"Rating should be filled and between 0 - 1000"`
}

type Item struct {
	ID           int64      `json:"id" db:"id"`
	UserID       int64      `json:"user_id" db:"user_id"`
	Name         string     `json:"name" db:"name" jsondb:"data>$.name>CHAR(50)" validate:"required,min=10,forbiddens=Sex_Gay_Lesbian" valerr:"Name should be filled and longer than 10 characters"`
	Rating       int        `json:"rating" db:"rating" jsondb:"data>$.rating>UNSIGNED" validate:"required,min=0,max=5" valerr:"Rating should be filled and between 0 - 5"`
	Category     string     `json:"category" db:"category" jsondb:"data>$.rating>CHAR(20)"`
	Image        string     `json:"image" db:"image" jsondb:"data>$.image>CHAR(20)" validate:"required,min=0,max=5" valerr:"Image required and should be URL"`
	Reputation   Reputation `json:"reputation" db:"reputation" jsondb:"data>$.reputation"`
	Price        int64      `json:"price" db:"price" jsondb:"data>$.price>UNSIGNED" validate:"required,numeric" valerr:"Price required and should be numeric"`
	Availability int64      `json:"availability" db:"availability" jsondb:"data>$.availability>UNSIGNED" validate:"required,numeric" valerr:"Availability required and should be numeric"`
}

func (u Item) GetTableNameAndAlias() string {
	return "items"
}

type ItemCreatePayload struct {
	UserID int64
	Data   string
}
