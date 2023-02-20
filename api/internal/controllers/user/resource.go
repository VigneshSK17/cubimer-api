package user

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/render"

	. "github.com/VigneshSK17/cubimer-api/api/internal/renderers"
)

type UsersResource struct{}
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

	var user User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	if err := user.CheckUser(); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, user)

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

	render.Status(r, http.StatusNoContent)
	render.JSON(w, r, editedUser)
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

func (rs UsersResource) ListScrambles(w http.ResponseWriter, r *http.Request) {

    var user User

    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        render.Render(w, r, ErrInvalidRequest(err))
        return
    }

    scrambles, err := user.GetAllScrambles()
    if err != nil {
        render.Render(w, r, ErrRender(err))
        return
    }

    render.Status(r, http.StatusCreated)    
    render.JSON(w, r, scrambles)

}

func (rs UsersResource) SaveScramble(w http.ResponseWriter, r *http.Request) {

    var newScramble NewScramble

    if err := json.NewDecoder(r.Body).Decode(&newScramble); err != nil {
        render.Render(w, r, ErrInvalidRequest(err))
        return
    }

    scramble, err := newScramble.InsertScramble()
    if err != nil {
        render.Render(w, r, ErrRender(err))
        return
    }

    render.Status(r, http.StatusCreated)
    render.JSON(w, r, *scramble)

}

func (rs UsersResource) DeleteScramble(w http.ResponseWriter, r *http.Request) {
    
    var scramble ModifyScramble

    if err := json.NewDecoder(r.Body).Decode(&scramble); err != nil {
        render.Render(w, r, ErrInvalidRequest(err))
        return
    }

    if err := scramble.DeleteScramble(); err != nil {
        render.Render(w, r, ErrRender(err))
        return
    }

    render.Status(r, http.StatusNoContent)
    render.JSON(w, r, scramble)

}

func (rs UsersResource) EditScramble(w http.ResponseWriter, r *http.Request) {

    var scramble ModifyScramble

    if err := json.NewDecoder(r.Body).Decode(&scramble); err != nil {
        render.Render(w, r, ErrInvalidRequest(err))
        return
    }

    modifiedScramble, err := scramble.ModifyScramble()
    if err != nil {
        render.Render(w, r, ErrRender(err))
        return
    }

    render.Status(r, http.StatusNoContent)
    render.JSON(w, r, *modifiedScramble)

}
