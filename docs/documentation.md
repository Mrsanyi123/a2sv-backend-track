Library Management System (Go)
Overview

The Library Management System is a console-based Go application that demonstrates:

Structs (Book and Member)

Interfaces (LibraryManager)

Methods, slices, maps

Package-based project structure

Console interaction

Concurrency using Goroutines, Channels, and Mutexes

It allows users to:

Add books and members

Borrow and return books

Reserve books concurrently

View available or borrowed books

Folder Structure
library_management/
├── main.go                       # Entry point
├── controllers/  
│   └── library_controller.go     # Handles console menu
├── models/
│   ├── book.go                   # Book struct
│   └── member.go                 # Member struct
├── services/
│   └── library_service.go        # LibraryManager interface, implementation, and reservation logic
├── concurrency/
│   └── reservation_worker.go     # Handles concurrent reservations via Goroutines and Channels
├── docs/
│   └── documentation.md          # Project documentation
└── go.mod                        # Go module definition

Models
Book

Represents a book in the library with:

ID — unique identifier

Title — book title

Author — author name

Status — "Available", "Borrowed"

ReservedBy — member ID who reserved the book (0 if none)

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

ReserveBook(bookID int, memberID int) error — allows concurrent reservations

Library Struct

Implements LibraryManager. Stores:

Books — map of all books (map[int]Book)

Members — map of all members (map[int]Member)

ReservationQueue — channel for queuing reservation requests

mu — sync.Mutex for thread-safe access

Concurrency Implementation

Reservation requests are sent to a buffered channel (ReservationQueue).

Multiple worker Goroutines process requests concurrently.

Mutexes prevent race conditions when updating book availability or member data.

Reservations are auto-cancelled after 5 seconds if not borrowed.

Borrowing a book is only allowed if it is either available or reserved by the same member.

Controllers

The controller handles console interaction with a menu:

Add Book

Add Member

Borrow Book

Return Book

List Available Books

List Borrowed Books

Reserve Book — allows a member to reserve a book concurrently

Simulate Concurrent Reservations — multiple members try reserving the same book at the same time (demonstrates concurrency)

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

Add Book → enter ID, title, author

Add Member → enter ID, name

Borrow Book → enter book ID and member ID

Return Book → enter book ID and member ID

List Available Books

List Borrowed Books → enter member ID

Reserve Book → enter book ID and member ID (auto-cancels if not borrowed in 5s)

Simulate Concurrent Reservations → test multiple members reserving the same book

Exit

Features

Structs, slices, and maps

Interface implementation

Simple console-based UI

Modular Go project structure

In-memory storage of books and members

Concurrent reservations using Goroutines and Channels

Thread-safe access with Mutexes

Auto-cancellation of unborrowed reservations

Simulation of multiple concurrent reservation requests