package controllers
import (
	"bookstore/pkg/models"
	"bookstore/pkg/utils"
	"encoding/json"
	"fmt"

	//"log"
	"net/http"
	"strconv"
	//"time"

	//"html/template"
	//"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	//"golang.org/x/crypto/bcrypt"
	//"reflect"
)
func CreateCategory(writer http.ResponseWriter, request *http.Request) {
	CreateCategory := &models.Category{}
	utils.ParseBody(request, CreateCategory)
	category := CreateCategory.CreateCategory()
	res, _ := json.Marshal(category)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(res)
}

func UpdateCategory(writer http.ResponseWriter, request *http.Request) {
	var updateCategory = &models.Category{}
	utils.ParseBody(request, updateCategory)
	//log.Println(updateBook)
	vars := mux.Vars(request)
	categoryId := vars["categoryID"]
	id, err := strconv.ParseInt(categoryId, 0, 0)
	if err != nil {
		fmt.Println("Parse Error: ", err.Error())
	}

	categoryDetails, db := models.GetCategoryById(id)
	
		if updateCategory.Name != "" {
			categoryDetails.Name = updateCategory.Name
		}
	
		db.Save(&categoryDetails)
		res, _ := json.Marshal(categoryDetails)
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		writer.Write(res)
	
}

func DeleteCategory(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	categoryId := vars["categoryID"]
	id, err := strconv.ParseInt(categoryId, 0, 0)
	if err != nil {
		fmt.Println("Parse Error: ", err.Error())
	}
	category := models.DeleteCategory(id)
	result, _ := json.Marshal(category)
	
	writer.WriteHeader(http.StatusNoContent)
	writer.Write(result)
}
func AdminUpdateBook(writer http.ResponseWriter, request *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(request, updateBook)
	//log.Println(updateBook)
	vars := mux.Vars(request)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Parse Error: ", err.Error())
	}

	bookDetails, db := models.GetBookById(id)
	
		if updateBook.Name != "" {
			bookDetails.Name = updateBook.Name
		}
	
		if updateBook.Author != "" {
			bookDetails.Author = updateBook.Author
		}
	
		if updateBook.Publication != "" {
			bookDetails.Publication = updateBook.Publication
		}
	
		if updateBook.Price != 0 {
			bookDetails.Price = updateBook.Price
		}
		
		if updateBook.CategoryID != 0 {
			bookDetails.CategoryID = updateBook.CategoryID
		}
	
		db.Save(&bookDetails)
		res, _ := json.Marshal(bookDetails)
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		writer.Write(res)
	
}

func AdminDeleteBook(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Parse Error: ", err.Error())
	}
		book := models.DeleteBook(id)
		result, _ := json.Marshal(book)
	
		writer.WriteHeader(http.StatusNoContent)
		writer.Write(result)


}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var id = params["id"]
	var user models.User
	db := models.GetDB()
	db.First(&user, id)
	db.Delete(&user)
	res, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}