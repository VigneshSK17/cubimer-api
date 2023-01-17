package main

import (
	"net/http"

	"github.com/VigneshSK17/cubimer-api/api/internal/controllers/user"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		test_resp := map[string]string{"name": "Cubimer API"}
		render.JSON(w, r, test_resp)
	})

	/** Users routes **/

	r.Route("/users", func(r chi.Router) {
		r.Post("/", user.UsersResource{}.Create)
		r.Patch("/", user.UsersResource{}.List)
		r.Delete("/", user.UsersResource{}.Delete)
		r.Put("/", user.UsersResource{}.Update)

		r.Get("/", user.UsersResource{}.Login)

	})

    /** Scrambles routes **/
    r.Route("/scrambles", func(r chi.Router) {
        r.Get("/", user.UsersResource{}.ListScrambles)
    })

	http.ListenAndServe("localhost:8080", r)

}
