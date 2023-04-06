package queries

const (
	GetUser = `
		{
			"select": [
				{"col": "u.*"}
			],
			"from": {
				"value": "users", "as": "u"
			},
			"where": {
				"and": [
					{"col":"id", "value":"u.id"},
					{"col":"email", "value":"u.email"},
					{"col:"-", "value":"u.deleted_at is null"}
				]
			}
	  	}
	`

	InsertUser = `
		insert into users(created_at, created_by, updated_at, updated_by, first_name, last_name, email, phone, password)
		values(NOW(),?,NOW(),?,?,?,?,?,?)
	`
)
