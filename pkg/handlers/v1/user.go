package handlers

import (
	// inputs "acw-crypto-risk-management/pkg/inputs/riskmgmt"
	// "acw-crypto-risk-management/pkg/middlewares"
	// "acw-crypto-risk-management/pkg/requests"
	// "acw-crypto-risk-management/pkg/services/riskmgmt"
	// "encoding/json"
	// "strings"
	"app-controller/pkg/errors"
	"app-controller/pkg/middlewares"
	"app-controller/pkg/model"
	"app-controller/pkg/services/contlr"
	"encoding/base64"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
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
	user , err := middlewares.Auth(c)
	if err != nil {
		return err
	}

	f, err := r.controllerService.GetUser(c.Request().Context(),user.Username)

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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(usr.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	hashedPasswordString := base64.StdEncoding.EncodeToString(hashedPassword)
	usr.Password = hashedPasswordString

	token , err := r.controllerService.AddUser(c.Request().Context(), usr)

	if err != nil {
		return errors.NewError(echo.ErrBadRequest.Code, errors.ErrBadRequest, err.Error(),nil,err)
	}

	return c.JSON(200, map[string]interface{}{
		"token": token,
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

	tokenString , err := r.controllerService.Login(c.Request().Context(), usr)

	if err != nil {
		return errors.NewError(echo.ErrBadRequest.Code, errors.ErrBadRequest, err.Error(),nil,err)
	}

	return c.JSON(200, map[string]interface{}{
		"token": tokenString,
	})
}

func (r *AppController) Authentication(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return errors.NewError(echo.ErrBadRequest.Code, errors.ErrBadRequest, "Authorization header is missing",nil,nil)
	}
	tokenString := strings.Split(authHeader, " ")[1]
	if tokenString == "" {
		return errors.NewError(echo.ErrBadRequest.Code, errors.ErrBadRequest, "Token is missing",nil,nil)
	}

	token, err := jwt.ParseWithClaims(tokenString, &model.UserToken{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("notdoingthatnowlol"), nil
	})

	if err != nil {
		return errors.NewError(echo.ErrUnauthorized.Code, errors.ErrCodeValidationFail, "Invalid token",nil,err)
	}

	if !token.Valid {
		return errors.NewError(echo.ErrUnauthorized.Code, errors.ErrCodeValidationFail, "Invalid token",nil,nil)
	}

	claims, ok := token.Claims.(*model.UserToken)
	if !ok {
		return errors.NewError(echo.ErrInternalServerError.Code, errors.ErrCodeInternalError, "Failed to extract claims from token",nil,nil)
	}


	message := "welcome, " + claims.Username

	return c.JSON(200, map[string]interface{}{
		"message": message,
		"info":claims,
	})
}

func (r *AppController) EditUser(c echo.Context) error {
	user , err := middlewares.Auth(c)
	if err != nil {
		return err
	}

	var usr model.User

	if err := c.Bind(&usr); err != nil {
		return err
	}
	if err := c.Validate(&usr); err != nil {
		return err
	}

	err = r.controllerService.EditUser(c.Request().Context(), usr,user.Username)

	if err != nil {
		return errors.NewError(echo.ErrInternalServerError.Code, errors.ErrCodeInternalError, "", nil , err)
	}

	return c.JSON(200, map[string]interface{}{
		"message": "success",
	})
}

func (r *AppController) EditProfile(c echo.Context) error {
	user , err := middlewares.Auth(c)
	if err != nil {
		return err
	}

	var ath model.ProfileAttachment

	if err := c.Bind(&ath); err != nil {
		return err
	}
	if err := c.Validate(&ath); err != nil {
		return err
	}

	err = r.controllerService.EditProfile(c.Request().Context(), ath, user.Username)

	if err != nil {
		return errors.NewError(echo.ErrInternalServerError.Code, errors.ErrCodeInternalError, "", nil , err)
	}

	return c.JSON(200, map[string]interface{}{
		"message": "success",
	})
}

func (r *AppController) InviteResponse(c echo.Context) error {
	usr , err := middlewares.Auth(c)
	if err != nil {
		return err
	}

	var rsp model.InviteResponse

	if err := c.Bind(&rsp); err != nil {
		return err
	}
	if err := c.Validate(&rsp); err != nil {
		return err
	}
	
	rsp.Respondent = usr.Username

	err = r.controllerService.InviteResponse(c.Request().Context(), rsp)

	if err != nil {
		return errors.NewError(echo.ErrInternalServerError.Code, errors.ErrCodeInternalError, "", err.Error() , err)
	}

	return c.JSON(200, map[string]interface{}{
		"message": "success",
	})
}