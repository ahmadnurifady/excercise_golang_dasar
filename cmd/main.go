package main

import "latihan-solid/cmd/application"

func main() {

	application.App()

	// repo := repository.NewBookRepository()
	// uc := usecase.NewBookUsecase(repo)
	// h := handler.NewBookHandler(uc)

	// repoUser := repository.NewUserRepository()
	// ucUser := usecase.NewUserRepository(repoUser)
	// hUser := handler.NewUserHandler(ucUser)

	// repoLoan := repository.NewLoaningRepository()
	// ucLoan := usecase.NewLoaningUsecase(repoLoan)
	// hLoan := handler.NewLoaningHandler(ucLoan)

	// loan1 := domain.Loaning{
	// 	Id: "",
	// 	Book: domain.Book{
	// 		ID:     1,
	// 		Title:  "1111",
	// 		Author: "1111",
	// 		Stock:  false,
	// 	},
	// 	User: domain.User{
	// 		Id:   1,
	// 		Name: "11111",
	// 	},
	// 	Available: false,
	// 	CreatedAt: time.Now(),
	// 	UpdatedAt: time.Now(),
	// }

	// result, err := hLoan.CreateLoan(loan1)
	// if err == nil {
	// 	fmt.Println(result)
	// }

	// bookReq := domain.Book{
	// 	ID:     1,
	// 	Title:  "Buku Ngoding Golang",
	// 	Author: "Bene",
	// 	Stock:  11,
	// }

	// book2 := domain.Book{
	// 	ID:     2,
	// 	Title:  "Buku Ngoding Golang",
	// 	Author: "Alex",
	// 	Stock:  12,
	// }

	// orang1 := domain.User{
	// 	Id:   1,
	// 	Name: "orang1 dalam negeri",
	// }

	// orang2 := domain.User{
	// 	Id:   2,
	// 	Name: "orang2 luar negeri",
	// }

	// _, err := hUser.Save(orang1)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// _, err2 := hUser.Save(orang2)
	// if err2 != nil {
	// 	fmt.Println(err2)
	// }

	// allUsers, _ := hUser.FindAll()
	// fmt.Println("All books in handler:")
	// for _, user := range allUsers {
	// 	fmt.Printf("%+v\n", user)
	// }

	// findUserById, _ := hUser.FindUserById(2)
	// fmt.Println(findUserById)

	// deleteUser, _ := hUser.DeleteUser(2)
	// fmt.Println(deleteUser)

	// allUsers1, _ := hUser.FindAll()
	// fmt.Println("All books in handler:")
	// for _, user := range allUsers1 {
	// 	fmt.Printf("%+v\n", user)
	// }

	// updateBook := domain.Book{
	// 	ID:     2,
	// 	Title:  "Buku Ngoding Golang update",
	// 	Author: "waluyo",
	// 	Stock:  11,
	// }

	// e, _ := json.Marshal(book2)
	// a, _ := json.Marshal(bookReq)

	// fmt.Println(string(e))
	// fmt.Println(a)

	// repo.Save(&bookReq)
	// repo.FindAll()
	// _, err := h.Save(bookReq)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// _, err2 := h.Save(book2)
	// if err2 != nil {
	// 	fmt.Println(err2)
	// }

	// allBooks, _ := repo.FindAll()
	// fmt.Println("All books in repository:")
	// for _, book := range allBooks {
	// 	fmt.Printf("%+v\n", book)
	// }

	// uptBook, errUpt := h.Update(2, &updateBook)
	// if errUpt != nil {
	// 	fmt.Println("SEHABIS DIUPDATE", uptBook)
	// }

	// allBooksForUpdate, _ := repo.FindAll()
	// fmt.Println("All books in repository:")
	// for _, book := range allBooksForUpdate {
	// 	fmt.Printf("%+v\n", book)
	// }

	// var reflectValue reflect.Value

	// reflectValue = reflect.ValueOf(10)

	// fmt.Println(reflectValue.Type())

	// reflectValue = reflect.ValueOf("anjayyy")

	// fmt.Println(reflectValue.Len())

	// fmt.Println(reflectValue.Type())

	// // findById, _ := repo.FindBookById(1)
	// // fmt.Println(findById)

	// repo.DeleteBook(1)

	// allBooks2, _ := repo.FindAll()
	// fmt.Println("All books in repository:")
	// for _, book := range allBooks2 {
	// 	fmt.Printf("%+v\n", book)
	// }
}
