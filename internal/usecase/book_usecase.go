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
	UpdateBook
}

type UpdateBook interface {
	Update(bookID int, updateData *domain.Book) (*domain.Book, error)
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

type BookUsecase struct {
	repo repository.BookRepositoryInterface
}

// DeleteBook implements BookUsecaseInterface.
func (uc *BookUsecase) DeleteBook(bookId int) (string, error) {
	msg, err := uc.repo.DeleteBook(bookId)
	if err != nil {
		return "", err
	}
	return msg, nil
}

// FindAll implements BookUsecaseInterface.
func (uc *BookUsecase) FindAll() ([]domain.Book, error) {
	books, err := uc.repo.FindAll()
	if err != nil {
		return books, err
	}

	return books, nil
}

// FindBookById implements BookUsecaseInterface.
func (uc *BookUsecase) FindBookById(bookId int) (domain.Book, error) {
	book, err := uc.repo.FindBookById(bookId)
	if err != nil {
		return book, err
	}

	return book, nil

}

// Save implements BookUsecaseInterface.
func (uc *BookUsecase) Save(bookRequest domain.Book) (domain.Book, error) {
	book, err := uc.repo.Save(&bookRequest)
	if err != nil {
		return *book, err
	}

	return *book, nil

}

// Update implements BookUsecaseInterface.
func (uc *BookUsecase) Update(bookID int, updateData *domain.Book) (*domain.Book, error) {
	book, err := uc.repo.Update(bookID, updateData)
	if err != nil {
		return book, err
	}

	return book, nil
}

func NewBookUsecase(repo repository.BookRepositoryInterface) BookUsecaseInterface {
	return &BookUsecase{
		repo: repo,
	}
}
