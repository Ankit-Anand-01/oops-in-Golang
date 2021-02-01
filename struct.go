// Golang program to illustrate the
// concept of custom types
package main

import (
	"fmt"
)

// declaring a struct
type Book struct {

	// defining struct variables
	name   string
	author string
	pages  int
}

// function to print book details
func (book Book) print_details() {

	fmt.Printf("%s was written by %s.", book.name, book.author)
	fmt.Printf("\nIt contains %d pages.\n", book.pages)
}

// main function
func main() {

	// declaring a struct instance
	book1 := Book{"It", "Stephen King", 1156}
	book2 := Book{"The Blue Umbrella", "Ruskin Bond", 56}
	book3 := Book{"The Monk Who Sold his Ferrari", "Robin Sharma", 224}

	// printing details of book1
	book1.print_details()
	//printing details of book2
	book2.print_details()
	book3.print_details()

}
