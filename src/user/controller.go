package user

import (
	"encoding/json"
	"github.com/go-chi/render"
	"net/http"

	. "github.com/VigneshSK17/cubimer-api/src/db"
	"golang.org/x/crypto/bcrypt"
)

type RegisterUserResponse struct {
	ID       int64
	Username string
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var createUser CreateUserParams

	if err := json.NewDecoder(r.Body).Decode(&createUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := hashPassword(&createUser); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user, err := Instance.Queries.CreateUser(ctx, createUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, RegisterUserResponse{
		ID:       user.ID,
		Username: user.Username,
	})
}

func hashPassword(user *CreateUserParams) error {
	b, err := bcrypt.GenerateFromPassword([]byte(user.Password), 17)
	if err != nil {
		return err
	}

	user.Password = string(b)
	return nil
}

func checkPassword(user User, providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
