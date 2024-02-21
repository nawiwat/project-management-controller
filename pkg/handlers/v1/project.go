package handlers

import (
	// "encoding/json"
	// "strings"
	// "time"
	"app-controller/pkg/middlewares"
	"github.com/labstack/echo/v4"
	"app-controller/pkg/errors"
	"app-controller/pkg/model"
)

func (r *AppController) GetProjects(c echo.Context) error {
	usr , err := middlewares.Auth(c)
	if err != nil {
		return err
	}

	f, err := r.controllerService.GetProjects(c.Request().Context(),usr.Username)

	if err != nil {
		return err
	}

	return c.JSON(200, map[string]interface{}{"data": f})
}

func (r *AppController) GetProjectInfo(c echo.Context) error {
	_ , err := middlewares.Auth(c)
	if err != nil {
		return err
	}
	var pj model.Project

	if err := c.Bind(&pj); err != nil {
		return err
	}
	if err := c.Validate(&pj); err != nil {
		return err
	}

	f, err := r.controllerService.GetProjectInfo(c.Request().Context(),pj.ID)

	if err != nil {
		return err
	}


	return c.JSON(200, map[string]interface{}{"data": f})
}

func (r *AppController) AddProject(c echo.Context) error {
	user , err := middlewares.Auth(c)
	if err != nil {
		return err
	}

	var prj model.Project

	if err := c.Bind(&prj); err != nil {
		return err
	}

	if err := c.Validate(&prj); err != nil {
		return err
	}

	err = r.controllerService.AddProject(c.Request().Context(), prj,user.Username)

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

func (r *AppController) EditProject(c echo.Context) error {
	_ , err := middlewares.Auth(c)
	if err != nil {
		return err
	}

	var prj model.Project

	if err := c.Bind(&prj); err != nil {
		return err
	}

	if err := c.Validate(&prj); err != nil {
		return err
	}

	err = r.controllerService.EditProject(c.Request().Context(), prj)

	if err != nil {
		return errors.NewError(echo.ErrInternalServerError.Code, errors.ErrCodeInternalError, "", prj , err)
	}

	return c.JSON(200, map[string]interface{}{
		"message": "success",
	})
}

func (r *AppController) DeleteProject(c echo.Context) error {
	_ , err := middlewares.Auth(c)
	if err != nil {
		return err
	}

	var prj model.Project

	if err := c.Bind(&prj); err != nil {
		return err
	}

	if err := c.Validate(&prj); err != nil {
		return err
	}

	err = r.controllerService.DeleteProject(c.Request().Context(), prj.ID)

	if err != nil {
		return errors.NewError(echo.ErrInternalServerError.Code, errors.ErrCodeInternalError, "", nil , err)
	}

	return c.JSON(200, map[string]interface{}{
		"message": "success",
	})
}