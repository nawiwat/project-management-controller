package handlers

import (
	// inputs "acw-crypto-risk-management/pkg/inputs/riskmgmt"
	// "acw-crypto-risk-management/pkg/middlewares"
	// "acw-crypto-risk-management/pkg/requests"
	// "acw-crypto-risk-management/pkg/services/riskmgmt"
	// "encoding/json"
	// "strings"
	// "time"

	"github.com/labstack/echo/v4"
	"app-controller/pkg/errors"
	"app-controller/pkg/services/contlr"
	"app-controller/pkg/model"
)

type AppController struct {
	controllerService contlr.ControllerService
}

func NewAppController(
	controllerService contlr.ControllerService,
) *AppController {
	return &AppController{
		controllerService,
	}
}

func (r *AppController) GetUsers(c echo.Context) error {
	f, err := r.controllerService.GetUsers(c.Request().Context())

	if err != nil {
		return err
	}

	return c.JSON(200, map[string]interface{}{"data": f})
}

func (r *AppController) AddUser(c echo.Context) error {
	var usr model.User

	if err := c.Bind(&usr); err != nil {
		return err
	}

	if err := c.Validate(&usr); err != nil {
		return err
	}

	err := r.controllerService.AddUser(c.Request().Context(), usr)

	if err != nil {
		return errors.NewError(echo.ErrInternalServerError.Code, errors.ErrCodeInternalError, "", usr, err)
	}

	return c.JSON(200, map[string]interface{}{
		"message": "success",
	})
}