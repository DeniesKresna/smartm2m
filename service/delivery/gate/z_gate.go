package gate

import (
	"net/http"
	"strings"

	"github.com/DeniesKresna/danatest/models"
	"github.com/DeniesKresna/danatest/service/usecase"
)

type Gate struct {
	StockUsecase usecase.IStockUsecase
	ListRoutes   []models.HTTPRoute
}

func InitGate(stockUsecase usecase.IStockUsecase) *Gate {
	gate := &Gate{
		StockUsecase: stockUsecase,
	}
	gate.InitRoutes()
	return gate
}

func (c *Gate) Get(path string, handl http.HandlerFunc) {
	fPath := strings.ToLower(strings.TrimSpace(path))
	newHttpRoute := models.HTTPRoute{
		Method:  http.MethodGet,
		Path:    fPath,
		Handler: handl,
	}
	listRoutes := append(c.ListRoutes, newHttpRoute)
	c.ListRoutes = listRoutes
}

func (c *Gate) Post(path string, handl http.HandlerFunc) {
	fPath := strings.ToLower(strings.TrimSpace(path))
	newHttpRoute := models.HTTPRoute{
		Method:  http.MethodPost,
		Path:    fPath,
		Handler: handl,
	}
	listRoutes := append(c.ListRoutes, newHttpRoute)
	c.ListRoutes = listRoutes
}

func (c *Gate) Put(path string, handl http.HandlerFunc) {
	fPath := strings.ToLower(strings.TrimSpace(path))
	newHttpRoute := models.HTTPRoute{
		Method:  http.MethodPut,
		Path:    fPath,
		Handler: handl,
	}
	listRoutes := append(c.ListRoutes, newHttpRoute)
	c.ListRoutes = listRoutes
}

func (c *Gate) Delete(path string, handl http.HandlerFunc) {
	fPath := strings.ToLower(strings.TrimSpace(path))
	newHttpRoute := models.HTTPRoute{
		Method:  http.MethodDelete,
		Path:    fPath,
		Handler: handl,
	}
	listRoutes := append(c.ListRoutes, newHttpRoute)
	c.ListRoutes = listRoutes
}
