package models

import (
	"github.com/labstack/gommon/log"
	"github.com/vukovlevi/vukovlevidev/db"
	"golang.org/x/crypto/bcrypt"
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

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 12)
    if err != nil {
        return err
    }

    _, err = db.DB.Exec(stmt, u.Username, string(hashedPassword), u.Role)
    if err != nil {
        Log.Warning("could not create user", "err", err)
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

    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
    if err != nil {
        return nil, err
    }

    log.Info("got user", "username", user.Username, "passwd", user.Password, "role", user.Role)
    return user, nil
}
