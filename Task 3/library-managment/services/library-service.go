package services

import (
	"errors"
	"library-management/models"
)

type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
}

type Library struct {
	Books   map[int]models.Book
	Members map[int]models.Member
}

func NewLibrary() *Library {
	return &Library{
		Books:   make(map[int]models.Book),
		Members: make(map[int]models.Member),
	}
}

func (l *Library) AddBook(book models.Book) {
	l.Books[book.ID] = book
}

func (l *Library) RemoveBook(bookID int) {
	delete(l.Books, bookID)
}

func (l *Library) BorrowBook(bookID int, memberID int) error {
	book, ok := l.Books[bookID]
	if !ok {
		return errors.New("book not found")
	}
	if book.Status == "Borrowed" {
		return errors.New("book already borrowed")
	}
	member, ok := l.Members[memberID]
	if !ok {
		return errors.New("member not found")
	}
	book.Status = "Borrowed"
	l.Books[bookID] = book
	member.BorrowedBooks = append(member.BorrowedBooks, book)
	l.Members[memberID] = member
	return nil
}

func (l *Library) ReturnBook(bookID int, memberID int) error {
	book, ok := l.Books[bookID]
	if !ok {
		return errors.New("book not found")
	}
	member, ok := l.Members[memberID]
	if !ok {
		return errors.New("member not found")
	}
	book.Status = "Available"
	l.Books[bookID] = book
	for i, b := range member.BorrowedBooks {
		if b.ID == bookID {
			member.BorrowedBooks = append(member.BorrowedBooks[:i], member.BorrowedBooks[i+1:]...)
			break
		}
	}
	l.Members[memberID] = member
	return nil
}

func (l *Library) ListAvailableBooks() []models.Book {
	var list []models.Book
	for _, b := range l.Books {
		if b.Status == "Available" {
			list = append(list, b)
		}
	}
	return list
}

func (l *Library) ListBorrowedBooks(memberID int) []models.Book {
	if member, ok := l.Members[memberID]; ok {
		return member.BorrowedBooks
	}
	return nil
}