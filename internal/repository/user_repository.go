package repository

import (
	"fmt"
	"latihan-solid/internal/domain"
)

type UserRepositoryInterface interface {
	SaveUser
	FindAllUser
	FindUserById
	DeleteUser
	UpdateUser
}

type UpdateUser interface {
	Update(userID int, updateData *domain.User) (*domain.User, error)
}

type SaveUser interface {
	Save(userRequest *domain.User) (*domain.User, error)
}

type FindAllUser interface {
	FindAll() ([]domain.User, error)
}

type FindUserById interface {
	FindUserById(userId int) (domain.User, error)
}

type DeleteUser interface {
	DeleteUser(userId int) (string, error)
}

type UserRepository struct {
	db map[int]domain.User
}

// Save implements UserRepositoryInterface.
func (repo *UserRepository) Save(userRequest *domain.User) (*domain.User, error) {
	if _, exists := repo.db[userRequest.Id]; exists {
		return nil, fmt.Errorf("user already exists")
	}

	repo.db[userRequest.Id] = *userRequest
	return userRequest, nil
}

// FindAll implements UserRepositoryInterface.
func (repo *UserRepository) FindAll() ([]domain.User, error) {
	var allUsers []domain.User

	for _, book := range repo.db {
		allUsers = append(allUsers, book)
	}

	return allUsers, nil
}

// FindUserById implements UserRepositoryInterface.
func (repo *UserRepository) FindUserById(userId int) (domain.User, error) {
	if _, exist := repo.db[userId]; !exist {
		return domain.User{}, fmt.Errorf("User already exists")
	}
	var user domain.User
	for _, val := range repo.db {
		if val.Id == userId {
			user = val
		}
	}
	return user, nil
}

// Update implements UserRepositoryInterface.
func (repo *UserRepository) Update(userID int, updateData *domain.User) (*domain.User, error) {
	panic("unimplemented")
}

// DeleteUser implements UserRepositoryInterface.
func (repo *UserRepository) DeleteUser(userId int) (string, error) {
	if _, exist := repo.db[userId]; !exist {
		return "", fmt.Errorf("User already exists")
	}

	delete(repo.db, userId)

	return "Success Delete User", nil
}

func NewUserRepository() UserRepositoryInterface {
	return &UserRepository{
		db: make(map[int]domain.User)}
}
