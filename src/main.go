package main

import (
	"github.com/VigneshSK17/cubimer-api/src/user"
	"log"
	"net/http"
	"time"

	"github.com/VigneshSK17/cubimer-api/src/db"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// TODO: Add UIUD to users

func main() {

	queries, err := db.Get()
	if err != nil {
		log.Fatalf("Could not set up database: %v", err)
	}
	db.Instance = queries

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/api", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Cubimer API"))
		})

		r.Post("/token", user.GenerateToken)

		r.Post("/user", user.RegisterUser)

		r.Route("/secured", func(r chi.Router) {
			r.Use(user.Auth)

			r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("pong"))
			})
		})
	})

	http.ListenAndServe(":8080", r)

}
