package controllers

import (
	"bufio"
	"fmt"
	"library_management/models"
	"library_management/services"
	"os"
	"strconv"
	"strings"
)

type LibraryController struct {
	Library services.LibraryManager
}

func NewLibraryController(library services.LibraryManager) *LibraryController {
	return &LibraryController{Library: library}
}

func (c *LibraryController) Run() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nLibrary Management System")
		fmt.Println("1. Add Book")
		fmt.Println("2. Remove Book")
		fmt.Println("3. Borrow Book")
		fmt.Println("4. Return Book")
		fmt.Println("5. List Available Books")
		fmt.Println("6. List Borrowed Books (by Member)")
		fmt.Println("7. Add Member")
		fmt.Println("0. Exit")

		fmt.Print("Enter your choice: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			c.addBook(reader)
		case "2":
			c.removeBook(reader)
		case "3":
			c.borrowBook(reader)
		case "4":
			c.returnBook(reader)
		case "5":
			c.listAvailableBooks()
		case "6":
			c.listBorrowedBooks(reader)
		case "7":
			c.addMember(reader)
		case "0":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func (c *LibraryController) addBook(reader *bufio.Reader) {
	fmt.Print("Enter Book ID: ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid ID. Please enter a number.")
		return
	}

	book := models.Book{ID: id, Status: "Available"}
	err = c.Library.AddBook(book)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Print("Enter Book Title: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Print("Enter Book Author: ")
	author, _ := reader.ReadString('\n')
	author = strings.TrimSpace(author)

	fmt.Println("Book added successfully!")
}

func (c *LibraryController) removeBook(reader *bufio.Reader) {
	fmt.Print("Enter Book ID to remove: ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid ID. Please enter a number.")
		return
	}

	c.Library.RemoveBook(id)
	fmt.Println("Book removed successfully!")
}

func (c *LibraryController) borrowBook(reader *bufio.Reader) {
	fmt.Print("Enter Book ID to borrow: ")
	bookIDStr, _ := reader.ReadString('\n')
	bookIDStr = strings.TrimSpace(bookIDStr)
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		fmt.Println("Invalid Book ID. Please enter a number.")
		return
	}

	fmt.Print("Enter Member ID: ")
	memberIDStr, _ := reader.ReadString('\n')
	memberIDStr = strings.TrimSpace(memberIDStr)
	memberID, err := strconv.Atoi(memberIDStr)
	if err != nil {
		fmt.Println("Invalid Member ID. Please enter a number.")
		return
	}

	err = c.Library.BorrowBook(bookID, memberID)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Book borrowed successfully!")
	}
}

func (c *LibraryController) returnBook(reader *bufio.Reader) {
	fmt.Print("Enter Book ID to return: ")
	bookIDStr, _ := reader.ReadString('\n')
	bookIDStr = strings.TrimSpace(bookIDStr)
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		fmt.Println("Invalid Book ID. Please enter a number.")
		return
	}

	fmt.Print("Enter Member ID: ")
	memberIDStr, _ := reader.ReadString('\n')
	memberIDStr = strings.TrimSpace(memberIDStr)
	memberID, err := strconv.Atoi(memberIDStr)
	if err != nil {
		fmt.Println("Invalid Member ID. Please enter a number.")
		return
	}

	err = c.Library.ReturnBook(bookID, memberID)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Book returned successfully!")
	}
}

func (c *LibraryController) listAvailableBooks() {
	availableBooks := c.Library.ListAvailableBooks()
	if len(availableBooks) == 0 {
		fmt.Println("No available books in the library.")
		return
	}

	fmt.Println("Available Books:")
	for _, book := range availableBooks {
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
	}
}

func (c *LibraryController) listBorrowedBooks(reader *bufio.Reader) {
	fmt.Print("Enter Member ID to list borrowed books: ")
	memberIDStr, _ := reader.ReadString('\n')
	memberIDStr = strings.TrimSpace(memberIDStr)
	memberID, err := strconv.Atoi(memberIDStr)
	if err != nil {
		fmt.Println("Invalid Member ID. Please enter a number.")
		return
	}

	borrowedBooks := c.Library.ListBorrowedBooks(memberID)
	if len(borrowedBooks) == 0 {
		fmt.Println("No books borrowed by this member.")
		return
	}

	fmt.Println("Borrowed Books:")
	for _, book := range borrowedBooks {
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
	}
}

func (c *LibraryController) addMember(reader *bufio.Reader) {
	fmt.Print("Enter Member ID: ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid ID. Please enter a number.")
		return
	}
	member := models.Member{ID: id, BorrowedBooks: []models.Book{}}
	err = c.Library.AddMember(member)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Print("Enter Member Name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Println("Member added successfully!")
}
