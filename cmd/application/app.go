package application

import (
	"bufio"
	"fmt"
	"latihan-solid/cmd/route"
	"latihan-solid/internal/domain"
	"latihan-solid/internal/handler"
	"latihan-solid/internal/repository"
	"latihan-solid/internal/usecase"
	"os"
)

func initializeBooks(h handler.BookHandlerInterface) {
	books := []domain.Book{
		{ID: 1, Title: "Buku Ngoding Golang", Author: "Bene", Stock: true},
		{ID: 2, Title: "Buku Ngoding Ruby", Author: "Alex", Stock: false},
		{ID: 3, Title: "Harri Petir", Author: "J.C. Bowling", Stock: true},
	}

	for _, book := range books {
		h.Save(book)

	}
}

func initializeUsers(h handler.UserHandlerInterface) {
	users := []domain.User{
		{Id: 1, Name: "admin"},
		{Id: 2, Name: "user"},
		{Id: 3, Name: "ahmad"}}

	for _, user := range users {
		h.Save(user)
	}
}

func App() {
	repo := repository.NewBookRepository()
	uc := usecase.NewBookUsecase(repo)
	h := handler.NewBookHandler(uc)

	repoUser := repository.NewUserRepository()
	ucUser := usecase.NewUserUsecase(repoUser)
	hUser := handler.NewUserHandler(ucUser)

	initializeBooks(h)
	initializeUsers(hUser)

	scaner := bufio.NewScanner(os.Stdin)
	continueProgram := true

	for continueProgram {
		fmt.Println("SILAHKAN LOGIN AKUN")
		fmt.Print("masukan nama anda : ")
		scaner.Scan()
		inputLogin := scaner.Text()
		fmt.Printf("\n")

		findUserByName, _ := hUser.FindUserByName(inputLogin)
		if findUserByName.Id == 0 && findUserByName.Name == "" {
			fmt.Println("AKUN ANDA TIDAK ADA")
			hUser.Save(domain.User{
				Id:   77,
				Name: inputLogin,
			})
			listUser, _ := hUser.FindAll()
			for _, users := range listUser {
				fmt.Println(users)
			}
		} else {
			fmt.Printf("SELAMAT DATANG %s\n", findUserByName.Name)
		}

		switch findUserByName.Name {
		case "exit":
			continueProgram = false
		default:
			fmt.Println("MASUK KE ROUTE USER PUBLIC")
			account, _ := hUser.FindUserByName(inputLogin)

			userRouter := route.NewUserRouter(&account)
			userRouter.Start()
		}
	}

}
