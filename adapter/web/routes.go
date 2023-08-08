package web

import (
	"github.com/ikhsan892/goceng"
	"github.com/ikhsan892/goceng/adapter/web/apis"
	"github.com/labstack/echo/v4"
)

func initRoutes(e *echo.Echo, app goceng.App) {
	apis.NewOrderAPI(app, e)

}
