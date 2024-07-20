package repository

import (
	"errors"
	"latihan-solid/internal/domain"
	"time"
)

type LoaningRepositoryInterface interface {
	CreateLoan
	ListAllAvailableBook
	UpdateAvailableBook
	ListBookLoanByUserName
}

type CreateLoan interface {
	CreateLoan(loanRequest *domain.Loaning) (*domain.Loaning, error)
}

type ListBookLoanByUserName interface {
	ListBookLoanByUserName(userName string) ([]domain.Loaning, error)
}

type ListAllAvailableBook interface {
	ListAllAvailableBook() ([]domain.Loaning, error)
}

type UpdateAvailableBook interface {
	UpdateAvailableBook(loaningId string) (domain.Loaning, error)
}

type LoaningRepository struct {
	db map[string]domain.Loaning
}

// createLoan implements LoaningRepositoryInterface.
func (repo *LoaningRepository) CreateLoan(loanRequest *domain.Loaning) (*domain.Loaning, error) {
	if _, exists := repo.db[loanRequest.Id]; exists {
		return nil, errors.New("peminjaman buku ini sedang dilakukan")
	}

	repo.db[loanRequest.Id] = *loanRequest

	return loanRequest, nil
}

// ListBookLoanByUserName implements LoaningRepositoryInterface.
func (repo *LoaningRepository) ListBookLoanByUserName(userName string) ([]domain.Loaning, error) {
	var bucketListLoan []domain.Loaning

	for _, loan := range repo.db {
		if loan.Peminjam.Name == userName {
			bucketListLoan = append(bucketListLoan, loan)
		}
	}

	return bucketListLoan, nil
}

// listAllAvailableBook implements LoaningRepositoryInterface.
func (repo *LoaningRepository) ListAllAvailableBook() ([]domain.Loaning, error) {
	var allLoaning []domain.Loaning

	for _, loan := range repo.db {
		allLoaning = append(allLoaning, loan)
	}

	return allLoaning, nil
}

// updateAvailableBook implements LoaningRepositoryInterface.
func (repo *LoaningRepository) UpdateAvailableBook(loaningId string) (domain.Loaning, error) {
	loaningExist, exists := repo.db[loaningId]
	if !exists {
		return domain.Loaning{}, errors.New("transaksi peminjaman tidak ditemukan")
	}

	// if loaningExist.Status == ""
	loaningExist.UpdatedAt = time.Now()

	repo.db[loaningId] = loaningExist

	return loaningExist, nil
}

func NewLoaningRepository() LoaningRepositoryInterface {
	return &LoaningRepository{
		db: make(map[string]domain.Loaning)}
}
