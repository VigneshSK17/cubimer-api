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
	Id       int64
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

    userDb, err := db.ConnectDB(0)
    if err != nil {
        return err
    }
    defer userDb.Close()

	result, err := userDb.Exec(query, u.Username, u.Password, u.Username)
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

    userDb, err := db.ConnectDB(0)
    if err != nil {
        return nil, err
    }
    defer userDb.Close()

	if err := userDb.Select(&users, query); err != nil {
		return nil, errors.New("Could not access users.")
	}

	return users, nil
}

func (u User) DeleteUser() error {
	const query string = `
        DELETE FROM users WHERE 
            id=?
            AND username=?
            AND password=?;
    `

    userDb, err := db.ConnectDB(0)
    if err != nil {
        return err
    }
    defer userDb.Close()

	if _, err := userDb.Exec(query, u.Id, u.Username, u.Password); err != nil {
		return errors.New("User to be deleted not found.")
	}

	return nil
}

func (u User) EditUser() error {
	const query string = `
        UPDATE users
        SET username = ?, password = ?
        WHERE id = ?;
    `

    userDb, err := db.ConnectDB(0)
    if err != nil {
        return err
    }
    defer userDb.Close()

	if _, err := userDb.Exec(query, u.Username, u.Password, u.Id); err != nil {
		return errors.New("User to be edited not found.")
	}

	return nil
}

// TODO: Store the passwords and such in a secure way
func (u *User) CheckUser() error {
	const query string = `
        SELECT id FROM users WHERE (
            username = ? AND password = ?
		);
    `

	var userId int64

    userDb, err := db.ConnectDB(0)
    if err != nil {
        return err
    }
    defer userDb.Close()

	row := userDb.QueryRow(query, u.Username, u.Password)

	if err := row.Scan(&userId); err != nil {
		return errors.New("Could not find user with given username and password.")
	}

	u.Id = userId
	return nil
}
