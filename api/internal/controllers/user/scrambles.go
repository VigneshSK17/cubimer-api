package user

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	. "github.com/VigneshSK17/cubimer-api/api/internal/controllers/scramble"
	. "github.com/VigneshSK17/cubimer-api/db"
)

type NewScramble struct {
    UserId int64
    Cube string
    ScrambleStr string `json:"scrambleStr"`
    Time int64
}

type ModifyScramble struct {
    UserId int64 `json:"userId"`
    ScrambleId int64 `json:"scrambleId"`
    ScrambleStr string `json:"scrambleStr"`
    Time int64
}

func (u *NewScramble) Bind(r *http.Request) error {
	return nil
}

func (u *ModifyScramble) Bind(r *http.Request) error {
	return nil
}

// Table name for a user
func (u User) GetTableName() string {
	return fmt.Sprintf("%s%d", u.Username, u.UserId)
}

// Initializes new table of scrambles for user
func (u User) CreateScramblesTable() error {
    queryStr := fmt.Sprintf(`
        CREATE TABLE %s (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            cube TEXT NOT NULL,
            scrambleStr TEXT NOT NULL,
            time INTEGER NOT NULL,
            createdAt datetime NOT NULL,
	        updatedAt datetime NOT NULL
        );`, u.GetTableName())

	tableName := u.GetTableName()

	if _, err := DB.Db.Exec(queryStr, tableName); err != nil {
		return err
	}

	return nil
}

func (u User) GetAllScrambles() ([]Scramble, error) {

    scrambles := []Scramble{}

    queryStr := `SELECT * FROM scrambles
        JOIN users ON scrambles.userId = users.userId
        WHERE scrambles.userId = ?;
    `

    if err := DB.Db.Select(&scrambles, queryStr, u.UserId); err != nil {
        return nil, errors.New("Could not access scrambles for the user.")
    }

    return scrambles, nil

}

func GetScrambleFromId(scrambleId int64) (*Scramble, error) {
    
    queryStr := `SELECT * FROM scrambles WHERE scrambleId = ?;`

    var scramble Scramble
    if err := DB.Db.Get(&scramble, queryStr, scrambleId); err != nil {
        return nil, errors.New("Could not access scramble of given id from user.")
    }

    return &scramble, nil
}


func (s NewScramble) InsertScramble() (*Scramble, error) {

    query := `
        INSERT INTO scrambles (cube, scrambleStr, time, createdAt, updatedAt, userId)
        VALUES (?, ?, ?, ?, ?, ?);
    `

    now := time.Now()

	result, err := DB.Db.Exec(query, s.Cube, s.ScrambleStr, s.Time, now, now, s.UserId)
	if err != nil {
		return nil, errors.New("Could not create new scramble.")
	}

	var newId int64
	if newId, err = result.LastInsertId(); err != nil {
		return nil, errors.New("Count not find the newest scramble created. Please try again.")
	}

    scramble := Scramble{
        ScrambleId: newId,
        UserId: s.UserId,
        Cube: CubeType(s.Cube),
        ScrambleStr: s.ScrambleStr,
        Time: s.Time,
        CreatedAt: now,
        UpdatedAt: now,
    }

    return &scramble, nil
}

func (s ModifyScramble) DeleteScramble() error {

    queryStr := `DELETE FROM scrambles WHERE scrambleId=?;`

    if _, err := DB.Db.Exec(queryStr, s.ScrambleId); err != nil {
        return errors.New("Scramble could not be deleted")
    }

    return nil
}

func (s ModifyScramble) ModifyScramble() (*Scramble, error) {

    now := time.Now()

    queryStr := `
        UPDATE scrambles
        SET scrambleStr = ?, time = ?, updatedAt = ?
        WHERE scrambleId = ?;
    `

    if _, err := DB.Db.Exec(queryStr, s.ScrambleStr, s.Time, now, s.ScrambleId); err != nil {
        return nil, errors.New("Could not modify scramble with given information.")
    }

    return GetScrambleFromId(s.ScrambleId)
}

