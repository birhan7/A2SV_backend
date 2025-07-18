package services

import (
	"fmt"
	"task3/library_management/models"
)

type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
}

type Library struct {
	Books   map[int]models.Book
	Members map[int]models.Member
}

// methods that can be called on Library type
func (lib *Library) AddBook(book models.Book) {
	_, ok := lib.Books[book.ID]

	if ok {
		fmt.Printf("Error: Book with id %d already found.\n", book.ID)
		return
	}

	lib.Books[book.ID] = book
	fmt.Printf("Book with id %d added successfully.\n", book.ID)
}

func (lib *Library) AddMember(member models.Member) {
	_, ok := lib.Members[member.ID]

	if ok {
		fmt.Printf("Error: Member with id %d already found.\n", member.ID)
		return
	}

	lib.Members[member.ID] = member
	fmt.Printf("Member with id %d added successfully.\n", member.ID)
}

func (lib *Library) RemoveBook(bookId int) {
	_, ok := lib.Books[bookId]

	if !ok {
		println("Error: Book not found")
		return
	}
	delete(lib.Books, bookId)
	fmt.Printf("Book with id %d removed successfully.\n", bookId)
}

func (lib *Library) BorrowBook(bookId, memberId int) error {
	book, ok := lib.Books[bookId]
	if !ok {
		return fmt.Errorf("book with id %d not found", bookId)
	}

	member, ok := lib.Members[memberId]
	if !ok {
		return fmt.Errorf("member with id %d not found", memberId)
	}

	if book.Status == "Borrowed" {
		return fmt.Errorf("book with id %d is already borrowed", bookId)
	}

	member.BorrowedBooks = append(member.BorrowedBooks, book)
	lib.Members[memberId] = member

	book.Status = "Borrowed"
	lib.Books[bookId] = book

	fmt.Printf("Book with id %d borrowed by member %d successfully.\n", bookId, memberId)
	return nil
}

func (lib *Library) ReturnBook(bookId, memberId int) error {
	book, ok := lib.Books[bookId]
	if !ok {
		return fmt.Errorf("book with id %d not found", bookId)
	}

	member, ok := lib.Members[memberId]
	if !ok {
		return fmt.Errorf("member with id %d not found", memberId)
	}

	if book.Status == "Available" {
		return fmt.Errorf("book with id %d is already available", bookId)
	}

	for i, b := range member.BorrowedBooks {
		if b.ID == bookId {
			member.BorrowedBooks = append(member.BorrowedBooks[:i], member.BorrowedBooks[i+1:]...)
			break
		}
	}

	book.Status = "Available"
	lib.Books[bookId] = book
	lib.Members[memberId] = member
	fmt.Printf("Book with id %d returned by member %d successfully.\n", bookId, memberId)

	return nil
}

func (lib *Library) ListAvailableBooks() []models.Book {
	var availableBooks []models.Book
	for _, book := range lib.Books {
		if book.Status == "Available" {
			availableBooks = append(availableBooks, book)
		}
	}
	return availableBooks
}

func (lib *Library) ListBorrowedBooks(memberID int) []models.Book {
	member, ok := lib.Members[memberID]
	if !ok {
		fmt.Printf("Error: Member with id %d not found.\n", memberID)
		return nil
	}
	return member.BorrowedBooks
}
