package auth

import (
	"bookstore/pkg/controllers"
	"bookstore/pkg/models"
	"context"
	"encoding/json"

	//"log"
	"net/http"

	//"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

//Exception struct
type Exception models.Exception

// JwtVerify Middleware function
func JwtVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		/*var header = r.Header.Get("Authorization") //Grab the token from the header
		const prefix = "Bearer "
        if strings.HasPrefix(header, prefix) {
            // Remove the prefix
            header = header[len(prefix):]
        }
		header = strings.TrimSpace(header)*/
		header := controllers.GetToken()
		//log.Println(header)
		if header == "" {
			//Token is missing, returns with error code 403 Unauthorized
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(Exception{Message: "Missing auth token"})
			return
		}
		tk := &models.Token{}

		_, err := jwt.ParseWithClaims(header, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})

		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(Exception{Message: err.Error()})
			return
		}

		ctx := context.WithValue(r.Context(), "user", tk)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
	
}
func AdminVerify(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Check if the user is authenticated as an admin
        if controllers.GetCurrentUserRole(r) {
            // If user is an admin, proceed with the next handler
            next.ServeHTTP(w, r)
            return
        }
        // If user is not an admin, return unauthorized error
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
    })
}