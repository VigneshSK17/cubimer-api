package scramble

import (
	"net/http"
    "math/rand"
	"time"
    "strings"
)

type CubeType string
const (
    ThreeByThree CubeType = "3x3" 
)

func (c CubeType) String() string {
    return string(c)
}

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

func GenScrambleOfficial(scrambleType CubeType) string {

    // Works for 2x2 and 3x3
    scramblesOpts := []string{"R","R'","R2","L","L'","L2","U","U'","U2","D","D'","D2","F","F'","F2","B","B'","B2"} 
    var scrambleArr []string

    // TODO: Modify based on cube being scrambled
    var moves int
    var prevJ int32 = -1

    // Makes sure scrambles are random
    rand.Seed(time.Now().UnixNano())

    switch(scrambleType) {
    case ThreeByThree:
        moves = 20
        scrambleArr = make([]string, moves)
    }

    for i := 0; i < moves; i++ {
        
        j := rand.Int31n(int32(len(scramblesOpts)))
        if prevJ == j {
            i += 1
        }

        scrambleArr[i] = scramblesOpts[j] 
    }

    return strings.Join(scrambleArr, " ")

}
