package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kirantiloh/gameroom/config"
	"github.com/kirantiloh/gameroom/internal/api/rooms"
)

func main() {
	r := chi.NewRouter()

	config.LoadConfig()
  roomRoutes := chi.NewRouter()
  rooms.SetupRoutes(roomRoutes)

  r.Mount("/rooms", roomRoutes)

	http.ListenAndServe(":8001", r)
}
