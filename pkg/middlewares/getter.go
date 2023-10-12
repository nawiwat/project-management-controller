package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// Getter use to embeded context get/set logic
type Getter struct{}

// SetLogger to set logger
func (g Getter) SetLogger(c echo.Context, logger *logrus.Entry) {
	c.Set("logger", logger)
}

// GetLogger to get logger
func (g Getter) GetLogger(c echo.Context) *logrus.Entry {
	logger, ok := c.Get("logger").(*logrus.Entry)
	if !ok {
		return logrus.NewEntry(logrus.New())
	}

	return logger
}
