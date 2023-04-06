package queries

const (
	GetItems = `
		{
			"select": [
				{"col": "i.*"}
			],
			"from": {
				"value": "items", "as": "i"
			},
			"where": {
				"and": [
					{"col":"id", "value":"i.id"},
					{"col":"user_id", "value:"i.user_id"}
				]
			}
	  	}
	`

	UpdateItems = `
		{
			"set": [
				{"col": "itemData", "value": "i.data"}
			],
			"from": {
				"value": "items", "as": "i"
			},
			"where": {
				"and": [
					{"col":"id", "value":"i.id"},
					{"col":"user_id", "value:"i.user_id"}
				]
			}
		}
	`

	InsertItem = `
		insert into items(user_id, data) values(?,?)
	`
)
