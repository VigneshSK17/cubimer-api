package main

import (
	"net/http"

	"github.com/VigneshSK17/cubimer-api/api/internal/controllers/user"
	"github.com/VigneshSK17/cubimer-api/db"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	_ "github.com/mattn/go-sqlite3"
)

// TODO: Actually add passwords
var testSchema = `
DROP TABLE IF EXISTS user;
CREATE TABLE user (
    user_id INTEGER PRIMARY KEY,
    username VARCHAR(250) NOT NULL,
    password VARCHAR(250) DEFAULT NULL
);
`

func main() {

    // db, err := sqlx.Connect("sqlite3", "../db/test.db")
    // if err != nil {
    //     log.Fatalln(err)
    // }
    // db.MustExec(testSchema)
    //
    // tx := db.MustBegin()
    // // TODO: Use NamedExec w/ struct instead
    // tx.MustExec("INSERT INTO user (username) VALUES ($1)", "vigsk17")
    // tx.Commit()

    // Initialize db
    db.ConnectDB()

    r := chi.NewRouter()

    r.Use(middleware.Logger)
    r.Use(render.SetContentType(render.ContentTypeJSON))

    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        test_resp := map[string]string{"name": "Cubimer API"}
        render.JSON(w, r, test_resp)
    })

    // r.Mount("/users", user.UsersResource{}.Routes())

    r.Post("/users", user.UsersResource{}.Create)
    r.Get("/users", user.UsersResource{}.List)

    http.ListenAndServe("localhost:8080", r)

}
