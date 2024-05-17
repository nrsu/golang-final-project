package routes

import (
	"bookstore/pkg/controllers"
	//"bookstore/pkg/utils"
	"github.com/gorilla/mux"
	"bookstore/pkg/auth"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {

	//reg log
	router.HandleFunc("/register", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/login", controllers.Login).Methods("POST")
	s := router.PathPrefix("/auth").Subrouter()
	s.Use(auth.JwtVerify)

	//users
	s.HandleFunc("/user", controllers.FetchUsers).Methods("GET")
	s.HandleFunc("/user/{id}", controllers.GetUser).Methods("GET")
	//s.HandleFunc("/user/{id}", controllers.UpdateUser).Methods("PUT")
	//s.HandleFunc("/user/{id}", controllers.DeleteUser).Methods("DELETE")

	//books
	s.HandleFunc("/book", controllers.CreateBook).Methods("POST")
	s.HandleFunc("/book", controllers.GetBook).Methods("GET")
	s.HandleFunc("/book/categories/{categoryId}", controllers.GetBooksByCategory).Methods("GET")
	s.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	s.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	s.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")

	//categories
	s.HandleFunc("/categories", controllers.GetCategories).Methods("GET")
	s.HandleFunc("/categories/{categoryID}", controllers.GetCategoriesByID).Methods("GET")

	//admin
	a := s.PathPrefix("/admin").Subrouter()
	a.Use(auth.AdminVerify)

	//books
	a.HandleFunc("/book", controllers.GetBook).Methods("GET") 
	a.HandleFunc("/book", controllers.CreateBook).Methods("POST") 
	a.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET") 
	a.HandleFunc("/book/{bookId}", controllers.AdminUpdateBook).Methods("PUT") 
	a.HandleFunc("/book/{bookId}", controllers.AdminDeleteBook).Methods("DELETE") 

	//users
	a.HandleFunc("/user", controllers.FetchUsers).Methods("GET") 
	a.HandleFunc("/user/{id}", controllers.GetUser).Methods("GET") 
	a.HandleFunc("/user/{id}", controllers.DeleteUser).Methods("DELETE")

	//categories
	a.HandleFunc("/categories", controllers.GetCategories).Methods("GET") 
	a.HandleFunc("/categories", controllers.CreateCategory).Methods("POST") 
	a.HandleFunc("/categories/{categoryID}", controllers.GetCategoriesByID).Methods("GET")
	a.HandleFunc("/categories/{categoryID}", controllers.UpdateCategory).Methods("PUT")
	a.HandleFunc("/categories/{categoryID}", controllers.DeleteCategory).Methods("DELETE")
	
}
