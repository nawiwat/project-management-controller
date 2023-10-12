package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

const (
	contextName = "app-controller"
)

func registerHealthCheckRouteV1(
	c *echo.Echo,
) {
	c.GET(fmt.Sprintf("/%s/info", contextName), func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"context": contextName,
			"build": map[string]interface{}{
				"version": os.Getenv("APP_VERSION"),
			},
		})
	})

	c.GET(fmt.Sprintf("/%s/health", contextName), func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status": "UP",
		})
	})
}
