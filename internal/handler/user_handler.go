package handler

import (
	"latihan-solid/internal/domain"
	"latihan-solid/internal/usecase"
)

type UserHandlerInterface interface {
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
	Save(userRequest domain.User) (domain.User, error)
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

type UserHandler struct {
	uc usecase.UserUsecaseInterface
}

// DeleteUser implements UserHandlerInterface.
func (h *UserHandler) DeleteUser(userId int) (string, error) {
	user, err := h.uc.DeleteUser(userId)
	if err != nil {
		return user, err
	}

	return user, nil
}

// FindAll implements UserHandlerInterface.
func (h *UserHandler) FindAll() ([]domain.User, error) {
	users, err := h.uc.FindAll()
	if err != nil {
		return users, err
	}

	return users, nil
}

// FindUserById implements UserHandlerInterface.
func (h *UserHandler) FindUserById(userId int) (domain.User, error) {
	user, err := h.uc.FindUserById(userId)
	if err != nil {
		return user, err
	}

	return user, nil
}

// Save implements UserHandlerInterface.
func (h *UserHandler) Save(userRequest domain.User) (domain.User, error) {
	user, err := h.uc.Save(userRequest)
	if err != nil {
		return user, err
	}

	return user, nil
}

// Update implements UserHandlerInterface.
func (h *UserHandler) Update(userID int, updateData *domain.User) (*domain.User, error) {
	panic("unimplemented")
}

func NewUserHandler(uc usecase.UserUsecaseInterface) UserHandlerInterface {
	return &UserHandler{
		uc: uc}
}
