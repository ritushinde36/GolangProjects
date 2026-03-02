package models

import (
	"books-crud-mysql/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

// initialize the database
func init() {
	config.Connect()
	db = config.GetDB()
	//this is to migrate you schema so that it is up to date
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBooksSQL() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetBooksSQL() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetBookByIDSQL(id int) (Book, *gorm.DB) {
	var book Book
	db := db.Where("ID=?", id).Find(&book)
	return book, db
}

func DeleteBookSQL(id int) Book {
	var book Book
	db.Where("ID=?", id).Delete(book)
	return book
}
