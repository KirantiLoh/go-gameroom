package auth

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authToken := r.Header.Get("Authorization")

		if authToken == "" {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Unauthorized",
			})
			return
		}

		claims, err := DecodeJWT(strings.TrimPrefix(authToken, "Bearer "))

		if err != nil {
      println(err.Error())
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Unauthorized",
			})
			return
		}

		ctx := context.WithValue(r.Context(), "user", claims.UserData)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
