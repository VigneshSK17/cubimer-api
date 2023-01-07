package user

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/render"
)

// TODO: Put these renderers somewhere else
// -----------------

type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "Invalid request.",
		ErrorText:      err.Error(),
	}
}

func ErrRender(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 422,
		StatusText:     "Error rendering response.",
		ErrorText:      err.Error(),
	}
}

// -----------------

type UsersResource struct{}

// func (rs UsersResource) Routes() chi.Router {
//     r := chi.NewRouter()
//     r.Use(UserCtx) // TODO: Make this open DB
//
//     r.Get("/", rs.List) // TODO: Restrict this to admin only
//     r.Post("/", rs.Create)
//
//     r.Route("/{username}", func(r chi.Router) {
//         r.Get("/", rs.Login) // Return user ID to access scrambles
//         r.Put("/", rs.Update)
//         r.Delete("/", rs.Delete)
//     })
//
//     return r
// }

func UserCtx(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

    })
}

func (rs UsersResource) List(w http.ResponseWriter, r *http.Request) {

    testUser := User{}

    users, err := testUser.GetAllUsers()
    if err != nil {
        render.Render(w, r, ErrRender(err))
        return
    }

    render.Status(r, http.StatusCreated)
    render.JSON(w, r, users)

}

func (rs UsersResource) Create(w http.ResponseWriter, r *http.Request) {
        
    var newUser User

    if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
        render.Render(w, r, ErrInvalidRequest(err))
        return
    }

    if err := newUser.InsertNewUser(); err != nil {
        render.Render(w, r, ErrRender(err))
        return
    }

    render.Status(r, http.StatusCreated)
    render.JSON(w, r, newUser)

}

func (rs UsersResource) Login(w http.ResponseWriter, r *http.Request) {
    
}

func (rs UsersResource) Update(w http.ResponseWriter, r *http.Request) {

    var editedUser User 

    if err := json.NewDecoder(r.Body).Decode(&editedUser); err != nil {
        render.Render(w, r, ErrInvalidRequest(err))
        return
    }

    if err := editedUser.EditUser(); err != nil {
        render.Render(w, r, ErrRender(err))
        return
    }
}

func (rs UsersResource) Delete(w http.ResponseWriter, r *http.Request) {
    
    var user User

    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        render.Render(w, r, ErrInvalidRequest(err))
        return
    }

    if err := user.DeleteUser(); err != nil {
        render.Render(w, r, ErrRender(err))
        return
    }

    render.Status(r, http.StatusNoContent)

}
