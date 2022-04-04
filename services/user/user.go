package user

import (
	"errors"
	_helper "project-todo-list/delivery/helper"
	"project-todo-list/entities"
	_userRepository "project-todo-list/repository/user"
)

type UserService struct {
	userRepository _userRepository.UserRepositoryInterface
}

func NewUserService(userRepo _userRepository.UserRepositoryInterface) UserServiceInterface {
	return &UserService{
		userRepository: userRepo,
	}
}

func (u UserService) CreateUser(user entities.User) (entities.User, error) {
	//TODO implement me
	password, err := _helper.HashPassword(user.Password)

	if err != nil {
		return user, err
	}

	user.Password = password
	newUser, err := u.userRepository.CreateUser(user)

	if err != nil {
		return newUser, nil
	}
	return newUser, nil
}

func (u UserService) FindUserByID(id int) (entities.User, error) {
	//TODO implement me
	user, err := u.userRepository.FindUserByID(id)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("no user found on with that ID")
	}

	return user, nil
}

func (u UserService) UpdateUser(user entities.User, id int) (entities.User, error) {
	//TODO implement me
	userId, err := u.userRepository.FindUserByID(id)

	password, err := _helper.HashPassword(user.Password)

	userId.Name = user.Name
	userId.Email = user.Email
	userId.Password = password

	if err != nil {
		return userId, err
	}

	userUpdate, err := u.userRepository.UpdateUser(userId)

	if err != nil {
		return userUpdate, err
	}

	return userUpdate, nil
}

func (u UserService) FindUserAll() ([]entities.User, error) {
	//TODO implement me
	users, err := u.userRepository.FindUserAll()
	if err != nil {
		return users, err
	}

	return users, nil
}

func (u UserService) DeleteUser(id int) error {
	//TODO implement me
	err := u.userRepository.DeleteUser(id)

	if err != nil {
		return err
	}

	return nil
}
