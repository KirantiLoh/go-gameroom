package rooms

import (
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/kirantiloh/gameroom/internal/room"
	"github.com/kirantiloh/gameroom/pkg/auth"
	"github.com/kirantiloh/gameroom/pkg/database"
)

func SetupRoutes(r *chi.Mux) {
	db, err := database.InitDB(os.Getenv("DATABASE_URL"))

	if err != nil {
		panic("Cannot create db conn")
	}

	r.Use(middleware.Logger)
  r.Use(auth.JWTMiddleware)

  roomRepo := room.NewRepository(db)
  roomService := room.NewService(roomRepo)

  roomHandler := room.NewHandler(roomService)

  r.Post("/", roomHandler.CreateRoomHandler)
  r.Get("/{uuid}", roomHandler.FindRoomByUUIDHandler)

}
