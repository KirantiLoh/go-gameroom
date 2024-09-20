package users

import (
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/kirantiloh/gameroom/internal/users"
	"github.com/kirantiloh/gameroom/pkg/database"
)

func SetupRoutes(r *chi.Mux) {
	db, err := database.InitDB(os.Getenv("DATABASE_URL"))

	if err != nil {
		panic("Cannot create db conn")
	}

	r.Use(middleware.Logger)

	userRepo := users.NewRepository(db)
	userService := users.NewService(userRepo)
	userHandler := users.NewHandler(userService)

	r.Route("/auth", func(r chi.Router) {
		r.Post("/login", userHandler.LoginHandler)
		r.Post("/register", userHandler.RegisterHandler)
		r.Get("/verify/{hash}", userHandler.VerifyAccountHandler)
	})

}
