package usecase

import (
	"latihan-solid/internal/domain"
	"latihan-solid/internal/repository"
)

type BookUsecaseInterface interface {
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

type BookRepository struct {
	repo repository.BookRepositoryInterface
}

// DeleteBook implements BookUsecaseInterface.
func (uc *BookRepository) DeleteBook(bookId int) (string, error) {
	msg, err := uc.repo.DeleteBook(bookId)
	if err != nil {
		return "", err
	}
	return msg, nil
}

// FindAll implements BookUsecaseInterface.
func (uc *BookRepository) FindAll() ([]domain.Book, error) {
	books, err := uc.repo.FindAll()
	if err != nil {
		return books, err
	}

	return books, nil
}

// FindBookById implements BookUsecaseInterface.
func (uc *BookRepository) FindBookById(bookId int) (domain.Book, error) {
	book, err := uc.repo.FindBookById(bookId)
	if err != nil {
		return book, err
	}

	return book, nil

}

// Save implements BookUsecaseInterface.
func (uc *BookRepository) Save(bookRequest domain.Book) (domain.Book, error) {
	book, err := uc.repo.Save(&bookRequest)
	if err != nil {
		return *book, err
	}

	return *book, nil

}

func NewBookUsecase(repo repository.BookRepositoryInterface) BookUsecaseInterface {
	return &BookRepository{
		repo: repo,
	}
}
