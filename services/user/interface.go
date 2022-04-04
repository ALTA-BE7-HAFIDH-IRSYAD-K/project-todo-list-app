package user

import "project-todo-list/entities"

type UserServiceInterface interface {
	CreateUser(user entities.User) (entities.User, error)
	FindUserByID(id int) (entities.User, error)
	UpdateUser(user entities.User, id int) (entities.User, error)
	FindUserAll() ([]entities.User, error)
	DeleteUser(id int) error
}
