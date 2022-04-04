package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"project-todo-list/delivery/helper"
	"project-todo-list/entities"
	_userService "project-todo-list/services/user"
	"strconv"
)

type UserHandler struct {
	userService _userService.UserServiceInterface
}

func NewUserHandler(userService _userService.UserServiceInterface) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (uh *UserHandler) CreateUserHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var userData entities.User

		err := c.Bind(&userData)
		if err != nil {
			c.JSON(http.StatusBadRequest, nil)
		}

		user, err := uh.userService.CreateUser(userData)

		if err != nil {
			c.JSON(http.StatusBadRequest, nil)
		}

		return c.JSON(http.StatusOK, helper.ResponseSuccess("success create user", user))
	}
}

func (uh *UserHandler) GetAllUserHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		users, err := uh.userService.FindUserAll()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get all users", users))
	}
}

func (uh *UserHandler) GetUserByIdHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("id not recognise"))
		}
		users, err := uh.userService.FindUserByID(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to fetch data"))
		}
		if users.ID == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("data not found"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get users by id", users))
	}
}

func (uh *UserHandler) UpdateUserHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var userData entities.User

		err := c.Bind(&userData)

		if err != nil {
			fmt.Println("err is", err)
		}

		idStr := c.Param("id")
		id, _ := strconv.Atoi(idStr)

		user, err := uh.userService.UpdateUser(userData, id)

		if err != nil {
			c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}

		return c.JSON(http.StatusOK, helper.ResponseSuccess("success update user", user))
	}
}

func (uh *UserHandler) DeleteUserHandler() echo.HandlerFunc {
	return func(c echo.Context) error {

		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("id not recognise"))
		}

		errDelete := uh.userService.DeleteUser(id)

		if errDelete != nil {
			c.JSON(http.StatusBadRequest, helper.ResponseSuccessWithoutData("failed delete user"))
		}

		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("success delete user"))
	}
}
