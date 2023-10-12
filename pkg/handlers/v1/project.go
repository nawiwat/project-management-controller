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
	"app-controller/pkg/model"
)

func (r *AppController) GetProjects(c echo.Context) error {
	f, err := r.controllerService.GetProjects(c.Request().Context())

	if err != nil {
		return err
	}

	return c.JSON(200, map[string]interface{}{"data": f})
}

func (r *AppController) AddProject(c echo.Context) error {
	var prj model.Project

	if err := c.Bind(&prj); err != nil {
		return err
	}

	if err := c.Validate(&prj); err != nil {
		return err
	}

	err := r.controllerService.AddProject(c.Request().Context(), prj)

	if err != nil {
		return errors.NewError(echo.ErrInternalServerError.Code, errors.ErrCodeInternalError, "", prj , err)
	}

	return c.JSON(200, map[string]interface{}{
		"message": "success",
	})
}

func (r *AppController) AddMember(c echo.Context) error {
	var mem model.Membership

	if err := c.Bind(&mem); err != nil {
		return err
	}

	if err := c.Validate(&mem); err != nil {
		return err
	}

	err := r.controllerService.AddMember(c.Request().Context(), mem)

	if err != nil {
		return errors.NewError(echo.ErrInternalServerError.Code, errors.ErrCodeInternalError, "", mem , err)
	}

	return c.JSON(200, map[string]interface{}{
		"message": "success",
	})
}