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
    Id int64
    Username string
    Password string
    Cube string
    ScrambleStr string `json:"scrambleStr"`
    Time int64
}

type GetScramble struct {
    UserId int64
    Username string
    Password string
    ScrambleId int64 `json:"scrambleId"`
}

type ModifyScramble struct {
    UserId int64 `json:"userId"`
    Username string
    Password string
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
	return fmt.Sprintf("%s%d", u.Username, u.Id)
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

    queryStr := fmt.Sprintf(`
        SELECT * FROM %s ORDER BY id DESC;
    `, u.GetTableName())

    if err := DB.Db.Select(&scrambles, queryStr); err != nil {
        return nil, errors.New("Could not access scrambles for the user.")
    }

    return scrambles, nil

}

func (s GetScramble) GetScrambleFromId() (*Scramble, error) {
    
    userTable := User{
        Id: s.UserId,
        Username: s.Username,
        Password: s.Password,
    }.GetTableName()

    queryStr := fmt.Sprintf(`SELECT * FROM %s WHERE id = ?;`, userTable)

    var scramble Scramble
    if err := DB.Db.Get(&scramble, queryStr, s.ScrambleId); err != nil {
        return nil, errors.New("Could not access scramble of given id from user.")
    }

    return &scramble, nil
}


func (s *NewScramble) InsertScramble() (*Scramble, error) {

    user := User{
        Id: s.Id,
        Username: s.Username,
        Password: s.Password,
    }
    
    query := fmt.Sprintf(`
        INSERT INTO %s (cube, scrambleStr, time, createdAt, updatedAt)
        VALUES (?, ?, ?, ?, ?);
    `, user.GetTableName())

    now := time.Now()

	result, err := DB.Db.Exec(query, s.Cube, s.ScrambleStr, s.Time, now, now)
	if err != nil {
		return nil, errors.New("Could not create new scramble.")
	}

	var newId int64
	if newId, err = result.LastInsertId(); err != nil {
		return nil, errors.New("Count not find the newest scramble created. Please try again.")
	}

    scramble := Scramble{
        Id: newId,
        Cube: CubeType(s.Cube),
        ScrambleStr: s.ScrambleStr,
        Time: s.Time,
        CreatedAt: now,
        UpdatedAt: now,
    }

    return &scramble, nil
}

func (s ModifyScramble) DeleteScramble() error {

    userTable := User{
        Id: s.UserId,
        Username: s.Username,
        Password: s.Password,
    }.GetTableName()

    queryStr := fmt.Sprintf(`DELETE FROM %s WHERE id=?;`, userTable)

    if _, err := DB.Db.Exec(queryStr, s.ScrambleId); err != nil {
        return errors.New("Scramble could not be deleted")
    }

    return nil
}

func (s ModifyScramble) ModifyScramble() (*Scramble, error) {

    userTable := User{
        Id: s.UserId,
        Username: s.Username,
        Password: s.Password,
    }.GetTableName()

    now := time.Now()

    queryStr := fmt.Sprintf(`
        UPDATE %s
        SET scrambleStr = ?, time = ?, updatedAt = ?
        WHERE id = ?;
    `, userTable)

    if _, err := DB.Db.Exec(queryStr, s.ScrambleStr, s.Time, now, s.ScrambleId); err != nil {
        return nil, errors.New("Could not modify scramble with given information.")
    }

    return GetScramble{
        UserId: s.UserId,
        Username: s.Username,
        Password: s.Password,
        ScrambleId: s.ScrambleId,
    }.GetScrambleFromId()
}

