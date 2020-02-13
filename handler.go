package main

import (
	"ft/ft"
	"net/http"

	"github.com/labstack/echo"
)

func UserShowHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		client, err := ft.GetConnection(ctx)
		if err != nil {
			return c.JSON(http.StatusOK, map[string]interface{}{`err`: err})
		}

		userID := c.Param("userID")
		user, err := ft.GetUser(ctx, client, userID)
		if err != nil {
			return c.JSON(http.StatusOK, map[string]interface{}{`err`: err})
		}
		return c.JSON(http.StatusOK, user)
	}
}
