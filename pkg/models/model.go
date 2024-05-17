package models

import (
	"bookstore/pkg/config"
	
	"github.com/dgrijalva/jwt-go"

	"gorm.io/gorm"
)

var db *gorm.DB

type Category struct {
    gorm.Model
    Name  string  `json:"name"`
}

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
	Price       int    `json:"price"` 
	CategoryID  uint    `json:"category_id"`
    Category    Category `json:"category" gorm:"foreignKey:CategoryID"`
	UserID    uint    `json:"user_id"`
	User      User   `json:"user" gorm:"foreignKey: UserID"`
}

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"type:varchar(100);unique_index"`
	Gender   string `json:"Gender"`
	Password string `json:"Password"`
	Admin    bool   `json:"Admin"`
}

type Token struct {
	UserID   uint
	Name     string
	Email    string
	Gender   string
	Password string
	Admin    bool
	*jwt.StandardClaims
}



func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Category{})
}

func GetDB() *gorm.DB{
	return db
}

func (b *Book) CreateBook(userID uint) *Book {
	b.UserID=userID
	db.Create(&b)
	return b
}

func (c *Category) CreateCategory() *Category {
	db.Create(&c)
	return c
}

/*func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}*/
func GetAllBooks(page, perPage int, column string, order string, moreLess string, price int) []Book {
    var books []Book
    offset := (page - 1) * perPage

	if order=="none" && moreLess=="none"{
		db.Preload("User").Preload("Category").Offset(offset).Limit(perPage).Find(&books)
	}

	if order!="none" && moreLess=="none"{
		db.Preload("User").Preload("Category").Offset(offset).Limit(perPage).Order(column + " " + order).Find(&books)
	}

	if order=="none" && moreLess!="none"{
		if moreLess=="less"{
			db.Preload("User").Preload("Category").Offset(offset).Limit(perPage).Where("price < ?", price).Find(&books)
		}else{
			db.Preload("User").Preload("Category").Offset(offset).Limit(perPage).Where("price > ?", price).Find(&books)
		}
	}

	if order!="none" && moreLess!="none"{
		if moreLess=="less"{
			db.Preload("User").Preload("Category").Offset(offset).Limit(perPage).Where("price < ?", price).Order(column + " " + order).Find(&books)
		}else{
			db.Preload("User").Preload("Category").Offset(offset).Limit(perPage).Where("price > ?", price).Order(column + " " + order).Find(&books)
		}
	}

    return books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Preload("User").Preload("Category").Where("ID=?", Id).Find(&getBook)
	return &getBook, db

}

func GetCategoryById(Id int64) (*Category, *gorm.DB) {
	var getCategory Category
	db := db.Where("ID=?", Id).Find(&getCategory)
	return &getCategory, db

}

func GetBooksByCategory(Id int64, page, perPage int, column string, order string, moreLess string, price int) []Book {
	var books []Book
	offset := (page - 1) * perPage

	if order=="none" && moreLess=="none"{
		db.Preload("User").Preload("Category").Offset(offset).Limit(perPage).Where("category_id=?", Id).Find(&books)
	}

	if order!="none" && moreLess=="none"{
		db.Preload("User").Preload("Category").Offset(offset).Limit(perPage).Order(column + " " + order).Where("category_id=?", Id).Find(&books)
	}

	if order=="none" && moreLess!="none"{
		if moreLess=="less"{
			db.Preload("User").Preload("Category").Offset(offset).Limit(perPage).Where("price < ?", price).Where("category_id=?", Id).Find(&books)
		}else{
			db.Preload("User").Preload("Category").Offset(offset).Limit(perPage).Where("price > ?", price).Where("category_id=?", Id).Find(&books)
		}
	}

	if order!="none" && moreLess!="none"{
		if moreLess=="less"{
			db.Preload("User").Preload("Category").Offset(offset).Limit(perPage).Where("price < ?", price).Where("category_id=?", Id).Order(column + " " + order).Find(&books)
		}else{
			db.Preload("User").Preload("Category").Offset(offset).Limit(perPage).Where("price > ?", price).Where("category_id=?", Id).Order(column + " " + order).Find(&books)
		}
	}
	//db.Preload("Category").Where("category_id=?", Id).Find(&books)
	return books

}

func GetAllCategories()[]Category{
	var categories []Category
	db.Find(&categories)
	return categories
}

func DeleteBook(Id int64) Book {
	var book Book
	db.Delete(&book, Id)
	return book

}
func DeleteCategory(Id int64) Category {
	var category Category
	db.Delete(&category, Id)
	return category

}

/*func GetBooksCheaperThan(price int) []Book {
    var books []Book
    db.Where("price < ?", price).Find(&books)
    return books
}*/

func (user *User) CreateUser() *User {
	//db.NewRecord(b)
	db.Create(&user)
	return user
}



