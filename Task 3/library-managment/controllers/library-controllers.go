package controllers

import (
	"fmt"
	"library-management/models"
	"library-management/services"
)

func Run(library *services.Library) {
	for {
		fmt.Println("\n1. Add Book")
		fmt.Println("2. Add Member")
		fmt.Println("3. Borrow Book")
		fmt.Println("4. Return Book")
		fmt.Println("5. List Available Books")
		fmt.Println("6. List Borrowed Books")
		fmt.Println("0. Exit")
		fmt.Print("Choice: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var id int
			var title, author string
			fmt.Print("Enter Book ID: ")
			fmt.Scan(&id)
			fmt.Print("Enter Book Title: ")
			fmt.Scan(&title)
			fmt.Print("Enter Book Author: ")
			fmt.Scan(&author)
			library.AddBook(models.Book{ID: id, Title: title, Author: author, Status: "Available"})
			fmt.Println("Book added successfully.")
		case 2:
			var id int
			var name string
			fmt.Print("Enter Member ID: ")
			fmt.Scan(&id)
			fmt.Print("Enter Member Name: ")
			fmt.Scan(&name)
			library.Members[id] = models.Member{ID: id, Name: name}
			fmt.Println("Member added successfully.")
		case 3:
			var bookID, memberID int
			fmt.Print("Book ID: ")
			fmt.Scan(&bookID)
			fmt.Print("Member ID: ")
			fmt.Scan(&memberID)
			err := library.BorrowBook(bookID, memberID)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Book borrowed!")
			}

		case 4:
			var bookID, memberID int
			fmt.Print("Book ID: ")
			fmt.Scan(&bookID)
			fmt.Print("Member ID: ")
			fmt.Scan(&memberID)
			err := library.ReturnBook(bookID, memberID)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Book returned!")
			}

		case 5:
			fmt.Println("\nAvailable Books:")
			for _, b := range library.ListAvailableBooks() {
				fmt.Printf("%d - %s by %s\n", b.ID, b.Title, b.Author)
			}

		case 6:
			var memberID int
			fmt.Print("Member ID: ")
			fmt.Scan(&memberID)
			fmt.Println("\nBorrowed Books:")
			for _, b := range library.ListBorrowedBooks(memberID) {
				fmt.Printf("%d - %s by %s\n", b.ID, b.Title, b.Author)
			}
		case 0:
			fmt.Println("Bye..")
			return
		default:
			fmt.Println("Invalid choice, try again.")
		}
	}	
}