package models

import(
	"github.com/jinzhu/gorm"
	"github.com/FranciscoGontijo/book-management-system/pkg/config"
)

var db *gorm.DB

type Book struct{
	gorm.Model
	Name string `gorm:"" json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}

func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book{
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book{
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	result := db.Where("ID=?", Id).First(&getBook)
	return &getBook, result
}

func DeleteBook(Id int64) (Book, *gorm.DB) {
	var book Book
	result := db.First(&book, Id)

	if result.Error != nil {
		return Book{}, result
	}

	deleteResult := db.Delete(&book)
	return book, deleteResult
}