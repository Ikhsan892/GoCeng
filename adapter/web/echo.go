package web

import (
	"github.com/ikhsan892/goceng"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"net/http"
)

type EchoWebAdapter struct {
	ec  *echo.Echo
	app goceng.App
}

func NewEcho(app goceng.App) *EchoWebAdapter {
	e := echo.New()
	logger := app.ZapLogger()

	e.Debug = false

	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info("request",
				zap.String("URI", v.URI),
				zap.Int("status", v.Status),
			)

			return nil
		},
	}))
	e.Use(middleware.RequestID())
	e.Use(middleware.Gzip())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"localhost"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"Content-Disposition"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderAuthorization, echo.HeaderContentType, "module", "Content-Range", "Accept-Language"},
	}))
	initRoutes(e)

	return &EchoWebAdapter{ec: e, app: app}
}

func (e *EchoWebAdapter) Init() error {
	s := http.Server{
		Addr:    ":3008",
		Handler: e.ec, // set Echo as handler
		//ReadTimeout: 30 * time.Second, // use custom timeouts
	}

	return s.ListenAndServe()
}
