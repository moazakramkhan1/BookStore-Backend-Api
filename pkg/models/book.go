package models

import (
	"github.com/jinzhu/gorm"
	"github.com/moaz/go-BookStoreProject/pkg/config"
)

var db gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = *config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(b)
	return b
}
func GetBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}
func GetBooksbyID(id int64) (*Book, *gorm.DB) {
	var getbook Book
	db := db.Where("ID=?", id).Find(&getbook)
	return &getbook, db
}
func DeleteBook(ID int64) Book {
	var deletedBook Book
	db.Where("ID=?", ID).Delete(&deletedBook)
	return deletedBook
}
