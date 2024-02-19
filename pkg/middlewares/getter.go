package middlewares

import (
	"app-controller/pkg/errors"
	"app-controller/pkg/model"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/dgrijalva/jwt-go"
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

func Auth(c echo.Context) (*model.UserToken , error) {
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return nil , errors.NewError(echo.ErrBadRequest.Code, errors.ErrBadRequest, "Authorization header is missing",nil,nil)
	}
	tokenString := strings.Split(authHeader, " ")[1]
	if tokenString == "" {
		return nil , errors.NewError(echo.ErrBadRequest.Code, errors.ErrBadRequest, "Token is missing",nil,nil)
	}

	token, err := jwt.ParseWithClaims(tokenString, &model.UserToken{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("notdoingthatnowlol"), nil
	})

	if err != nil {
		return nil , errors.NewError(echo.ErrUnauthorized.Code, errors.ErrCodeValidationFail, "Invalid token",nil,err)
	}

	if !token.Valid {
		return nil , errors.NewError(echo.ErrUnauthorized.Code, errors.ErrCodeValidationFail, "Invalid token",nil,nil)
	}

	claims , ok := token.Claims.(*model.UserToken)
	if !ok {
		return nil , errors.NewError(echo.ErrInternalServerError.Code, errors.ErrCodeInternalError, "Failed to extract claims from token",nil,nil)
	}

	return claims , nil
}
