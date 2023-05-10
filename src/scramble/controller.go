package scramble

import (
	"encoding/json"
	. "github.com/VigneshSK17/cubimer-api/src/db"
	"github.com/go-chi/render"
	"net/http"
)

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
	type UserId struct {
		UserID int64 `json:"UserID"`
	}

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

func UpdateScramble(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var updateScramble UpdateScrambleParams

	if err := json.NewDecoder(r.Body).Decode(&updateScramble); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	scramble, err := Instance.Queries.UpdateScramble(ctx, updateScramble)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, scramble)
}

func DeleteScramble(w http.ResponseWriter, r *http.Request) {
	type ScrambleId struct {
		ScrambleID int64 `json:"ScrambleID"`
	}

	ctx := r.Context()
	var scrambleId ScrambleId

	if err := json.NewDecoder(r.Body).Decode(&scrambleId); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if _, err := Instance.Queries.DeleteScramble(ctx, scrambleId.ScrambleID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.Status(r, http.StatusNoContent)
}
