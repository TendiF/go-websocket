package utils 

import (
	"os"
	"net/http"
	// "log"
	"strings"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
)
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		type MyCustomClaims struct {
			IdUser string `json:"user_id"`
			jwt.StandardClaims
		}

		if r.URL.Path == "/user/login" {
				next.ServeHTTP(w, r)
				return
		}

		authorizationHeader := r.Header.Get("Authorization")
		if !strings.Contains(authorizationHeader, "Bearer") {
				http.Error(w, "Invalid token", http.StatusBadRequest)
				return
		}

		tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)

		token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			http.Error(w, "Error Parse token", http.StatusBadRequest)
		}

		if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
			context.Set(r, "user_id", claims.IdUser)
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Not Allowed", http.StatusMethodNotAllowed)
		}

})
}