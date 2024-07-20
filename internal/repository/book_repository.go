package repository

import (
	"errors"
	"fmt"
	"latihan-solid/internal/domain"
)

type BookRepositoryInterface interface {
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
	Save(bookRequest *domain.Book) (*domain.Book, error)
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
	db map[int]domain.Book
}

// DeleteBook implements BookRepositoryInterface.
func (repo *BookRepository) DeleteBook(bookId int) (string, error) {
	if _, exist := repo.db[bookId]; !exist {
		return "", fmt.Errorf("buku dengan id: %d tidak ditemukan", bookId)
	}

	delete(repo.db, bookId)

	return "Success Delete Book", nil
}

// findAll implements BookRepositoryInterface.
func (repo *BookRepository) FindAll() ([]domain.Book, error) {
	var allBooks []domain.Book

	for _, book := range repo.db {
		allBooks = append(allBooks, book)
	}

	return allBooks, nil
}

// findBookById implements BookRepositoryInterface.
func (repo *BookRepository) FindBookById(bookId int) (domain.Book, error) {
	if _, exist := repo.db[bookId]; !exist {
		return domain.Book{}, fmt.Errorf("buku dengan id: %d tidak ditemukan", bookId)
	}
	var book domain.Book
	for _, val := range repo.db {
		if val.ID == bookId {
			book = val
		}
	}
	return book, nil
}

// save implements BookRepositoryInterface.
func (repo *BookRepository) Save(bookRequest *domain.Book) (*domain.Book, error) {
	if _, exists := repo.db[bookRequest.ID]; exists {
		return &domain.Book{}, errors.New("buku sudah terdaftar")
	}

	repo.db[bookRequest.ID] = *bookRequest
	return bookRequest, nil
}

// Update implements BookRepositoryInterface.
func (repo *BookRepository) Update(bookID int, updateData *domain.Book) (*domain.Book, error) {
	existingBook, exists := repo.db[bookID]

	if !exists {
		return &domain.Book{}, fmt.Errorf("buku dengan id: %d tidak ditemukan", bookID)
	}

	repo.db[bookID] = *updateData
	return &existingBook, nil

}

func NewBookRepository() BookRepositoryInterface {
	return &BookRepository{
		db: make(map[int]domain.Book)}
}
