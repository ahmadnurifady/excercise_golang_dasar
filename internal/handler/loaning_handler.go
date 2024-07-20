package handler

import (
	"latihan-solid/internal/domain"
	"latihan-solid/internal/usecase"
)

type LoaningHandlerInterface interface {
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

type LoaningHandler struct {
	uc usecase.LoaningUsecaseInterface
}

// createLoan implements LoaningHandlerInterface.
func (h *LoaningHandler) CreateLoan(loanRequest domain.Loaning) (domain.Loaning, error) {
	loan, err := h.uc.CreateLoan(loanRequest)
	if err != nil {
		return domain.Loaning{}, err
	}

	return loan, nil
}

// ListBookLoanByUserName implements LoaningHandlerInterface.
func (h *LoaningHandler) ListBookLoanByUserName(userName string) ([]domain.Loaning, error) {
	loans, err := h.uc.ListBookLoanByUserName(userName)
	if err != nil {
		return []domain.Loaning{}, err
	}

	return loans, nil
}

// listAllAvailableBook implements LoaningHandlerInterface.
func (h *LoaningHandler) ListAllAvailableBook() ([]domain.Loaning, error) {
	loans, err := h.uc.ListAllAvailableBook()
	if err != nil {
		return []domain.Loaning{}, err
	}

	return loans, nil
}

// updateAvailableBook implements LoaningHandlerInterface.
func (h *LoaningHandler) UpdateAvailableBook(loaningId string) (domain.Loaning, error) {
	loan, err := h.uc.UpdateAvailableBook(loaningId)
	if err != nil {
		return domain.Loaning{}, err
	}

	return loan, nil
}

func NewLoaningHandler(uc usecase.LoaningUsecaseInterface) LoaningHandlerInterface {
	return &LoaningHandler{
		uc: uc,
	}
}
