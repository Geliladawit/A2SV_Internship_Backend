# Library Management System Documentation

## Overview

This is a simple console-based library management system implemented in Go. It allows users to:

*   Add new books to the library.
*   Remove existing books from the library.
*   Borrow books.
*   Return books.
*   List all available books.
*   List all books borrowed by a specific member.

## Folder Structure

*   `main.go`: Entry point of the application.
*   `controllers/`: Contains the `library_controller.go` file, which handles console input and invokes the appropriate service methods.
*   `models/`: Contains the `book.go` and `member.go` files, which define the `Book` and `Member` structs, respectively.
*   `services/`: Contains the `library_service.go` file, which contains the business logic and data manipulation functions.
*   `docs/`: Contains the `documentation.md` file.
*   `go.mod`: Defines the module and its dependencies.

## Usage

1.  Compile the application: `go build`
2.  Run the application: `./library_management`

The application will present a menu of options. Follow the prompts to interact with the library management system.

## Data Structures

*   **Book:** Represents a book in the library.
    *   `ID`: Unique identifier for the book (integer).
    *   `Title`: Title of the book (string).
    *   `Author`: Author of the book (string).
    *   `Status`: Status of the book ("Available" or "Borrowed").
*   **Member:** Represents a member of the library.
    *   `ID`: Unique identifier for the member (integer).
    *   `Name`: Name of the member (string).
    *   `BorrowedBooks`: A slice of `Book` structs representing the books currently borrowed by the member.

## Interface

*   **LibraryManager:** Defines the interface for interacting with the library service.
    *   `AddBook(book Book) error`: Adds a new book to the library. Returns an error if a book with the same ID already exists.
    *   `RemoveBook(bookID int)`: Removes a book from the library by its ID.
    *   `BorrowBook(bookID int, memberID int) error`: Allows a member to borrow a book if it is available.
    *   `ReturnBook(bookID int, memberID int) error`: Allows a member to return a borrowed book.
    *   `ListAvailableBooks() []Book`: Lists all available books in the library.
    *   `ListBorrowedBooks(memberID int) []Book`: Lists all books borrowed by a specific member.
    *   `AddMember(member Member) error`: Adds a new member to the library. Returns an error if a member with the same ID already exists.

## Error Handling

The application includes error handling for scenarios such as:

*   Invalid user input (e.g., non-numeric input when a number is expected).
*   Attempting to borrow a book that is already borrowed.
*   Attempting to return a book that is not borrowed.
*   Attempting to add a book or member with an ID that already exists.  The system now prevents duplicate IDs and displays an error message if a user tries to add a book or member with an ID that is already in use.
*   Book or member not found.