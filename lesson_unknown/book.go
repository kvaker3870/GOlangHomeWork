package book

import (
	"fmt"
	"unicode"
)

type Book struct {
	BookId int
	Title  string
	Author string
}

type BookShelf struct {
	Books map[rune][]Book
}

func setBookOnShelf(book Book, shelf *BookShelf) {
	if shelf.Books == nil {
		shelf.Books = make(map[rune][]Book)
	}

	bookRune := []rune(book.Title)[0]
	shelf.Books[bookRune] = append(shelf.Books[bookRune], book)
}

func (s *BookShelf) getBooksFromBookshelfByAlphabet(letter rune) {

	letter = unicode.ToUpper(letter)
	fmt.Println(s.Books[letter])
}
