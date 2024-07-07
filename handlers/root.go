package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Root(c echo.Context) error {
	return c.Render(http.StatusOK, "root", map[string]interface{}{
		"name": "Your Name",
	})
}
