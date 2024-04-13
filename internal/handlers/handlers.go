package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/soloviev1d/avito-banner-service/internal/structs"
)

type Banner struct {
	tagId     int
	featureId int
}

func GetUserBanner(c echo.Context) error {
	var (
		tagIdParam     = c.QueryParam("tag_id")
		featureIdParam = c.QueryParam("feature_id")
		hotDataParam   = c.QueryParam("use_last_revision")
		token          = c.Request().Header.Get("token")
	)

	tagId, err := strconv.Atoi(tagIdParam)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			structs.NewInvalidType(featureIdParam, err),
		)
	}
	featureId, err := strconv.Atoi(featureIdParam)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			structs.NewInvalidType(featureIdParam, err),
		)
	}
	hotData, err := strconv.ParseBool(hotDataParam)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			structs.NewInvalidType(hotDataParam, err),
		)
	}
}
