package books

import (
	"strconv"

	"github.com/pallinder/go-randomdata"
)

var Books []Book

type Book struct {
	BookId      int    `json:"id"`
	Title       string `json:"title"`
	Publication string `json:"publication"`
	Ratings     int    `json:"ratings"`
	Author      Author `json:"author"`
}

type Author struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func New(bookid, ratings int, title, publication, firstName, lastName string) Book {
	return Book{
		BookId:      bookid,
		Title:       title,
		Publication: publication,
		Ratings:     ratings,
		Author: Author{
			FirstName: firstName,
			LastName:  lastName,
		},
	}
}

func AddBooks(count int) {
	for count != 0 {
		bookid, _ := strconv.Atoi(randomdata.Digits(5))
		ratings, _ := strconv.Atoi(randomdata.Digits(1))
		book := New(bookid, ratings, randomdata.SillyName(), randomdata.Adjective(), randomdata.FirstName(2), randomdata.LastName())
		Books = append(Books, book)
		count--
	}
}

// func (bookdata Book) Display_book() {
// 	fmt.Println("------------------------------")
// 	fmt.Printf("Book ID - %v\n", bookdata.BookId)
// 	fmt.Printf("Book Title - %v\n", bookdata.Title)
// 	fmt.Printf("Book Publication - %v\n", bookdata.Publication)
// 	fmt.Printf("Book Ratings - %v\n", bookdata.Ratings)
// 	fmt.Printf("Book Author First Name - %v\n", bookdata.Author.FirstName)
// 	fmt.Printf("Book Author Last Name - %v\n", bookdata.Author.LastName)
// 	fmt.Println("------------------------------")
// }
