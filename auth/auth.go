package auth

import (
	"errors"

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
