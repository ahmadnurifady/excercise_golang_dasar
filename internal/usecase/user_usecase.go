package usecase

import (
	"latihan-solid/internal/domain"
	"latihan-solid/internal/repository"
)

type UserUsecaseInterface interface {
	SaveUser
	FindAllUser
	FindUserByName
	DeleteUser
	UpdateUser
}

type UpdateUser interface {
	Update(userID int, updateData *domain.User) (*domain.User, error)
}

type SaveUser interface {
	Save(userRequest domain.User) (domain.User, error)
}

type FindAllUser interface {
	FindAll() ([]domain.User, error)
}

type FindUserByName interface {
	FindUserByName(userName string) (domain.User, error)
}

type DeleteUser interface {
	DeleteUser(userId int) (string, error)
}

type UserUsecase struct {
	repo repository.UserRepositoryInterface
}

// DeleteUser implements UserUsecaseInterface.
func (uc *UserUsecase) DeleteUser(userId int) (string, error) {
	msg, err := uc.repo.DeleteUser(userId)
	if err != nil {
		return "", err
	}
	return msg, nil
}

// FindAll implements UserUsecaseInterface.
func (uc *UserUsecase) FindAll() ([]domain.User, error) {
	users, err := uc.repo.FindAll()
	if err != nil {
		return []domain.User{}, err
	}

	return users, nil
}

// FindUserById implements UserUsecaseInterface.
func (uc *UserUsecase) FindUserByName(userName string) (domain.User, error) {
	user, err := uc.repo.FindUserByName(userName)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

// Save implements UserUsecaseInterface.
func (uc *UserUsecase) Save(userRequest domain.User) (domain.User, error) {
	book, err := uc.repo.Save(&userRequest)
	if err != nil {
		return domain.User{}, err
	}

	return *book, nil
}

// Update implements UserUsecaseInterface.
func (uc *UserUsecase) Update(userID int, updateData *domain.User) (*domain.User, error) {
	panic("unimplemented")
}

func NewUserUsecase(repo repository.UserRepositoryInterface) UserUsecaseInterface {
	return &UserUsecase{
		repo: repo}
}
