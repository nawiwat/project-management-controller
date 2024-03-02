package handlers

import (
	// "encoding/json"
	// "strings"
	// "time"

	"app-controller/pkg/errors"
	"app-controller/pkg/middlewares"
	"app-controller/pkg/model"

	"github.com/labstack/echo/v4"
)

func (r *AppController) CreateTask (c echo.Context) error {
	_ , err := middlewares.Auth(c)
	if err != nil {
		return err
	}

	var tsk model.Task
	if err := c.Bind(&tsk); err != nil {
		return err
	}
	if err := c.Validate(&tsk); err != nil {
		return err
	}

	out , err := r.controllerService.CreateTask(c.Request().Context(), tsk)

	if err != nil {
		return errors.NewError(echo.ErrInternalServerError.Code, errors.ErrCodeInternalError, "", nil , err)
	}

	return c.JSON(200, map[string]interface{}{"data": out})
}

func (r *AppController) GetTask (c echo.Context) error {
	_ , err := middlewares.Auth(c)
	if err != nil {
		return err
	}

	var tsk model.Task
	if err := c.Bind(&tsk); err != nil {
		return err
	}
	if err := c.Validate(&tsk); err != nil {
		return err
	}

	out , err := r.controllerService.GetTask(c.Request().Context(), tsk.ProjectId)

	if err != nil {
		return errors.NewError(echo.ErrInternalServerError.Code, errors.ErrCodeInternalError, "", nil , err)
	}

	return c.JSON(200, map[string]interface{}{"data": out})
}

func (r *AppController) UpdateTask(c echo.Context) error {
	_ , err := middlewares.Auth(c)
	if err != nil {
		return err
	}

	var tsk []model.Task
	if err := c.Bind(&tsk); err != nil {
		return err
	}
	// if err := c.Validate(&tsk); err != nil {
	// 	return err
	// }

	out , err := r.controllerService.UpdateTask(c.Request().Context(), tsk)

	if err != nil {
		return errors.NewError(echo.ErrInternalServerError.Code, errors.ErrCodeInternalError, "", nil , err)
	}

	return c.JSON(200, map[string]interface{}{"data": out})
}

func (r *AppController) DeleteTask(c echo.Context) error {
	_ , err := middlewares.Auth(c)
	if err != nil {
		return err
	}

	var tsk model.Task
	if err := c.Bind(&tsk); err != nil {
		return err
	}
	if err := c.Validate(&tsk); err != nil {
		return err
	}

	err = r.controllerService.DeleteTask(c.Request().Context(), tsk.ID)

	if err != nil {
		return errors.NewError(echo.ErrInternalServerError.Code, errors.ErrCodeInternalError, "", nil , err)
	}

	return c.JSON(200, map[string]interface{}{
		"message": "success",
	})
}