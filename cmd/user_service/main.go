package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kirantiloh/gameroom/config"
	"github.com/kirantiloh/gameroom/internal/api/users"
)

func main() {
	r := chi.NewRouter()

	config.LoadConfig()

	userRoutes := chi.NewRouter()
	users.SetupRoutes(r)

	r.Mount("/api", userRoutes)

	http.ListenAndServe(":8080", r)
}
