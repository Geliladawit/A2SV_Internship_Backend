# Library Management System Documentation

## Overview

This is a console-based library management system implemented in Go. It allows users to:

*   Add new books to the library.
*   Remove existing books from the library.
*   Borrow books.
*   Return books.
*   Reserve Books
*   List all available books.
*   List all books borrowed by a specific member.

## Folder Structure

*   `main.go`: Entry point of the application.
*   `controllers/`: Contains the `library_controller.go` file, which handles console input and invokes the appropriate service methods.  It also initializes the concurrent reservation workers.
*   `models/`: Contains the `book.go` and `member.go` files, which define the `Book` and `Member` structs, respectively. The `Book` struct now includes a `Status` field that can be "Available", "Borrowed", or "Reserved".
*   `services/`: Contains the `library_service.go` file, which contains the business logic and data manipulation functions for the core library operations. It manages the book and member data, and it handles the queuing of reservation requests.
*   `concurrency/`: Contains the `reservation_worker.go` file, which defines the `ReservationWorker` and handles concurrent book reservation requests using Goroutines and Channels.
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
    *   `Status`: Status of the book ("Available", "Borrowed", or "Reserved").
*   **Member:** Represents a member of the library.
    *   `ID`: Unique identifier for the member (integer).
    *   `Name`: Name of the member (string).
    *   `BorrowedBooks`: A slice of `Book` structs representing the books currently borrowed by the member.

## Interface

*   **LibraryManager:** Defines the interface for interacting with the library service.
    *   `AddBook(book Book) error`: Adds a new book to the library. An error is returned if a book with the same ID already exists.
    *   `RemoveBook(bookID int)`: Removes a book from the library by its ID.
    *   `BorrowBook(bookID int, memberID int) error`: Allows a member to borrow a book if it is available.
    *   `ReturnBook(bookID int, memberID int) error`: Allows a member to return a borrowed book.
    *   `ReserveBook(bookID int, memberID int) error`: Allows a member to reserve a book if it is available. This operation is handled concurrently.
    *   `ListAvailableBooks() []Book`: Lists all available books in the library.
    *   `ListBorrowedBooks(memberID int) []Book`: Lists all books borrowed by a specific member.
     *   `AddMember(member Member) error`: Adds a new member to the library. An error is returned if a member with the same ID already exists.

## Concurrency Implementation

The system utilizes Goroutines, Channels, and Mutexes to handle concurrent book reservation requests.

*   **Goroutines:** Multiple `ReservationWorker` Goroutines are launched at application startup to handle incoming reservation requests concurrently.  These workers are initialized in the `controllers/library_controller.go` file when the `Run()` method is invoked.  The number of workers can be configured.

*   **Channels:** The `Library` struct contains a `ReservationQueue` channel, which is a buffered channel used to pass reservation requests from the `ReserveBook` method in `library_service.go` to the `ReservationWorker` Goroutines.

*   **Mutexes:** A `sync.Mutex` (named `Mu` in the `Library` struct) is used to protect the `Books` and `Members` maps from race conditions. The mutex is locked and unlocked by the `ReservationWorker` and the add functions before and after accessing or modifying the data. This ensures that only one Goroutine can update a book's or member's status at a time.

*   **Reservation Workers:** The `concurrency/reservation_worker.go` file defines the `ReservationWorker` type, which is responsible for processing individual reservation requests. Each worker has its own Goroutine that listens for requests on the `ReservationQueue`.

### Auto-Cancellation

If a reserved book is not borrowed within 5 seconds of being reserved, the system automatically cancels the reservation, making the book available again. This is implemented using `time.AfterFunc` within the `processReservation` method of the `ReservationWorker`.

## Error Handling

The application includes error handling for scenarios such as:

*   Invalid user input (e.g., non-numeric input when a number is expected).
*   Attempting to borrow a book that is already borrowed or reserved.
*   Attempting to return a book that is not borrowed.
*   Attempting to reserve a book that is already borrowed or reserved.
*   Attempting to add a book or member with an ID that already exists.
*   Book or member not found.
*   Errors encountered during concurrent operations are reported back to the user.

## Notes

*   For simplicity, the current implementation does not track *which* member reserved a book. Therefore, any member can borrow a reserved book.  A more robust implementation would store the reserving member's ID and verify it during the borrow operation.
*   The number of reservation workers can be adjusted to optimize performance based on the expected level of concurrency.