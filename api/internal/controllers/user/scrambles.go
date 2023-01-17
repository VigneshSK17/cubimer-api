package user

import (
	"fmt"

	. "github.com/VigneshSK17/cubimer-api/api/internal/controllers/scramble"
	"github.com/VigneshSK17/cubimer-api/db"
)

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
