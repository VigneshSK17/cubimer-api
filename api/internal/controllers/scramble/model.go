package scramble

import (
	"net/http"
	"time"
)

type CubeType string
const (
    ThreeByThree CubeType = "3x3" 
)

type Scramble struct {
    Id int64
    Cube CubeType
    ScrambleStr string `db:"scrambleStr"`
    Time int64 // milliseconds
    CreatedAt time.Time `db:"createdAt"`
    UpdatedAt time.Time `db:"updatedAt"`
}   

func (s *Scramble) Bind(r *http.Request) error {
    return nil
}
