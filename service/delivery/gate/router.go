package gate

func (c *Gate) InitRoutes() {
	c.Post("/stock", Protected(c.AddStock))
	c.Get("/stock/{id}", c.GetStockByID)
	c.Post("/stock-bulk", c.CreateStockBulk)
}
