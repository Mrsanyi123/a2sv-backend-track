Library Management System (Go)
Overview

The Library Management System is a simple console-based Go application that demonstrates:

Structs (Book and Member)

Interfaces (LibraryManager)

Methods, slices, maps

Package-based project structure

Console interaction

It allows users to add books and members, borrow and return books, and view available or borrowed books.

Folder Structure
library_management/
├── main.go # Entry point
├── controllers/  
│ └── library_controller.go # Handles console menu
├── models/
│ ├── book.go # Book struct
│ └── member.go # Member struct
├── services/
│ └── library_service.go # LibraryManager interface and implementation
├── docs/
│ └── documentation.md # This documentation
└── go.mod # Go module definition

Models
Book

Represents a book in the library with:

ID — unique identifier

Title — book title

Author — author name

Status — "Available" or "Borrowed"

Member

Represents a library member with:

ID — unique identifier

Name — member name

BorrowedBooks — list of books currently borrowed

Services
LibraryManager Interface

Defines library operations:

AddBook(book Book)

RemoveBook(bookID int)

BorrowBook(bookID int, memberID int)

ReturnBook(bookID int, memberID int)

ListAvailableBooks() []Book

ListBorrowedBooks(memberID int) []Book

Library Struct

Implements LibraryManager. Stores:

Books — map of all books (map[int]Book)

Members — map of all members (map[int]Member)

Controllers

The controller handles console interaction with a menu:

Add Book

Add Member

Borrow Book

Return Book

List Available Books

List Borrowed Books

Exit

It calls the corresponding service methods to perform operations.

How to Run

Make sure Go is installed (go version).

Open terminal in project root (library_management/).

Initialize module (if needed):

go mod init library_management
go mod tidy

Run the application:

go run .

Follow the menu to interact with the library.

Usage Example

1. Add Book -> enter ID, title, author
2. Add Member -> enter ID, name
3. Borrow Book -> enter book ID and member ID
4. Return Book -> enter book ID and member ID
5. List Available Books
6. List Borrowed Books -> enter member ID
7. Exit

Features

Structs, slices, and maps

Interface implementation

Simple console-based UI

Modular Go project structure

In-memory storage of books and members
