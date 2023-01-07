package user

import (
	"errors"
	"net/http"

	"github.com/VigneshSK17/cubimer-api/db"
)

// TODO: Create custom error type

// TODO: Modify ID to be more secure and UUID
const CreateQuery string = `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        username TEXT NOT NULL,
        password TEXT NOT NULL
    );
`

type User struct {
    Id int64
    Username string
    Password string
}

func (u *User) Bind(r *http.Request) error {
    return nil
}

func (u *User) InsertNewUser() error {
    const query string = `
        INSERT INTO users (username, password)
        VALUES (?, ?);
    `

    // TODO: Fix connecting to db
    db.ConnectDB()
    defer db.DB.Close()

    result, err := db.DB.Exec(query, u.Username, u.Password, u.Username)
    if err != nil {
        return errors.New("Could not create new user.")
    }

    var newId int64
    if newId, err = result.LastInsertId(); err != nil {
        return errors.New("Count not find the newest user created. Please try again.")
    }

    u.Id = newId

    return nil
}

func (u User) GetAllUsers() ([]User, error) {
    const query string = `
        SELECT * FROM users ORDER BY id ASC;
    `
    users := []User{}

    // TODO: Fix connecting to db
    db.ConnectDB()
    defer db.DB.Close()

    if err := db.DB.Select(&users, query); err != nil {
        return nil, errors.New("Could not access users.")
    }

    return users, nil
}
