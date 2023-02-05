package gate

func (c *Gate) InitRoutes() {
	c.Post("/stock", c.AddStock)
	c.Get("/stock/{id}", c.GetStockByID)
	c.Post("/stock-bulk", c.CreateStockBulk)
}
