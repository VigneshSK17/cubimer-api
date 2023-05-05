package main

import (
	"context"
	"fmt"
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

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Cubimer API"))
	})

	r.Get("/test", func(w http.ResponseWriter, r *http.Request) {
		// TODO: Move to top
		if err != nil {
			fmt.Println(err.Error())
			w.Write([]byte("Could not access db"))
		} else {
			_, err := queries.Queries.ListUsers(context.Background())
			if err != nil {
				fmt.Println(err.Error())
				w.Write([]byte("Could not access users"))
			} else {
				w.Write([]byte("Can access users"))
				// w.Write([]byte{byte(len(users))})
			}
		}
	})

	http.ListenAndServe(":8080", r)

}
