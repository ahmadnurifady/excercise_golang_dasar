package usecase

import (
	"latihan-solid/internal/domain"
	"latihan-solid/internal/repository"
)

type LoaningUsecaseInterface interface {
	CreateLoan
	ListAllAvailableBook
	UpdateAvailableBook
	ListBookLoanByUserName
}

type CreateLoan interface {
	CreateLoan(loanRequest domain.Loaning) (domain.Loaning, error)
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

type LoaningUsecase struct {
	repo repository.LoaningRepositoryInterface
}

// createLoan implements LoaningUsecaseInterface.
func (uc *LoaningUsecase) CreateLoan(loanRequest domain.Loaning) (domain.Loaning, error) {
	loan, err := uc.repo.CreateLoan(&loanRequest)
	if err != nil {
		return domain.Loaning{}, err
	}

	return *loan, nil
}

// ListBookLoanByUserName implements LoaningUsecaseInterface.
func (uc *LoaningUsecase) ListBookLoanByUserName(userName string) ([]domain.Loaning, error) {
	loans, err := uc.repo.ListBookLoanByUserName(userName)
	if err != nil {
		return []domain.Loaning{}, err
	}

	return loans, nil
}

// listAllAvailableBook implements LoaningUsecaseInterface.
func (uc *LoaningUsecase) ListAllAvailableBook() ([]domain.Loaning, error) {
	loans, err := uc.repo.ListAllAvailableBook()
	if err != nil {
		return []domain.Loaning{}, err
	}

	return loans, nil
}

// updateAvailableBook implements LoaningUsecaseInterface.
func (uc *LoaningUsecase) UpdateAvailableBook(loaningId string) (domain.Loaning, error) {
	loan, err := uc.repo.UpdateAvailableBook(loaningId)
	if err != nil {
		return domain.Loaning{}, err
	}

	return loan, nil
}

// DeleteBook implements BookUsecaseInterface.

func NewLoaningUsecase(repo repository.LoaningRepositoryInterface) LoaningUsecaseInterface {
	return &LoaningUsecase{
		repo: repo,
	}
}
