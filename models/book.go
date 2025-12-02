package models


import "sync"


// Book represents a book in the library.
type Book struct {
	ID int
	Title string
	Author string
	Status string // "Available", "Borrowed"
	Available bool
	ReservedBy int // memberID who reserved, 0 if none
	mu sync.Mutex
}


func (b *Book) IsAvailable() bool {
b.mu.Lock()
defer b.mu.Unlock()
return b.Available && b.ReservedBy == 0
}