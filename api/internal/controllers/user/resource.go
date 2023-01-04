package user

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type usersResource struct{}

func (rs usersResource) Routes() chi.Router {
    r := chi.NewRouter()
    r.Use(UserCtx) // TODO: Make this open DB

    r.Get("/", rs.List) // TODO: Restrict this to admin only
    r.Post("/", rs.Create)

    r.Route("/{username}", func(r chi.Router) {
        r.Get("/", rs.Login) // Return user ID to access scrambles
        r.Put("/", rs.Update)
        r.Delete("/", rs.Delete)
    })

    return r
}

func UserCtx(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

    })
}

func (rs usersResource) List(w http.ResponseWriter, r *http.Request) {
    
}

func (rs usersResource) Create(w http.ResponseWriter, r *http.Request) {
    
}

func (rs usersResource) Login(w http.ResponseWriter, r *http.Request) {
    
}

func (rs usersResource) Update(w http.ResponseWriter, r *http.Request) {
    
}

func (rs usersResource) Delete(w http.ResponseWriter, r *http.Request) {
    
}
