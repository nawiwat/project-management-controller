package server

import (
	"app-controller/pkg/handlers/v1"
	"app-controller/pkg/services/contlr"
	"fmt"

	"github.com/labstack/echo/v4"
)

func registerRouterV1(
	c *echo.Echo,
	controllerService contlr.ControllerService,
) {
	g := c.Group(fmt.Sprintf("/%s/v1", contextName))

	registerControllerV1(g, controllerService)
}

func registerControllerV1(
	c *echo.Group,
	controllerService contlr.ControllerService,
) {
	h := handlers.NewAppController(controllerService)
    
	c.POST("/login",h.Login)
	c.GET("/",h.Authentication)

	c.GET("/users", h.GetUsers)
	c.GET("/user", h.GetUser)
	c.POST("/user", h.AddUser)
	c.POST("/user/edit",h.EditUser)
	c.POST("/user/profile",h.EditProfile)
	//c.POST("/user/invite",h.UserInvitation)

	c.GET("/projects",h.GetProjects)
	c.GET("/project",h.GetProjectInfo)
	c.POST("/project",h.AddProject)
	c.POST("/project/edit",h.EditProject)
	c.DELETE("/project",h.DeleteProject)

	c.POST("/member",h.AddMember)
	//c.POST("/projects/invite",h.InviteUser)


	//c.POST("/column",h.AddBoardColumn)

}
