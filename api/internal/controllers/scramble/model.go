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

// TODO: Get this to scrape website for scrambles
// TODO: Get image from scrambling site
// https://www.worldcubeassociation.org/regulations/history/files/scrambles/scramble_cube.htm
func GenScramble() {

}
