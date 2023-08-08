package apis

import (
	"github.com/ikhsan892/goceng"
	"github.com/labstack/echo/v4"
)

type OrderAPI struct {
	app goceng.App
}

func NewOrderAPI(app goceng.App, e *echo.Echo) {
	o := &OrderAPI{app: app}

	e.GET("/create-order", o.CreateOrder)
}

func (o OrderAPI) CreateOrder(c echo.Context) error {
	return c.String(200, "test")
}
