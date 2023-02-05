package queries

const (
	QueryCreateStock = `
		insert into stocks (name, price, availability, is_active, created_at, updated_at)
		values(?,?,?,?,NOW(),NOW())
	`

	QueryGetStockByID = `
		select * from stocks where id = ?
	`
)
