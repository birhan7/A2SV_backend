Library Management System - Documentation

Overview

This is a simple console-based Library Management System built with Go. The project demonstrates the use of structs, interfaces, methods, slices, and maps in Go. It supports basic library operations such as adding/removing books, borrowing/returning books, and listing available or borrowed books by a member.

Folder Structure

library_management/
├── main.go
├── controllers/
│ └── library_controller.go
├── models/
│ ├── book.go
│ └── member.go
├── services/
│ └── library_service.go
├── docs/
│ └── documentation.md
└── go.mod

Structs

Book

Defined in models/book.go

type Book struct {
ID int
Title string
Author string
Status string // "Available" or "Borrowed"
}

Member

Defined in models/member.go

type Member struct {
ID int
Name string
BorrowedBooks []Book
}

Interfaces

LibraryManager

Defined in services/library_service.go

type LibraryManager interface {
AddBook(book Book)
RemoveBook(bookID int)
BorrowBook(bookID int, memberID int) error
ReturnBook(bookID int, memberID int) error
ListAvailableBooks() []Book
ListBorrowedBooks(memberID int) []Book
}

Implementation

Library

Implemented in services/library_service.go

type Library struct {
Books map[int]Book
Members map[int]Member
}

The Library struct implements the LibraryManager interface with the following methods:

AddBook(book Book): Adds a new book to the library.

RemoveBook(bookID int): Removes a book by ID.

BorrowBook(bookID int, memberID int): Allows a member to borrow a book if it is available.

ReturnBook(bookID int, memberID int): Allows a member to return a book.

ListAvailableBooks() []Book: Lists all books that are currently available.

ListBorrowedBooks(memberID int) []Book: Lists all books borrowed by the given member.

Console Interaction

Implemented in controllers/library_controller.go. Functions are provided for user interaction:

addBookPrompt() - Prompts user to enter book details and adds it to the library.

removeBookPrompt() - Prompts user to enter a book ID and removes it.

borrowBookPrompt() - Prompts for book and member ID, attempts to borrow the book.

returnBookPrompt() - Prompts for book and member ID, attempts to return the book.

listAvailableBooksPrompt() - Displays all available books.

listBorrowedBooksPrompt() - Displays all books borrowed by a specific member.

Each function interacts with the service layer (Library) and prints appropriate messages.

Error Handling

Implemented robustly across all service methods:

Verifies if the book or member exists.

Checks if a book is already borrowed.

Prevents duplicate borrowing.

Validates returning books that are not borrowed.

Appropriate error values are returned and logged in the controller layer.

Running the Application

Navigate to the library_management/ folder.

Run the application:

go run main.go

Follow the console prompts to interact with the system.

Example Usage

Add Book: Enter book ID, title, and author.

Add Member: Member is pre-loaded or can be extended.

Borrow Book: Enter book ID and member ID.

Return Book: Enter book ID and member ID.

List Books: Show either available or borrowed by member.

Authors

Birhan Aklilu
