package main

import "fmt"

type BookManager interface {
	AddBook(title, author string)
	ListBooks() []Book
}

type Book struct {
	Title  string
	Author string
}

type Library struct {
	Books []*Book
}

func (l *Library) ListBooks() []Book {
	var books []Book
	fmt.Println("\nLIST BOOKS:")
	for i, book := range l.Books {
		fmt.Printf("%d. %s oleh %s\n", i+1, book.Title, book.Author)
		books = append(books, *book)
	}
	return books
}

func (l *Library) AddBook(title, author string) {
	newBook := &Book{Title: title, Author: author}

	l.Books = append(l.Books, newBook)
	fmt.Printf("Buku '%s' oleh %s berhasil ditambahkan ke perpustakaan.\n", title, author)
}

func main(){
	myListBooks := &Library{}

	myListBooks.AddBook("Golang", "Mps")
	myListBooks.AddBook("Js", "Yuhu")
	
	myListBooks.ListBooks()
}