package controllers

import (
	"bookstore/pkg/models"
	"bookstore/pkg/utils"
	"encoding/json"
	"fmt"
	"log"

	//"log"
	"net/http"
	"strconv"
	"time"

	//"html/template"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	//"reflect"
)

//var NewBook models.Book
/*func GetBook(writer http.ResponseWriter, request *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(res)
}*/


func GetCategories(writer http.ResponseWriter, request *http.Request) {
	newCategories := models.GetAllCategories()
	res, _ := json.Marshal(newCategories)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(res)
}

func GetCategoriesByID(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	categoryId := vars["categoryID"]
	id, err := strconv.ParseInt(categoryId, 0, 0)
	if err != nil {
		fmt.Println("Parse Error: ", err.Error())
	}
	categoryDetails, _ := models.GetCategoryById(id)
	res, _ := json.Marshal(categoryDetails)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(res)
}

func GetBook(writer http.ResponseWriter, request *http.Request) {
    page, err := strconv.Atoi(request.URL.Query().Get("page"))
    if err != nil || page < 1 {
        page = 1
    }
    perPage, err := strconv.Atoi(request.URL.Query().Get("perPage"))
    if err != nil || perPage < 1 {
        perPage = 10 
    }

	sortDirection := request.URL.Query().Get("sort")
    if sortDirection != "asc" && sortDirection != "desc" {
        sortDirection = "none"
    }

	moreLess := request.URL.Query().Get("filtering")
    if moreLess != "more" && moreLess != "less" {
        moreLess = "none"
    }
	
	priceStr := request.URL.Query().Get("price")
    price, err := strconv.Atoi(priceStr)
    if err != nil {
        price=100
    }

	field := request.URL.Query().Get("field")
    if field != "name" && field != "author" && field != "publication" && field != "price"{
        field = "price"
    }
	
    newBooks := models.GetAllBooks(page, perPage, field, sortDirection, moreLess, price)

    res, err := json.Marshal(newBooks)
    if err != nil {
        http.Error(writer, err.Error(), http.StatusInternalServerError)
        return
    }
    writer.Header().Set("Content-Type", "application/json")
    writer.WriteHeader(http.StatusOK)
    writer.Write(res)
}

func GetBooksByCategory(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	categoryId := vars["categoryId"]
	id, err := strconv.ParseInt(categoryId, 0, 0)
	page, err := strconv.Atoi(request.URL.Query().Get("page"))
    if err != nil || page < 1 {
        page = 1
    }
    perPage, err := strconv.Atoi(request.URL.Query().Get("perPage"))
    if err != nil || perPage < 1 {
        perPage = 10 
    }

	sortDirection := request.URL.Query().Get("sort")
    if sortDirection != "asc" && sortDirection != "desc" {
        sortDirection = "none"
    }

	moreLess := request.URL.Query().Get("filtering")
    if moreLess != "more" && moreLess != "less" {
        moreLess = "none"
    }
	
	priceStr := request.URL.Query().Get("price")
    price, err := strconv.Atoi(priceStr)
    if err != nil {
        price=100
    }

	field := request.URL.Query().Get("field")
    if field != "name" && field != "author" && field != "publication" && field != "price"{
        field = "price"
    }
	if err != nil {
		fmt.Println("Parse Error: ", err.Error())
	}
	log.Println(GetCurrentUserID(request))
	books := models.GetBooksByCategory(id, page, perPage, field, sortDirection, moreLess, price)
	res, _ := json.Marshal(books)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(res)
}

func GetBookById(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Parse Error: ", err.Error())
	}
	bookDetails, _ := models.GetBookById(id)
	res, _ := json.Marshal(bookDetails)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(res)
}

func CreateBook(writer http.ResponseWriter, request *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(request, CreateBook)
	userID := GetCurrentUserID(request)
	book := CreateBook.CreateBook(userID)
	res, _ := json.Marshal(book)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(res)
}

func DeleteBook(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Parse Error: ", err.Error())
	}
	bookDetails, _ := models.GetBookById(id)
	if(CurrentUsersBook(bookDetails, request)){
		book := models.DeleteBook(id)
		result, _ := json.Marshal(book)
	
		writer.WriteHeader(http.StatusNoContent)
		writer.Write(result)
	}else{
		http.NotFound(writer, request)
	}


}
func CurrentUsersBook(book *models.Book, request *http.Request) bool{
	userID := GetCurrentUserID(request)
	if(book.UserID == userID){
		return true
	}else{
		return false
	}
}

func GetCurrentUserRole(r *http.Request) (bool) {
	// Extract user information from request context
	claims:= r.Context().Value("user").(*models.Token)
		return claims.Admin
}

func UpdateBook(writer http.ResponseWriter, request *http.Request) {
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
	if(CurrentUsersBook(bookDetails, request)){
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
	}else{
		http.NotFound(writer, request)
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request){
	user := &models.User{}
	json.NewDecoder(r.Body).Decode(user)
	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		json.NewEncoder(w).Encode(err)
	}
	user.Password = string(pass)
	createdUser := user.CreateUser()
	res, _ := json.Marshal(createdUser)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

var token_string = ""

func Login(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user) 
	if err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Invalid request"}
		json.NewEncoder(w).Encode(resp)
		return
	}
	resp := FindOne(user.Email, user.Password)
	token_string = resp["token"].(string)
	res, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetToken() string{
	return token_string
}


func FindOne(email, password string) map[string]interface{} {
	user := &models.User{}
	db := models.GetDB()
	if err := db.Where("Email = ?", email).First(user).Error; err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Email address not found"}
		return resp
	}
	expiresAt := time.Now().Add(time.Minute * 100000).Unix()
	errf := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		var resp = map[string]interface{}{"status": false, "message": "Invalid login credentials. Please try again"}
		return resp
	}

	tk := &models.Token{
		UserID: user.ID,
		Name:   user.Name,
		Email:  user.Email,
		Gender: user.Gender,
		Password: user.Password,
		Admin: user.Admin,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)

	tokenString, error := token.SignedString([]byte("secret"))
	if error != nil {
		fmt.Println(error)
	}

	var resp = map[string]interface{}{"status": false, "message": "logged in"}
	resp["token"] = tokenString //Store the token in the response
	resp["user"] = user
	return resp
}





//FetchUser function
func FetchUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db := models.GetDB()
	db.Preload("auths").Find(&users)

	res, _ := json.Marshal(users)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

/*func UpdateUser(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	params := mux.Vars(r)
	db := models.GetDB()
	var id = params["id"]
	db.First(&user, id)
	json.NewDecoder(r.Body).Decode(user)
	db.Save(&user)
	json.NewEncoder(w).Encode(&user)
}*/

func GetCurrentUserEmail(r *http.Request) (string) {
	// Extract user information from request context
	claims:= r.Context().Value("user").(*models.Token)
		return claims.Email
}

func GetCurrentUserID(r *http.Request) (uint) {
	// Extract user information from request context
	claims:= r.Context().Value("user").(*models.Token)
		return claims.UserID
}



func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var id = params["id"]
	var user models.User
	db := models.GetDB()
	db.First(&user, id)
	res, _ := json.Marshal(&user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
