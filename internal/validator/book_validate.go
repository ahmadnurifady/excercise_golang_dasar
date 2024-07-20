package validator

import (
	"fmt"
	"latihan-solid/internal/domain"
	"reflect"
	"strings"
)

func ValidateBook(book *domain.Book) error {
	var reflectValue reflect.Value

	if book == nil {
		return fmt.Errorf("book cannot be nil")
	}

	reflectValue = reflect.ValueOf(book.ID)
	if reflectValue.Int() <= 0 {
		return fmt.Errorf("invalid book ID: must be positive")
	} else if reflectValue.Kind() != reflect.Int {
		return fmt.Errorf("invalid book ID: must be number")
	}

	if strings.TrimSpace(book.Title) == "" {
		return fmt.Errorf("book title cannot be empty")
	}

	if len(book.Title) > 100 {
		return fmt.Errorf("book title too long: maximum 100 characters")
	}

	if strings.TrimSpace(book.Author) == "" {
		return fmt.Errorf("book author cannot be empty")
	}

	if len(book.Author) > 50 {
		return fmt.Errorf("author name too long: maximum 50 characters")
	}

	return nil
}
