package main

import (
	"fmt"
	"latihan-solid/internal/domain"
	"latihan-solid/internal/handler"
	"latihan-solid/internal/repository"
	"latihan-solid/internal/usecase"
	"reflect"
)

func main() {
	repo := repository.NewBookRepository()
	uc := usecase.NewBookUsecase(repo)
	h := handler.NewBookHandler(uc)

	bookReq := domain.Book{
		ID:     1,
		Title:  "Buku Ngoding Golang",
		Author: "Bene",
		Stock:  11,
	}

	book2 := domain.Book{
		ID:     2,
		Title:  "Buku Ngoding Golang",
		Author: "Alex",
		Stock:  12,
	}

	// e, _ := json.Marshal(book2)
	// a, _ := json.Marshal(bookReq)

	// fmt.Println(string(e))
	// fmt.Println(a)

	// repo.Save(&bookReq)
	// repo.FindAll()
	_, err := h.Save(bookReq)
	if err != nil {
		fmt.Println(err)
	}
	_, err2 := h.Save(book2)
	if err2 != nil {
		fmt.Println(err2)
	}

	allBooks, _ := repo.FindAll()
	fmt.Println("All books in repository:")
	for _, book := range allBooks {
		fmt.Printf("%+v\n", book)
	}

	var reflectValue reflect.Value

	reflectValue = reflect.ValueOf(10)

	fmt.Println(reflectValue.Type())

	reflectValue = reflect.ValueOf("anjayyy")

	fmt.Println(reflectValue.Len())

	fmt.Println(reflectValue.Type())

	// // findById, _ := repo.FindBookById(1)
	// // fmt.Println(findById)

	// repo.DeleteBook(1)

	// allBooks2, _ := repo.FindAll()
	// fmt.Println("All books in repository:")
	// for _, book := range allBooks2 {
	// 	fmt.Printf("%+v\n", book)
	// }
}
