package controllers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"task3/library_management/models"
	"task3/library_management/services"
)

var reader = bufio.NewReader(os.Stdin)

func readLine(prompt string) string {
	fmt.Println(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func CreateBook() models.Book {
	var book models.Book
	fmt.Println("Enter Book ID:")
	fmt.Scanln(&book.ID)

	book.Title = readLine("Enter Book Title:")
	book.Author = readLine("Enter Book Author:")

	book.Status = "Available"
	return book
}

func CreateMember() models.Member {
	var member models.Member
	fmt.Println("Enter Member ID:")
	fmt.Scanln(&member.ID)

	member.Name = readLine("Enter Member Name:")
	member.BorrowedBooks = []models.Book{}
	return member
}

func Menu(lib *services.Library) {
	for {
		fmt.Println("\nConsole-Based Library Management System")
		fmt.Println("1. Add Book")
		fmt.Println("2. Remove Book")
		fmt.Println("3. Borrow Book")
		fmt.Println("4. Return Book")
		fmt.Println("5. List Available Books")
		fmt.Println("6. List Borrowed Books")
		fmt.Println("7. Add Member")
		fmt.Println("8. Exit")
		fmt.Print("Enter command: ")

		var command int
		fmt.Scanln(&command)

		switch command {
		case 1:
			book := CreateBook()
			lib.AddBook(book)
			fmt.Println("Book added successfully.")
		case 2:
			var bookID int
			fmt.Println("Enter Book ID to remove:")
			fmt.Scanln(&bookID)
			lib.RemoveBook(bookID)
			fmt.Println("Book removed successfully.")
		case 3:
			var bookID, memberID int
			fmt.Println("Enter Book ID to borrow:")
			fmt.Scanln(&bookID)
			fmt.Println("Enter Member ID:")
			fmt.Scanln(&memberID)
			if err := lib.BorrowBook(bookID, memberID); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Book borrowed successfully.")
			}
		case 4:
			var bookID, memberID int
			fmt.Println("Enter Book ID to return:")
			fmt.Scanln(&bookID)
			fmt.Println("Enter Member ID:")
			fmt.Scanln(&memberID)
			if err := lib.ReturnBook(bookID, memberID); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Book returned successfully.")
			}
		case 5:
			fmt.Println("Available Books:")
			for _, book := range lib.ListAvailableBooks() {
				fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
			}
		case 6:
			var memberID int
			fmt.Println("Enter Member ID to list borrowed books:")
			fmt.Scanln(&memberID)
			fmt.Println("Borrowed Books:")
			for _, book := range lib.ListBorrowedBooks(memberID) {
				fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
			}
		case 7:
			member := CreateMember()
			lib.AddMember(member)
			fmt.Println("Member added successfully.")
		case 8:
			fmt.Println("Exiting system...")
			return
		default:
			fmt.Println("Invalid command.")
		}
	}
}
