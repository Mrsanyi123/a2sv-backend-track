package services

import (
	"errors"
	"fmt"
	"library-management/models"
	"sync"
	"time"
)

type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
	ReserveBook(bookID int, memberID int) error
}

// ReservationRequest used to pass reservation requests to workers
type ReservationRequest struct {
	BookID   int
	MemberID int
	Result   chan error
}

type Library struct {
	Books   map[int]models.Book
	Members map[int]models.Member

	mu sync.Mutex // protects Books & Members

	ReservationQueue chan ReservationRequest
}

// NewLibrary constructs a Library and starts reservation workers
func NewLibrary() *Library {
	lib := &Library{
		Books:            make(map[int]models.Book),
		Members:          make(map[int]models.Member),
		ReservationQueue: make(chan ReservationRequest, 50), // buffered
	}

	// start several workers to process reservations concurrently
	workerCount := 3
	for i := 0; i < workerCount; i++ {
		go lib.reservationWorker(i + 1)
	}

	return lib
}

func (l *Library) AddBook(book models.Book) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.Books[book.ID] = book
}

func (l *Library) RemoveBook(bookID int) {
	l.mu.Lock()
	defer l.mu.Unlock()
	delete(l.Books, bookID)
}

func (l *Library) BorrowBook(bookID int, memberID int) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	book, ok := l.Books[bookID]
	if !ok {
		return errors.New("book not found")
	}
	if book.Status == "Borrowed" {
		return errors.New("book already borrowed")
	}
	// If reserved, must be reserved by this member to borrow
	if book.ReservedBy != 0 && book.ReservedBy != memberID {
		return errors.New("book reserved by another member")
	}
	member, ok := l.Members[memberID]
	if !ok {
		return errors.New("member not found")
	}

	// Mark borrowed and clear reservation
	book.Status = "Borrowed"
	book.ReservedBy = 0
	l.Books[bookID] = book

	member.BorrowedBooks = append(member.BorrowedBooks, book)
	l.Members[memberID] = member
	return nil
}

func (l *Library) ReturnBook(bookID int, memberID int) error {
	l.mu.Lock()
	defer l.mu.Unlock()

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

	// remove from borrowed list
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
	l.mu.Lock()
	defer l.mu.Unlock()

	var list []models.Book
	for _, b := range l.Books {
		if b.Status == "Available" && b.ReservedBy == 0 {
			list = append(list, b)
		}
	}
	return list
}

func (l *Library) ListBorrowedBooks(memberID int) []models.Book {
	l.mu.Lock()
	defer l.mu.Unlock()

	if member, ok := l.Members[memberID]; ok {
		return member.BorrowedBooks
	}
	return nil
}

// ReserveBook enqueues a reservation request. The actual reservation is processed by a worker.
// If reservation succeeds, it spawns a 5-second timer goroutine to auto-cancel if not borrowed.
func (l *Library) ReserveBook(bookID int, memberID int) error {
	resCh := make(chan error, 1)
	req := ReservationRequest{
		BookID:   bookID,
		MemberID: memberID,
		Result:   resCh,
	}

	// enqueue request
	l.ReservationQueue <- req

	// wait for worker to respond with result
	err := <-resCh
	if err != nil {
		return err
	}

	// start auto-cancel timer: if not borrowed in 5 seconds, cancel the reservation
	go func(bID, mID int) {
		timer := time.NewTimer(5 * time.Second)
		<-timer.C
		l.autoCancelReservation(bID, mID)
	}(bookID, memberID)

	return nil
}

// reservationWorker processes reservation requests from the channel
func (l *Library) reservationWorker(workerID int) {
	fmt.Printf("[worker %d] started\n", workerID)
	for req := range l.ReservationQueue {
		// perform reservation under lock
		l.mu.Lock()
		book, ok := l.Books[req.BookID]
		if !ok {
			l.mu.Unlock()
			req.Result <- errors.New("book not found")
			continue
		}

		if book.Status == "Borrowed" {
			l.mu.Unlock()
			req.Result <- errors.New("book already borrowed")
			continue
		}

		if book.ReservedBy != 0 {
			l.mu.Unlock()
			req.Result <- errors.New("book already reserved")
			continue
		}

		// reserve
		book.ReservedBy = req.MemberID
		l.Books[req.BookID] = book
		l.mu.Unlock()

		fmt.Printf("[worker %d] reserved book=%d for member=%d\n", workerID, req.BookID, req.MemberID)
		req.Result <- nil
	}
}

// autoCancelReservation cancels reservation if it's still reserved and not borrowed
func (l *Library) autoCancelReservation(bookID int, memberID int) {
	l.mu.Lock()
	defer l.mu.Unlock()

	book, ok := l.Books[bookID]
	if !ok {
		return
	}

	// if still reserved by same member and not borrowed, clear reservation
	if book.ReservedBy == memberID && book.Status != "Borrowed" {
		book.ReservedBy = 0
		l.Books[bookID] = book
		fmt.Printf("[auto-cancel] reservation cancelled for book=%d member=%d\n", bookID, memberID)
	}
}
