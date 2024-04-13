package server

import (
	"github.com/labstack/echo/v4"
	"github.com/soloviev1d/avito-banner-service/internal/handlers"
)

func ListenAndServe() {
	e := echo.New()
	e.GET("/user_banner", handlers.GetUserBanner)
	e.Logger.Fatal(e.Start(":8080"))
}
