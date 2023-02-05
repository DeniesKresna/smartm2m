package gate

func (c *Gate) InitRoutes() {
	c.Post("/stock", c.AddStock)
	c.Get("/stock/{id}", c.GetStockByID)
}
