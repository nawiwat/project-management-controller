package handlers

import (
	// "encoding/json"
	// "strings"
	// "time"

	//"app-controller/pkg/errors"
	"app-controller/pkg/middlewares"
	"app-controller/pkg/model"

	"github.com/labstack/echo/v4"
)

func (r *AppController) EditTask(c echo.Context) error {
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


	//err := r.controllerService.AddBoardColumn(c.Request().Context(), col)

	// if err != nil {
	// 	return errors.NewError(echo.ErrInternalServerError.Code, errors.ErrCodeInternalError, "", col , err)
	// }

	return c.JSON(200, map[string]interface{}{
		"message": "success",
	})
}