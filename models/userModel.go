package models

import (
	"log/slog"

	"github.com/vukovlevi/vukovlevidev/auth"
	"github.com/vukovlevi/vukovlevidev/db"
)

type Role string

type User struct {
    Id int
    Username string
    Password string
    Role Role
}

func (u *User) CreateUser() error {
    stmt := `INSERT INTO users (username, password, role) VALUES (?, ?, ?);`

    hashedPassword, err := auth.CreatePasswordHash(u.Password)
    if err != nil {
        slog.Error("could not create hashed password", "err", err)
        return err
    }

    _, err = db.DB.Exec(stmt, u.Username, string(hashedPassword), u.Role)
    if err != nil {
        slog.Error("could not create user is db", "err", err)
        return err
    }
    return nil
}

func GetUserByUsernameAndPassword(username, password string) (*User, error) {
    stmt := `SELECT * FROM users WHERE username = ?;`
    row := db.DB.QueryRow(stmt, username)

    user := new(User)
    err := row.Scan(&user.Id, &user.Username, &user.Password, &user.Role)
    if err != nil {
        return nil, err
    }

    err = auth.CheckPassword(user.Password, password)
    if err != nil {
        return nil, err
    }

    return user, nil
}
