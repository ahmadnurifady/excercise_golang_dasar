package handler

import (
	"latihan-solid/internal/domain"
	"latihan-solid/internal/usecase"
	"latihan-solid/internal/validator"
)

type BookHandlerInterface interface {
	SaveBook
	FindAllBook
	FindBookById
	DeleteBook
}

type SaveBook interface {
	Save(bookRequest domain.Book) (domain.Book, error)
}

type FindAllBook interface {
	FindAll() ([]domain.Book, error)
}

type FindBookById interface {
	FindBookById(bookId int) (domain.Book, error)
}

type DeleteBook interface {
	DeleteBook(bookId int) (string, error)
}

type BookHandler struct {
	uc usecase.BookUsecaseInterface
}

// DeleteBook implements BookHandlerInterface.
func (h *BookHandler) DeleteBook(bookId int) (string, error) {
	book, err := h.uc.DeleteBook(bookId)
	if err != nil {
		return book, err
	}

	return book, nil
}

// FindAll implements BookHandlerInterface.
func (h *BookHandler) FindAll() ([]domain.Book, error) {
	books, err := h.uc.FindAll()
	if err != nil {
		return books, err
	}

	return books, nil
}

// FindBookById implements BookHandlerInterface.
func (h *BookHandler) FindBookById(bookId int) (domain.Book, error) {
	book, err := h.uc.FindBookById(bookId)
	if err != nil {
		return book, err
	}

	return book, nil
}

// Save implements BookHandlerInterface.
func (h *BookHandler) Save(bookRequest domain.Book) (domain.Book, error) {

	errValidate := validator.ValidateBook(&bookRequest)
	if errValidate != nil {
		return domain.Book{}, errValidate
	}

	book, err := h.uc.Save(bookRequest)
	if err != nil {
		return book, err
	}

	return book, nil
}

func NewBookHandler(uc usecase.BookUsecaseInterface) BookHandlerInterface {
	return &BookHandler{
		uc: uc,
	}
}
