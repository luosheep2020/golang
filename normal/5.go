package main

import "fmt"

type book struct {
	title   string
	author  string
	subject string
	bookId  int
}

func main() {
	fmt.Print(book{
		title:   "三国演义",
		author:  "罗贯中",
		subject: "语文",
		bookId:  0,
	})

	fmt.Print(book{bookId: 0})
}
