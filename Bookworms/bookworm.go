package main

import (
	"encoding/json"
	"os"
	"sort"
)

type Bookworm struct {
	Name  string `json:"name"`
	Books []Book `json:"books"`
}
type Book struct {
	Author string `json:"author"`
	Title  string `json:"title"`
}

// loadBookworms loads the bookworms from the given file
func loadBookworms(filename string) ([]Bookworm, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var bookworms []Bookworm

	err = json.NewDecoder(f).
		Decode(&bookworms)
	if err != nil {
		return nil, err
	}

	return bookworms, nil
}

func findCommonBooks(bookworms []Bookworm) []Book {
	bookOnShelves := booksCount(bookworms)

	var commonBooks []Book
	for book, count := range bookOnShelves {
		if count > 1 {
			commonBooks = append(commonBooks, book)
		}
	}

	return sortBooks(commonBooks)
}

// booksCount registers all the books and their occurrences from bookworms shelves
func booksCount(bookworms []Bookworm) map[Book]uint {
	count := make(map[Book]uint)

	for _, bookworm := range bookworms {
		for _, book := range bookworm.Books {
			count[book]++
		}
	}

	return count
}

// sortBooks sorts the books by author and title
//func sortBooks(books []Book) []Book {
//sort.Slice(books, func(i, j int) bool {
//	if books[i].Author != books[j].Author {
//		return books[i].Author < books[j].Author
//	}
//	return books[i].Title < books[j].Title
//})

//return books
//}

// byAuthor is a type that implements sort.Interface for sorting books by author
type byAuthor []Book

// Len implements sort.Interface by returning the length of BookByAuthor.
func (b byAuthor) Len() int { return len(b) }

// Swap implements sort.Interface and swaps two books.
func (b byAuthor) Swap(i, j int) { b[i], b[j] = b[j], b[i] }

// Less implements sort.Interface and returns BookByAuthor sorted by Author and then Title.
func (b byAuthor) Less(i, j int) bool {
	if b[i].Author != b[j].Author {
		return b[i].Author < b[j].Author
	}
	return b[i].Title < b[j].Title
}

// sortBooks sorts the books by Author and then Title in alphabetical order.
func sortBooks(books []Book) []Book {
	sort.Sort(byAuthor(books))
	return books
}
