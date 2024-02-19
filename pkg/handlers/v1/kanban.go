package handlers

import (
	// "encoding/json"
	// "strings"
	// "time"

	"github.com/labstack/echo/v4"
	"app-controller/pkg/errors"
	"app-controller/pkg/model"
)

func (r *AppController) AddBoardColumn(c echo.Context) error {
	var col model.BoardColumn
	if err := c.Bind(&col); err != nil {
		return err
	}
	if err := c.Validate(&col); err != nil {
		return err
	}


	err := r.controllerService.AddBoardColumn(c.Request().Context(), col)

	if err != nil {
		return errors.NewError(echo.ErrInternalServerError.Code, errors.ErrCodeInternalError, "", col , err)
	}

	return c.JSON(200, map[string]interface{}{
		"message": "success",
	})
}