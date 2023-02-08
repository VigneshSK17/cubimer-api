package db

import (
	"context"
    "fmt"
	"os"

	"github.com/jmoiron/sqlx"
)

// type DB_URL int64
//
// const (
//     User DB_URL = iota
//     Scramble
// )
//
// func (d DB_URL) String() string {
//     urls := [...]string{
//         "../db/users.db",
//         "../db/scrambles.db",
//     }
//
//     if len(urls) < int(d) {
//         return ""
//     }
//     return urls[d]
// }

// func ConnectDB(dbType DB_URL) (*sqlx.DB, error){
//
//     db, err := sqlx.ConnectContext(context.Background(), "sqlite3", dbType.String())
//
//     if err != nil {
//         switch dbType {
//         case 0:
//             return nil, errors.New("Could not access users database.")
//         case 1:
//             return nil, errors.New("Could not access scrambles database.")
//         }
//     }
//
//     return db, nil
// }

type Dbinstance struct {
    Db *sqlx.DB
}

var DB Dbinstance

func ConnectDB(){

    dsn := fmt.Sprintf(
        "host=db user%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=America/New York",
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
    )

    db, err := sqlx.ConnectContext(context.Background(), "postgres", dsn)

    if err != nil {
        fmt.Fprintln(os.Stderr, err)
    }

    DB = Dbinstance {
        Db: db,
    }
}
