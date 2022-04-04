package routes

import (
	"github.com/labstack/echo/v4"
	_handler "project-todo-list/delivery/handler"
)

func UserPath(e *echo.Echo, uh *_handler.UserHandler) {
	e.POST("/api/v1/users", uh.CreateUserHandler())
	e.GET("/api/v1/users", uh.GetAllUserHandler())
	e.GET("/api/v1/users/:id", uh.GetUserByIdHandler())
	e.PUT("/api/v1/users/:id", uh.UpdateUserHandler())
	e.DELETE("/api/v1/users/:id", uh.DeleteUserHandler())
}
