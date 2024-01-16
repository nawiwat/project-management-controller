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

func (r *AppController) GetUser(c echo.Context) error {
	var usr model.User
	if err := c.Bind(&usr); err != nil {
		return err
	}
	if err := c.Validate(&usr); err != nil {
		return err
	}

	f, err := r.controllerService.GetUser(c.Request().Context(),usr.ID)

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


func (r *AppController) Login(c echo.Context) error {
	var usr model.UserLogin

	if err := c.Bind(&usr); err != nil {
		return err
	}

	if err := c.Validate(&usr); err != nil {
		return err
	}

	var urtoken model.UserToken

	urtoken.Token = "wVYrxaeNa9OxdnULvde1Au5m5w63"

	if usr.Email == "user1@themenate.net" && usr.Password == "2005ipo" {
		return c.JSON(200, map[string]interface{}{
			"data": urtoken,
		})
	}

	return c.JSON(400, map[string]interface{}{
		"message": "user_name or password is invalid",
	})
}

func (r *AppController) Authentication(c echo.Context) error {
	var usr model.UserToken

	usr.Token = c.Request().Header.Get("Authorization")

	if usr.Token == "Bearer wVYrxaeNa9OxdnULvde1Au5m5w63" {
		return c.JSON(200, map[string]interface{}{
			"message": "success",
		})
	}

	return c.JSON(401, map[string]interface{}{
		"message": "invalid token",
	})
}
