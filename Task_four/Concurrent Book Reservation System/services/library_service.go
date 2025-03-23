package services

import (
	"fmt"
	"library_management/models"
	"sync"
	"time"
)

type LibraryManager interface {
	AddBook(book models.Book) error
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
	ReserveBook(bookID int, memberID int) error
	AddMember(member models.Member) error
}

type Library struct {
	Books            map[int]models.Book
	Members          map[int]models.Member
	ReservationQueue chan ReservationRequest
	Mu               sync.Mutex
}

type ReservationRequest struct {
	BookID    int
	MemberID  int
	Result    chan error
	Timestamp time.Time
}

func NewLibrary() *Library {
	return &Library{
		Books:            make(map[int]models.Book),
		Members:          make(map[int]models.Member),
		ReservationQueue: make(chan ReservationRequest, 100),
		Mu:               sync.Mutex{},
	}
}

func (l *Library) ReserveBook(bookID int, memberID int) error {
	resultChan := make(chan error, 1)
	request := ReservationRequest{
		BookID:    bookID,
		MemberID:  memberID,
		Result:    resultChan,
		Timestamp: time.Now(),
	}

	l.ReservationQueue <- request

	err := <-resultChan
	return err
}

func (l *Library) AddBook(book models.Book) error {
	l.Mu.Lock()
	defer l.Mu.Unlock()

	if _, exists := l.Books[book.ID]; exists {
		return fmt.Errorf("book with ID %d already exists", book.ID)
	}

	l.Books[book.ID] = book
	return nil
}

func (l *Library) RemoveBook(bookID int) {
	l.Mu.Lock()
	defer l.Mu.Unlock()
	delete(l.Books, bookID)
}

func (l *Library) BorrowBook(bookID int, memberID int) error {
	l.Mu.Lock()
	defer l.Mu.Unlock()

	book, ok := l.Books[bookID]
	if !ok {
		return fmt.Errorf("book with ID %d not found", bookID)
	}

	if book.Status == "Borrowed" {
		return fmt.Errorf("book with ID %d is already borrowed", bookID)
	}

	if book.Status == "Reserved" {
		return fmt.Errorf("book with ID %d is already reserved", bookID)
	}

	member, ok := l.Members[memberID]
	if !ok {
		return fmt.Errorf("member with ID %d not found", memberID)
	}

	book.Status = "Borrowed"
	l.Books[bookID] = book

	borrowedBookList := member.BorrowedBooks
	borrowedBookList = append(borrowedBookList, book)
	member.BorrowedBooks = borrowedBookList
	l.Members[memberID] = member
	return nil
}

func (l *Library) ReturnBook(bookID int, memberID int) error {
	l.Mu.Lock()
	defer l.Mu.Unlock()

	book, ok := l.Books[bookID]
	if !ok {
		return fmt.Errorf("book with ID %d not found", bookID)
	}
	if book.Status == "Available" {
		return fmt.Errorf("book with ID %d is already available", bookID)
	}

	member, ok := l.Members[memberID]
	if !ok {
		return fmt.Errorf("member with ID %d not found", memberID)
	}

	book.Status = "Available"
	l.Books[bookID] = book

	var updatedBorrowedBooks []models.Book
	for _, borrowedBook := range member.BorrowedBooks {
		if borrowedBook.ID != bookID {
			updatedBorrowedBooks = append(updatedBorrowedBooks, borrowedBook)
		}
	}
	member.BorrowedBooks = updatedBorrowedBooks
	l.Members[memberID] = member

	return nil
}

func (l *Library) ListAvailableBooks() []models.Book {
	l.Mu.Lock()
	defer l.Mu.Unlock()
	availableBooks := []models.Book{}
	for _, book := range l.Books {
		if book.Status == "Available" {
			availableBooks = append(availableBooks, book)
		}
	}
	return availableBooks
}

func (l *Library) ListBorrowedBooks(memberID int) []models.Book {
	l.Mu.Lock()
	defer l.Mu.Unlock()
	member, ok := l.Members[memberID]
	if !ok {
		return []models.Book{}
	}
	return member.BorrowedBooks
}

func (l *Library) AddMember(member models.Member) error {
	l.Mu.Lock()
	defer l.Mu.Unlock()

	if _, exists := l.Members[member.ID]; exists {
		return fmt.Errorf("member with ID %d already exists", member.ID)
	}

	l.Members[member.ID] = member
	return nil
}
