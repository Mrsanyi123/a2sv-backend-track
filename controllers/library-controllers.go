package controllers

import (
	"fmt"
	"library-management/models"
	"library-management/services"
	"time"
)

// Run - interactive console that now includes Reserve and Simulation
func Run(library *services.Library) {
	for {
		fmt.Println("\n1. Add Book")
		fmt.Println("2. Add Member")
		fmt.Println("3. Borrow Book")
		fmt.Println("4. Return Book")
		fmt.Println("5. List Available Books")
		fmt.Println("6. List Borrowed Books")
		fmt.Println("7. Reserve Book")                  // new
		fmt.Println("8. Simulate Concurrent Reservations") // new: test
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
		case 7:
			// Reserve book
			var bookID, memberID int
			fmt.Print("Book ID: ")
			fmt.Scan(&bookID)
			fmt.Print("Member ID: ")
			fmt.Scan(&memberID)
			err := library.ReserveBook(bookID, memberID)
			if err != nil {
				fmt.Println("Error reserving:", err)
			} else {
				fmt.Println("Reserved successfully. If not borrowed within 5s, it will auto-cancel.")
			}
		case 8:
			// Simulate concurrent reservations: multiple members try to reserve the same book
			var bookID int
			fmt.Print("Book ID to stress-test: ")
			fmt.Scan(&bookID)

			// Create some sample members if not present
			memberIDs := []int{201, 202, 203, 204}
			for _, mid := range memberIDs {
				if _, ok := library.Members[mid]; !ok {
					library.Members[mid] = models.Member{ID: mid, Name: fmt.Sprintf("Member%d", mid)}
				}
			}

			fmt.Println("Starting concurrent reservation attempts...")
			for _, mid := range memberIDs {
				go func(m int) {
					err := library.ReserveBook(bookID, m)
					if err != nil {
						fmt.Printf("member %d: reserve failed: %v\n", m, err)
						return
					}
					fmt.Printf("member %d: reserved successfully\n", m)
					// optional: try to borrow after a short sleep for one member to demonstrate borrow before auto-cancel
					if m == memberIDs[0] {
						time.Sleep(2 * time.Second)
						if err := library.BorrowBook(bookID, m); err != nil {
							fmt.Printf("member %d: borrow failed: %v\n", m, err)
						} else {
							fmt.Printf("member %d: borrowed the book\n", m)
						}
					}
				}(mid)
			}

			// allow time for goroutines to run and auto-cancel to happen
			fmt.Println("Waiting 7 seconds for the simulation to finish...")
			time.Sleep(7 * time.Second)
			fmt.Println("Simulation complete.")
		case 0:
			fmt.Println("Bye..")
			return
		default:
			fmt.Println("Invalid choice, try again.")
		}
	}
}
