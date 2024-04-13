package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/soloviev1d/avito-banner-service/internal/cache"
	"github.com/soloviev1d/avito-banner-service/internal/database"
	"github.com/soloviev1d/avito-banner-service/internal/structs"
)

const (
	unauth = iota
	user
	admin
)

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

	if len(token) != 36 {
		return c.JSON(
			http.StatusBadRequest,
			structs.ErrorResponse{
				Error: fmt.Sprintf(
					"Токен должен иметь длину 36 символов, получено %d",
					len(token),
				),
			},
		)
	}

	accessLevel := cache.UserAccessLevel(token)
	if cache.UserAccessLevel(token) == unauth {
		return c.String(http.StatusUnauthorized, "Пользователь не авторизован")
	}
	if accessLevel == admin {
		if hotData {
			banner := database.GetUserBanner(tagId, featureId)
			if banner == nil {
				return c.String(http.StatusNotFound, "Баннер для пользователя не найден")
			}
			cache.BannerCache.Data[cache.AssembleCacheKey(tagId, featureId)] = banner
			return c.JSON(http.StatusOK, banner.ToBanner())
		} else {
			banner, ok := cache.BannerCache.Data[cache.AssembleCacheKey(tagId, featureId)]
			if !ok {
				return c.String(http.StatusNotFound, "Баннер для пользователя не найден")
			}
			return c.JSON(http.StatusOK, banner.ToBanner())
		}
	} else {
		coldBanner, ok := cache.BannerCache.Data[cache.AssembleCacheKey(tagId, featureId)]
		if !ok && !hotData {
			return c.String(http.StatusNotFound, "Баннер для пользователя не найден")
		}
		if !hotData && coldBanner.IsActive {
			return c.JSON(http.StatusOK, coldBanner.ToBanner())
		}
		if !hotData && !coldBanner.IsActive {
			return c.String(http.StatusForbidden, "Пользователь не имеет доступа")
		}
		if hotData {
			banner := database.GetUserBanner(tagId, featureId)
			if banner.IsActive {
				return c.JSON(http.StatusOK, banner.ToBanner())
			} else {
				return c.String(http.StatusForbidden, "Пользователь не имеет доступа")
			}
		}
	}
	return c.String(http.StatusInternalServerError, "Unreachable")
}
