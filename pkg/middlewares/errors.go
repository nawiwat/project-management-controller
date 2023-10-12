package middlewares

import (
	"app-controller/pkg/errors"

	"github.com/labstack/echo/v4"
)

type httpResponse struct {
	Code   string      `json:"code"`
	Reason string      `json:"reason,omitempty"`
	Fields interface{} `json:"fields,omitempty"`
}

func httpErrorHandling(err error, c echo.Context) {
	logger := Getter{}.GetLogger(c)
	appErr, ok := errors.AssertAppError(err)

	logger.Infof(appErr.Error())

	if ok {
		if err := c.JSON(appErr.HTTPCode, httpResponse{
			Code:   string(appErr.Code),
			Reason: appErr.Reason,
			Fields: appErr.Fields,
		}); err != nil {
			c.Logger().Errorf("fail to construct response from error: %s", err.Error())
		}
	} else {
		if err := c.JSON(appErr.HTTPCode, httpResponse{
			Code:   string(appErr.Code),
			Reason: err.Error(),
			Fields: appErr.Fields,
		}); err != nil {
			c.Logger().Errorf("fail to construct response from error: %s", err.Error())
		}
	}
}

// RegisterHTTPErrorHandling to register handler
func RegisterHTTPErrorHandling(c *echo.Echo) {
	c.HTTPErrorHandler = httpErrorHandling
}
