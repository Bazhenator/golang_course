package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"task_7/models"
	u "task_7/utils"

	"github.com/dgrijalva/jwt-go"
)

var JwtAuthentication = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		notAuth := []string{"/user/new", "/user/login", "/users"}
		requestPath := r.URL.Path
		for _, value := range notAuth {
			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		tokenHeader := r.Header.Get("Authorization")

		if tokenHeader == "" {
			response := u.NewError(http.StatusForbidden, 403, "Missing authentication token")
			u.JSONError(w, response)
			return
		}

		splitted := strings.Split(tokenHeader, " ")
		if len(splitted) != 2 {
			response := u.NewError(http.StatusForbidden, 403, "Invalid/Malformed authentication token")
			u.JSONError(w, response)
			return
		}

		tokenPart := splitted[1]
		tk := &models.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token_password")), nil
		})

		if err != nil {
			response := u.NewError(http.StatusForbidden, 403, "Malformed authentication token")
			u.JSONError(w, response)
			return
		}

		if !token.Valid {
			response := u.NewError(http.StatusForbidden, 403, "Token is not valid.")
			u.JSONError(w, response)
			return
		}

		fmt.Sprintf("User %", tk.UserId)
		ctx := context.WithValue(r.Context(), "user", tk.UserId)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
