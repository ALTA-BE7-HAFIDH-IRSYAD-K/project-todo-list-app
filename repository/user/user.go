package user

import (
	"gorm.io/gorm"
	"project-todo-list/entities"
)

type UserRepository struct {
	db *gorm.DB
}

func UserNewRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) CreateUser(user entities.User) (entities.User, error) {
	err := ur.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (ur *UserRepository) FindUserAll() ([]entities.User, error) {
	var users []entities.User
	tx := ur.db.Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return users, nil
}

func (ur *UserRepository) FindUserByID(id int) (entities.User, error) {
	var user entities.User
	tx := ur.db.Find(&user, id)
	if tx.Error != nil {
		return user, tx.Error
	}
	if tx.RowsAffected == 0 {
		return user, nil
	}
	return user, nil
}

func (ur *UserRepository) UpdateUser(user entities.User) (entities.User, error) {
	//TODO implement me
	tx := ur.db.Save(&user)

	if tx.Error != nil {
		return user, tx.Error
	}
	if tx.RowsAffected == 0 {
		return user, nil
	}

	return user, nil
}

func (ur *UserRepository) DeleteUser(id int) error {
	tx := ur.db.Delete(&entities.User{}, id)

	if tx.RowsAffected <= 0 {
		return tx.Error
	}

	return nil
}
