package main

import (
	"example/assignment2"
	"fmt"
)

func main() {

	book1 := assignment2.CreateBook("Winter", "Harini", 500, 2)
	book2 := assignment2.CreateBook("Summer", "Bhavani", 400, 10)
	book3 := assignment2.CreateBook("Spring", "Swetha", 380, 0)
	book1.DisplayBooks()
	book2.DisplayBooks()
	book3.DisplayBooks()
	fmt.Println(book1.Borrow())
	fmt.Println(book2.Borrow())
	fmt.Println(book3.Borrow())
	book1.ReturnBook()
	book2.ReturnBook()
	book3.ReturnBook()
	assignment2.SwapTitles(book1, book2)
	//assignment1.Stud()
}
