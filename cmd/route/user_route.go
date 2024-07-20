package route

import (
	"fmt"
	"latihan-solid/internal/domain"
	"latihan-solid/internal/handler"
	"latihan-solid/internal/repository"
	"latihan-solid/internal/usecase"
	"time"

	"github.com/gofrs/uuid"
)

type UserRouter struct {
	bookHandler    handler.BookHandlerInterface
	userHandler    handler.UserHandlerInterface
	loaningHandler handler.LoaningHandlerInterface
	account        *domain.User
}

func NewUserRouter(account *domain.User) *UserRouter {
	repo := repository.NewBookRepository()
	uc := usecase.NewBookUsecase(repo)
	h := handler.NewBookHandler(uc)

	repoUser := repository.NewUserRepository()
	ucUser := usecase.NewUserUsecase(repoUser)
	hUser := handler.NewUserHandler(ucUser)

	repoLoan := repository.NewLoaningRepository()
	ucLoan := usecase.NewLoaningUsecase(repoLoan)
	hLoan := handler.NewLoaningHandler(ucLoan)

	return &UserRouter{
		bookHandler:    h,
		userHandler:    hUser,
		loaningHandler: hLoan,
		account:        account,
	}
}

func (r *UserRouter) Start() {
	r.initializeBooks()

	continueProgram := true
	for continueProgram {
		var inputMenu int
		r.displayMenu()
		fmt.Scan(&inputMenu)
		fmt.Printf("\n")

		switch inputMenu {
		case 1:
			r.displayAllBooks()
		case 2:
			r.handleBookLoan()
		case 3:
			r.handleBookReturn()
		case 4:
			r.displayAllLoans()
		case 5:
			fmt.Println("LOGOUT AKUN")
			continueProgram = false
		default:
			fmt.Println("MENU YANG DIPILIH TIDAK ADA, SILAHKAN PILIH MENU YANG TERSEDIA")
		}
	}
}

func (r *UserRouter) displayMenu() {
	fmt.Println("\nSelamat datang di aplikasi perpustakaan")
	fmt.Println("Silakan pilih menu yang diinginkan:")
	fmt.Println("1. Lihat semua buku")
	fmt.Println("2. Melakukan peminjaman buku")
	fmt.Println("3. Melakukan pengembalian buku")
	fmt.Println("4. Lihat semua transaksi peminjaman")
	fmt.Println("5. Logout")
}

func (r *UserRouter) displayAllLoans() {
	fmt.Println("MENU LIHAT SEMUA TRANSAKSI PEMINJAMAN")
	allLoans, _ := r.loaningHandler.ListAllAvailableBook()

	for _, loan := range allLoans {
		fmt.Println(loan)
	}

}

func (r *UserRouter) displayAllBooks() {
	fmt.Println("MENU LIHAT SEMUA BUKU")
	allBooks, err := r.bookHandler.FindAll()
	if err != nil {
		fmt.Println("Error saat mengambil daftar buku:", err)
		return
	}
	for _, book := range allBooks {
		status := "dapat dipinjam"
		if !book.Stock {
			status = "tidak dapat dipinjam"
		}
		fmt.Printf("ID: %d, Judul: %s, Penulis: %s, Status: %s\n", book.ID, book.Title, book.Author, status)
	}
	fmt.Printf("\n")
}

func (r *UserRouter) handleBookLoan() {
	fmt.Println("MENU PEMINJAMAN BUKU")

	var idBook int
	fmt.Println("Masukan id buku : ")
	fmt.Scan(&idBook)
	if idBook == 0 {
		fmt.Println("Input tidak valid:")
		return
	}

	findBookById, err := r.bookHandler.FindBookById(idBook)
	if err != nil {
		fmt.Println(err)
		return
	} else if findBookById.Stock == false {
		fmt.Println("Buku ini sedang dipinjam, silahkan pilih buku lainnya")
		return
	}

	r.bookHandler.Update(findBookById.ID, &domain.Book{
		ID:     findBookById.ID,
		Title:  findBookById.Title,
		Author: findBookById.Author,
		Stock:  false,
	})

	uuidGenerate, _ := uuid.NewV4()
	createLoan, err := r.loaningHandler.CreateLoan(domain.Loaning{
		Id:        uuidGenerate.String(),
		Book:      findBookById,
		Peminjam:  *r.account,
		Status:    "DIPINJAM",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		fmt.Println("Gagal meminjam buku:", err.Error())
	} else {
		fmt.Println("BUKU BERHASIL DIPINJAM")
		fmt.Println(createLoan)
	}
}

func (r *UserRouter) handleBookReturn() {
	fmt.Println("MENU PENGEMBALIAN BUKU")

	listAllLoanUser, err := r.loaningHandler.ListBookLoanByUserName(r.account.Name)
	if err != nil {
		fmt.Println("Error saat mengambil daftar pinjaman:", err)
		return
	}

	if len(listAllLoanUser) == 0 {
		fmt.Println("Anda tidak memiliki buku yang dipinjam saat ini.")
		return
	}

	fmt.Println("Buku yang sedang Anda pinjam:")
	for _, loan := range listAllLoanUser {
		if loan.Book.Stock == true {
			fmt.Printf("ID Peminjaman: %s, Judul: %s\n", loan.Id, loan.Book.Title)
		}

	}

	var inputLoanId string
	fmt.Println("Masukan id peminjaman: ")
	_, err = fmt.Scan(&inputLoanId)
	if err != nil {
		fmt.Println("Input tidak valid:", err)
		return
	}

	loaning, err := r.loaningHandler.UpdateAvailableBook(inputLoanId)
	if err != nil {
		fmt.Println("Gagal mengembalikan buku:", err)
	} else {
		fmt.Println("BUKU BERHASIL DIKEMBALIKAN")
		r.bookHandler.Update(loaning.Book.ID, &domain.Book{
			ID:     loaning.Book.ID,
			Title:  loaning.Book.Title,
			Author: loaning.Book.Author,
			Stock:  true,
		})
	}
	r.loaningHandler.CreateLoan(domain.Loaning{
		Id:        inputLoanId,
		Book:      loaning.Book,
		Peminjam:  loaning.Peminjam,
		Status:    "DIKEMBALIKAN",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
}

func (r *UserRouter) initializeBooks() {
	books := []domain.Book{
		{ID: 1, Title: "Buku Ngoding Golang", Author: "Bene", Stock: true},
		{ID: 2, Title: "Buku Ngoding Ruby", Author: "Alex", Stock: false},
		{ID: 3, Title: "Harri Petir", Author: "J.C. Bowling", Stock: true},
	}

	for _, book := range books {
		_, err := r.bookHandler.Save(book)
		if err != nil {
			fmt.Printf("Error saat menyimpan buku: %s\n", err)
		}
	}
}
