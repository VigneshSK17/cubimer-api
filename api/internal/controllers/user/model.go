package user

import (
	"errors"
	"net/http"

	. "github.com/VigneshSK17/cubimer-api/db"
)

// TODO: Create custom error type

type User struct {
	UserId   int64
	Username string
	Password string
}

func (u *User) Bind(r *http.Request) error {
	return nil
}

func (u *User) InsertNewUser() error {

	const query string = "INSERT INTO users (username, password) VALUES ($1, $2) RETURNING userId;"
	var newId int

	if err := DB.Db.QueryRowx(query, u.Username, u.Password).Scan(&newId); err != nil {
		return errors.New("Could not create new user.")
	}

	u.UserId = int64(newId)

	return nil
}

func (u User) GetAllUsers() ([]User, error) {
	const query string = `
        SELECT * FROM users ORDER BY userId ASC;
    `
	users := []User{}

	if err := DB.Db.Select(&users, query); err != nil {
		return nil, errors.New("Could not access users.")
	}

	return users, nil
}

// TODO: Fix this
func (u *User) DeleteUser() error {
	query := `DELETE FROM users WHERE userId=$1`

	if _, err := DB.Db.Exec(query, u.UserId); err != nil {
		// return errors.New("User to be deleted not found.")
		return err
	}

	return nil
}

func (u User) EditUser() error {
	const query string = `
        UPDATE users 
        SET username = :username, password = :password
        WHERE userId = :userid;
    `
    // var editedUser User

	// if err := DB.Db.Get(&editedUser, query, u.Username, u.Password, u.UserId); err != nil {
	if _, err := DB.Db.NamedExec(query, u); err != nil {
		// return errors.New("User to be edited not found.")
		return err
	}

	return nil
}

// TODO: Store the passwords and such in a secure way
func (u *User) CheckUser() error {
	const query string = `
        SELECT userId FROM users WHERE (
            username = $1 AND password = $2
		);
    `

	var userId int64

	if err := DB.Db.QueryRowx(query, u.Username, u.Password).Scan(&userId); err != nil {
		return errors.New("Could not find user with given username and password.")
	}

	u.UserId = int64(userId)
	return nil
}
