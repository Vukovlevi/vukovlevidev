package auth

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"log"
	"log/slog"
	"time"

	"github.com/vukovlevi/vukovlevidev/db"
	"golang.org/x/crypto/bcrypt"
)

const (
    cost = 12
)

func ValidatePassword(password, passwordRepeat string) error {
    if len(password) < 8 {
        return errors.New("a jelszó túl rövid")
    }

    if password != passwordRepeat {
        return errors.New("a jelszavak nem egyeznek")
    }

    return nil
}

func CreatePasswordHash(password string) ([]byte, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), cost)
    if err != nil {
        return nil, err
    }

    return hashedPassword, nil
}

func CheckPassword(hash, password string) error {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    if err != nil {
        return err
    }
    return nil
}

func generateToken(length int) string {
    randomBytes := make([]byte, length, length)
    if _, err := rand.Read(randomBytes); err != nil {
        log.Fatalf("failed to generate random bytes: %s", err)
    }

    return base64.URLEncoding.EncodeToString(randomBytes)
}

func AuthenticateUser(userId int) (string, error) {
    stmt := `DELETE FROM sessions WHERE userId = ?;`
    _, err := db.DB.Exec(stmt, userId)
    if err != nil {
        slog.Error("could not delete previous session of user", "err", err)
        return "", err
    }

    sessionToken := generateToken(64)
    stmt = `INSERT INTO sessions VALUES (?, ?, ?);`
    _, err = db.DB.Exec(stmt, sessionToken, userId, time.Now().Add(time.Hour * 24))
    if err != nil {
        slog.Error("could not create new session for user", "err", err)
        return "", err
    }

    return sessionToken, nil
}

func AuthorizeAdminUser(sessionToken string) error {
    stmt := `SELECT role FROM users INNER JOIN sessions ON users.id = sessions.userId WHERE session_token = ? AND expires > datetime("now");`
    role := ""
    row := db.DB.QueryRow(stmt, sessionToken)

    err := row.Scan(&role)
    if err != nil {
        return err
    }

    if role != "admin" {
        return errors.New("not admin user")
    }

    return nil
}
