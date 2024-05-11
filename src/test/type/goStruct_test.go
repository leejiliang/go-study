package _type

import "testing"

type Book struct {
	title   string
	author  string
	subject string
	bookId  int
}

func TestBookInit(t *testing.T) {
	book1 := Book{title: "Go Programming Language", author: "Go Author", subject: "Go Programming Language", bookId: 1}
	book2 := Book{"title", "Go Author", "Go Programming Language", 1}
	t.Log(book1)
	t.Log(book2)

	t.Log("book1 title is : ", book1.title)
	book1.title = "Go Programming Language v2"
	t.Log("book1 title is : ", book1.title)
	t.Log("book1 title is : ", getTitle(book1))
	t.Log("book1 author is : ", getAuthor(&book1))
}

func getTitle(book Book) string {
	return book.title
}

func getAuthor(book *Book) string {
	return book.author
}
