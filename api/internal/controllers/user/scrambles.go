package user

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	. "github.com/VigneshSK17/cubimer-api/api/internal/controllers/scramble"
	"github.com/VigneshSK17/cubimer-api/db"
)

type NewScramble struct {
    Id int64
    Username string
    Password string
    Cube string
    ScrambleStr string `json:"scrambleStr"`
    Time int64
}

type ModifyScramble struct {
    UserId int64 `json:"userId"`
    Username string
    Password string
    ScrambleId int64 `json:"scrambleId"`
    ScrambleStr *string `json:"scrambleStr"`
    Time *int64
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

	// TODO: Fix connecting to db
	db.ConnectScrambleDB()
	defer db.DB.Close()

	tableName := u.GetTableName()

	if _, err := db.DB.Exec(queryStr, tableName); err != nil {
		// return errors.New("A table for the user given could not be created");
		return err
	}

	return nil
}

func (u User) GetAllScrambles() ([]Scramble, error) {

    scrambles := []Scramble{}

    queryStr := fmt.Sprintf(`
        SELECT * FROM %s ORDER BY id DESC;
    `, u.GetTableName())

	// TODO: Fix connecting to db
	db.ConnectScrambleDB()
	defer db.DB.Close()

    if err := db.DB.Select(&scrambles, queryStr); err != nil {
        // return nil, errors.New("Could not access scrambles for the user.")
        return nil, err
    }

    return scrambles, nil

}

func (s *NewScramble) InsertScramble() (Scramble, error) {

    user := User{
        Id: s.Id,
        Username: s.Username,
        Password: s.Password,
    }
    
    query := fmt.Sprintf(`
        INSERT INTO %s (cube, scrambleStr, time, createdAt, updatedAt)
        VALUES (?, ?, ?, ?, ?);
    `, user.GetTableName())

	// TODO: Fix connecting to db
	db.ConnectScrambleDB()
	defer db.DB.Close()

    now := time.Now()

	result, err := db.DB.Exec(query, s.Cube, s.ScrambleStr, s.Time, now, now)
	if err != nil {
		return Scramble{}, errors.New("Could not create new scramble.")
	}

	var newId int64
	if newId, err = result.LastInsertId(); err != nil {
		return Scramble{}, errors.New("Count not find the newest scramble created. Please try again.")
	}

    scramble := Scramble{
        Id: newId,
        Cube: CubeType(s.Cube),
        ScrambleStr: s.ScrambleStr,
        Time: s.Time,
        CreatedAt: now,
        UpdatedAt: now,
    }

    return scramble, nil
}

func (s ModifyScramble) DeleteScramble() error {

    userTable := User{
        Id: s.UserId,
        Username: s.Username,
        Password: s.Password,
    }.GetTableName()

    queryStr := fmt.Sprintf(`DELETE FROM %s WHERE id=?;`, userTable)

    // TODO: Fix connecting to db
    db.ConnectScrambleDB()
    defer db.DB.Close()

    if _, err := db.DB.Exec(queryStr, s.ScrambleId); err != nil {
        return errors.New("Scramble could not be deleted")
    }

    return nil
}
