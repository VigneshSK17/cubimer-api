package scramble

import (
	"encoding/json"
	. "github.com/VigneshSK17/cubimer-api/src/db"
	"github.com/go-chi/render"
	"net/http"
)

// TODO: Fix by checking save scramble userid capitalization
type UserId struct {
	UserID int64 `json:"UserID"`
}

func SaveScramble(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var createScramble CreateScrambleParams

	if err := json.NewDecoder(r.Body).Decode(&createScramble); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	scramble, err := Instance.Queries.CreateScramble(ctx, createScramble)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, scramble)
}

func GetScramblesByUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var userId UserId

	if err := json.NewDecoder(r.Body).Decode(&userId); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	scrambles, err := Instance.Queries.GetScramblesByUser(ctx, userId.UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, scrambles)
}
