package scramble

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	. "github.com/VigneshSK17/cubimer-api/api/internal/renderers"
)

type ScramblesResource struct{}

func (rs ScramblesResource) GetScramble(w http.ResponseWriter, r *http.Request) {
    
    cubeTypeStr := chi.URLParam(r, "cubeType") 
    var scrambleJson map[string]string

    switch cubeTypeStr {
    case "3x3":
        scrambleJson = map[string]string{"scramble": GenScrambleOfficial(ThreeByThree)}
    default:
        render.Render(w, r, ErrInvalidRequest(errors.New("Could not find cube type given.")))
        return
    }

    render.Status(r, http.StatusCreated)
    render.JSON(w, r, scrambleJson)
     
}
