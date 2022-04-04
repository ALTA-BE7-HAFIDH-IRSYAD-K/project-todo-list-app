package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	_config "project-todo-list/config"
	_userHandler "project-todo-list/delivery/handler"
	"project-todo-list/delivery/routes"
	"project-todo-list/driver"
	_userRepo "project-todo-list/repository/user"
	_userService "project-todo-list/services/user"
)

func main() {
	config := _config.GetConfig()
	db := driver.InitDB(config)

	userRepo := _userRepo.UserNewRepository(db)
	userService := _userService.NewUserService(userRepo)
	userHandler := _userHandler.NewUserHandler(userService)

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())

	routes.UserPath(e, userHandler)

	log.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))

}
