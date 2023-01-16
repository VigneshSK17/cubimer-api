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
    ScrambleStr string
    Time int64 // milliseconds
    CreatedAt time.Time
    UpdatedAt time.Time


func (s *Scramble) Bind(r *http.Request) error {
    return nil
}
