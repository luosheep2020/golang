package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	bookId  int
}
func main() {
	book := Books{
		title:   "Go语言教程",
		author:  "罗铭扬",
		subject: "计算机科学",
		bookId:  1,
	}
	var books *Books
	books=&book
	fmt.Print(book)
	fmt.Print(books)
}