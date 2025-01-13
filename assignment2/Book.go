package assignment2

import "fmt"

type Book struct {
	Title           string
	Author          string
	Pages           int
	CopiesAvailable int
}

func CreateBook(Title1, Author1 string, Pages1, CopiesAvailable1 int) *Book {
	return &Book{
		Title:           Title1,
		Author:          Author1,
		Pages:           Pages1,
		CopiesAvailable: CopiesAvailable1,
	}
}

func (b Book) DisplayBooks() {
	fmt.Printf("Book Title : %s\n", b.Title)
	fmt.Printf("Book TAuthor : %s\n", b.Author)
	fmt.Printf("Book Title : %d\n", b.Pages)
	fmt.Printf("Book Title : %d\n", b.CopiesAvailable)
	fmt.Println()
}

var message string

func (b *Book) Borrow() string {
	if b.CopiesAvailable > 0 {
		b.CopiesAvailable = b.CopiesAvailable - 1
		fmt.Println()
		message = "Borrow Successful"
	} else {
		fmt.Println()
		message = "Borrow Unsuccessful"
	}
	fmt.Println(b.Title, b.Author, b.Pages, b.CopiesAvailable)
	return message
}

func (b *Book) ReturnBook() {
	fmt.Println("\nincrements by 1: ")
	b.CopiesAvailable += 1
	fmt.Println(b.Title, b.Author, b.Pages, b.CopiesAvailable)
}

func SwapTitles(a, b *Book) {
	fmt.Println("\nBefore Swapping: ", a.Title, b.Title)
	temp := a.Title
	a.Title = b.Title
	b.Title = temp
	fmt.Println("After Swapping: ", a.Title, b.Title)
}
