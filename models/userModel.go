package models

import "github.com/vukovlevi/vukovlevidev/db"

type Role string

type User struct {
    Id int
    Username string
    Password string
    Role Role
}

func (u *User) CreateUser() error {
    stmt := `INSERT INTO users (username, password, role) VALUES (?, ?, ?);`
    _, err := db.DB.Exec(stmt, u.Username, u.Password, u.Role)
    if err != nil {
        Log.Warning("could not create user", "err", err)
        return err
    }
    return nil
}

func GetUserByUsernameAndPassword(username, password string) (*User, error) {
    stmt := `SELECT id, username, role FROM users WHERE username = ? AND password = ?;`
    row := db.DB.QueryRow(stmt, username, password)

    user := new(User)
    err := row.Scan(&user.Id, &user.Username, &user.Role)
    if err != nil {
        return nil, err
    }
    return user, nil
}
